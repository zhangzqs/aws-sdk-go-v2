// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are enabled and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are enabled. If multiple options are enabled and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign in. If an MFA type is enabled for a user, the user will be prompted for MFA
// during all sign in attempts, unless device tracking is turned on and the device
// has been trusted. If you would like MFA to be applied selectively based on the
// assessed risk level of sign in attempts, disable MFA for users and turn on
// Adaptive Authentication for the user pool.
func (c *Client) SetUserMFAPreference(ctx context.Context, params *SetUserMFAPreferenceInput, optFns ...func(*Options)) (*SetUserMFAPreferenceOutput, error) {
	if params == nil {
		params = &SetUserMFAPreferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserMFAPreference", params, optFns, addOperationSetUserMFAPreferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserMFAPreferenceInput struct {

	// The access token for the user.
	//
	// This member is required.
	AccessToken *string

	// The SMS text message multi-factor authentication (MFA) settings.
	SMSMfaSettings *types.SMSMfaSettingsType

	// The time-based one-time password software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType
}

type SetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetUserMFAPreferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserMFAPreference{}, middleware.After)
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
	if err = addOpSetUserMFAPreferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserMFAPreference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetUserMFAPreference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUserMFAPreference",
	}
}
