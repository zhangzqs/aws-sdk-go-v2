// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// This examples serializes a streaming media-typed blob shape in the request body.
// This examples uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) StreamingTraitsWithMediaType(ctx context.Context, params *StreamingTraitsWithMediaTypeInput, optFns ...func(*Options)) (*StreamingTraitsWithMediaTypeOutput, error) {
	if params == nil {
		params = &StreamingTraitsWithMediaTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StreamingTraitsWithMediaType", params, optFns, addOperationStreamingTraitsWithMediaTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StreamingTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamingTraitsWithMediaTypeInput struct {

	// This value conforms to the media type: text/plain
	Blob io.Reader

	Foo *string
}

type StreamingTraitsWithMediaTypeOutput struct {

	// This value conforms to the media type: text/plain
	Blob io.ReadCloser

	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStreamingTraitsWithMediaTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStreamingTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamingTraitsWithMediaType{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStreamingTraitsWithMediaType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStreamingTraitsWithMediaType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StreamingTraitsWithMediaType",
	}
}
