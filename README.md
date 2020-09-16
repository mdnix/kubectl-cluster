# kubectl-cluster

The kubectl-cluster plugin lets you run kubectl commands across a number of specified clusters.
It is possible to run any kubectl command since this plugin is basically a wrapper around kubectl.

## Install

Linux example:

```bash
git clone https://github.com/mdnix/kubectl-cluster.git
make compile
cp _out/linux/kubectl-cluster /usr/local/bin/
```

## Usage

### Define the cluster config

```bash
cat /home/marco/.clusters

clusters:
  - name: minikube
    config: /home/marco/.kube/config-mini
    tags: test
  - name: okd
    config: /home/marco/.kube/config-okd
    tags: prod
```

### List available clusters

```bash
$ kubectl cluster list
```

### Get Pods of all clusters using the current context

```bash
$ kubectl cluster run get pods 
```

### Get Pods of all clusters, overriding the current namespace of the context

**_NOTE:_**
When using flags for regular kubectl commands "--" has to be added to signify the end of command options


```bash
$ kubectl cluster run -- get pods --all-namespaces
```

### Get Pods of specified clusters using the current context

```bash
$ kubectl cluster run --targets okd,minikube get pods 
```


### Get Pods of clusters containing the specified tag in the config

```bash
$ kubectl cluster run --targets okd,minikube get pods 
```


