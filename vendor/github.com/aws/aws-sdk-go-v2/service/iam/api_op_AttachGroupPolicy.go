// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AttachGroupPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name (friendly name, not ARN) of the group to attach the policy to.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM policy you want to attach.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachGroupPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachGroupPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachGroupPolicyInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachGroupPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachGroupPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachGroupPolicy = "AttachGroupPolicy"

// AttachGroupPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Attaches the specified managed policy to the specified IAM group.
//
// You use this API to attach a managed policy to a group. To embed an inline
// policy in a group, use PutGroupPolicy.
//
// For more information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using AttachGroupPolicyRequest.
//    req := client.AttachGroupPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AttachGroupPolicy
func (c *Client) AttachGroupPolicyRequest(input *AttachGroupPolicyInput) AttachGroupPolicyRequest {
	op := &aws.Operation{
		Name:       opAttachGroupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachGroupPolicyInput{}
	}

	req := c.newRequest(op, input, &AttachGroupPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AttachGroupPolicyRequest{Request: req, Input: input, Copy: c.AttachGroupPolicyRequest}
}

// AttachGroupPolicyRequest is the request type for the
// AttachGroupPolicy API operation.
type AttachGroupPolicyRequest struct {
	*aws.Request
	Input *AttachGroupPolicyInput
	Copy  func(*AttachGroupPolicyInput) AttachGroupPolicyRequest
}

// Send marshals and sends the AttachGroupPolicy API request.
func (r AttachGroupPolicyRequest) Send(ctx context.Context) (*AttachGroupPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachGroupPolicyResponse{
		AttachGroupPolicyOutput: r.Request.Data.(*AttachGroupPolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachGroupPolicyResponse is the response type for the
// AttachGroupPolicy API operation.
type AttachGroupPolicyResponse struct {
	*AttachGroupPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachGroupPolicy request.
func (r *AttachGroupPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
