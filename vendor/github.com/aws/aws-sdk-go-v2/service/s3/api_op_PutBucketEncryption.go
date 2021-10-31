// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/checksum"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutBucketEncryptionInput struct {
	_ struct{} `type:"structure" payload:"ServerSideEncryptionConfiguration"`

	// Specifies default encryption for a bucket using server-side encryption with
	// Amazon S3-managed keys (SSE-S3) or customer master keys stored in AWS KMS
	// (SSE-KMS). For information about the Amazon S3 default encryption feature,
	// see Amazon S3 Default Bucket Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies the default server-side-encryption configuration.
	//
	// ServerSideEncryptionConfiguration is a required field
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `locationName:"ServerSideEncryptionConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketEncryptionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.ServerSideEncryptionConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerSideEncryptionConfiguration"))
	}
	if s.ServerSideEncryptionConfiguration != nil {
		if err := s.ServerSideEncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ServerSideEncryptionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketEncryptionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketEncryptionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ServerSideEncryptionConfiguration != nil {
		v := s.ServerSideEncryptionConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "ServerSideEncryptionConfiguration", v, metadata)
	}
	return nil
}

func (s *PutBucketEncryptionInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutBucketEncryptionInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutBucketEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketEncryptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketEncryption = "PutBucketEncryption"

// PutBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This implementation of the PUT operation uses the encryption subresource
// to set the default encryption state of an existing bucket.
//
// This implementation of the PUT operation sets default encryption for a bucket
// using server-side encryption with Amazon S3-managed keys SSE-S3 or AWS KMS
// customer master keys (CMKs) (SSE-KMS). For information about the Amazon S3
// default encryption feature, see Amazon S3 Default Bucket Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html).
//
// This operation requires AWS Signature Version 4. For more information, see
// Authenticating Requests (AWS Signature Version 4) (sig-v4-authenticating-requests.html).
//
// To use this operation, you must have permissions to perform the s3:PutEncryptionConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Related Resources
//
//    * GetBucketEncryption
//
//    * DeleteBucketEncryption
//
//    // Example sending a request using PutBucketEncryptionRequest.
//    req := client.PutBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketEncryption
func (c *Client) PutBucketEncryptionRequest(input *PutBucketEncryptionInput) PutBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opPutBucketEncryption,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &PutBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &PutBucketEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	req.Handlers.Build.PushBackNamed(aws.NamedHandler{
		Name: "contentMd5Handler",
		Fn:   checksum.AddBodyContentMD5Handler,
	})

	return PutBucketEncryptionRequest{Request: req, Input: input, Copy: c.PutBucketEncryptionRequest}
}

// PutBucketEncryptionRequest is the request type for the
// PutBucketEncryption API operation.
type PutBucketEncryptionRequest struct {
	*aws.Request
	Input *PutBucketEncryptionInput
	Copy  func(*PutBucketEncryptionInput) PutBucketEncryptionRequest
}

// Send marshals and sends the PutBucketEncryption API request.
func (r PutBucketEncryptionRequest) Send(ctx context.Context) (*PutBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketEncryptionResponse{
		PutBucketEncryptionOutput: r.Request.Data.(*PutBucketEncryptionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketEncryptionResponse is the response type for the
// PutBucketEncryption API operation.
type PutBucketEncryptionResponse struct {
	*PutBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketEncryption request.
func (r *PutBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
