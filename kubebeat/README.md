# POC Documentation

I generated this repo using the [beats development guide](https://www.elastic.co/guide/en/beats/devguide/current/newbeat-generate.html).
The kube-api call is based on the [k8s go-client example](https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration).

The interesting files are:
* `beater/kubebeat.go` - the beats logic
* `kubebeat.yml` - the beats config
* `Dockerfile` - runs the beat dockerized with debug flags
* `pod.yaml` - deploy the beat


## Table of contents
- [POC Documentation](#poc-documentation)
  - [Table of contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Running the Kubebeat](#running-the-kubebeat)
    - [Clean up](#clean-up)
    - [Remote Debugging](#remote-debugging)
- [{Beat}](#beat)
  - [Getting Started with {Beat}](#getting-started-with-beat)
    - [Requirements](#requirements)
    - [Init Project](#init-project)
    - [Build](#build)
    - [Run](#run)
    - [Test](#test)
    - [Update](#update)
    - [Cleanup](#cleanup)
    - [Clone](#clone)
  - [Packaging](#packaging)
  - [Build Elastic-Agent Docker with pre-packaged kubebeat](#build-elastic-agent-docker-with-pre-packaged-kubebeat)


## Prerequisites
**Please make sure that you run the following instructions within the `kubebeat` directory.**

1. Elasticsearch with the default username & password (`elastic` & `changeme`) running on the default port (`http://localhost:9200`)
2. Kibana with running on the default port (`http://localhost:5601`)
3. Minikube cluster running locally (`minikube start`)

First initialize the git submodule:

4. Clone the git submodule of the CIS rules:
```
    $ git submodule update --init
```
5. Comment the Rego code that uses data.yaml (Temporary fix) - go to compliance/cis_k8s/cis_k8s.rego and comment the following line of code:

    ```
    data.activated_rules.cis_k8s[rule_id]
   ```

    $ patch kubebeat.yml kubebeat_minikube.yml.patch

Then package it to a docker image using the provided Dockerfile to run it on Kubernetes:

Running a [Minikube](https://minikube.sigs.k8s.io/docs/) cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube:

    eval $(minikube docker-env)
    docker build -t kubebeat .

If you are not using Minikube, you should build this image and push it to a registry that your Kubernetes cluster can pull from.

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions:

    $ kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default

Then, run the image in a Pod with a single instance Deployment:

    $ kubectl apply -f pod.yml

The example now sends requests to the Kubernetes API and sends to elastic events with pod information from the cluster every 5 seconds.

To validate check the logs:

    $ kubectl logs -f kubebeat-demo

Now go and check out the data on your Kibana! Make sure to add an index pattern `kubebeat*`

note: when changing the fields kibana will reject events dent from the kubebeat for not matching the existing scheme. make sure to delete the index when changing the event fields in your code.

### Clean up

To stop this example and clean up the pod, run:

    kubectl delete pod kubebeat-demo

### Remote Debugging

Build binary:

    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l"

Then use the patch file to change the configuration for Minikube (Only Once):

    patch kubebeat.yml kubebeat_minikube.yml.patch

Then package it to a docker image using the provided Dockerfile to run it on Kubernetes:

Running a [Minikube](https://minikube.sigs.k8s.io/docs/) cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube:

    eval $(minikube docker-env)
    docker build -f Dockerfile.debug -t kubebeat .

If you are not using Minikube, you should build this image and push it to a registry that your Kubernetes cluster can pull from.

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions(Only Once):

    kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default

After running the pod, expose the relevant ports:

    kubectl port-forward kubebeat-demo 40000:40000 8080:8080

The app will wait for the debugger to connect before starting

    ❯ kubectl logs -f kubebeat-demo

    API server listening at: [::]:40000
Use your favorite IDE to connect to the debugger on `localhost:40000` (for example [Goland](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#step-3-create-the-remote-run-debug-configuration-on-the-client-computer))

# {Beat}

Welcome to {Beat}.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/elastic/beats/v7/kubebeat`

## Getting Started with {Beat}

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with {Beat} and also install the
dependencies, run the following command:

```
make update
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push {Beat} in the git repository, run the following commands:

```
git remote set-url origin https://github.com/elastic/beats/v7/kubebeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for {Beat} run the command below. This will generate a binary
in the same directory with the name kubebeat.

```
make
```


### Run

To run {Beat} with debugging output enabled, run:

```
./kubebeat -c kubebeat.yml -e -d "*"
```


### Test

To test {Beat}, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  {Beat} source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone {Beat} from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/elastic/beats/v7/kubebeat
git clone https://github.com/elastic/beats/v7/kubebeat ${GOPATH}/src/github.com/elastic/beats/v7/kubebeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

## Build Elastic-Agent Docker with pre-packaged kubebeat


**1.Build Elastic-Agent Docker**

1. initialise git submodule for rego rules:
```
$ git submodule update --init
```
2. Access the Elastic-Agent dir
```
$ cd x-pack/elastic-agent
```
3. Build the elastic-agent docker( You might need to increase docker engine resources on your docker-engine)
```
$ DEV=true SNAPSHOT=true PLATFORMS=linux/amd64 TYPES=docker mage -v package # It takes a while on the first execution.
```
4. One build is finished, Verify the image is present on your machine

```
$ docker image ls | grep elastic-agent                                                                                                                          10076
docker.elastic.co/beats/elastic-agent-complete   8.1.0-SNAPSHOT           b73bbbc00e04   6 hours ago    1.6GB
docker.elastic.co/beats-ci/elastic-agent-cloud   8.1.0-SNAPSHOT           66f9b9d41737   6 hours ago    878MB
docker.elastic.co/beats/elastic-agent            8.1.0-SNAPSHOT           f089c673b70b   7 hours ago    554MB
docker.elastic.co/beats/elastic-agent-ubi8       8.1.0-SNAPSHOT           7140d1b7bc0e   7 hours ago    324MB
docker.elastic.co/beats/elastic-agent-complete   <none>                   182ee0acd1c0   2 weeks ago    1.64GB
docker.elastic.co/beats/elastic-agent            <none>                   e0ecbfa4a14e   2 weeks ago    595MB
docker.elastic.co/beats/elastic-agent-complete   8.0.0-SNAPSHOT           f081a7fcdc96   2 weeks ago    1.64GB
docker.elastic.co/beats/elastic-agent            <none>                   47584f609ed0   3 weeks ago    583MB
docker.elastic.co/beats/elastic-agent            <none>                   a7636fffbf82   4 weeks ago    580MB
docker.elastic.co/beats/elastic-agent-complete   7.15.1                   a0a9dfa8e527   6 weeks ago    1.61GB
```

</br>

**2. Deploy on Minikube:** 
1. Navigate to the kubebeat directory
```
$ cd beats/kubebeat 
```
2. Make a local copy of the yaml(it's set in gitignore) & insert your ES details(host,user,pass) 
```
$ cp  kubebeat/deploy/k8s/k8sbeat-agent-standalone-ds.yaml kubebeat/deploy/k8s/k8sbeat-agent-standalone-ds-local.yaml 
```
3. Load locally built image to minikube
```
$ minikube image load docker.elastic.co/beats/elastic-agent:8.1.0-SNAPSHOT
```
4. Deploy elastic-agent on minikube under kube-system namespace as a daemon set
```
$ kubectl apply -f kubebeat/deploy/k8s/k8sbeat-agent-standalone-ds-local.yaml 
```
5.  Validate agent pod is in running state
```
$ kubectl get po --selector="app=elastic-agent" -n kube-system 
```
6. Get log output from elastic-agent pod
```
$ kubectl logs -f  --selector="app=elastic-agent" -n kube-system 
```