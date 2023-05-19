# seed

## installation

### code

initialize
```shell
go run main.go initialize \
  --billing-account-name <billing-account-name> \
  --billing-profile-name <billing-profile-name> \
  --invoice-section-name <invoice-section-name> \
  --subscription-id <subscription-id> \
  --resource-group-name <resource-group-name> \
  --resource-group-location <resource-group-location> \
  --kubernetes-cluster-name <kubernetes-cluster-name> \
  --kubernetes-cluster-location <kubernetes-cluster-location> \
  --kubernetes-default-node-pool-name <kubernetes-default-node-pool-name> \
  --kubernetes-default-node-pool-vm-size <kubernetes-default-node-pool-vm-size> \
  --kubernetes-default-node-pool-count <kubernetes-default-node-pool-count> \
  --kubernetes-cluster-identity-type <kubernetes-cluster-identity-type>
```

generate
```shell
go run main.go azure generate seed
```

### seed

initialize
```shell
seed azure initialize \
  --billing-account-name <billing-account-name> \
  --billing-profile-name <billing-profile-name> \
  --invoice-section-name <invoice-section-name> \
  --subscription-id <subscription-id> \
  --resource-group-name <resource-group-name> \
  --resource-group-location <resource-group-location> \
  --kubernetes-cluster-name <kubernetes-cluster-name> \
  --kubernetes-cluster-location <kubernetes-cluster-location> \
  --kubernetes-default-node-pool-name <kubernetes-default-node-pool-name> \
  --kubernetes-default-node-pool-vm-size <kubernetes-default-node-pool-vm-size> \
  --kubernetes-default-node-pool-count <kubernetes-default-node-pool-count> \
  --kubernetes-cluster-identity-type <kubernetes-cluster-identity-type>
```

generate
```shell
seed azure generate seed
```

### terraform cdk

deploy
```shell
cdktf deploy
```

destroy
```shell
cdktf destroy
```

### clusterctl

prerequisites
```shell
```shell
az ad sp create-for-rbac --role contributor --scopes="/subscriptions/f2fd252d-9490-475d-a4dc-2974e02847b9"
```

```shell
{
  "appId": "7a3e84b4-05d9-4c3f-95df-b04b6c5517a9",
  "displayName": "azure-cli-2023-05-17-20-33-23",
  "password": "",
  "tenant": "201ed03f-5b3a-4b83-a14d-1e792e8861ea"
}
```

```shell
./script/azure_cluster_install
```

```shell
secret/cluster-identity-secret created
Fetching providers
Installing cert-manager Version="v1.11.1"
Waiting for cert-manager to be available...
Installing Provider="cluster-api" Version="v1.4.2" TargetNamespace="capi-system"
Installing Provider="bootstrap-kubeadm" Version="v1.4.2" TargetNamespace="capi-kubeadm-bootstrap-system"
Installing Provider="control-plane-kubeadm" Version="v1.4.2" TargetNamespace="capi-kubeadm-control-plane-system"
Installing Provider="infrastructure-azure" Version="v1.9.2" TargetNamespace="capz-system"

Your management cluster has been initialized successfully!

You can now create your first workload cluster by running the following:

  clusterctl generate cluster [name] --kubernetes-version [version] | kubectl apply -f -
```

```shell
clusterctl generate cluster capi-quickstart --flavor development \
  --kubernetes-version v1.25.9 \
  --control-plane-machine-count=1 \
  --worker-machine-count=1 \
  > capi-quickstart.yaml
```

```shell
clusterctl get kubeconfig capi-quickstart > capi-quickstart.kubeconfig
```

### kubectl

```shell
kubectl get cluster
```

```shell
clusterctl describe cluster capi-quickstart
```

```shell
kubectl get kubeadmcontrolplane
```

```shell
kubectl --kubeconfig=./capi-quickstart.kubeconfig get nodes
```
### seed

```shell
seed generate example-cluster
```

### helm

#### https://github.com/projectcalico/calico
cni - calico
```shell
helm repo add projectcalico https://docs.tigera.io/calico/charts --kubeconfig=./capi-quickstart.kubeconfig
```

```shell
helm install calico projectcalico/tigera-operator --kubeconfig=capi-quickstart.kubeconfig -f https://raw.githubusercontent.com/kubernetes-sigs/cluster-api-provider-azure/main/templates/addons/calico/values.yaml --namespace tigera-operator --create-namespace
```


### teardown

```shell
kubectl delete -f capi-quickstart.yaml 
```

```shell
cdktf destroy
```

## seed cli

### generate

```shell
seed generate
```

### initialize

```shell
seed initialize
```