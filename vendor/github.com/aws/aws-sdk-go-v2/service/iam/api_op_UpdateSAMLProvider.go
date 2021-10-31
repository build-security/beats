// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	// An XML document generated by an identity provider (IdP) that supports SAML
	// 2.0. The document includes the issuer's name, expiration information, and
	// keys that can be used to validate the SAML authentication response (assertions)
	// that are received from the IdP. You must generate the metadata document using
	// the identity management software that is used as your organization's IdP.
	//
	// SAMLMetadataDocument is a required field
	SAMLMetadataDocument *string `min:"1000" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the SAML provider to update.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// SAMLProviderArn is a required field
	SAMLProviderArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSAMLProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSAMLProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSAMLProviderInput"}

	if s.SAMLMetadataDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("SAMLMetadataDocument"))
	}
	if s.SAMLMetadataDocument != nil && len(*s.SAMLMetadataDocument) < 1000 {
		invalidParams.Add(aws.NewErrParamMinLen("SAMLMetadataDocument", 1000))
	}

	if s.SAMLProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SAMLProviderArn"))
	}
	if s.SAMLProviderArn != nil && len(*s.SAMLProviderArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("SAMLProviderArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful UpdateSAMLProvider request.
type UpdateSAMLProviderOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the SAML provider that was updated.
	SAMLProviderArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s UpdateSAMLProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSAMLProvider = "UpdateSAMLProvider"

// UpdateSAMLProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the metadata document for an existing SAML provider resource object.
//
// This operation requires Signature Version 4 (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
//    // Example sending a request using UpdateSAMLProviderRequest.
//    req := client.UpdateSAMLProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateSAMLProvider
func (c *Client) UpdateSAMLProviderRequest(input *UpdateSAMLProviderInput) UpdateSAMLProviderRequest {
	op := &aws.Operation{
		Name:       opUpdateSAMLProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSAMLProviderInput{}
	}

	req := c.newRequest(op, input, &UpdateSAMLProviderOutput{})

	return UpdateSAMLProviderRequest{Request: req, Input: input, Copy: c.UpdateSAMLProviderRequest}
}

// UpdateSAMLProviderRequest is the request type for the
// UpdateSAMLProvider API operation.
type UpdateSAMLProviderRequest struct {
	*aws.Request
	Input *UpdateSAMLProviderInput
	Copy  func(*UpdateSAMLProviderInput) UpdateSAMLProviderRequest
}

// Send marshals and sends the UpdateSAMLProvider API request.
func (r UpdateSAMLProviderRequest) Send(ctx context.Context) (*UpdateSAMLProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSAMLProviderResponse{
		UpdateSAMLProviderOutput: r.Request.Data.(*UpdateSAMLProviderOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSAMLProviderResponse is the response type for the
// UpdateSAMLProvider API operation.
type UpdateSAMLProviderResponse struct {
	*UpdateSAMLProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSAMLProvider request.
func (r *UpdateSAMLProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
