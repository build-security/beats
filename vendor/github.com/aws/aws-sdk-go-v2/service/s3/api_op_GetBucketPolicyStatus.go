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

type GetBucketPolicyStatusInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon S3 bucket whose policy status you want to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketPolicyStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketPolicyStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketPolicyStatusInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketPolicyStatusInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketPolicyStatusInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketPolicyStatusInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketPolicyStatusInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketPolicyStatusOutput struct {
	_ struct{} `type:"structure" payload:"PolicyStatus"`

	// The policy status for the specified bucket.
	PolicyStatus *PolicyStatus `type:"structure"`
}

// String returns the string representation
func (s GetBucketPolicyStatusOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketPolicyStatusOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PolicyStatus != nil {
		v := s.PolicyStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PolicyStatus", v, metadata)
	}
	return nil
}

const opGetBucketPolicyStatus = "GetBucketPolicyStatus"

// GetBucketPolicyStatusRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Retrieves the policy status for an Amazon S3 bucket, indicating whether the
// bucket is public. In order to use this operation, you must have the s3:GetBucketPolicyStatus
// permission. For more information about Amazon S3 permissions, see Specifying
// Permissions in a Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
//
// For more information about when Amazon S3 considers a bucket public, see
// The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status).
//
// The following operations are related to GetBucketPolicyStatus:
//
//    * Using Amazon S3 Block Public Access (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
//
//    * GetPublicAccessBlock
//
//    * PutPublicAccessBlock
//
//    * DeletePublicAccessBlock
//
//    // Example sending a request using GetBucketPolicyStatusRequest.
//    req := client.GetBucketPolicyStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketPolicyStatus
func (c *Client) GetBucketPolicyStatusRequest(input *GetBucketPolicyStatusInput) GetBucketPolicyStatusRequest {
	op := &aws.Operation{
		Name:       opGetBucketPolicyStatus,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?policyStatus",
	}

	if input == nil {
		input = &GetBucketPolicyStatusInput{}
	}

	req := c.newRequest(op, input, &GetBucketPolicyStatusOutput{})

	return GetBucketPolicyStatusRequest{Request: req, Input: input, Copy: c.GetBucketPolicyStatusRequest}
}

// GetBucketPolicyStatusRequest is the request type for the
// GetBucketPolicyStatus API operation.
type GetBucketPolicyStatusRequest struct {
	*aws.Request
	Input *GetBucketPolicyStatusInput
	Copy  func(*GetBucketPolicyStatusInput) GetBucketPolicyStatusRequest
}

// Send marshals and sends the GetBucketPolicyStatus API request.
func (r GetBucketPolicyStatusRequest) Send(ctx context.Context) (*GetBucketPolicyStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketPolicyStatusResponse{
		GetBucketPolicyStatusOutput: r.Request.Data.(*GetBucketPolicyStatusOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketPolicyStatusResponse is the response type for the
// GetBucketPolicyStatus API operation.
type GetBucketPolicyStatusResponse struct {
	*GetBucketPolicyStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketPolicyStatus request.
func (r *GetBucketPolicyStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
