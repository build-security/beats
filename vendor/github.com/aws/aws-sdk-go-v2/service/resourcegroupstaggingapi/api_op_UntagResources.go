// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UntagResourcesInput struct {
	_ struct{} `type:"structure"`

	// A list of ARNs. An ARN (Amazon Resource Name) uniquely identifies a resource.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// ResourceARNList is a required field
	ResourceARNList []string `min:"1" type:"list" required:"true"`

	// A list of the tag keys that you want to remove from the specified resources.
	//
	// TagKeys is a required field
	TagKeys []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourcesInput"}

	if s.ResourceARNList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARNList"))
	}
	if s.ResourceARNList != nil && len(s.ResourceARNList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARNList", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}
	if s.TagKeys != nil && len(s.TagKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagKeys", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagResourcesOutput struct {
	_ struct{} `type:"structure"`

	// Details of resources that could not be untagged. An error code, status code,
	// and error message are returned for each failed item.
	FailedResourcesMap map[string]FailureInfo `type:"map"`
}

// String returns the string representation
func (s UntagResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagResources = "UntagResources"

// UntagResourcesRequest returns a request value for making API operation for
// AWS Resource Groups Tagging API.
//
// Removes the specified tags from the specified resources. When you specify
// a tag key, the action removes both that key and its associated value. The
// operation succeeds even if you attempt to remove tags from a resource that
// were already removed. Note the following:
//
//    * To remove tags from a resource, you need the necessary permissions for
//    the service that the resource belongs to as well as permissions for removing
//    tags. For more information, see this list (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
//    * You can only tag resources that are located in the specified Region
//    for the AWS account.
//
//    // Example sending a request using UntagResourcesRequest.
//    req := client.UntagResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/UntagResources
func (c *Client) UntagResourcesRequest(input *UntagResourcesInput) UntagResourcesRequest {
	op := &aws.Operation{
		Name:       opUntagResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagResourcesInput{}
	}

	req := c.newRequest(op, input, &UntagResourcesOutput{})

	return UntagResourcesRequest{Request: req, Input: input, Copy: c.UntagResourcesRequest}
}

// UntagResourcesRequest is the request type for the
// UntagResources API operation.
type UntagResourcesRequest struct {
	*aws.Request
	Input *UntagResourcesInput
	Copy  func(*UntagResourcesInput) UntagResourcesRequest
}

// Send marshals and sends the UntagResources API request.
func (r UntagResourcesRequest) Send(ctx context.Context) (*UntagResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourcesResponse{
		UntagResourcesOutput: r.Request.Data.(*UntagResourcesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourcesResponse is the response type for the
// UntagResources API operation.
type UntagResourcesResponse struct {
	*UntagResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResources request.
func (r *UntagResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
