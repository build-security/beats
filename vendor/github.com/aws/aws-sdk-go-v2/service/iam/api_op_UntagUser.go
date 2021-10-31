// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type UntagUserInput struct {
	_ struct{} `type:"structure"`

	// A list of key names as a simple array of strings. The tags with matching
	// keys are removed from the specified user.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`

	// The name of the IAM user from which you want to remove tags.
	//
	// This parameter accepts (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consist of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UntagUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagUserInput"}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
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

type UntagUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagUser = "UntagUser"

// UntagUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Removes the specified tags from the user. For more information about tagging,
// see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using UntagUserRequest.
//    req := client.UntagUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UntagUser
func (c *Client) UntagUserRequest(input *UntagUserInput) UntagUserRequest {
	op := &aws.Operation{
		Name:       opUntagUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagUserInput{}
	}

	req := c.newRequest(op, input, &UntagUserOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UntagUserRequest{Request: req, Input: input, Copy: c.UntagUserRequest}
}

// UntagUserRequest is the request type for the
// UntagUser API operation.
type UntagUserRequest struct {
	*aws.Request
	Input *UntagUserInput
	Copy  func(*UntagUserInput) UntagUserRequest
}

// Send marshals and sends the UntagUser API request.
func (r UntagUserRequest) Send(ctx context.Context) (*UntagUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagUserResponse{
		UntagUserOutput: r.Request.Data.(*UntagUserOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagUserResponse is the response type for the
// UntagUser API operation.
type UntagUserResponse struct {
	*UntagUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagUser request.
func (r *UntagUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
