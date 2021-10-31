// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AttachUserPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy you want to attach.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachUserPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachUserPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachUserPolicyInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachUserPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachUserPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachUserPolicy = "AttachUserPolicy"

// AttachUserPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Attaches the specified managed policy to the specified user.
//
// You use this API to attach a managed policy to a user. To embed an inline
// policy in a user, use PutUserPolicy.
//
// For more information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using AttachUserPolicyRequest.
//    req := client.AttachUserPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AttachUserPolicy
func (c *Client) AttachUserPolicyRequest(input *AttachUserPolicyInput) AttachUserPolicyRequest {
	op := &aws.Operation{
		Name:       opAttachUserPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachUserPolicyInput{}
	}

	req := c.newRequest(op, input, &AttachUserPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AttachUserPolicyRequest{Request: req, Input: input, Copy: c.AttachUserPolicyRequest}
}

// AttachUserPolicyRequest is the request type for the
// AttachUserPolicy API operation.
type AttachUserPolicyRequest struct {
	*aws.Request
	Input *AttachUserPolicyInput
	Copy  func(*AttachUserPolicyInput) AttachUserPolicyRequest
}

// Send marshals and sends the AttachUserPolicy API request.
func (r AttachUserPolicyRequest) Send(ctx context.Context) (*AttachUserPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachUserPolicyResponse{
		AttachUserPolicyOutput: r.Request.Data.(*AttachUserPolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachUserPolicyResponse is the response type for the
// AttachUserPolicy API operation.
type AttachUserPolicyResponse struct {
	*AttachUserPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachUserPolicy request.
func (r *AttachUserPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
