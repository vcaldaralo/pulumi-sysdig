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
// Secure process runtime rules can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Secure/ruleProcess:RuleProcess example 12345
// ```
type RuleProcess struct {
	pulumi.CustomResourceState

	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defines if the process name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrOutput `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of processes to match.
	Processes pulumi.StringArrayOutput `pulumi:"processes"`
	// A list of tags for this rule.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Current version of the resource in Sysdig Secure.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRuleProcess registers a new resource with the given unique name, arguments, and options.
func NewRuleProcess(ctx *pulumi.Context,
	name string, args *RuleProcessArgs, opts ...pulumi.ResourceOption) (*RuleProcess, error) {
	if args == nil {
		args = &RuleProcessArgs{}
	}

	var resource RuleProcess
	err := ctx.RegisterResource("sysdig:Secure/ruleProcess:RuleProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleProcess gets an existing RuleProcess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleProcessState, opts ...pulumi.ResourceOption) (*RuleProcess, error) {
	var resource RuleProcess
	err := ctx.ReadResource("sysdig:Secure/ruleProcess:RuleProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleProcess resources.
type ruleProcessState struct {
	// The description of Secure rule. By default is empty.
	Description *string `pulumi:"description"`
	// Defines if the process name matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name *string `pulumi:"name"`
	// List of processes to match.
	Processes []string `pulumi:"processes"`
	// A list of tags for this rule.
	Tags []string `pulumi:"tags"`
	// Current version of the resource in Sysdig Secure.
	Version *int `pulumi:"version"`
}

type RuleProcessState struct {
	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrInput
	// Defines if the process name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringPtrInput
	// List of processes to match.
	Processes pulumi.StringArrayInput
	// A list of tags for this rule.
	Tags pulumi.StringArrayInput
	// Current version of the resource in Sysdig Secure.
	Version pulumi.IntPtrInput
}

func (RuleProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleProcessState)(nil)).Elem()
}

type ruleProcessArgs struct {
	// The description of Secure rule. By default is empty.
	Description *string `pulumi:"description"`
	// Defines if the process name matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// The name of the Secure rule. It must be unique.
	Name *string `pulumi:"name"`
	// List of processes to match.
	Processes []string `pulumi:"processes"`
	// A list of tags for this rule.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a RuleProcess resource.
type RuleProcessArgs struct {
	// The description of Secure rule. By default is empty.
	Description pulumi.StringPtrInput
	// Defines if the process name matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput
	// The name of the Secure rule. It must be unique.
	Name pulumi.StringPtrInput
	// List of processes to match.
	Processes pulumi.StringArrayInput
	// A list of tags for this rule.
	Tags pulumi.StringArrayInput
}

func (RuleProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleProcessArgs)(nil)).Elem()
}

type RuleProcessInput interface {
	pulumi.Input

	ToRuleProcessOutput() RuleProcessOutput
	ToRuleProcessOutputWithContext(ctx context.Context) RuleProcessOutput
}

func (*RuleProcess) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleProcess)(nil))
}

func (i *RuleProcess) ToRuleProcessOutput() RuleProcessOutput {
	return i.ToRuleProcessOutputWithContext(context.Background())
}

func (i *RuleProcess) ToRuleProcessOutputWithContext(ctx context.Context) RuleProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleProcessOutput)
}

func (i *RuleProcess) ToRuleProcessPtrOutput() RuleProcessPtrOutput {
	return i.ToRuleProcessPtrOutputWithContext(context.Background())
}

func (i *RuleProcess) ToRuleProcessPtrOutputWithContext(ctx context.Context) RuleProcessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleProcessPtrOutput)
}

type RuleProcessPtrInput interface {
	pulumi.Input

	ToRuleProcessPtrOutput() RuleProcessPtrOutput
	ToRuleProcessPtrOutputWithContext(ctx context.Context) RuleProcessPtrOutput
}

type ruleProcessPtrType RuleProcessArgs

func (*ruleProcessPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleProcess)(nil))
}

func (i *ruleProcessPtrType) ToRuleProcessPtrOutput() RuleProcessPtrOutput {
	return i.ToRuleProcessPtrOutputWithContext(context.Background())
}

func (i *ruleProcessPtrType) ToRuleProcessPtrOutputWithContext(ctx context.Context) RuleProcessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleProcessPtrOutput)
}

// RuleProcessArrayInput is an input type that accepts RuleProcessArray and RuleProcessArrayOutput values.
// You can construct a concrete instance of `RuleProcessArrayInput` via:
//
//          RuleProcessArray{ RuleProcessArgs{...} }
type RuleProcessArrayInput interface {
	pulumi.Input

	ToRuleProcessArrayOutput() RuleProcessArrayOutput
	ToRuleProcessArrayOutputWithContext(context.Context) RuleProcessArrayOutput
}

type RuleProcessArray []RuleProcessInput

func (RuleProcessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleProcess)(nil)).Elem()
}

func (i RuleProcessArray) ToRuleProcessArrayOutput() RuleProcessArrayOutput {
	return i.ToRuleProcessArrayOutputWithContext(context.Background())
}

func (i RuleProcessArray) ToRuleProcessArrayOutputWithContext(ctx context.Context) RuleProcessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleProcessArrayOutput)
}

// RuleProcessMapInput is an input type that accepts RuleProcessMap and RuleProcessMapOutput values.
// You can construct a concrete instance of `RuleProcessMapInput` via:
//
//          RuleProcessMap{ "key": RuleProcessArgs{...} }
type RuleProcessMapInput interface {
	pulumi.Input

	ToRuleProcessMapOutput() RuleProcessMapOutput
	ToRuleProcessMapOutputWithContext(context.Context) RuleProcessMapOutput
}

type RuleProcessMap map[string]RuleProcessInput

func (RuleProcessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleProcess)(nil)).Elem()
}

func (i RuleProcessMap) ToRuleProcessMapOutput() RuleProcessMapOutput {
	return i.ToRuleProcessMapOutputWithContext(context.Background())
}

func (i RuleProcessMap) ToRuleProcessMapOutputWithContext(ctx context.Context) RuleProcessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleProcessMapOutput)
}

type RuleProcessOutput struct{ *pulumi.OutputState }

func (RuleProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleProcess)(nil))
}

func (o RuleProcessOutput) ToRuleProcessOutput() RuleProcessOutput {
	return o
}

func (o RuleProcessOutput) ToRuleProcessOutputWithContext(ctx context.Context) RuleProcessOutput {
	return o
}

func (o RuleProcessOutput) ToRuleProcessPtrOutput() RuleProcessPtrOutput {
	return o.ToRuleProcessPtrOutputWithContext(context.Background())
}

func (o RuleProcessOutput) ToRuleProcessPtrOutputWithContext(ctx context.Context) RuleProcessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleProcess) *RuleProcess {
		return &v
	}).(RuleProcessPtrOutput)
}

type RuleProcessPtrOutput struct{ *pulumi.OutputState }

func (RuleProcessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleProcess)(nil))
}

func (o RuleProcessPtrOutput) ToRuleProcessPtrOutput() RuleProcessPtrOutput {
	return o
}

func (o RuleProcessPtrOutput) ToRuleProcessPtrOutputWithContext(ctx context.Context) RuleProcessPtrOutput {
	return o
}

func (o RuleProcessPtrOutput) Elem() RuleProcessOutput {
	return o.ApplyT(func(v *RuleProcess) RuleProcess {
		if v != nil {
			return *v
		}
		var ret RuleProcess
		return ret
	}).(RuleProcessOutput)
}

type RuleProcessArrayOutput struct{ *pulumi.OutputState }

func (RuleProcessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleProcess)(nil))
}

func (o RuleProcessArrayOutput) ToRuleProcessArrayOutput() RuleProcessArrayOutput {
	return o
}

func (o RuleProcessArrayOutput) ToRuleProcessArrayOutputWithContext(ctx context.Context) RuleProcessArrayOutput {
	return o
}

func (o RuleProcessArrayOutput) Index(i pulumi.IntInput) RuleProcessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleProcess {
		return vs[0].([]RuleProcess)[vs[1].(int)]
	}).(RuleProcessOutput)
}

type RuleProcessMapOutput struct{ *pulumi.OutputState }

func (RuleProcessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RuleProcess)(nil))
}

func (o RuleProcessMapOutput) ToRuleProcessMapOutput() RuleProcessMapOutput {
	return o
}

func (o RuleProcessMapOutput) ToRuleProcessMapOutputWithContext(ctx context.Context) RuleProcessMapOutput {
	return o
}

func (o RuleProcessMapOutput) MapIndex(k pulumi.StringInput) RuleProcessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RuleProcess {
		return vs[0].(map[string]RuleProcess)[vs[1].(string)]
	}).(RuleProcessOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleProcessOutput{})
	pulumi.RegisterOutputType(RuleProcessPtrOutput{})
	pulumi.RegisterOutputType(RuleProcessArrayOutput{})
	pulumi.RegisterOutputType(RuleProcessMapOutput{})
}
