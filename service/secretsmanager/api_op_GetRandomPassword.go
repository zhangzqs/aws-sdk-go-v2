// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a random password of the specified complexity. This operation is
// intended for use in the Lambda rotation function. Per best practice, we
// recommend that you specify the maximum length and include every character type
// that the system you are generating a password for can support. Minimum
// permissions To run this command, you must have the following permissions:
//
// *
// secretsmanager:GetRandomPassword
func (c *Client) GetRandomPassword(ctx context.Context, params *GetRandomPasswordInput, optFns ...func(*Options)) (*GetRandomPasswordOutput, error) {
	if params == nil {
		params = &GetRandomPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRandomPassword", params, optFns, addOperationGetRandomPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRandomPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRandomPasswordInput struct {

	// A string that includes characters that should not be included in the generated
	// password. The default is that all characters from the included sets can be used.
	ExcludeCharacters *string

	// Specifies that the generated password should not include lowercase letters. The
	// default if you do not include this switch parameter is that lowercase letters
	// can be included.
	ExcludeLowercase bool

	// Specifies that the generated password should not include digits. The default if
	// you do not include this switch parameter is that digits can be included.
	ExcludeNumbers bool

	// Specifies that the generated password should not include punctuation characters.
	// The default if you do not include this switch parameter is that punctuation
	// characters can be included. The following are the punctuation characters that
	// can be included in the generated password if you don't explicitly exclude them
	// with ExcludeCharacters or ExcludePunctuation: ! " # $ % & ' ( ) * + , - . / : ;
	// < = > ? @ [ \ ] ^ _ ` { | } ~
	ExcludePunctuation bool

	// Specifies that the generated password should not include uppercase letters. The
	// default if you do not include this switch parameter is that uppercase letters
	// can be included.
	ExcludeUppercase bool

	// Specifies that the generated password can include the space character. The
	// default if you do not include this switch parameter is that the space character
	// is not included.
	IncludeSpace bool

	// The desired length of the generated password. The default value if you do not
	// include this parameter is 32 characters.
	PasswordLength int64

	// A boolean value that specifies whether the generated password must include at
	// least one of every allowed character type. The default value is True and the
	// operation requires at least one of every character type.
	RequireEachIncludedType bool
}

type GetRandomPasswordOutput struct {

	// A string with the generated password.
	RandomPassword *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRandomPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRandomPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRandomPassword{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRandomPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRandomPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "GetRandomPassword",
	}
}
