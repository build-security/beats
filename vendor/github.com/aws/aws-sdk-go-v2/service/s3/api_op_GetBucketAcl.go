// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type GetBucketAclInput struct {
	_ struct{} `type:"structure"`

	// Specifies the S3 bucket whose ACL is being requested.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketAclInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketAclInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAclInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketAclInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketAclInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketAclOutput struct {
	_ struct{} `type:"structure"`

	// A list of grants.
	Grants []Grant `locationName:"AccessControlList" locationNameList:"Grant" type:"list"`

	// Container for the bucket owner's display name and ID.
	Owner *Owner `type:"structure"`
}

// String returns the string representation
func (s GetBucketAclOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAclOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Grants != nil {
		v := s.Grants

		metadata := protocol.Metadata{ListLocationName: "Grant"}
		ls0 := e.List(protocol.BodyTarget, "AccessControlList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Owner != nil {
		v := s.Owner

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Owner", v, metadata)
	}
	return nil
}

const opGetBucketAcl = "GetBucketAcl"

// GetBucketAclRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This implementation of the GET operation uses the acl subresource to return
// the access control list (ACL) of a bucket. To use GET to return the ACL of
// the bucket, you must have READ_ACP access to the bucket. If READ_ACP permission
// is granted to the anonymous user, you can return the ACL of the bucket without
// using an authorization header.
//
// Related Resources
//
//    *
//
//    // Example sending a request using GetBucketAclRequest.
//    req := client.GetBucketAclRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAcl
func (c *Client) GetBucketAclRequest(input *GetBucketAclInput) GetBucketAclRequest {
	op := &aws.Operation{
		Name:       opGetBucketAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?acl",
	}

	if input == nil {
		input = &GetBucketAclInput{}
	}

	req := c.newRequest(op, input, &GetBucketAclOutput{})

	return GetBucketAclRequest{Request: req, Input: input, Copy: c.GetBucketAclRequest}
}

// GetBucketAclRequest is the request type for the
// GetBucketAcl API operation.
type GetBucketAclRequest struct {
	*aws.Request
	Input *GetBucketAclInput
	Copy  func(*GetBucketAclInput) GetBucketAclRequest
}

// Send marshals and sends the GetBucketAcl API request.
func (r GetBucketAclRequest) Send(ctx context.Context) (*GetBucketAclResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketAclResponse{
		GetBucketAclOutput: r.Request.Data.(*GetBucketAclOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketAclResponse is the response type for the
// GetBucketAcl API operation.
type GetBucketAclResponse struct {
	*GetBucketAclOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketAcl request.
func (r *GetBucketAclResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
