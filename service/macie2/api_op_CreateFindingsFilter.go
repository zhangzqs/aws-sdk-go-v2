// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and defines the criteria and other settings for a findings filter.
func (c *Client) CreateFindingsFilter(ctx context.Context, params *CreateFindingsFilterInput, optFns ...func(*Options)) (*CreateFindingsFilterOutput, error) {
	if params == nil {
		params = &CreateFindingsFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFindingsFilter", params, optFns, addOperationCreateFindingsFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFindingsFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFindingsFilterInput struct {

	// The action to perform on findings that meet the filter criteria
	// (findingCriteria). Valid values are: ARCHIVE, suppress (automatically archive)
	// the findings; and, NOOP, don't perform any action on the findings.
	//
	// This member is required.
	Action types.FindingsFilterAction

	// The criteria to use to filter findings.
	//
	// This member is required.
	FindingCriteria *types.FindingCriteria

	// A custom name for the filter. The name must contain at least 3 characters and
	// can contain as many as 64 characters. We strongly recommend that you avoid
	// including any sensitive data in the name of a filter. Other users of your
	// account might be able to see the filter's name, depending on the actions that
	// they're allowed to perform in Amazon Macie.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	ClientToken *string

	// A custom description of the filter. The description can contain as many as 512
	// characters. We strongly recommend that you avoid including any sensitive data in
	// the description of a filter. Other users of your account might be able to see
	// the filter's description, depending on the actions that they're allowed to
	// perform in Amazon Macie.
	Description *string

	// The position of the filter in the list of saved filters on the Amazon Macie
	// console. This value also determines the order in which the filter is applied to
	// findings, relative to other filters that are also applied to the findings.
	Position int32

	// A map of key-value pairs that specifies the tags to associate with the filter. A
	// findings filter can have a maximum of 50 tags. Each tag consists of a tag key
	// and an associated tag value. The maximum length of a tag key is 128 characters.
	// The maximum length of a tag value is 256 characters.
	Tags map[string]string
}

type CreateFindingsFilterOutput struct {

	// The Amazon Resource Name (ARN) of the filter that was created.
	Arn *string

	// The unique identifier for the filter that was created.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFindingsFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFindingsFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFindingsFilter{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFindingsFilterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFindingsFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFindingsFilter(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFindingsFilter struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFindingsFilter) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFindingsFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFindingsFilterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFindingsFilterInput ")
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
func addIdempotencyToken_opCreateFindingsFilterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFindingsFilter{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFindingsFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateFindingsFilter",
	}
}
