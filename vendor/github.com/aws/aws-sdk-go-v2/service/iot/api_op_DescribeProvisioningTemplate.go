// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a provisioning template. Requires permission to
// access the DescribeProvisioningTemplate (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeProvisioningTemplate(ctx context.Context, params *DescribeProvisioningTemplateInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateOutput, error) {
	if params == nil {
		params = &DescribeProvisioningTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProvisioningTemplate", params, optFns, c.addOperationDescribeProvisioningTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningTemplateInput struct {

	// The name of the provisioning template.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

type DescribeProvisioningTemplateOutput struct {

	// The date when the provisioning template was created.
	CreationDate *time.Time

	// The default fleet template version ID.
	DefaultVersionId *int32

	// The description of the provisioning template.
	Description *string

	// True if the provisioning template is enabled, otherwise false.
	Enabled bool

	// The date when the provisioning template was last modified.
	LastModifiedDate *time.Time

	// Gets information about a pre-provisioned hook.
	PreProvisioningHook *types.ProvisioningHook

	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string

	// The ARN of the provisioning template.
	TemplateArn *string

	// The JSON formatted contents of the provisioning template.
	TemplateBody *string

	// The name of the provisioning template.
	TemplateName *string

	// The type you define in a provisioning template. You can create a template with
	// only one type. You can't change the template type after its creation. The
	// default value is FLEET_PROVISIONING . For more information about provisioning
	// template, see: Provisioning template (https://docs.aws.amazon.com/iot/latest/developerguide/provision-template.html)
	// .
	Type types.TemplateType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProvisioningTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProvisioningTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProvisioningTemplate{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeProvisioningTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDescribeProvisioningTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "DescribeProvisioningTemplate",
	}
}
