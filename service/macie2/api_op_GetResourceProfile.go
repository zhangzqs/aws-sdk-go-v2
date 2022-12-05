// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) sensitive data discovery statistics and the sensitivity
// score for an S3 bucket.
func (c *Client) GetResourceProfile(ctx context.Context, params *GetResourceProfileInput, optFns ...func(*Options)) (*GetResourceProfileOutput, error) {
	if params == nil {
		params = &GetResourceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceProfile", params, optFns, c.addOperationGetResourceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceProfileInput struct {

	// The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type GetResourceProfileOutput struct {

	// The date and time, in UTC and extended ISO 8601 format, when Amazon Macie most
	// recently recalculated sensitive data discovery statistics and details for the
	// bucket. If the bucket's sensitivity score is calculated automatically, this
	// includes the score.
	ProfileUpdatedAt *time.Time

	// The current sensitivity score for the bucket, ranging from -1 (no analysis due
	// to an error) to 100 (sensitive). By default, this score is calculated
	// automatically based on the amount of data that Amazon Macie has analyzed in the
	// bucket and the amount of sensitive data that Macie has found in the bucket.
	SensitivityScore int32

	// Specifies whether the bucket's current sensitivity score was set manually. If
	// this value is true, the score was manually changed to 100. If this value is
	// false, the score was calculated automatically by Amazon Macie.
	SensitivityScoreOverridden bool

	// The sensitive data discovery statistics for the bucket. The statistics capture
	// the results of automated sensitive data discovery activities that Amazon Macie
	// has performed for the bucket.
	Statistics *types.ResourceStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceProfile{}, middleware.After)
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
	if err = addOpGetResourceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetResourceProfile",
	}
}