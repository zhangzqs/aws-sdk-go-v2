// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a resource policy that manages access for a model group. For information
// about resource policies, see Identity-based policies and resource-based policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)
// in the AWS Identity and Access Management User Guide..
func (c *Client) GetModelPackageGroupPolicy(ctx context.Context, params *GetModelPackageGroupPolicyInput, optFns ...func(*Options)) (*GetModelPackageGroupPolicyOutput, error) {
	if params == nil {
		params = &GetModelPackageGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModelPackageGroupPolicy", params, optFns, addOperationGetModelPackageGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelPackageGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelPackageGroupPolicyInput struct {

	// The name of the model group for which to get the resource policy.
	//
	// This member is required.
	ModelPackageGroupName *string
}

type GetModelPackageGroupPolicyOutput struct {

	// The resource policy for the model group.
	//
	// This member is required.
	ResourcePolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetModelPackageGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetModelPackageGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetModelPackageGroupPolicy{}, middleware.After)
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
	if err = addOpGetModelPackageGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelPackageGroupPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModelPackageGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "GetModelPackageGroupPolicy",
	}
}
