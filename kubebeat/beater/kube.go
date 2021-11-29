package beater

import (
	"context"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/policy"
)

const ()

type KubeFetcher struct {
	watcher kubernetes.Watcher
}

func NewKubeFetcher(kubeconfig string, interval time.Duration) (Fetcher, error) {
	client, err := kubernetes.GetKubernetesClient(kubeconfig, kubernetes.KubeClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not get k8s client: %w", err)
	}

	logp.Info("Client initiated.")

	psp, err := client.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	fmt.Println("RAW PSP, ERR", psp, err)

	watchOptions := kubernetes.WatchOptions{
		SyncTimeout: interval,
		Namespace:   "kube-system",
	}

	watcher, err := kubernetes.NewWatcher(client, &policy.PodSecurityPolicy{}, watchOptions, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create watcher: %w", err)
	}

	if err := watcher.Start(); err != nil {
		return nil, fmt.Errorf("could not start watcher: %w", err)
	}

	logp.Info("Watcher initiated.")

	fmt.Println("KUBE API REGISTERED")

	return &KubeFetcher{
		watcher: watcher,
	}, nil
}

func (f *KubeFetcher) Fetch() ([]interface{}, error) {
	fmt.Println("KUBE API FETCH")
	psp := f.watcher.Store().List()

	fmt.Println("PSP COUNT", len(psp))

	for _, p := range psp {
		psp, ok := p.(*policy.PodSecurityPolicy)

		fmt.Println("PSP OK?", psp, ok)

		if !ok {
			logp.Info("could not convert to PodSecurityPolicy")
			continue
		}
		psp.SetManagedFields(nil)
		psp.Kind = "PodSecurityPolicy" // see https://github.com/kubernetes/kubernetes/issues/3030. TODO(yashtewari): Does this still apply for PodSecurityPolicy?
	}

	return psp, nil
}

func (f *KubeFetcher) Stop() {
	f.watcher.Stop()
}
