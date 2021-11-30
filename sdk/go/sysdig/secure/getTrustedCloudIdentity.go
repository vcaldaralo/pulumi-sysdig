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
	CloudProvider string `pulumi:"cloudProvider"`
}

// A collection of values returned by GetTrustedCloudIdentity.
type GetTrustedCloudIdentityResult struct {
	AwsAccountId            string `pulumi:"awsAccountId"`
	AwsRoleName             string `pulumi:"awsRoleName"`
	AzureServicePrincipalId string `pulumi:"azureServicePrincipalId"`
	AzureTenantId           string `pulumi:"azureTenantId"`
	CloudProvider           string `pulumi:"cloudProvider"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
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

func (o GetTrustedCloudIdentityResultOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AwsAccountId }).(pulumi.StringOutput)
}

func (o GetTrustedCloudIdentityResultOutput) AwsRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AwsRoleName }).(pulumi.StringOutput)
}

func (o GetTrustedCloudIdentityResultOutput) AzureServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.AzureServicePrincipalId }).(pulumi.StringOutput)
}

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

func (o GetTrustedCloudIdentityResultOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedCloudIdentityResult) string { return v.Identity }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrustedCloudIdentityResultOutput{})
}