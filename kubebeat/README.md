# POC Documentation

I generated this repo using the [beats development guide](https://www.elastic.co/guide/en/beats/devguide/current/newbeat-generate.html).
The kube-api call is based on the [k8s go-client example](https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration).

The interesting files are:
* `beater/kubebeat.go` - the beats logic
* `kubebeat.yml` - the beats config
* `Dockerfile` - runs the beat dockerized with debug flags
* `pod.yaml` - deploy the beat


## Running this example

This example assumes you have:
1. Elasticsearch with the default username & password (`elastic` & `changeme`) running on the default port (`http://localhost:9200`)
2. Kibana with running on the default port (`http://localhost:5601`)
3. Minikube cluster running locally (`minikube start`)

First compile the application for Linux:

    GOOS=linux go build

Then use the patch file to change the configuration for Minikube (or change the configuration according to your setup):

    patch kubebeat.yml kubebeat_minikube.yml.patch

Then package it to a docker image using the provided Dockerfile to run it on Kubernetes:

Running a [Minikube](https://minikube.sigs.k8s.io/docs/) cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube:

    eval $(minikube docker-env)
    docker build -t kubebeat .

If you are not using Minikube, you should build this image and push it to a registry that your Kubernetes cluster can pull from.

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions:

    kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default

Then, run the image in a Pod with a single instance Deployment:

    kubectl apply -f pod.yml

The example now sends requests to the Kubernetes API and sends to elastic events with pod information from the cluster every 5 seconds.

To validate check the logs:

    kubectl logs -f kubebeat-demo

Now go and check out the data on your Kibana! Make sure to add an index pattern `kubebeat*`

note: when changing the fields kibana will reject events dent from the kubebeat for not matching the existing scheme. make sure to delete the index when changing the event fields in your code.

### Clean up

To stop this example and clean up the pod, run:

    kubectl delete pod kubebeat-demo

### Open questions

1. Could we use some code from `kube-mgmt`/`gatekeeper`/`metricbeat` to do the kube-api querying and data management?
2. How should we integrate this to the agent?
3. ... many more

### Remote Debugging

Build binary:

    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l"

Build docker image:

    docker build -f Dockerfile.debug -t kubebeat .

After running the pod, expose the relevant ports:

    kubectl port-forward kubebeat-demo 40000:40000 8080:8080

The app will wait for the debugger to connect before starting

    ❯ kubectl logs -f kubebeat-demo

    API server listening at: [::]:40000
Use your favorite IDE to connect to the debugger on `localhost:40000` (for example [Goland](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#step-3-create-the-remote-run-debug-configuration-on-the-client-computer))


### Build with the agent (Hacky way)
This method's target is to make the agent think kubebeat is osquerybeat and then run it upon the osquery integration activation through the kibana fleet manager.
The next immediate step would be to create a local setup with our own integration, this way we wouldn't have to fool the agent to think our kubebeat is osquerybeat.

How it's done:
1. Copy the kubebeat release tar.gz and it's sha (changing the name in the sha file) to the downloads directory of the agent.
2. Create a key and sign the tar.gz with the key.
3. Configure the agent to use the generated key to validate the tar.gz.

currently, step 3 doesn't work (Only tried with a fleet-managed agent) and results with:

    2021-11-17T10:55:22.454Z	ERROR	log/reporter.go:36	2021-11-17T10:55:22Z - message: Application: osquerybeat--8.0.0-SNAPSHOT[41f91bfa-7915-4388-be70-71cc51cf97b2]: State changed to FAILED: operation 'operation-verify' failed to verify osquerybeat.8.0.0-SNAPSHOT: 3 errors occurred:
    * check detached signature: openpgp: signature made by unknown entity
    * check detached signature: openpgp: invalid signature: hash tag doesn't match
    * fetching asc file from https://artifacts.elastic.co/downloads/beats/osquerybeat/osquerybeat-8.0.0-SNAPSHOT-linux-x86_64.tar.gz.asc: call to 'https://artifacts.elastic.co/downloads/beats/osquerybeat/osquerybeat-8.0.0-SNAPSHOT-linux-x86_64.tar.gz.asc' returned unsuccessful status code: 404

    - type: 'ERROR' - sub_type: 'FAILED'

#### Detailed walkthrough:

Build binary:

    PLATFORMS=linux/amd64 TYPES=tgz SNAPSHOT=true make release

Change the beat name in the file 'kubebeat/build/distributions/kubebeat-8.1.0-SNAPSHOT-linux-x86_64.tar.gz.sha512':


original

    abcdSOMESHA  kubebeat-8.1.0-SNAPSHOT-linux-x86_64.tar.gz

new

    abcdSOMESHA  osquerybeat-8.0.0-SNAPSHOT-linux-x86_64.tar.gz


Build docker image:

    docker build -f agent_tools/Dockerfile.agent -t kubebeatagent .

Deploy the agent to the cluster (fleet-managed):

    kubectl apply -f agent_tools/my-elastic-agent-kubebeat-managed-kubernetes.yaml

Configure the agent to send events to the local running elasticsearch host

![agent_tools/fleet_settings.png](agent_tools/fleet_settings.png)

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
