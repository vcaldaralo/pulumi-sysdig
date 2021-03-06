// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Secure Cloud Accounts can be imported using the `account_id`, e.g.
//
// ```sh
//  $ pulumi import sysdig:Secure/cloudAccount:CloudAccount sample 123456789012
// ```
type CloudAccount struct {
	pulumi.CustomResourceState

	// The unique identifier of the cloud account. e.g. for AWS: `123456789012`,
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A human friendly alias for `accountId`.
	Alias pulumi.StringPtrOutput `pulumi:"alias"`
	// The cloud provider in which the account exists. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	ExternalId    pulumi.StringOutput `pulumi:"externalId"`
	// Whether or not a role is provisioned withing this account, that Sysdig has permission to AssumeRole in order to run Benchmarks. Default: `false`.
	RoleEnabled pulumi.BoolPtrOutput `pulumi:"roleEnabled"`
	// The name of the role Sysdig will have permission to AssumeRole if `roleEnaled` is set to `true`. Default: `SysdigCloudBench`.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
}

// NewCloudAccount registers a new resource with the given unique name, arguments, and options.
func NewCloudAccount(ctx *pulumi.Context,
	name string, args *CloudAccountArgs, opts ...pulumi.ResourceOption) (*CloudAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.CloudProvider == nil {
		return nil, errors.New("invalid value for required argument 'CloudProvider'")
	}
	var resource CloudAccount
	err := ctx.RegisterResource("sysdig:Secure/cloudAccount:CloudAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudAccount gets an existing CloudAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudAccountState, opts ...pulumi.ResourceOption) (*CloudAccount, error) {
	var resource CloudAccount
	err := ctx.ReadResource("sysdig:Secure/cloudAccount:CloudAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudAccount resources.
type cloudAccountState struct {
	// The unique identifier of the cloud account. e.g. for AWS: `123456789012`,
	AccountId *string `pulumi:"accountId"`
	// A human friendly alias for `accountId`.
	Alias *string `pulumi:"alias"`
	// The cloud provider in which the account exists. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider *string `pulumi:"cloudProvider"`
	ExternalId    *string `pulumi:"externalId"`
	// Whether or not a role is provisioned withing this account, that Sysdig has permission to AssumeRole in order to run Benchmarks. Default: `false`.
	RoleEnabled *bool `pulumi:"roleEnabled"`
	// The name of the role Sysdig will have permission to AssumeRole if `roleEnaled` is set to `true`. Default: `SysdigCloudBench`.
	RoleName *string `pulumi:"roleName"`
}

type CloudAccountState struct {
	// The unique identifier of the cloud account. e.g. for AWS: `123456789012`,
	AccountId pulumi.StringPtrInput
	// A human friendly alias for `accountId`.
	Alias pulumi.StringPtrInput
	// The cloud provider in which the account exists. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider pulumi.StringPtrInput
	ExternalId    pulumi.StringPtrInput
	// Whether or not a role is provisioned withing this account, that Sysdig has permission to AssumeRole in order to run Benchmarks. Default: `false`.
	RoleEnabled pulumi.BoolPtrInput
	// The name of the role Sysdig will have permission to AssumeRole if `roleEnaled` is set to `true`. Default: `SysdigCloudBench`.
	RoleName pulumi.StringPtrInput
}

func (CloudAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountState)(nil)).Elem()
}

type cloudAccountArgs struct {
	// The unique identifier of the cloud account. e.g. for AWS: `123456789012`,
	AccountId string `pulumi:"accountId"`
	// A human friendly alias for `accountId`.
	Alias *string `pulumi:"alias"`
	// The cloud provider in which the account exists. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider string `pulumi:"cloudProvider"`
	// Whether or not a role is provisioned withing this account, that Sysdig has permission to AssumeRole in order to run Benchmarks. Default: `false`.
	RoleEnabled *bool `pulumi:"roleEnabled"`
	// The name of the role Sysdig will have permission to AssumeRole if `roleEnaled` is set to `true`. Default: `SysdigCloudBench`.
	RoleName *string `pulumi:"roleName"`
}

// The set of arguments for constructing a CloudAccount resource.
type CloudAccountArgs struct {
	// The unique identifier of the cloud account. e.g. for AWS: `123456789012`,
	AccountId pulumi.StringInput
	// A human friendly alias for `accountId`.
	Alias pulumi.StringPtrInput
	// The cloud provider in which the account exists. Currently supported providers are `aws`, `gcp` and `azure`
	CloudProvider pulumi.StringInput
	// Whether or not a role is provisioned withing this account, that Sysdig has permission to AssumeRole in order to run Benchmarks. Default: `false`.
	RoleEnabled pulumi.BoolPtrInput
	// The name of the role Sysdig will have permission to AssumeRole if `roleEnaled` is set to `true`. Default: `SysdigCloudBench`.
	RoleName pulumi.StringPtrInput
}

func (CloudAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountArgs)(nil)).Elem()
}

type CloudAccountInput interface {
	pulumi.Input

	ToCloudAccountOutput() CloudAccountOutput
	ToCloudAccountOutputWithContext(ctx context.Context) CloudAccountOutput
}

func (*CloudAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudAccount)(nil))
}

func (i *CloudAccount) ToCloudAccountOutput() CloudAccountOutput {
	return i.ToCloudAccountOutputWithContext(context.Background())
}

func (i *CloudAccount) ToCloudAccountOutputWithContext(ctx context.Context) CloudAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountOutput)
}

func (i *CloudAccount) ToCloudAccountPtrOutput() CloudAccountPtrOutput {
	return i.ToCloudAccountPtrOutputWithContext(context.Background())
}

