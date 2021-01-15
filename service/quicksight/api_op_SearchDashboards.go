// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for dashboards that belong to a user.
func (c *Client) SearchDashboards(ctx context.Context, params *SearchDashboardsInput, optFns ...func(*Options)) (*SearchDashboardsOutput, error) {
	if params == nil {
		params = &SearchDashboardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDashboards", params, optFns, addOperationSearchDashboardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDashboardsInput struct {

	// The ID of the AWS account that contains the user whose dashboards you're
	// searching for.
	//
	// This member is required.
	AwsAccountId *string

	// The filters to apply to the search. Currently, you can search only by user name,
	// for example, "Filters": [ { "Name": "QUICKSIGHT_USER", "Operator":
	// "StringEquals", "Value": "arn:aws:quicksight:us-east-1:1:user/default/UserName1"
	// } ]
	//
	// This member is required.
	Filters []types.DashboardSearchFilter

	// The maximum number of results to be returned per request.
	MaxResults int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type SearchDashboardsOutput struct {

	// The list of dashboards owned by the user specified in Filters in your request.
	DashboardSummaryList []types.DashboardSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchDashboardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchDashboards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchDashboards{}, middleware.After)
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
	if err = addOpSearchDashboardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDashboards(options.Region), middleware.Before); err != nil {
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

// SearchDashboardsAPIClient is a client that implements the SearchDashboards
// operation.
type SearchDashboardsAPIClient interface {
	SearchDashboards(context.Context, *SearchDashboardsInput, ...func(*Options)) (*SearchDashboardsOutput, error)
}

var _ SearchDashboardsAPIClient = (*Client)(nil)

// SearchDashboardsPaginatorOptions is the paginator options for SearchDashboards
type SearchDashboardsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchDashboardsPaginator is a paginator for SearchDashboards
type SearchDashboardsPaginator struct {
	options   SearchDashboardsPaginatorOptions
	client    SearchDashboardsAPIClient
	params    *SearchDashboardsInput
	nextToken *string
	firstPage bool
}

// NewSearchDashboardsPaginator returns a new SearchDashboardsPaginator
func NewSearchDashboardsPaginator(client SearchDashboardsAPIClient, params *SearchDashboardsInput, optFns ...func(*SearchDashboardsPaginatorOptions)) *SearchDashboardsPaginator {
	options := SearchDashboardsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &SearchDashboardsInput{}
	}

	return &SearchDashboardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchDashboardsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SearchDashboards page.
func (p *SearchDashboardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchDashboardsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.SearchDashboards(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchDashboards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchDashboards",
	}
}
