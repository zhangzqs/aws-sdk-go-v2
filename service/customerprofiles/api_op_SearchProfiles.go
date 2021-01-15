// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for profiles within a specific domain name using name, phone number,
// email address, account number, or a custom defined index.
func (c *Client) SearchProfiles(ctx context.Context, params *SearchProfilesInput, optFns ...func(*Options)) (*SearchProfilesOutput, error) {
	if params == nil {
		params = &SearchProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProfiles", params, optFns, addOperationSearchProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProfilesInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A searchable identifier of a customer profile. The predefined keys you can use
	// to search include: _account, _profileId, _fullName, _phone, _email,
	// _ctrContactId, _marketoLeadId, _salesforceAccountId, _salesforceContactId,
	// _zendeskUserId, _zendeskExternalId, _serviceNowSystemId.
	//
	// This member is required.
	KeyName *string

	// A list of key values.
	//
	// This member is required.
	Values []string

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous SearchProfiles API call.
	NextToken *string
}

type SearchProfilesOutput struct {

	// The list of SearchProfiles instances.
	Items []types.Profile

	// The pagination token from the previous SearchProfiles API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchProfiles{}, middleware.After)
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
	if err = addOpSearchProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProfiles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "SearchProfiles",
	}
}
