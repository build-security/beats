// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/checksum"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type DeleteObjectsInput struct {
	_ struct{} `type:"structure" payload:"Delete"`

	// The bucket name containing the objects to delete.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies whether you want to delete this object even if it has a Governance-type
	// Object Lock in place. You must have sufficient permissions to perform this
	// operation.
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// Container for the request.
	//
	// Delete is a required field
	Delete *Delete `locationName:"Delete" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device. Required to
	// permanently delete a versioned object if versioning is configured with MFA
	// delete enabled.
	MFA *string `location:"header" locationName:"x-amz-mfa" type:"string"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Delete == nil {
		invalidParams.Add(aws.NewErrParamRequired("Delete"))
	}
	if s.Delete != nil {
		if err := s.Delete.Validate(); err != nil {
			invalidParams.AddNested("Delete", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteObjectsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BypassGovernanceRetention != nil {
		v := *s.BypassGovernanceRetention

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bypass-governance-retention", protocol.BoolValue(v), metadata)
	}
	if s.MFA != nil {
		v := *s.MFA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mfa", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Delete != nil {
		v := s.Delete

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Delete", v, metadata)
	}
	return nil
}

func (s *DeleteObjectsInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteObjectsInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteObjectsOutput struct {
	_ struct{} `type:"structure"`

	// Container element for a successful delete. It identifies the object that
	// was successfully deleted.
	Deleted []DeletedObject `type:"list" flattened:"true"`

	// Container for a failed delete operation that describes the object that Amazon
	// S3 attempted to delete and the error it encountered.
	Errors []Error `locationName:"Error" type:"list" flattened:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Deleted != nil {
		v := s.Deleted

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Deleted", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Errors != nil {
		v := s.Errors

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Error", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opDeleteObjects = "DeleteObjects"

// DeleteObjectsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation enables you to delete multiple objects from a bucket using
// a single HTTP request. If you know the object keys that you want to delete,
// then this operation provides a suitable alternative to sending individual
// delete requests, reducing per-request overhead.
//
// The request contains a list of up to 1000 keys that you want to delete. In
// the XML, you provide the object key names, and optionally, version IDs if
// you want to delete a specific version of the object from a versioning-enabled
// bucket. For each key, Amazon S3 performs a delete operation and returns the
// result of that delete, success, or failure, in the response. Note that if
// the object specified in the request is not found, Amazon S3 returns the result
// as deleted.
//
// The operation supports two modes for the response: verbose and quiet. By
// default, the operation uses verbose mode in which the response includes the
// result of deletion of each key in your request. In quiet mode the response
// includes only keys where the delete operation encountered an error. For a
// successful deletion, the operation does not return any information about
// the delete in the response body.
//
// When performing this operation on an MFA Delete enabled bucket, that attempts
// to delete any versioned objects, you must include an MFA token. If you do
// not provide one, the entire request will fail, even if there are non-versioned
// objects you are trying to delete. If you provide an invalid token, whether
// there are versioned keys in the request or not, the entire Multi-Object Delete
// request will fail. For information about MFA Delete, see MFA Delete (https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete).
//
// Finally, the Content-MD5 header is required for all Multi-Object Delete requests.
// Amazon S3 uses the header value to ensure that your request body has not
// been altered in transit.
//
// The following operations are related to DeleteObjects:
//
//    * CreateMultipartUpload
//
//    * UploadPart
//
//    * CompleteMultipartUpload
//
//    * ListParts
//
//    * AbortMultipartUpload
//
//    // Example sending a request using DeleteObjectsRequest.
//    req := client.DeleteObjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObjects
func (c *Client) DeleteObjectsRequest(input *DeleteObjectsInput) DeleteObjectsRequest {
	op := &aws.Operation{
		Name:       opDeleteObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}?delete",
	}

	if input == nil {
		input = &DeleteObjectsInput{}
	}

	req := c.newRequest(op, input, &DeleteObjectsOutput{})

	req.Handlers.Build.PushBackNamed(aws.NamedHandler{
		Name: "contentMd5Handler",
		Fn:   checksum.AddBodyContentMD5Handler,
	})

	return DeleteObjectsRequest{Request: req, Input: input, Copy: c.DeleteObjectsRequest}
}

// DeleteObjectsRequest is the request type for the
// DeleteObjects API operation.
type DeleteObjectsRequest struct {
	*aws.Request
	Input *DeleteObjectsInput
	Copy  func(*DeleteObjectsInput) DeleteObjectsRequest
}

// Send marshals and sends the DeleteObjects API request.
func (r DeleteObjectsRequest) Send(ctx context.Context) (*DeleteObjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteObjectsResponse{
		DeleteObjectsOutput: r.Request.Data.(*DeleteObjectsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteObjectsResponse is the response type for the
// DeleteObjects API operation.
type DeleteObjectsResponse struct {
	*DeleteObjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteObjects request.
func (r *DeleteObjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}