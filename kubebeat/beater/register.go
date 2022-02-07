package beater

import (
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/kubebeat/resources/fetchers"
)

func init() {

}

func ConfigFetchers(registry resources.FetchersRegistry, cfg config.Config) error {
	for _, fcfg := range cfg.Fetchers {
		gen := resources.BaseFetcherConfig{}
		err := fcfg.Unpack(&gen)
		if err != nil {
			return err
		}

		var f resources.Fetcher
		c := make([]resources.FetcherCondition, 0)
		switch gen.Fetcher {
		case fetchers.KubeAPIType:
			apiCfg := fetchers.KubeApiFetcherConfig{}
			err = fcfg.Unpack(&apiCfg)
			if err != nil {
				return err
			}
			f, err = fetchers.NewKubeFetcher(apiCfg)

		case fetchers.FileSystemType:
			fileCfg := fetchers.FileFetcherConfig{}
			err = fcfg.Unpack(&fileCfg)
			if err != nil {
				return err
			}
			f = fetchers.NewFileFetcher(fileCfg)
		case fetchers.ProcessType:
			procCfg := fetchers.ProcessFetcherConfig{}
			err = fcfg.Unpack(&procCfg)
			if err != nil {
				return err
			}
			f = fetchers.NewProcessesFetcher(procCfg)
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
