package kube

import (
	"fmt"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	client kubernetes.Interface
	once   sync.Once
)

// GetKubeClient Creates the kubernetes client using the kubeconfig path
func GetKubeClient(kubeconfig string) (kubernetes.Interface, error) {
	var err error

	once.Do(func() {
		var config *rest.Config
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return
		}
		// create the clientset
		client, err = kubernetes.NewForConfig(config)
	})

	if err != nil {
		return nil, fmt.Errorf("error initiating Kubernetes client: %w", err)
	}

	return client, nil
}
