/*
Copyright 2020 DevSpace Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	apiserver "github.com/loft-sh/apiserver/pkg/server"
	"github.com/loft-sh/kiosk/pkg/apis"
	"github.com/loft-sh/kiosk/pkg/apis/tenancy"
	tenancyv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/tenancy/v1alpha1"
	_ "github.com/loft-sh/kiosk/pkg/apiserver/registry"
	"github.com/loft-sh/kiosk/pkg/leaderelection"
	"github.com/loft-sh/kiosk/pkg/manager/blockingcacheclient"
	"github.com/loft-sh/kiosk/pkg/manager/quota/controller"
	"github.com/loft-sh/kiosk/pkg/openapi"
	"github.com/loft-sh/kiosk/pkg/store/apiservice"
	"github.com/loft-sh/kiosk/pkg/store/crd"
	"github.com/loft-sh/kiosk/pkg/store/validatingwebhookconfiguration"
	"github.com/loft-sh/kiosk/pkg/util/certhelper"
	"github.com/loft-sh/kiosk/pkg/util/log"
	"github.com/loft-sh/kiosk/pkg/util/secret"
	"github.com/loft-sh/kiosk/pkg/watch"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/klog"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	"math/rand"
	"os"
	"path/filepath"
	client2 "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"strconv"
	"time"

	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	"github.com/loft-sh/kiosk/pkg/manager/controllers"
	"github.com/loft-sh/kiosk/pkg/manager/webhooks"
	"k8s.io/apimachinery/pkg/runtime"
	genericfeatures "k8s.io/apiserver/pkg/features"
	featureutil "k8s.io/apiserver/pkg/util/feature"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	// +kubebuilder:scaffold:imports

	// Make sure dep tools picks up these dependencies
	_ "github.com/go-openapi/loads"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	// API extensions are not in the above scheme set,
	// and must thus be added separately.
	_ = apiextensionsv1beta1.AddToScheme(scheme)
	_ = apiextensionsv1.AddToScheme(scheme)
	_ = apiregistrationv1.AddToScheme(scheme)

	_ = tenancy.AddToScheme(scheme)
	_ = tenancyv1alpha1.AddToScheme(scheme)
	_ = configv1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// set global logger
	if os.Getenv("DEBUG") == "true" {
		ctrl.SetLogger(log.NewLog(0))
	} else {
		ctrl.SetLogger(log.NewLog(2))
	}

	// retrieve in cluster config
	config := ctrl.GetConfigOrDie()

	// set qps, burst & timeout
	config.QPS = 80
	config.Burst = 100
	config.Timeout = 0

	// create a new temporary client
	uncachedClient, err := client2.New(config, client2.Options{Scheme: scheme})
	if err != nil {
		setupLog.Error(err, "unable to create client")
		os.Exit(1)
	}

	// Make sure the certificates are there
	err = secret.EnsureCertSecrets(context.Background(), uncachedClient)
	if err != nil {
		setupLog.Error(err, "unable to generate certificates")
		os.Exit(1)
	}

	// Make sure the needed crds are installed in the cluster
	err = installCRDS(uncachedClient)
	if err != nil {
		klog.Fatal(err)
	}

	// create the manager
	mgr, err := ctrl.NewManager(config, ctrl.Options{
		NewClient:          blockingcacheclient.NewCacheClient,
		Scheme:             scheme,
		MetricsBindAddress: ":8080",
		CertDir:            certhelper.WebhookCertFolder,
		LeaderElection:     false,
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Inject the cached, uncached client and scheme
	injectClient(mgr.GetClient(), uncachedClient, scheme)

	// Add required indices
	err = controllers.AddManagerIndices(mgr.GetCache())
	if err != nil {
		setupLog.Error(err, "unable to set manager indices")
		os.Exit(1)
	}

	stopChan := make(chan struct{})
	ctx := signals.SetupSignalHandler()
	ctrlCtx := controllers.NewControllerContext(mgr, stopChan)

	// Register webhooks
	err = webhooks.Register(ctrlCtx)
	if err != nil {
		setupLog.Error(err, "unable to register webhooks")
		os.Exit(1)
	}

	// Start controller context
	go ctrlCtx.Start()

	// Start the local manager
	go func() {
		setupLog.Info("starting manager")
		err = mgr.Start(ctx)
		if err != nil {
			panic(err)
		}
	}()

	// Make sure the manager is synced
	mgr.GetCache().WaitForCacheSync(ctx)

	// Start watch registries
	err = watch.NamespaceRegistry.Start(ctx, mgr.GetCache(), &corev1.Namespace{})
	if err != nil {
		klog.Fatalf("start namespace informer: %v", err)
	}
	err = watch.AccountRegistry.Start(ctx, mgr.GetCache(), &configv1alpha1.Account{})
	if err != nil {
		klog.Fatalf("start account informer: %v", err)
	}

	// Start the api server
	go func() {
		if os.Getenv("SERVER_SIDE_APPLY_ENABLED") != "true" {
			err := featureutil.DefaultMutableFeatureGate.Set(string(genericfeatures.ServerSideApply) + "=false")
			if err != nil {
				panic(err)
			}
		}

		err = apiserver.StartAPIServer(&apiserver.StartOptions{
			Apis:                  apis.GetAllApiBuilders(),
			GetOpenAPIDefinitions: openapi.GetOpenAPIDefinitions,
			TweakServerOptions: func(o *apiserver.ServerOptions) {
				o.DisableWebhooks = os.Getenv("DISABLE_WEBHOOKS") == "true"
				o.RecommendedOptions.SecureServing.ServerCert.CertKey.CertFile = filepath.Join(certhelper.APIServiceCertFolder, "tls.crt")
				o.RecommendedOptions.SecureServing.ServerCert.CertKey.KeyFile = filepath.Join(certhelper.APIServiceCertFolder, "tls.key")

				var err error
				apiServicePort := 8443
				apiServicePortEnv := os.Getenv("APISERVICE_PORT")
				if apiServicePortEnv != "" {
					apiServicePort, err = strconv.Atoi(apiServicePortEnv)
					if err != nil {
						panic(fmt.Sprintf("parsing api service port %s: %v", apiServicePortEnv, err))
					}
				}

				o.RecommendedOptions.SecureServing.BindPort = apiServicePort
			},
		})
		if err != nil {
			panic(err)
		}
	}()

	// start leader election for controllers
	go func() {
		err = leaderelection.StartLeaderElection(ctx, scheme, config, func() error {
			// setup ValidatingWebhookConfiguration
			if os.Getenv("UPDATE_WEBHOOK") != "false" {
				err = validatingwebhookconfiguration.EnsureValidatingWebhookConfiguration(context.Background(), mgr.GetClient())
				if err != nil {
					setupLog.Error(err, "unable to set up validating webhook configuration")
					os.Exit(1)
				}
			}

			// setup APIService
			if os.Getenv("UPDATE_APISERVICE") != "false" {
				err = apiservice.EnsureAPIService(context.Background(), mgr.GetClient())
				if err != nil {
					setupLog.Error(err, "unable to set up apiservice")
					os.Exit(1)
				}
			}

			// Register generic controllers
			err = controllers.Register(mgr)
			if err != nil {
				return errors.Wrap(err, "unable to register controller")
			}

			// Register quota controller
			err = controller.Register(ctrlCtx)
			if err != nil {
				return errors.Wrap(err, "unable to register quota controller")
			}

			return nil
		})
		if err != nil {
			klog.Fatalf("Error starting leader election: %v", err)
		}
	}()

	// Wait till stopChan is closed
	<-stopChan
}

func installCRDS(uncachedClient client2.Client) error {
	klog.Info("Installing crds...")

	builder := crd.NewBuilder(uncachedClient)
	_, err := builder.CreateCRDs(context.Background(), apis.TypeDefinitions...)
	return err
}

func injectClient(cachedClient client2.Client, uncachedClient client2.Client, scheme *runtime.Scheme) {
	tenancy.CachedClient = cachedClient
	tenancy.UncachedClient = uncachedClient
	tenancy.Scheme = scheme
}
