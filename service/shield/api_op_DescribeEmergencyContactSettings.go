// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of email addresses and phone numbers that the DDoS Response Team (DRT)
// can use to contact you if you have proactive engagement enabled, for escalations
// to the DRT and to initiate proactive customer support.
func (c *Client) DescribeEmergencyContactSettings(ctx context.Context, params *DescribeEmergencyContactSettingsInput, optFns ...func(*Options)) (*DescribeEmergencyContactSettingsOutput, error) {
	if params == nil {
		params = &DescribeEmergencyContactSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEmergencyContactSettings", params, optFns, addOperationDescribeEmergencyContactSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEmergencyContactSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEmergencyContactSettingsInput struct {
}

type DescribeEmergencyContactSettingsOutput struct {

	// A list of email addresses and phone numbers that the DDoS Response Team (DRT)
	// can use to contact you if you have proactive engagement enabled, for escalations
	// to the DRT and to initiate proactive customer support.
	EmergencyContactList []types.EmergencyContact

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEmergencyContactSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEmergencyContactSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEmergencyContactSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEmergencyContactSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEmergencyContactSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DescribeEmergencyContactSettings",
	}
}
