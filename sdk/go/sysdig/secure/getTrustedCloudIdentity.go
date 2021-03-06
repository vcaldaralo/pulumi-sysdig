// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTrustedCloudIdentity(ctx *pulumi.Context, args *GetTrustedCloudIdentityArgs, opts ...pulumi.InvokeOption) (*GetTrustedCloudIdentityResult, error) {
	var rv GetTrustedCloudIdentityResult
	err := ctx.Invoke("sysdig:Secure/getTrustedCloudIdentity:GetTrustedCloudIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetTrustedCloudIdentity.
type GetTrustedCloudIdentityArgs struct {
	// The cloud provider in which the trusted identity will be used. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider string `pulumi:"cloudProvider"`
}

// A collection of values returned by GetTrustedCloudIdentity.
type GetTrustedCloudIdentityResult struct {
	// If `identity` is an AWS ARN, this attribute contains the AWS Account ID to which the ARN belongs, otherwise it contains the empty string. `cloudProvider` must be equal to `aws` or `gcp`.
	AwsAccountId string `pulumi:"awsAccountId"`
	// If `identity` is a AWS IAM Role ARN, this attribute contains the name of the role, otherwise it contains the empty string. `cloudProvider` must be equal to `aws` or `gcp`.
	AwsRoleName string `pulumi:"awsRoleName"`
	// If `identity` contains credentials for an Azure Service Principal, this attribute contains the service principal's ID. `cloudProvider` must be equal to `azure`.
	AzureServicePrincipalId string `pulumi:"azureServicePrincipalId"`
	// If `identity` contains credentials for an Azure Service Principal, this attribute contains the service principal's Tenant ID. `cloudProvider` must be equal to `azure`.
	AzureTenantId string `pulumi:"azureTenantId"`
	CloudProvider string `pulumi:"cloudProvider"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Sysdig's identity (User/Role/etc) that should be used to create a trust relationship allowing Sysdig access to your cloud account.
	Identity string `pulumi:"identity"`
}

func GetTrustedCloudIdentityOutput(ctx *pulumi.Context, args GetTrustedCloudIdentityOutputArgs, opts ...pulumi.InvokeOption) GetTrustedCloudIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrustedCloudIdentityResult, error) {
			args := v.(GetTrustedCloudIdentityArgs)
			r, err := GetTrustedCloudIdentity(ctx, &args, opts...)
			return *r, err
		}).(GetTrustedCloudIdentityResultOutput)
}

// A collection of arguments for invoking GetTrustedCloudIdentity.
type GetTrustedCloudIdentityOutputArgs struct {
	// The cloud provider in which the trusted identity will be used. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider pulumi.StringInput `pulumi:"cloudProvider"`
}

func (GetTrustedCloudIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrustedCloudIdentityArgs)(nil)).Elem()
}

// A collection of values returned by GetTrustedCloudIdentity.
type GetTrustedCloudIdentityResultOutput struct{ *pulumi.OutputState }

func (GetTrustedCloudIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrustedCloudIdentityResult)(nil)).Elem()
}

func (o GetTrustedCloudIdentityResultOutput) ToGetTrustedCloudIdentityResultOutput() GetTrustedCloudIdentityResultOutput {
	return o
}

func (o GetTrustedCloudIdentityResultOutput) ToGetTrustedCloudIdentityResultOutputWithContext(ctx context.Context) GetTrustedCloudIdentityResultOutput {
	return o
}

// If `identity` is an AWS ARN, this attribute contains the AWS Account ID to which the ARN belongs, otherwise it contains the empty string. `cloudProvider` must be equal to `aws` or `gcp`.
func (o GetTrustedCloudIdentityResultOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AwsAccountId }).(pulumi.StringOutput)
}

// If `identity` is a AWS IAM Role ARN, this attribute contains the name of the role, otherwise it contains the empty string. `cloudProvider` must be equal to `aws` or `gcp`.
func (o GetTrustedCloudIdentityResultOutput) AwsRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AwsRoleName }).(pulumi.StringOutput)
}

// If `identity` contains credentials for an Azure Service Principal, this attribute contains the service principal's ID. `cloudProvider` must be equal to `azure`.
func (o GetTrustedCloudIdentityResultOutput) AzureServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AzureServicePrincipalId }).(pulumi.StringOutput)
}

// If `identity` contains credentials for an Azure Service Principal, this attribute contains the service principal's Tenant ID. `cloudProvider` must be equal to `azure`.
func (o GetTrustedCloudIdentityResultOutput) AzureTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AzureTenantId }).(pulumi.StringOutput)
}

func (o GetTrustedCloudIdentityResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTrustedCloudIdentityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Sysdig's identity (User/Role/etc) that should be used to create a trust relationship allowing Sysdig access to your cloud account.
func (o GetTrustedCloudIdentityResultOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.Identity }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrustedCloudIdentityResultOutput{})
}
