package validatingwebhookconfiguration

import (
	"context"
	"io/ioutil"

	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	"github.com/loft-sh/kiosk/pkg/util/certhelper"
	"github.com/loft-sh/kiosk/pkg/util/clienthelper"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// ValidatingWebhookConfigurationName is the name of the validating webhook configuration
	ValidatingWebhookConfigurationName = "kiosk"
)

// EnsureValidatingWebhookConfiguration makes sure the validating webhook configuration is up and correct
func EnsureValidatingWebhookConfiguration(ctx context.Context, client client.Client) error {
	config := &admissionregistrationv1.ValidatingWebhookConfiguration{}
	err := client.Get(ctx, types.NamespacedName{Name: ValidatingWebhookConfigurationName}, config)
	if err != nil {
		if kerrors.IsNotFound(err) == false {
			return err
		}

		config.Name = ValidatingWebhookConfigurationName
		err = prepareValidatingWebhookConfiguration(config)
		if err != nil {
			return err
		}

		return client.Create(ctx, config)
	}

	err = prepareValidatingWebhookConfiguration(config)
	if err != nil {
		return err
	}

	return client.Update(ctx, config)
}

func prepareValidatingWebhookConfiguration(config *admissionregistrationv1.ValidatingWebhookConfiguration) error {
	caBundleData, err := ioutil.ReadFile(filepath.Join(certhelper.WebhookCertFolder, "ca.crt"))
	if err != nil {
		return err
	}

	failPolicy := admissionregistrationv1.Fail
	namespaceScope := admissionregistrationv1.NamespacedScope
	quotaPath := "/quota"
	validatePath := "/validate"
	namespace, err := clienthelper.CurrentNamespace()
	if err != nil {
		return err
	}

	sideEffects := admissionregistrationv1.SideEffectClassNone
	config.Webhooks = []admissionregistrationv1.ValidatingWebhook{
		{
			Name:          "accountquota.kiosk.sh",
			FailurePolicy: &failPolicy,
			SideEffects:   &sideEffects,
			ClientConfig: admissionregistrationv1.WebhookClientConfig{
				Service: &admissionregistrationv1.ServiceReference{
					Namespace: namespace,
					Name:      certhelper.WebhookServiceName,
					Path:      &quotaPath,
				},
				CABundle: caBundleData,
			},
			NamespaceSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "kiosk.sh/account",
						Operator: metav1.LabelSelectorOpExists,
					},
				},
			},
			Rules: []admissionregistrationv1.RuleWithOperations{
				{
					Operations: []admissionregistrationv1.OperationType{admissionregistrationv1.Create, admissionregistrationv1.Update},
					Rule: admissionregistrationv1.Rule{
						APIGroups:   []string{"*"},
						APIVersions: []string{"*"},
						Resources:   []string{"*"},
						Scope:       &namespaceScope,
					},
				},
			},
			AdmissionReviewVersions: []string{admissionregistrationv1.SchemeGroupVersion.Version, "v1beta1"},
		},
		{
			Name:          "config.kiosk.sh",
			FailurePolicy: &failPolicy,
			SideEffects:   &sideEffects,
			ClientConfig: admissionregistrationv1.WebhookClientConfig{
				Service: &admissionregistrationv1.ServiceReference{
					Namespace: namespace,
					Name:      certhelper.WebhookServiceName,
					Path:      &validatePath,
				},
				CABundle: caBundleData,
			},
			Rules: []admissionregistrationv1.RuleWithOperations{
				{
					Operations: []admissionregistrationv1.OperationType{admissionregistrationv1.Create, admissionregistrationv1.Update},
					Rule: admissionregistrationv1.Rule{
						APIGroups:   []string{configv1alpha1.SchemeGroupVersion.Group},
						APIVersions: []string{configv1alpha1.SchemeGroupVersion.Version},
						Resources:   []string{"*"},
					},
				},
			},
			AdmissionReviewVersions: []string{admissionregistrationv1.SchemeGroupVersion.Version, "v1beta1"},
		},
	}

	delete(config.Annotations, "cert-manager.io/inject-ca-from")
	return nil
}
