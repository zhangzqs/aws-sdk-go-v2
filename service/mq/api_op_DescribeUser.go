// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about an ActiveMQ user.
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	if params == nil {
		params = &DescribeUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUser", params, optFns, addOperationDescribeUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string

	// The username of the ActiveMQ user. This value can contain only alphanumeric
	// characters, dashes, periods, underscores, and tildes (- . _ ~). This value must
	// be 2-100 characters long.
	//
	// This member is required.
	Username *string
}

type DescribeUserOutput struct {

	// Required. The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// Enables access to the the ActiveMQ Web Console for the ActiveMQ user.
	ConsoleAccess bool

	// The list of groups (20 maximum) to which the ActiveMQ user belongs. This value
	// can contain only alphanumeric characters, dashes, periods, underscores, and
	// tildes (- . _ ~). This value must be 2-100 characters long.
	Groups []string

	// The status of the changes pending for the ActiveMQ user.
	Pending *types.UserPendingChanges

	// Required. The username of the ActiveMQ user. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 2-100 characters long.
	Username *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeUser",
	}
}
