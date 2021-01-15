// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a transit gateway Connect peer from a device and link.
func (c *Client) DisassociateTransitGatewayConnectPeer(ctx context.Context, params *DisassociateTransitGatewayConnectPeerInput, optFns ...func(*Options)) (*DisassociateTransitGatewayConnectPeerOutput, error) {
	if params == nil {
		params = &DisassociateTransitGatewayConnectPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTransitGatewayConnectPeer", params, optFns, addOperationDisassociateTransitGatewayConnectPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTransitGatewayConnectPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTransitGatewayConnectPeerInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Resource Name (ARN) of the transit gateway Connect peer.
	//
	// This member is required.
	TransitGatewayConnectPeerArn *string
}

type DisassociateTransitGatewayConnectPeerOutput struct {

	// The transit gateway Connect peer association.
	TransitGatewayConnectPeerAssociation *types.TransitGatewayConnectPeerAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateTransitGatewayConnectPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateTransitGatewayConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateTransitGatewayConnectPeer{}, middleware.After)
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
	if err = addOpDisassociateTransitGatewayConnectPeerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTransitGatewayConnectPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateTransitGatewayConnectPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DisassociateTransitGatewayConnectPeer",
	}
}
