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

// Retrieves a subset of information about all the findings filters for an account.
func (c *Client) ListFindingsFilters(ctx context.Context, params *ListFindingsFiltersInput, optFns ...func(*Options)) (*ListFindingsFiltersOutput, error) {
	if params == nil {
		params = &ListFindingsFiltersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindingsFilters", params, optFns, addOperationListFindingsFiltersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingsFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingsFiltersInput struct {

	// The maximum number of items to include in each page of a paginated response.
	MaxResults int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string
}

type ListFindingsFiltersOutput struct {

	// An array of objects, one for each filter that's associated with the account.
	FindingsFilterListItems []types.FindingsFilterListItem

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFindingsFiltersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindingsFilters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindingsFilters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindingsFilters(options.Region), middleware.Before); err != nil {
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

// ListFindingsFiltersAPIClient is a client that implements the ListFindingsFilters
// operation.
type ListFindingsFiltersAPIClient interface {
	ListFindingsFilters(context.Context, *ListFindingsFiltersInput, ...func(*Options)) (*ListFindingsFiltersOutput, error)
}

var _ ListFindingsFiltersAPIClient = (*Client)(nil)

// ListFindingsFiltersPaginatorOptions is the paginator options for
// ListFindingsFilters
type ListFindingsFiltersPaginatorOptions struct {
	// The maximum number of items to include in each page of a paginated response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingsFiltersPaginator is a paginator for ListFindingsFilters
type ListFindingsFiltersPaginator struct {
	options   ListFindingsFiltersPaginatorOptions
	client    ListFindingsFiltersAPIClient
	params    *ListFindingsFiltersInput
	nextToken *string
	firstPage bool
}

// NewListFindingsFiltersPaginator returns a new ListFindingsFiltersPaginator
func NewListFindingsFiltersPaginator(client ListFindingsFiltersAPIClient, params *ListFindingsFiltersInput, optFns ...func(*ListFindingsFiltersPaginatorOptions)) *ListFindingsFiltersPaginator {
	options := ListFindingsFiltersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFindingsFiltersInput{}
	}

	return &ListFindingsFiltersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingsFiltersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFindingsFilters page.
func (p *ListFindingsFiltersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingsFiltersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListFindingsFilters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFindingsFilters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListFindingsFilters",
	}
}
