package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/osquery/osquery-go/plugin/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generate uses the api to retrieve information on all pods
func Pods(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	err := Init()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to kubernetes API server: %s", err))
	}

	client := GetClient()

	pods, err := client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println("could not get pods from k8s api")
		return nil, err
	}
	rows := make([]map[string]string, len(pods.Items))
	for i, pod := range pods.Items {
		sc, err := json.Marshal(pod.Spec.SecurityContext)
		if err != nil {
			return nil, err
		}

		rows[i] = map[string]string{
			"uid":              string(pod.UID), // UID is an alias to string
			"name":             pod.Name,
			"namespace":        pod.Namespace,
			"ip":               pod.Status.PodIP,
			"phase":            string(pod.Status.Phase),
			"service_account":  pod.Spec.ServiceAccountName,
			"node_name":        pod.Spec.NodeName,
			"security_context": string(sc),
		}
	}
	return rows, nil
}
