// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describe the attributes of a custom routing accelerator.
func (c *Client) DescribeCustomRoutingAcceleratorAttributes(ctx context.Context, params *DescribeCustomRoutingAcceleratorAttributesInput, optFns ...func(*Options)) (*DescribeCustomRoutingAcceleratorAttributesOutput, error) {
	if params == nil {
		params = &DescribeCustomRoutingAcceleratorAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomRoutingAcceleratorAttributes", params, optFns, addOperationDescribeCustomRoutingAcceleratorAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomRoutingAcceleratorAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomRoutingAcceleratorAttributesInput struct {

	// The Amazon Resource Name (ARN) of the custom routing accelerator to describe the
	// attributes for.
	//
	// This member is required.
	AcceleratorArn *string
}

type DescribeCustomRoutingAcceleratorAttributesOutput struct {

	// The attributes of the custom routing accelerator.
	AcceleratorAttributes *types.CustomRoutingAcceleratorAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCustomRoutingAcceleratorAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCustomRoutingAcceleratorAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCustomRoutingAcceleratorAttributes{}, middleware.After)
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
	if err = addOpDescribeCustomRoutingAcceleratorAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomRoutingAcceleratorAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCustomRoutingAcceleratorAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "DescribeCustomRoutingAcceleratorAttributes",
	}
}
