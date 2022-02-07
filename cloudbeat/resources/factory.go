package resources

import (
	"errors"

	"github.com/elastic/beats/v7/cloudbeat/config"
	"github.com/elastic/beats/v7/libbeat/common"
)

var Factories = factories{m: make(map[string]FetcherFactory)}

type FetcherFactory func(common.Config) (Fetcher, error)

type factories struct {
	m map[string]FetcherFactory
}

func (fa *factories) ListFetcher(name string, f FetcherFactory) {
	_, ok := fa.m[name]
	if ok {
		// log
	}

	fa.m[name] = f
}

func (fa *factories) CreateFetcher(name string, c common.Config) (Fetcher, error) {
	factory, ok := fa.m[name]
	if !ok {
		return nil, errors.New("fetcher factory could not be found")
	}

	return factory(c)
}

func ConfigFetchers(registry FetchersRegistry, cfg config.Config) error {
	for _, fcfg := range cfg.Fetchers {
		gen := BaseFetcherConfig{}
		err := fcfg.Unpack(&gen)
		if err != nil {
			return err
		}

		c := make([]FetcherCondition, 0)
		f, err := Factories.CreateFetcher(gen.Fetcher, *fcfg)
		if err != nil {
			return err
		}

		registry.Register(gen.Fetcher, f, c...)
	}

	return nil
}

// {
// 	kubebeat:
// 		fetchers: [
// 			{
// 				type: file_fetcher
// 				patterns: []
// 				period:
// 			}
// 			{
// 				type: process_fetcher
// 				period: 10s
// 			}
// 			{
// 				type: kubeapi_fetcher
// 				requiredResources: [
// 					{
// 						resource_name: POD
// 						resource_namespace: ""
// 					}
// 					{
// 						resource_name: Deployment
// 						resource_namespace: "kube-system"
// 					}
// 				]
// 			}
// 		]
// }
