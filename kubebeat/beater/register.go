package beater

import (
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/libbeat/common"
)

func RegisterFetchers(data *resources.Data, cfg *common.Config) error {
	conf := config.Config{}
	err := cfg.Unpack(&conf)
	if err != nil {
		return err
	}

	for _, fcfg := range conf.Fetchers {
		gen := BaseFetcherConfig{}
		err := fcfg.Unpack(&gen)
		if err != nil {
			return err
		}

		var x interface{}
		switch gen.Fetcher {
		case "kubeApi":
			x = KubeApiFetcherConfig{}
		case "amir":
		}

		err = fcfg.Unpack(&x)
		if err != nil {
			return err
		}
	}

	return nil
}

type BaseFetcherConfig struct {
	Fetcher string `config:"fetcher"`
}

type FileFetcherConfig struct {
	Patterns []string `config:"patterns"`
	Period   int      `config:"period"`
}

type ProcessFetcherConfig struct {
	Period int `config:"period"`
}

type KubeApiFetcherConfig struct {
	Period    int                            `config:"period"`
	Resources []KubeApiFetcherConfigResource `config:"resources"`
}

type KubeApiFetcherConfigResource struct {
	Name      string `config:"name"`
	Namespace string `config:"namespace"`
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
