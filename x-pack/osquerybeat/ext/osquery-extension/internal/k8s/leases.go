package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/osquery/osquery-go/plugin/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generate uses the api to retrieve information on all leases
func Leases(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	err := Init()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to kubernetes API server: %s", err))
	}

	client := GetClient()

	leases, err := client.CoordinationV1().Leases("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println("could not get leases from k8s api")
		return nil, err
	}
	rows := make([]map[string]string, len(leases.Items))
	for i, lease := range leases.Items {
		hi, err := json.Marshal(lease.Spec.HolderIdentity)
		ld, err := json.Marshal(lease.Spec.LeaseDurationSeconds)
		at, err := json.Marshal(lease.Spec.AcquireTime)
		rt, err := json.Marshal(lease.Spec.RenewTime)
		lt, err := json.Marshal(lease.Spec.LeaseTransitions)

		if err != nil {
			return nil, err
		}

		rows[i] = map[string]string{
			"uid":                  string(lease.UID), // UID is an alias to string
			"name":                 lease.Name,
			"namespace":            lease.Namespace,
			"holderIdentity":       string(hi),
			"leaseDurationSeconds": string(ld),
			"acquireTime":          string(at),
			"renewTime":            string(rt),
			"leaseTransitions":     string(lt),
		}
	}
	return rows, nil
}
