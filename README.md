# seed

## installation

```shell

```

### kind

```shell

```

```shell
kind create cluster
```

### clusterctl

```shell
export CLUSTER_TOPOLOGY=true
```

```shell
clusterctl init --infrastructure docker
```

```shell
export SERVICE_CIDR=["10.96.0.0/12"]
```

```shell
export POD_CIDR=["192.168.0.0/16"]
```

```shell
export SERVICE_DOMAIN="k8s.test"
```

```shell
clusterctl generate cluster capi-quickstart --flavor development \
  --kubernetes-version v1.25.9 \
  --control-plane-machine-count=1 \
  --worker-machine-count=1 \
  > capi-quickstart.yaml
```

### seed

```shellr
seed generate example-cluster
```

### terraform cdk

```shell
cdktf deploy
```

### helm

```shell
helm install 
```