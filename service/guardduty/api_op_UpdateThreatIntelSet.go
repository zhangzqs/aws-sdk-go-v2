// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the ThreatIntelSet specified by the ThreatIntelSet ID.
func (c *Client) UpdateThreatIntelSet(ctx context.Context, params *UpdateThreatIntelSetInput, optFns ...func(*Options)) (*UpdateThreatIntelSetOutput, error) {
	if params == nil {
		params = &UpdateThreatIntelSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThreatIntelSet", params, optFns, addOperationUpdateThreatIntelSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThreatIntelSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThreatIntelSetInput struct {

	// The detectorID that specifies the GuardDuty service whose ThreatIntelSet you
	// want to update.
	//
	// This member is required.
	DetectorId *string

	// The unique ID that specifies the ThreatIntelSet that you want to update.
	//
	// This member is required.
	ThreatIntelSetId *string

	// The updated Boolean value that specifies whether the ThreateIntelSet is active
	// or not.
	Activate bool

	// The updated URI of the file that contains the ThreateIntelSet.
	Location *string

	// The unique ID that specifies the ThreatIntelSet that you want to update.
	Name *string
}

type UpdateThreatIntelSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateThreatIntelSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThreatIntelSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThreatIntelSet{}, middleware.After)
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
	if err = addOpUpdateThreatIntelSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThreatIntelSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateThreatIntelSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateThreatIntelSet",
	}
}
