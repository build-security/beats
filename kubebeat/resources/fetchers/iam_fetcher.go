package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elastic/beats/v7/kubebeat/resources"
)

const IAMType = "aws-iam"

type IAMFetcher struct {
	iamProvider *IAMProvider
	cfg         IAMFetcherConfig
}

type IAMFetcherConfig struct {
	RoleName string `config:"roleName"`
}

func NewIAMFetcher(awsCfg aws.Config, cfg IAMFetcherConfig) (resources.Fetcher, error) {
	iam := NewIAMProvider(awsCfg)

	return &IAMFetcher{
		cfg:         cfg,
		iamProvider: iam,
	}, nil
}

func (f IAMFetcher) Fetch() ([]resources.FetcherResult, error) {
	results := make([]resources.FetcherResult, 0)
	ctx := context.Background()
	result, err := f.iamProvider.GetIAMRolePermissions(ctx, f.cfg.RoleName)
	results = append(results, resources.FetcherResult{
		Type:     IAMType,
		Resource: result,
	})

	return results, err
}

func (f IAMFetcher) Stop() {
}
