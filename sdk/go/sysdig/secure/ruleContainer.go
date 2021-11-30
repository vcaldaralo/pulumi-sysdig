// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Secure container runtime rules can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Secure/ruleContainer:RuleContainer example 12345
// ```
type RuleContainer struct {
	pulumi.CustomResourceState

	// List of containers to match.
	Containers pulumi.StringArrayOutput `pulumi:"containers"`
	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defines if the image name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrOutput `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of tags for this rule.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Current version of the resource in Sysdig Secure.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRuleContainer registers a new resource with the given unique name, arguments, and options.
func NewRuleContainer(ctx *pulumi.Context,
	name string, args *RuleContainerArgs, opts ...pulumi.ResourceOption) (*RuleContainer, error) {
	if args == nil {
		args = &RuleContainerArgs{}
	}

	var resource RuleContainer
	err := ctx.RegisterResource("sysdig:Secure/ruleContainer:RuleContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleContainer gets an existing RuleContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleContainerState, opts ...pulumi.ResourceOption) (*RuleContainer, error) {
	var resource RuleContainer
	err := ctx.ReadResource("sysdig:Secure/ruleContainer:RuleContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleContainer resources.
type ruleContainerState struct {
	// List of containers to match.
	Containers []string `pulumi:"containers"`
	// The description of Secure rule. By default is empty.
	Description *string `pulumi:"description"`
	// Defines if the image name matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name *string `pulumi:"name"`
	// A list of tags for this rule.
	Tags []string `pulumi:"tags"`
	// Current version of the resource in Sysdig Secure.
	Version *int `pulumi:"version"`
}

type RuleContainerState struct {
	// List of containers to match.
	Containers pulumi.StringArrayInput
	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrInput
	// Defines if the image name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringPtrInput
	// A list of tags for this rule.
	Tags pulumi.StringArrayInput
	// Current version of the resource in Sysdig Secure.
	Version pulumi.IntPtrInput
}

func (RuleContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleContainerState)(nil)).Elem()
}

type ruleContainerArgs struct {
	// List of containers to match.
	Containers []string `pulumi:"containers"`
	// The description of Secure rule. By default is empty.
	Description *string `pulumi:"description"`
	// Defines if the image name matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name *string `pulumi:"name"`
	// A list of tags for this rule.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a RuleContainer resource.
type RuleContainerArgs struct {
	// List of containers to match.
	Containers pulumi.StringArrayInput
	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrInput
	// Defines if the image name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringPtrInput
	// A list of tags for this rule.
	Tags pulumi.StringArrayInput
}

func (RuleContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleContainerArgs)(nil)).Elem()
}

type RuleContainerInput interface {
	pulumi.Input

	ToRuleContainerOutput() RuleContainerOutput
	ToRuleContainerOutputWithContext(ctx context.Context) RuleContainerOutput
}

func (*RuleContainer) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleContainer)(nil))
}

func (i *RuleContainer) ToRuleContainerOutput() RuleContainerOutput {
	return i.ToRuleContainerOutputWithContext(context.Background())
}

func (i *RuleContainer) ToRuleContainerOutputWithContext(ctx context.Context) RuleContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleContainerOutput)
}

func (i *RuleContainer) ToRuleContainerPtrOutput() RuleContainerPtrOutput {
	return i.ToRuleContainerPtrOutputWithContext(context.Background())
}

func (i *RuleContainer) ToRuleContainerPtrOutputWithContext(ctx context.Context) RuleContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleContainerPtrOutput)
}

type RuleContainerPtrInput interface {
	pulumi.Input

	ToRuleContainerPtrOutput() RuleContainerPtrOutput
	ToRuleContainerPtrOutputWithContext(ctx context.Context) RuleContainerPtrOutput
}

type ruleContainerPtrType RuleContainerArgs

func (*ruleContainerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleContainer)(nil))
}

func (i *ruleContainerPtrType) ToRuleContainerPtrOutput() RuleContainerPtrOutput {
	return i.ToRuleContainerPtrOutputWithContext(context.Background())
}

func (i *ruleContainerPtrType) ToRuleContainerPtrOutputWithContext(ctx context.Context) RuleContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleContainerPtrOutput)
}

// RuleContainerArrayInput is an input type that accepts RuleContainerArray and RuleContainerArrayOutput values.
// You can construct a concrete instance of `RuleContainerArrayInput` via:
//
//          RuleContainerArray{ RuleContainerArgs{...} }
type RuleContainerArrayInput interface {
	pulumi.Input

	ToRuleContainerArrayOutput() RuleContainerArrayOutput
	ToRuleContainerArrayOutputWithContext(context.Context) RuleContainerArrayOutput
}

type RuleContainerArray []RuleContainerInput

func (RuleContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleContainer)(nil)).Elem()
}

func (i RuleContainerArray) ToRuleContainerArrayOutput() RuleContainerArrayOutput {
	return i.ToRuleContainerArrayOutputWithContext(context.Background())
}

func (i RuleContainerArray) ToRuleContainerArrayOutputWithContext(ctx context.Context) RuleContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleContainerArrayOutput)
}

// RuleContainerMapInput is an input type that accepts RuleContainerMap and RuleContainerMapOutput values.
// You can construct a concrete instance of `RuleContainerMapInput` via:
//
//          RuleContainerMap{ "key": RuleContainerArgs{...} }
type RuleContainerMapInput interface {
	pulumi.Input

	ToRuleContainerMapOutput() RuleContainerMapOutput
	ToRuleContainerMapOutputWithContext(context.Context) RuleContainerMapOutput
}

type RuleContainerMap map[string]RuleContainerInput

func (RuleContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleContainer)(nil)).Elem()
}

func (i RuleContainerMap) ToRuleContainerMapOutput() RuleContainerMapOutput {
	return i.ToRuleContainerMapOutputWithContext(context.Background())
}

func (i RuleContainerMap) ToRuleContainerMapOutputWithContext(ctx context.Context) RuleContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleContainerMapOutput)
}

type RuleContainerOutput struct{ *pulumi.OutputState }

func (RuleContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleContainer)(nil))
}

func (o RuleContainerOutput) ToRuleContainerOutput() RuleContainerOutput {
	return o
}

func (o RuleContainerOutput) ToRuleContainerOutputWithContext(ctx context.Context) RuleContainerOutput {
	return o
}

func (o RuleContainerOutput) ToRuleContainerPtrOutput() RuleContainerPtrOutput {
	return o.ToRuleContainerPtrOutputWithContext(context.Background())
}

func (o RuleContainerOutput) ToRuleContainerPtrOutputWithContext(ctx context.Context) RuleContainerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleContainer) *RuleContainer {
		return &v
	}).(RuleContainerPtrOutput)
}

type RuleContainerPtrOutput struct{ *pulumi.OutputState }

func (RuleContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleContainer)(nil))
}

func (o RuleContainerPtrOutput) ToRuleContainerPtrOutput() RuleContainerPtrOutput {
	return o
}

func (o RuleContainerPtrOutput) ToRuleContainerPtrOutputWithContext(ctx context.Context) RuleContainerPtrOutput {
	return o
}

func (o RuleContainerPtrOutput) Elem() RuleContainerOutput {
	return o.ApplyT(func(v *RuleContainer) RuleContainer {
		if v != nil {
			return *v
		}
		var ret RuleContainer
		return ret
	}).(RuleContainerOutput)
}

type RuleContainerArrayOutput struct{ *pulumi.OutputState }

func (RuleContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleContainer)(nil))
}

func (o RuleContainerArrayOutput) ToRuleContainerArrayOutput() RuleContainerArrayOutput {
	return o
}

func (o RuleContainerArrayOutput) ToRuleContainerArrayOutputWithContext(ctx context.Context) RuleContainerArrayOutput {
	return o
}

func (o RuleContainerArrayOutput) Index(i pulumi.IntInput) RuleContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleContainer {
		return vs[0].([]RuleContainer)[vs[1].(int)]
	}).(RuleContainerOutput)
}

type RuleContainerMapOutput struct{ *pulumi.OutputState }

func (RuleContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RuleContainer)(nil))
}

func (o RuleContainerMapOutput) ToRuleContainerMapOutput() RuleContainerMapOutput {
	return o
}

func (o RuleContainerMapOutput) ToRuleContainerMapOutputWithContext(ctx context.Context) RuleContainerMapOutput {
	return o
}

func (o RuleContainerMapOutput) MapIndex(k pulumi.StringInput) RuleContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RuleContainer {
		return vs[0].(map[string]RuleContainer)[vs[1].(string)]
	}).(RuleContainerOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleContainerOutput{})
	pulumi.RegisterOutputType(RuleContainerPtrOutput{})
	pulumi.RegisterOutputType(RuleContainerArrayOutput{})
	pulumi.RegisterOutputType(RuleContainerMapOutput{})
}