func (i *CloudAccount) ToCloudAccountPtrOutputWithContext(ctx context.Context) CloudAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountPtrOutput)
}

type CloudAccountPtrInput interface {
	pulumi.Input

	ToCloudAccountPtrOutput() CloudAccountPtrOutput
	ToCloudAccountPtrOutputWithContext(ctx context.Context) CloudAccountPtrOutput
}

type cloudAccountPtrType CloudAccountArgs

func (*cloudAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccount)(nil))
}

func (i *cloudAccountPtrType) ToCloudAccountPtrOutput() CloudAccountPtrOutput {
	return i.ToCloudAccountPtrOutputWithContext(context.Background())
}

func (i *cloudAccountPtrType) ToCloudAccountPtrOutputWithContext(ctx context.Context) CloudAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountPtrOutput)
}

// CloudAccountArrayInput is an input type that accepts CloudAccountArray and CloudAccountArrayOutput values.
// You can construct a concrete instance of `CloudAccountArrayInput` via:
//
//          CloudAccountArray{ CloudAccountArgs{...} }
type CloudAccountArrayInput interface {
	pulumi.Input

	ToCloudAccountArrayOutput() CloudAccountArrayOutput
	ToCloudAccountArrayOutputWithContext(context.Context) CloudAccountArrayOutput
}

type CloudAccountArray []CloudAccountInput

func (CloudAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccount)(nil)).Elem()
}

func (i CloudAccountArray) ToCloudAccountArrayOutput() CloudAccountArrayOutput {
	return i.ToCloudAccountArrayOutputWithContext(context.Background())
}

func (i CloudAccountArray) ToCloudAccountArrayOutputWithContext(ctx context.Context) CloudAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountArrayOutput)
}

// CloudAccountMapInput is an input type that accepts CloudAccountMap and CloudAccountMapOutput values.
// You can construct a concrete instance of `CloudAccountMapInput` via:
//
//          CloudAccountMap{ "key": CloudAccountArgs{...} }
type CloudAccountMapInput interface {
	pulumi.Input

	ToCloudAccountMapOutput() CloudAccountMapOutput
	ToCloudAccountMapOutputWithContext(context.Context) CloudAccountMapOutput
}

type CloudAccountMap map[string]CloudAccountInput

func (CloudAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccount)(nil)).Elem()
}

func (i CloudAccountMap) ToCloudAccountMapOutput() CloudAccountMapOutput {
	return i.ToCloudAccountMapOutputWithContext(context.Background())
}

func (i CloudAccountMap) ToCloudAccountMapOutputWithContext(ctx context.Context) CloudAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountMapOutput)
}

type CloudAccountOutput struct{ *pulumi.OutputState }

func (CloudAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudAccount)(nil))
}

func (o CloudAccountOutput) ToCloudAccountOutput() CloudAccountOutput {
	return o
}

func (o CloudAccountOutput) ToCloudAccountOutputWithContext(ctx context.Context) CloudAccountOutput {
	return o
}

func (o CloudAccountOutput) ToCloudAccountPtrOutput() CloudAccountPtrOutput {
	return o.ToCloudAccountPtrOutputWithContext(context.Background())
}

func (o CloudAccountOutput) ToCloudAccountPtrOutputWithContext(ctx context.Context) CloudAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudAccount) *CloudAccount {
		return &v
	}).(CloudAccountPtrOutput)
}

type CloudAccountPtrOutput struct{ *pulumi.OutputState }

func (CloudAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccount)(nil))
}

func (o CloudAccountPtrOutput) ToCloudAccountPtrOutput() CloudAccountPtrOutput {
	return o
}

func (o CloudAccountPtrOutput) ToCloudAccountPtrOutputWithContext(ctx context.Context) CloudAccountPtrOutput {
	return o
}

func (o CloudAccountPtrOutput) Elem() CloudAccountOutput {
	return o.ApplyT(func(v *CloudAccount) CloudAccount {
		if v != nil {
			return *v
		}
		var ret CloudAccount
		return ret
	}).(CloudAccountOutput)
}

type CloudAccountArrayOutput struct{ *pulumi.OutputState }

func (CloudAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudAccount)(nil))
}

func (o CloudAccountArrayOutput) ToCloudAccountArrayOutput() CloudAccountArrayOutput {
	return o
}

func (o CloudAccountArrayOutput) ToCloudAccountArrayOutputWithContext(ctx context.Context) CloudAccountArrayOutput {
	return o
}

func (o CloudAccountArrayOutput) Index(i pulumi.IntInput) CloudAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudAccount {
		return vs[0].([]CloudAccount)[vs[1].(int)]
	}).(CloudAccountOutput)
}

type CloudAccountMapOutput struct{ *pulumi.OutputState }

func (CloudAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudAccount)(nil))
}

func (o CloudAccountMapOutput) ToCloudAccountMapOutput() CloudAccountMapOutput {
	return o
}

func (o CloudAccountMapOutput) ToCloudAccountMapOutputWithContext(ctx context.Context) CloudAccountMapOutput {
	return o
}

func (o CloudAccountMapOutput) MapIndex(k pulumi.StringInput) CloudAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudAccount {
		return vs[0].(map[string]CloudAccount)[vs[1].(string)]
	}).(CloudAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudAccountOutput{})
	pulumi.RegisterOutputType(CloudAccountPtrOutput{})
	pulumi.RegisterOutputType(CloudAccountArrayOutput{})
	pulumi.RegisterOutputType(CloudAccountMapOutput{})
}
