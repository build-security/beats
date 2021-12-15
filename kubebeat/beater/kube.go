package beater

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	kubeSystemNamespace = "kube-system"
)

var (
	watcherlock sync.Once
)

type KubeFetcher struct {
	kubeconfig string
	interval   time.Duration
	watchers   []kubernetes.Watcher
}

func NewKubeFetcher(kubeconfig string, interval time.Duration) Fetcher {
	return &KubeFetcher{
		kubeconfig: kubeconfig,
		interval:   interval,
		watchers:   make([]kubernetes.Watcher, 0),
	}
}

func (f *KubeFetcher) initWatcher(w kubernetes.Watcher, err error) error {
	if err != nil {
		return fmt.Errorf("could not create watcher: %w", err)
	}

	if err := w.Start(); err != nil {
		return fmt.Errorf("could not start watcher: %w", err)
	}

	f.watchers = append(f.watchers, w)

	return nil
}

func (f *KubeFetcher) initWatchers() error {
	client, err := kubernetes.GetKubernetesClient(f.kubeconfig, kubernetes.KubeClientOptions{})
	if err != nil {
		return fmt.Errorf("could not get k8s client: %w", err)
	}

	logp.Info("Kubernetes client initiated.")

	psp, err := client.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	fmt.Println("RAW PSP, ERR", psp, err)

	watchOptions := kubernetes.WatchOptions{
		SyncTimeout: f.interval,
		Namespace:   kubeSystemNamespace,
	}

	for _, r := range []kubernetes.Resource{
		&kubernetes.Pod{},
		&kubernetes.Secret{},
		&kubernetes.Role{},
		&kubernetes.RoleBinding{},
		&kubernetes.ClusterRole{},
		&kubernetes.ClusterRoleBinding{},
		&kubernetes.PodSecurityPolicy{},
		&kubernetes.NetworkPolicy{},
	} {
		f.initWatcher(kubernetes.NewWatcher(client, r, watchOptions, nil))
	}

	logp.Info("Kubernetes Watchers initiated.")

	return nil
}

func (f *KubeFetcher) Fetch() ([]interface{}, error) {
	var err error
	watcherlock.Do(func() {
		err = f.initWatchers()
	})
	if err != nil {
		return nil, err
	}

	ret := make([]interface{}, 0)

	for _, w := range f.watchers {
		resources := w.Store().List()

		for _, r := range resources {
			o, ok := r.(runtime.Object)

			if !ok {
				logp.L().Errorf("bad resource: %#v does not implement runtime.Object", r)
				continue
			}

			// o.SetManagedFields(nil) -- TODO(yashtewari): What was this supposed to do?
			addTypeInformationToObject(o) // see https://github.com/kubernetes/kubernetes/issues/3030. TODO(yashtewari): Does this still apply for PodSecurityPolicy?
		}

		ret = append(ret, resources...)
	}

	return ret, nil
}

func (f *KubeFetcher) Stop() {
	for _, w := range f.watchers {
		w.Stop()
	}
}

// addTypeInformationToObject adds TypeMeta information to a runtime.Object based upon the loaded scheme.Scheme
// inspired by: https://github.com/kubernetes/cli-runtime/blob/v0.19.2/pkg/printers/typesetter.go#L41
func addTypeInformationToObject(obj runtime.Object) error {
	gvks, _, err := scheme.Scheme.ObjectKinds(obj)
	if err != nil {
		return fmt.Errorf("missing apiVersion or kind and cannot assign it; %w", err)
	}

	for _, gvk := range gvks {
		if len(gvk.Kind) == 0 {
			continue
		}
		if len(gvk.Version) == 0 || gvk.Version == runtime.APIVersionInternal {
			continue
		}
		obj.GetObjectKind().SetGroupVersionKind(gvk)
		break
	}

	return nil
}
