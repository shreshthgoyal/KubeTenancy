# KubeTenancy

```

alias kt=kiosk

sudo chmod 666 /var/run/docker.sock
export KUBECONFIG="C:\Users\sg\.kube\config"


kind delete cluster

 kind create cluster

kubectl get all

kubectl auth can-i "*" "*" --all-namespaces
kubectl auth can-i "*" namespace
kubectl auth can-i "*" clusterrole
kubectl auth can-i "*" crd

kubectl create namespace kt
helm install kiosk --repo https://charts.devspace.sh/ kiosk --namespace kt --atomic

 kubectl get pod -n kt

kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/account.yaml

kubectl get accounts --as=john

kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/rbac-creator.yaml
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/space.yaml --as=john
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/space.yaml --as=mary
kubectl get space --as=john
kubectl get space johns-space --as=mary
kubectl apply -n johns-space --as=john -f https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/application/deployment.yaml

kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/account-space-limit.yaml
kubectl get space --as=john
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/space-2.yaml --as=john
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/space-3.yaml --as=john
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/accountquota.yaml
kubectl apply -f https://raw.githubusercontent.com/shreshthgoyal/KubeTenancy/main/yaml/template-manifests.yaml
kubectl apply -n johns-space --as=john -f https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/application/deployment.yaml
kubectl patch -n kube-system deployment metrics-server --type=json   -p '[{"op":"add","path":"/spec/template/spec/containers/0/args/-","value":"--kubelet-insecure-tls"}]'
kubectl top node

```
