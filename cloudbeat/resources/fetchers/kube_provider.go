package fetchers

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	"k8s.io/apimachinery/pkg/api/meta"
)

type ExtK8sResource struct {
	kubernetes.Resource
}

func GetKubeData(watchers []kubernetes.Watcher) []PolicyResource {
	ret := make([]PolicyResource, 0)

	for _, watcher := range watchers {
		rs := watcher.Store().List()

		for _, r := range rs {
			nullifyManagedFields(r)
			resource, ok := r.(kubernetes.Resource)

			if !ok {
				logp.L().Errorf("Bad resource: %#v does not implement kubernetes.Resource", r)
				continue
			}

			err := addTypeInformationToKubeResource(resource)
			if err != nil {
				logp.L().Errorf("Bad resource: %w", err)
				continue
			} // See https://github.com/kubernetes/kubernetes/issues/3030

			ret = append(ret, ExtK8sResource{resource})
		}
	}

	return ret
}

func (res ExtK8sResource) GetID() string {
	accessor, err := meta.Accessor(res)
	if err != nil {
		// Some err occur while trying to get metadata - return obj without id
		fmt.Errorf("missing required metadata fields; %w", err)
		return ""
	}

	uid := accessor.GetUID()
	return string(uid)
}

// nullifyManagedFields ManagedFields field contains fields with dot that prevent from elasticsearch to index
// the events.
func nullifyManagedFields(resource interface{}) {
	switch val := resource.(type) {
	case *kubernetes.Pod:
		val.ManagedFields = nil
	case *kubernetes.Secret:
		val.ManagedFields = nil
	case *kubernetes.Role:
		val.ManagedFields = nil
	case *kubernetes.RoleBinding:
		val.ManagedFields = nil
	case *kubernetes.ClusterRole:
		val.ManagedFields = nil
	case *kubernetes.ClusterRoleBinding:
		val.ManagedFields = nil
	case *kubernetes.PodSecurityPolicy:
		val.ManagedFields = nil
	case *kubernetes.NetworkPolicy:
		val.ManagedFields = nil
	}
}
