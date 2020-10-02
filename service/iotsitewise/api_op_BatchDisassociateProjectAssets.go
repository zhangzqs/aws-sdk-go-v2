// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a group (batch) of assets from an AWS IoT SiteWise Monitor
// project.
func (c *Client) BatchDisassociateProjectAssets(ctx context.Context, params *BatchDisassociateProjectAssetsInput, optFns ...func(*Options)) (*BatchDisassociateProjectAssetsOutput, error) {
	stack := middleware.NewStack("BatchDisassociateProjectAssets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchDisassociateProjectAssetsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addIdempotencyToken_opBatchDisassociateProjectAssetsMiddleware(stack, options)
	addOpBatchDisassociateProjectAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateProjectAssets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "BatchDisassociateProjectAssets",
			Err:           err,
		}
	}
	out := result.(*BatchDisassociateProjectAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateProjectAssetsInput struct {

	// The IDs of the assets to be disassociated from the project.
	//
	// This member is required.
	AssetIds []*string

	// The ID of the project from which to disassociate the assets.
	//
	// This member is required.
	ProjectId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type BatchDisassociateProjectAssetsOutput struct {

	// A list of associated error information, if any.
	Errors []*types.AssetErrorDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchDisassociateProjectAssetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisassociateProjectAssets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisassociateProjectAssets{}, middleware.After)
}

type idempotencyToken_initializeOpBatchDisassociateProjectAssets struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpBatchDisassociateProjectAssets) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpBatchDisassociateProjectAssets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*BatchDisassociateProjectAssetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *BatchDisassociateProjectAssetsInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opBatchDisassociateProjectAssetsMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpBatchDisassociateProjectAssets{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opBatchDisassociateProjectAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchDisassociateProjectAssets",
	}
}