package beater

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/elastic/beats/v7/libbeat/logp"
)

type IamProvider struct {
	client *iam.Client
}

func NewIamProvider(cfg aws.Config) *IamProvider {
	svc := iam.New(cfg)
	return &IamProvider{
		client: svc,
	}
}

func (f IamProvider) GetIamRolePermissions(ctx context.Context, roleName string) (interface{}, error) {

	results := make([]interface{}, 0)
	policiesIdentifiers, err := f.getAllRolePolicies(svc, ctx, roleName)
	if err != nil {
		logp.Err("Failed to list role %s policies - %+v", roleName, err)
		return nil, err
	}

	for _, policyId := range policiesIdentifiers {

		input := &iam.GetRolePolicyInput{
			PolicyName: policyId.PolicyName,
			RoleName:   &roleName,
		}
		req := f.client.GetRolePolicyRequest(input)
		policy, err := req.Send(ctx)
		if err != nil {
			logp.Err("Failed to get policy %s - %+v", *policyId.PolicyName, err)
			continue
		}
		results = append(results, policy)
	}

	return results, nil
}

func (f IamProvider) getAllRolePolicies(ctx context.Context, roleName string) ([]iam.AttachedPolicy, error) {

	input := &iam.ListAttachedRolePoliciesInput{
		RoleName: &roleName,
	}

	req := f.client.ListAttachedRolePoliciesRequest(input)
	allPolicies, err := req.Send(ctx)
	if err != nil {
		logp.Err("Failed to list role %s policies - %+v", roleName, err)
		return nil, err
	}

	return allPolicies.AttachedPolicies, err
}
