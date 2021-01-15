// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the paths to the specified product. A path is how the user has access to a
// specified product, and is necessary when provisioning a product. A path also
// determines the constraints put on the product.
func (c *Client) ListLaunchPaths(ctx context.Context, params *ListLaunchPathsInput, optFns ...func(*Options)) (*ListLaunchPathsOutput, error) {
	if params == nil {
		params = &ListLaunchPathsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLaunchPaths", params, optFns, addOperationListLaunchPathsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLaunchPathsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLaunchPathsInput struct {

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListLaunchPathsOutput struct {

	// Information about the launch path.
	LaunchPathSummaries []types.LaunchPathSummary

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLaunchPathsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLaunchPaths{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLaunchPaths{}, middleware.After)
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
	if err = addOpListLaunchPathsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLaunchPaths(options.Region), middleware.Before); err != nil {
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

// ListLaunchPathsAPIClient is a client that implements the ListLaunchPaths
// operation.
type ListLaunchPathsAPIClient interface {
	ListLaunchPaths(context.Context, *ListLaunchPathsInput, ...func(*Options)) (*ListLaunchPathsOutput, error)
}

var _ ListLaunchPathsAPIClient = (*Client)(nil)

// ListLaunchPathsPaginatorOptions is the paginator options for ListLaunchPaths
type ListLaunchPathsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLaunchPathsPaginator is a paginator for ListLaunchPaths
type ListLaunchPathsPaginator struct {
	options   ListLaunchPathsPaginatorOptions
	client    ListLaunchPathsAPIClient
	params    *ListLaunchPathsInput
	nextToken *string
	firstPage bool
}

// NewListLaunchPathsPaginator returns a new ListLaunchPathsPaginator
func NewListLaunchPathsPaginator(client ListLaunchPathsAPIClient, params *ListLaunchPathsInput, optFns ...func(*ListLaunchPathsPaginatorOptions)) *ListLaunchPathsPaginator {
	options := ListLaunchPathsPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListLaunchPathsInput{}
	}

	return &ListLaunchPathsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLaunchPathsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListLaunchPaths page.
func (p *ListLaunchPathsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLaunchPathsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListLaunchPaths(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListLaunchPaths(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListLaunchPaths",
	}
}
