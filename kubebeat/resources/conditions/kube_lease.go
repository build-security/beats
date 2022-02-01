package conditions

import (
	"context"
	"fmt"
	"os"
	"strings"

	"k8s.io/client-go/kubernetes"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PodNameEnvar           = "POD_NAME"
	DefaultLeaderLeaseName = "elastic-agent-cluster-leader"
)

type LeaderLeaseProvider interface {
	IsLeader() (bool, error)
}

type leaseInfo struct {
	ctx    context.Context
	client kubernetes.Interface
}

func NewLeaseInfo(ctx context.Context, client kubernetes.Interface) *leaseInfo {
	return &leaseInfo{ctx, client}
}

func (l *leaseInfo) IsLeader() (bool, error) {
	leases, err := l.client.CoordinationV1().Leases("kube-system").List(l.ctx, v1.ListOptions{})
	if err != nil {
		return false, err
	}

	for _, lease := range leases.Items {
		if lease.Name == DefaultLeaderLeaseName {
			podid := lastPart(*lease.Spec.HolderIdentity)

			if podid == l.currentPodID() {
				return true, nil
			}

			return false, nil
		}
	}

	return false, fmt.Errorf("could not find lease %v in Kube leases", DefaultLeaderLeaseName)
}

func (l *leaseInfo) currentPodID() string {
	pod := os.Getenv(PodNameEnvar)

	return lastPart(pod)
}

func lastPart(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}
