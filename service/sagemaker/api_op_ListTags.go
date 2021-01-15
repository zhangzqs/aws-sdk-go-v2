// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the tags for the specified Amazon SageMaker resource.
func (c *Client) ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) {
	if params == nil {
		params = &ListTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTags", params, optFns, addOperationListTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsInput struct {

	// The Amazon Resource Name (ARN) of the resource whose tags you want to retrieve.
	//
	// This member is required.
	ResourceArn *string

	// Maximum number of tags to return.
	MaxResults *int32

	// If the response to the previous ListTags request is truncated, Amazon SageMaker
	// returns this token. To retrieve the next set of tags, use it in the subsequent
	// request.
	NextToken *string
}

type ListTagsOutput struct {

	// If response is truncated, Amazon SageMaker includes a token in the response. You
	// can use this token in your subsequent request to fetch next set of tokens.
	NextToken *string

	// An array of Tag objects, each with a tag key and a value.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTags{}, middleware.After)
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
	if err = addOpListTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTags(options.Region), middleware.Before); err != nil {
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

// ListTagsAPIClient is a client that implements the ListTags operation.
type ListTagsAPIClient interface {
	ListTags(context.Context, *ListTagsInput, ...func(*Options)) (*ListTagsOutput, error)
}

var _ ListTagsAPIClient = (*Client)(nil)

// ListTagsPaginatorOptions is the paginator options for ListTags
type ListTagsPaginatorOptions struct {
	// Maximum number of tags to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTagsPaginator is a paginator for ListTags
type ListTagsPaginator struct {
	options   ListTagsPaginatorOptions
	client    ListTagsAPIClient
	params    *ListTagsInput
	nextToken *string
	firstPage bool
}

// NewListTagsPaginator returns a new ListTagsPaginator
func NewListTagsPaginator(client ListTagsAPIClient, params *ListTagsInput, optFns ...func(*ListTagsPaginatorOptions)) *ListTagsPaginator {
	options := ListTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTagsInput{}
	}

	return &ListTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTagsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTags page.
func (p *ListTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTagsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListTags(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTags",
	}
}
