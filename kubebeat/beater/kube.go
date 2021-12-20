package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	"k8s.io/apimachinery/pkg/runtime"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	kubeSystemNamespace = "kube-system"
	allNamespaces       = "" // The Kube API treats this as "all namespaces"
)

var (
	// vanillaClusterResources represents those resources that are required for a vanilla
	// Kubernetes cluster.
	vanillaClusterResources = []requiredResource{
		{
			&kubernetes.Pod{},
			kubeSystemNamespace,
		},
		{

			&kubernetes.Secret{},
			allNamespaces,
		},
		{
			&kubernetes.Role{},
			allNamespaces,
		},
		{
			&kubernetes.RoleBinding{},
			allNamespaces,
		},
		{
			&kubernetes.ClusterRole{},
			allNamespaces,
		},
		{
			&kubernetes.ClusterRoleBinding{},
			allNamespaces,
		},
		{
			&kubernetes.PodSecurityPolicy{},
			allNamespaces,
		},
		// TODO(yashtewari): Problem: github.com/elastic/beats/vendor/k8s.io/apimachinery/pkg/api/errors/errors.go#401
		// > "the server could not find the requested resource"
		// {
		// 	&kubernetes.NetworkPolicy{},
		// 	allNamespaces,
		// },
	}
)

type requiredResource struct {
	resource  kubernetes.Resource
	namespace string
}

type KubeFetcher struct {
	kubeconfig string
	interval   time.Duration
	watchers   []kubernetes.Watcher
}

func NewKubeFetcher(kubeconfig string, interval time.Duration) (Fetcher, error) {
	f := &KubeFetcher{
		kubeconfig: kubeconfig,
		interval:   interval,
		watchers:   make([]kubernetes.Watcher, 0),
	}

	if err := f.initWatchers(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *KubeFetcher) initWatcher(client k8s.Interface, r requiredResource) error {
	w, err := kubernetes.NewWatcher(client, r.resource, kubernetes.WatchOptions{
		SyncTimeout: f.interval,
		Namespace:   r.namespace,
	}, nil)
	if err != nil {
		return fmt.Errorf("could not create watcher: %w", err)
	}

	// TODO(yashtewari): it appears that Start never returns in case of certain failures, for example
	// if the configured Client's Role does not have the necessary permissions to list the Resource
	// being watched. This needs to be handled.
	//
	// When such a failure happens, kubebeat won't shut down gracefuly, i.e. Stop will not work. This
	// happens due to a context.TODO present in the libbeat dependency. It needs to accept context
	// from the caller instead.
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

	f.watchers = make([]kubernetes.Watcher, 0)

	for _, r := range vanillaClusterResources {
		err := f.initWatcher(client, r)
		if err != nil {
			return err
		}
	}

	logp.Info("Kubernetes Watchers initiated.")

	return nil
}

func (f *KubeFetcher) Fetch() ([]interface{}, error) {
	ret := make([]interface{}, 0)

	for _, w := range f.watchers {
		resources := w.Store().List()

		for _, r := range resources {
			o, ok := r.(runtime.Object)

			if !ok {
				logp.L().Errorf("Bad resource: %#v does not implement runtime.Object", r)
				continue
			}

			addTypeInformationToObject(o) // See https://github.com/kubernetes/kubernetes/issues/3030
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
