// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets face detection results for a Amazon Rekognition Video analysis started by
// StartFaceDetection. Face detection with Amazon Rekognition Video is an
// asynchronous operation. You start face detection by calling StartFaceDetection
// which returns a job identifier (JobId). When the face detection operation
// finishes, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to
// StartFaceDetection. To get the results of the face detection operation, first
// check that the status value published to the Amazon SNS topic is SUCCEEDED. If
// so, call GetFaceDetection and pass the job identifier (JobId) from the initial
// call to StartFaceDetection. GetFaceDetection returns an array of detected faces
// (Faces) sorted by the time the faces were detected. Use MaxResults parameter to
// limit the number of labels returned. If there are more results than specified in
// MaxResults, the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetFaceDetection and populate the NextToken request parameter with
// the token value returned from the previous call to GetFaceDetection.
func (c *Client) GetFaceDetection(ctx context.Context, params *GetFaceDetectionInput, optFns ...func(*Options)) (*GetFaceDetectionOutput, error) {
	if params == nil {
		params = &GetFaceDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFaceDetection", params, optFns, addOperationGetFaceDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFaceDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFaceDetectionInput struct {

	// Unique identifier for the face detection job. The JobId is returned from
	// StartFaceDetection.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more faces to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of faces.
	NextToken *string
}

type GetFaceDetectionOutput struct {

	// An array of faces detected in the video. Each element contains a detected face's
	// details and the time, in milliseconds from the start of the video, the face was
	// detected.
	Faces []types.FaceDetection

	// The current status of the face detection job.
	JobStatus types.VideoJobStatus

	// If the response is truncated, Amazon Rekognition returns this token that you can
	// use in the subsequent request to retrieve the next set of faces.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFaceDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFaceDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFaceDetection{}, middleware.After)
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
	if err = addOpGetFaceDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFaceDetection(options.Region), middleware.Before); err != nil {
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

// GetFaceDetectionAPIClient is a client that implements the GetFaceDetection
// operation.
type GetFaceDetectionAPIClient interface {
	GetFaceDetection(context.Context, *GetFaceDetectionInput, ...func(*Options)) (*GetFaceDetectionOutput, error)
}

var _ GetFaceDetectionAPIClient = (*Client)(nil)

// GetFaceDetectionPaginatorOptions is the paginator options for GetFaceDetection
type GetFaceDetectionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFaceDetectionPaginator is a paginator for GetFaceDetection
type GetFaceDetectionPaginator struct {
	options   GetFaceDetectionPaginatorOptions
	client    GetFaceDetectionAPIClient
	params    *GetFaceDetectionInput
	nextToken *string
	firstPage bool
}

// NewGetFaceDetectionPaginator returns a new GetFaceDetectionPaginator
func NewGetFaceDetectionPaginator(client GetFaceDetectionAPIClient, params *GetFaceDetectionInput, optFns ...func(*GetFaceDetectionPaginatorOptions)) *GetFaceDetectionPaginator {
	options := GetFaceDetectionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetFaceDetectionInput{}
	}

	return &GetFaceDetectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFaceDetectionPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetFaceDetection page.
func (p *GetFaceDetectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFaceDetectionOutput, error) {
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

	result, err := p.client.GetFaceDetection(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFaceDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetFaceDetection",
	}
}
