package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const IAMType = "aws-iam"

type IAMFetcher struct {
	iamProvider *IAMProvider
	roleName    string
}

type IAMFetchResult struct {
	obj interface{}
}

func NewIAMFetcher(cfg aws.Config, roleName string) (Fetcher, error) {
	iam := NewIAMProvider(cfg)

	return &IAMFetcher{
		iamProvider: iam,
		roleName:    roleName,
	}, nil
}

func (f IAMFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	result, err := f.iamProvider.GetIAMRolePermissions(ctx, f.roleName)
	results = append(results, IAMFetchResult{result})

	return results, err
}

func (f IAMFetcher) Stop() {
}

func (res IAMFetchResult) GetID() string {
	return ""
}
