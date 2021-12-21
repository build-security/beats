package beater

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"strings"
	"time"
)

type AwsKubeFetcher struct {
	clusterName string
	cfg         aws.Config
	kubeClient  k8s.Interface
	ecr         *ECRProvider
	eks         *EKSProvider
	elb         *ELBProvider
}

func NewAwsKubeFetcher(kubeconfig string, clusterName string) (Fetcher, error) {

	kubernetesClient, err := kubernetes.GetKubernetesClient(kubeconfig, kubernetes.KubeClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("fail to get k8sclient client: %s", err.Error())
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, fmt.Errorf("fail init aws config: %s", err.Error())
	}

	ecr := NewEcrProvider(cfg)
	eks := NewEksProvider(cfg)
	elb := NewELBProvider(cfg)

	return &AwsKubeFetcher{
		cfg:         cfg,
		ecr:         ecr,
		kubeClient:  kubernetesClient,
		eks:         eks,
		elb:         elb,
		clusterName: clusterName,
	}, nil
}

func dump(items []interface{}) {
	logp.Info("Started dump ")

	for i := 0; i < len(items); i++ {
		logp.Info("%v", items[i])
	}

	logp.Info("Finished dump ")

}

func (f AwsKubeFetcher) Fetch() ([]interface{}, error) {

	results := make([]interface{}, 0)

	//ecrCtx, ecrCtxCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	//defer ecrCtxCancel()
	//repositories, err := f.GetECRInformation(ecrCtx)
	//results = append(results, repositories)

	logp.Info("ClusterCtx result started.")

	clusterCtx, clusterCtxCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer clusterCtxCancel()
	data, err := f.GetClusterDescription(clusterCtx)
	results = append(results, data)
	dump(results)
	logp.Info("ClusterCtx result finished")

	//kubeCtx, kubeCtxCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	//defer kubeCtxCancel()
	//lbCtx, lbctxCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	//defer lbctxCancel()
	//lbData, err := f.GetLoadBalancerDescriptions(kubeCtx, lbCtx)
	//results = append(results, lbData)

	//nodeCtx, nodeCtxCancel := context.WithTimeout(context.TODO(), 30*time.Second)
	//defer nodeCtxCancel()
	//nodeData, err := f.GetNodeDescription(nodeCtx)
	//results = append(results, nodeData)

	return results, err
}

// 2.1.1 Enable audit Logs (Manual)
// 5.3.1 - Ensure Kubernetes Secrets are encrypted using Customer Master Keys (CMKs) managed in AWS KMS (Automated)
// 5.4.1 - Restrict Access to the Control Plane Endpoint (Manual)
// 5.4.2 - Ensure clusters are created with Private Endpoint Enabled and Public Access Disabled (Manual)
func (f AwsKubeFetcher) GetClusterDescription(ctx context.Context) (*eks.Cluster, error) {

	// https://github.com/kubernetes/client-go/issues/530
	// Currently we could not auto-detected the cluster name
	// TODO - leader election
	result, err := f.eks.DescribeCluster(ctx, f.clusterName)

	return result.Cluster, err
}

// EKS benchmark 5.1.1 -  Ensure Image Vulnerability Scanning using Amazon ECR image scanning or a third party provider (Manual)
func (f AwsKubeFetcher) GetECRInformation(ctx context.Context) ([]ecr.Repository, error) {

	// TODO - Need to use leader election

	// TODO - Currently we do not know how to extract the ECR repository out of the image
	// When we would know, we need to scan all the pods and gets their images
	// Otherwise it will get repositories that are not associated with this cluster
	repositories, err := f.ecr.DescribeAllECRRepositories(ctx)

	return repositories, err
}

// EKS benchmark 5.4.5 -  Encrypt traffic to HTTPS load balancers with TLS certificates (Manual)
func (f AwsKubeFetcher) GetLoadBalancerDescriptions(kubectx context.Context, lbctx context.Context) ([]elasticloadbalancing.LoadBalancerDescription, error) {

	// TODO - leader election
	// Running on all namespaces
	services, err := f.kubeClient.CoreV1().Services("").List(kubectx, metav1.ListOptions{})
	if err != nil {
		logp.Err("Failed to get all services  - %+v", err)
		return nil, err
	}

	loadBalancers := make([]string, 0)
	for _, service := range services.Items {

		for _, ingress := range service.Status.LoadBalancer.Ingress {
			if strings.Contains(ingress.Hostname, "amazonaws.com") {
				// TODO - Needs to be refactored
				lbName := strings.Split(ingress.Hostname, "-")[0]
				loadBalancers = append(loadBalancers, lbName)
			}
		}
	}
	result, err := f.elb.DescribeLoadBalancer(lbctx, loadBalancers)

	return result, err
}

// EKS benchmark 5.4.3 Ensure clusters are created with Private Nodes (Manual)
func (f AwsKubeFetcher) GetNodeDescription(ctx context.Context) ([]v1.Node, error) {

	// TODO - leader election
	nodeList, err := f.kubeClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		logp.Err("Failed to get all nodes information  - %+v", err)
		return nil, err
	}

	return nodeList.Items, nil
}

func (f AwsKubeFetcher) Stop() {

}
