package beater

import (
	"context"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterHelper struct {
	clusterId string
}

func newClusterHelper() *ClusterHelper {
	return &ClusterHelper{clusterId: getClusterIdFromClient()}
}

func (c ClusterHelper) ClusterId() string {
	return c.clusterId
}

func getClusterIdFromClient() string {
	client, _ := kubernetes.GetKubernetesClient("", kubernetes.KubeClientOptions{})
	n, _ := client.CoreV1().Namespaces().Get(context.Background(), "kube-system", metav1.GetOptions{})
	return string(n.ObjectMeta.UID)
}
