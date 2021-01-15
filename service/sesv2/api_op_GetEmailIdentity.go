// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about a specific identity, including the identity's
// verification status, sending authorization policies, its DKIM authentication
// status, and its custom Mail-From settings.
func (c *Client) GetEmailIdentity(ctx context.Context, params *GetEmailIdentityInput, optFns ...func(*Options)) (*GetEmailIdentityOutput, error) {
	if params == nil {
		params = &GetEmailIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEmailIdentity", params, optFns, addOperationGetEmailIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to return details about an email identity.
type GetEmailIdentityInput struct {

	// The email identity that you want to retrieve details for.
	//
	// This member is required.
	EmailIdentity *string
}

// Details about an email identity.
type GetEmailIdentityOutput struct {

	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes *types.DkimAttributes

	// The feedback forwarding configuration for the identity. If the value is true,
	// you receive email notifications when bounce or complaint events occur. These
	// notifications are sent to the address that you specified in the Return-Path
	// header of the original email. You're required to have a method of tracking
	// bounces and complaints. If you haven't set up another mechanism for receiving
	// bounce or complaint notifications (for example, by setting up an event
	// destination), you receive an email notification when these events occur (even if
	// this setting is disabled).
	FeedbackForwardingStatus bool

	// The email identity type.
	IdentityType types.IdentityType

	// An object that contains information about the Mail-From attributes for the email
	// identity.
	MailFromAttributes *types.MailFromAttributes

	// A map of policy names to policies.
	Policies map[string]string

	// An array of objects that define the tags (keys and values) that are associated
	// with the email identity.
	Tags []types.Tag

	// Specifies whether or not the identity is verified. You can only send email from
	// verified email addresses or domains. For more information about verifying
	// identities, see the Amazon Pinpoint User Guide
	// (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEmailIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEmailIdentity{}, middleware.After)
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
	if err = addOpGetEmailIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEmailIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEmailIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetEmailIdentity",
	}
}
