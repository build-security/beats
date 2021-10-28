// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteMessageInput struct {
	_ struct{} `type:"structure"`

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The receipt handle associated with the message to delete.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMessageInput"}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}

	if s.ReceiptHandle == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReceiptHandle"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMessageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMessageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMessage = "DeleteMessage"

// DeleteMessageRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Deletes the specified message from the specified queue. To select the message
// to delete, use the ReceiptHandle of the message (not the MessageId which
// you receive when you send the message). Amazon SQS can delete a message from
// a queue even if a visibility timeout setting causes the message to be locked
// by another consumer. Amazon SQS automatically deletes messages left in a
// queue longer than the retention period configured for the queue.
//
// The ReceiptHandle is associated with a specific instance of receiving a message.
// If you receive a message more than once, the ReceiptHandle is different each
// time you receive a message. When you use the DeleteMessage action, you must
// provide the most recently received ReceiptHandle for the message (otherwise,
// the request succeeds, but the message might not be deleted).
//
// For standard queues, it is possible to receive a message even after you delete
// it. This might happen on rare occasions if one of the servers which stores
// a copy of the message is unavailable when you send the request to delete
// the message. The copy remains on the server and might be returned to you
// during a subsequent receive request. You should ensure that your application
// is idempotent, so that receiving a message more than once does not cause
// issues.
//
//    // Example sending a request using DeleteMessageRequest.
//    req := client.DeleteMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessage
func (c *Client) DeleteMessageRequest(input *DeleteMessageInput) DeleteMessageRequest {
	op := &aws.Operation{
		Name:       opDeleteMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMessageInput{}
	}

	req := c.newRequest(op, input, &DeleteMessageOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteMessageRequest{Request: req, Input: input, Copy: c.DeleteMessageRequest}
}

// DeleteMessageRequest is the request type for the
// DeleteMessage API operation.
type DeleteMessageRequest struct {
	*aws.Request
	Input *DeleteMessageInput
	Copy  func(*DeleteMessageInput) DeleteMessageRequest
}

// Send marshals and sends the DeleteMessage API request.
func (r DeleteMessageRequest) Send(ctx context.Context) (*DeleteMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMessageResponse{
		DeleteMessageOutput: r.Request.Data.(*DeleteMessageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMessageResponse is the response type for the
// DeleteMessage API operation.
type DeleteMessageResponse struct {
	*DeleteMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMessage request.
func (r *DeleteMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
