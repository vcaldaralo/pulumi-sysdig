// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAction struct {
	// Captures with Sysdig the stream of system calls:
	Captures []PolicyActionCapture `pulumi:"captures"`
	// The action applied to container when this Policy is
	// triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
	// no action will be applied at the container level.
	Container *string `pulumi:"container"`
}

// PolicyActionInput is an input type that accepts PolicyActionArgs and PolicyActionOutput values.
// You can construct a concrete instance of `PolicyActionInput` via:
//
//          PolicyActionArgs{...}
type PolicyActionInput interface {
	pulumi.Input

	ToPolicyActionOutput() PolicyActionOutput
	ToPolicyActionOutputWithContext(context.Context) PolicyActionOutput
}

type PolicyActionArgs struct {
	// Captures with Sysdig the stream of system calls:
	Captures PolicyActionCaptureArrayInput `pulumi:"captures"`
	// The action applied to container when this Policy is
	// triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
	// no action will be applied at the container level.
	Container pulumi.StringPtrInput `pulumi:"container"`
}

func (PolicyActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAction)(nil)).Elem()
}

func (i PolicyActionArgs) ToPolicyActionOutput() PolicyActionOutput {
	return i.ToPolicyActionOutputWithContext(context.Background())
}

func (i PolicyActionArgs) ToPolicyActionOutputWithContext(ctx context.Context) PolicyActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyActionOutput)
}

// PolicyActionArrayInput is an input type that accepts PolicyActionArray and PolicyActionArrayOutput values.
// You can construct a concrete instance of `PolicyActionArrayInput` via:
//
//          PolicyActionArray{ PolicyActionArgs{...} }
type PolicyActionArrayInput interface {
	pulumi.Input

	ToPolicyActionArrayOutput() PolicyActionArrayOutput
	ToPolicyActionArrayOutputWithContext(context.Context) PolicyActionArrayOutput
}

type PolicyActionArray []PolicyActionInput

func (PolicyActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAction)(nil)).Elem()
}

func (i PolicyActionArray) ToPolicyActionArrayOutput() PolicyActionArrayOutput {
	return i.ToPolicyActionArrayOutputWithContext(context.Background())
}

func (i PolicyActionArray) ToPolicyActionArrayOutputWithContext(ctx context.Context) PolicyActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyActionArrayOutput)
}

type PolicyActionOutput struct{ *pulumi.OutputState }

func (PolicyActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAction)(nil)).Elem()
}

func (o PolicyActionOutput) ToPolicyActionOutput() PolicyActionOutput {
	return o
}

func (o PolicyActionOutput) ToPolicyActionOutputWithContext(ctx context.Context) PolicyActionOutput {
	return o
}

// Captures with Sysdig the stream of system calls:
func (o PolicyActionOutput) Captures() PolicyActionCaptureArrayOutput {
	return o.ApplyT(func(v PolicyAction) []PolicyActionCapture { return v.Captures }).(PolicyActionCaptureArrayOutput)
}

// The action applied to container when this Policy is
// triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
// no action will be applied at the container level.
func (o PolicyActionOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAction) *string { return v.Container }).(pulumi.StringPtrOutput)
}

type PolicyActionArrayOutput struct{ *pulumi.OutputState }

func (PolicyActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAction)(nil)).Elem()
}

func (o PolicyActionArrayOutput) ToPolicyActionArrayOutput() PolicyActionArrayOutput {
	return o
}

func (o PolicyActionArrayOutput) ToPolicyActionArrayOutputWithContext(ctx context.Context) PolicyActionArrayOutput {
	return o
}

func (o PolicyActionArrayOutput) Index(i pulumi.IntInput) PolicyActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAction {
		return vs[0].([]PolicyAction)[vs[1].(int)]
	}).(PolicyActionOutput)
}

type PolicyActionCapture struct {
	// Captures the system calls for the amount
	// of seconds after the policy was triggered.
	SecondsAfterEvent int `pulumi:"secondsAfterEvent"`
	// Captures the system calls during the
	// amount of seconds before the policy was triggered.
	SecondsBeforeEvent int `pulumi:"secondsBeforeEvent"`
}

// PolicyActionCaptureInput is an input type that accepts PolicyActionCaptureArgs and PolicyActionCaptureOutput values.
// You can construct a concrete instance of `PolicyActionCaptureInput` via:
//
//          PolicyActionCaptureArgs{...}
type PolicyActionCaptureInput interface {
	pulumi.Input

	ToPolicyActionCaptureOutput() PolicyActionCaptureOutput
	ToPolicyActionCaptureOutputWithContext(context.Context) PolicyActionCaptureOutput
}

type PolicyActionCaptureArgs struct {
	// Captures the system calls for the amount
	// of seconds after the policy was triggered.
	SecondsAfterEvent pulumi.IntInput `pulumi:"secondsAfterEvent"`
	// Captures the system calls during the
	// amount of seconds before the policy was triggered.
	SecondsBeforeEvent pulumi.IntInput `pulumi:"secondsBeforeEvent"`
}

func (PolicyActionCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyActionCapture)(nil)).Elem()
}

func (i PolicyActionCaptureArgs) ToPolicyActionCaptureOutput() PolicyActionCaptureOutput {
	return i.ToPolicyActionCaptureOutputWithContext(context.Background())
}

func (i PolicyActionCaptureArgs) ToPolicyActionCaptureOutputWithContext(ctx context.Context) PolicyActionCaptureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyActionCaptureOutput)
}

// PolicyActionCaptureArrayInput is an input type that accepts PolicyActionCaptureArray and PolicyActionCaptureArrayOutput values.
// You can construct a concrete instance of `PolicyActionCaptureArrayInput` via:
//
//          PolicyActionCaptureArray{ PolicyActionCaptureArgs{...} }
type PolicyActionCaptureArrayInput interface {
	pulumi.Input

	ToPolicyActionCaptureArrayOutput() PolicyActionCaptureArrayOutput
	ToPolicyActionCaptureArrayOutputWithContext(context.Context) PolicyActionCaptureArrayOutput
}

type PolicyActionCaptureArray []PolicyActionCaptureInput

func (PolicyActionCaptureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyActionCapture)(nil)).Elem()
}

func (i PolicyActionCaptureArray) ToPolicyActionCaptureArrayOutput() PolicyActionCaptureArrayOutput {
	return i.ToPolicyActionCaptureArrayOutputWithContext(context.Background())
}

func (i PolicyActionCaptureArray) ToPolicyActionCaptureArrayOutputWithContext(ctx context.Context) PolicyActionCaptureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyActionCaptureArrayOutput)
}

type PolicyActionCaptureOutput struct{ *pulumi.OutputState }

func (PolicyActionCaptureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyActionCapture)(nil)).Elem()
}

func (o PolicyActionCaptureOutput) ToPolicyActionCaptureOutput() PolicyActionCaptureOutput {
	return o
}

func (o PolicyActionCaptureOutput) ToPolicyActionCaptureOutputWithContext(ctx context.Context) PolicyActionCaptureOutput {
	return o
}

// Captures the system calls for the amount
// of seconds after the policy was triggered.
func (o PolicyActionCaptureOutput) SecondsAfterEvent() pulumi.IntOutput {
	return o.ApplyT(func(v PolicyActionCapture) int { return v.SecondsAfterEvent }).(pulumi.IntOutput)
}

// Captures the system calls during the
// amount of seconds before the policy was triggered.
func (o PolicyActionCaptureOutput) SecondsBeforeEvent() pulumi.IntOutput {
	return o.ApplyT(func(v PolicyActionCapture) int { return v.SecondsBeforeEvent }).(pulumi.IntOutput)
}

type PolicyActionCaptureArrayOutput struct{ *pulumi.OutputState }

func (PolicyActionCaptureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyActionCapture)(nil)).Elem()
}

func (o PolicyActionCaptureArrayOutput) ToPolicyActionCaptureArrayOutput() PolicyActionCaptureArrayOutput {
	return o
}

func (o PolicyActionCaptureArrayOutput) ToPolicyActionCaptureArrayOutputWithContext(ctx context.Context) PolicyActionCaptureArrayOutput {
	return o
}

func (o PolicyActionCaptureArrayOutput) Index(i pulumi.IntInput) PolicyActionCaptureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyActionCapture {
		return vs[0].([]PolicyActionCapture)[vs[1].(int)]
	}).(PolicyActionCaptureOutput)
}

type RuleFalcoException struct {
	// Contains comparison operators that align 1-1 with the items in the fields property.
	Comps []string `pulumi:"comps"`
	// Contains one or more fields that will extract a value from the syscall/k8s_audit events.
	Fields []string `pulumi:"fields"`
	// The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
	Name string `pulumi:"name"`
	// Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
	Values string `pulumi:"values"`
}

// RuleFalcoExceptionInput is an input type that accepts RuleFalcoExceptionArgs and RuleFalcoExceptionOutput values.
// You can construct a concrete instance of `RuleFalcoExceptionInput` via:
//
//          RuleFalcoExceptionArgs{...}
type RuleFalcoExceptionInput interface {
	pulumi.Input

	ToRuleFalcoExceptionOutput() RuleFalcoExceptionOutput
	ToRuleFalcoExceptionOutputWithContext(context.Context) RuleFalcoExceptionOutput
}

type RuleFalcoExceptionArgs struct {
	// Contains comparison operators that align 1-1 with the items in the fields property.
	Comps pulumi.StringArrayInput `pulumi:"comps"`
	// Contains one or more fields that will extract a value from the syscall/k8s_audit events.
	Fields pulumi.StringArrayInput `pulumi:"fields"`
	// The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
	Name pulumi.StringInput `pulumi:"name"`
	// Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
	Values pulumi.StringInput `pulumi:"values"`
}

func (RuleFalcoExceptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFalcoException)(nil)).Elem()
}

func (i RuleFalcoExceptionArgs) ToRuleFalcoExceptionOutput() RuleFalcoExceptionOutput {
	return i.ToRuleFalcoExceptionOutputWithContext(context.Background())
}

func (i RuleFalcoExceptionArgs) ToRuleFalcoExceptionOutputWithContext(ctx context.Context) RuleFalcoExceptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFalcoExceptionOutput)
}

// RuleFalcoExceptionArrayInput is an input type that accepts RuleFalcoExceptionArray and RuleFalcoExceptionArrayOutput values.
// You can construct a concrete instance of `RuleFalcoExceptionArrayInput` via:
//
//          RuleFalcoExceptionArray{ RuleFalcoExceptionArgs{...} }
type RuleFalcoExceptionArrayInput interface {
	pulumi.Input

	ToRuleFalcoExceptionArrayOutput() RuleFalcoExceptionArrayOutput
	ToRuleFalcoExceptionArrayOutputWithContext(context.Context) RuleFalcoExceptionArrayOutput
}

type RuleFalcoExceptionArray []RuleFalcoExceptionInput

func (RuleFalcoExceptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFalcoException)(nil)).Elem()
}

func (i RuleFalcoExceptionArray) ToRuleFalcoExceptionArrayOutput() RuleFalcoExceptionArrayOutput {
	return i.ToRuleFalcoExceptionArrayOutputWithContext(context.Background())
}

func (i RuleFalcoExceptionArray) ToRuleFalcoExceptionArrayOutputWithContext(ctx context.Context) RuleFalcoExceptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFalcoExceptionArrayOutput)
}

type RuleFalcoExceptionOutput struct{ *pulumi.OutputState }

func (RuleFalcoExceptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFalcoException)(nil)).Elem()
}

func (o RuleFalcoExceptionOutput) ToRuleFalcoExceptionOutput() RuleFalcoExceptionOutput {
	return o
}

func (o RuleFalcoExceptionOutput) ToRuleFalcoExceptionOutputWithContext(ctx context.Context) RuleFalcoExceptionOutput {
	return o
}

// Contains comparison operators that align 1-1 with the items in the fields property.
func (o RuleFalcoExceptionOutput) Comps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleFalcoException) []string { return v.Comps }).(pulumi.StringArrayOutput)
}

// Contains one or more fields that will extract a value from the syscall/k8s_audit events.
func (o RuleFalcoExceptionOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleFalcoException) []string { return v.Fields }).(pulumi.StringArrayOutput)
}

// The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
func (o RuleFalcoExceptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RuleFalcoException) string { return v.Name }).(pulumi.StringOutput)
}

// Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
func (o RuleFalcoExceptionOutput) Values() pulumi.StringOutput {
	return o.ApplyT(func(v RuleFalcoException) string { return v.Values }).(pulumi.StringOutput)
}

type RuleFalcoExceptionArrayOutput struct{ *pulumi.OutputState }

func (RuleFalcoExceptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFalcoException)(nil)).Elem()
}

func (o RuleFalcoExceptionArrayOutput) ToRuleFalcoExceptionArrayOutput() RuleFalcoExceptionArrayOutput {
	return o
}

func (o RuleFalcoExceptionArrayOutput) ToRuleFalcoExceptionArrayOutputWithContext(ctx context.Context) RuleFalcoExceptionArrayOutput {
	return o
}

func (o RuleFalcoExceptionArrayOutput) Index(i pulumi.IntInput) RuleFalcoExceptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleFalcoException {
		return vs[0].([]RuleFalcoException)[vs[1].(int)]
	}).(RuleFalcoExceptionOutput)
}

type RuleFilesystemReadOnly struct {
	// Defines if the path matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// List of paths to match.
	Paths []string `pulumi:"paths"`
}

// RuleFilesystemReadOnlyInput is an input type that accepts RuleFilesystemReadOnlyArgs and RuleFilesystemReadOnlyOutput values.
// You can construct a concrete instance of `RuleFilesystemReadOnlyInput` via:
//
//          RuleFilesystemReadOnlyArgs{...}
type RuleFilesystemReadOnlyInput interface {
	pulumi.Input

	ToRuleFilesystemReadOnlyOutput() RuleFilesystemReadOnlyOutput
	ToRuleFilesystemReadOnlyOutputWithContext(context.Context) RuleFilesystemReadOnlyOutput
}

type RuleFilesystemReadOnlyArgs struct {
	// Defines if the path matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput `pulumi:"matching"`
	// List of paths to match.
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (RuleFilesystemReadOnlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFilesystemReadOnly)(nil)).Elem()
}

func (i RuleFilesystemReadOnlyArgs) ToRuleFilesystemReadOnlyOutput() RuleFilesystemReadOnlyOutput {
	return i.ToRuleFilesystemReadOnlyOutputWithContext(context.Background())
}

func (i RuleFilesystemReadOnlyArgs) ToRuleFilesystemReadOnlyOutputWithContext(ctx context.Context) RuleFilesystemReadOnlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFilesystemReadOnlyOutput)
}

// RuleFilesystemReadOnlyArrayInput is an input type that accepts RuleFilesystemReadOnlyArray and RuleFilesystemReadOnlyArrayOutput values.
// You can construct a concrete instance of `RuleFilesystemReadOnlyArrayInput` via:
//
//          RuleFilesystemReadOnlyArray{ RuleFilesystemReadOnlyArgs{...} }
type RuleFilesystemReadOnlyArrayInput interface {
	pulumi.Input

	ToRuleFilesystemReadOnlyArrayOutput() RuleFilesystemReadOnlyArrayOutput
	ToRuleFilesystemReadOnlyArrayOutputWithContext(context.Context) RuleFilesystemReadOnlyArrayOutput
}

type RuleFilesystemReadOnlyArray []RuleFilesystemReadOnlyInput

func (RuleFilesystemReadOnlyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFilesystemReadOnly)(nil)).Elem()
}

func (i RuleFilesystemReadOnlyArray) ToRuleFilesystemReadOnlyArrayOutput() RuleFilesystemReadOnlyArrayOutput {
	return i.ToRuleFilesystemReadOnlyArrayOutputWithContext(context.Background())
}

func (i RuleFilesystemReadOnlyArray) ToRuleFilesystemReadOnlyArrayOutputWithContext(ctx context.Context) RuleFilesystemReadOnlyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFilesystemReadOnlyArrayOutput)
}

type RuleFilesystemReadOnlyOutput struct{ *pulumi.OutputState }

func (RuleFilesystemReadOnlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFilesystemReadOnly)(nil)).Elem()
}

func (o RuleFilesystemReadOnlyOutput) ToRuleFilesystemReadOnlyOutput() RuleFilesystemReadOnlyOutput {
	return o
}

func (o RuleFilesystemReadOnlyOutput) ToRuleFilesystemReadOnlyOutputWithContext(ctx context.Context) RuleFilesystemReadOnlyOutput {
	return o
}

// Defines if the path matches or not with the provided list. Default is true.
func (o RuleFilesystemReadOnlyOutput) Matching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleFilesystemReadOnly) *bool { return v.Matching }).(pulumi.BoolPtrOutput)
}

// List of paths to match.
func (o RuleFilesystemReadOnlyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleFilesystemReadOnly) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type RuleFilesystemReadOnlyArrayOutput struct{ *pulumi.OutputState }

func (RuleFilesystemReadOnlyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFilesystemReadOnly)(nil)).Elem()
}

func (o RuleFilesystemReadOnlyArrayOutput) ToRuleFilesystemReadOnlyArrayOutput() RuleFilesystemReadOnlyArrayOutput {
	return o
}

func (o RuleFilesystemReadOnlyArrayOutput) ToRuleFilesystemReadOnlyArrayOutputWithContext(ctx context.Context) RuleFilesystemReadOnlyArrayOutput {
	return o
}

func (o RuleFilesystemReadOnlyArrayOutput) Index(i pulumi.IntInput) RuleFilesystemReadOnlyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleFilesystemReadOnly {
		return vs[0].([]RuleFilesystemReadOnly)[vs[1].(int)]
	}).(RuleFilesystemReadOnlyOutput)
}

type RuleFilesystemReadWrite struct {
	// Defines if the path matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// List of paths to match.
	Paths []string `pulumi:"paths"`
}

// RuleFilesystemReadWriteInput is an input type that accepts RuleFilesystemReadWriteArgs and RuleFilesystemReadWriteOutput values.
// You can construct a concrete instance of `RuleFilesystemReadWriteInput` via:
//
//          RuleFilesystemReadWriteArgs{...}
type RuleFilesystemReadWriteInput interface {
	pulumi.Input

	ToRuleFilesystemReadWriteOutput() RuleFilesystemReadWriteOutput
	ToRuleFilesystemReadWriteOutputWithContext(context.Context) RuleFilesystemReadWriteOutput
}

type RuleFilesystemReadWriteArgs struct {
	// Defines if the path matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput `pulumi:"matching"`
	// List of paths to match.
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (RuleFilesystemReadWriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFilesystemReadWrite)(nil)).Elem()
}

func (i RuleFilesystemReadWriteArgs) ToRuleFilesystemReadWriteOutput() RuleFilesystemReadWriteOutput {
	return i.ToRuleFilesystemReadWriteOutputWithContext(context.Background())
}

func (i RuleFilesystemReadWriteArgs) ToRuleFilesystemReadWriteOutputWithContext(ctx context.Context) RuleFilesystemReadWriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFilesystemReadWriteOutput)
}

// RuleFilesystemReadWriteArrayInput is an input type that accepts RuleFilesystemReadWriteArray and RuleFilesystemReadWriteArrayOutput values.
// You can construct a concrete instance of `RuleFilesystemReadWriteArrayInput` via:
//
//          RuleFilesystemReadWriteArray{ RuleFilesystemReadWriteArgs{...} }
type RuleFilesystemReadWriteArrayInput interface {
	pulumi.Input

	ToRuleFilesystemReadWriteArrayOutput() RuleFilesystemReadWriteArrayOutput
	ToRuleFilesystemReadWriteArrayOutputWithContext(context.Context) RuleFilesystemReadWriteArrayOutput
}

type RuleFilesystemReadWriteArray []RuleFilesystemReadWriteInput

func (RuleFilesystemReadWriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFilesystemReadWrite)(nil)).Elem()
}

func (i RuleFilesystemReadWriteArray) ToRuleFilesystemReadWriteArrayOutput() RuleFilesystemReadWriteArrayOutput {
	return i.ToRuleFilesystemReadWriteArrayOutputWithContext(context.Background())
}

func (i RuleFilesystemReadWriteArray) ToRuleFilesystemReadWriteArrayOutputWithContext(ctx context.Context) RuleFilesystemReadWriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleFilesystemReadWriteArrayOutput)
}

type RuleFilesystemReadWriteOutput struct{ *pulumi.OutputState }

func (RuleFilesystemReadWriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleFilesystemReadWrite)(nil)).Elem()
}

func (o RuleFilesystemReadWriteOutput) ToRuleFilesystemReadWriteOutput() RuleFilesystemReadWriteOutput {
	return o
}

func (o RuleFilesystemReadWriteOutput) ToRuleFilesystemReadWriteOutputWithContext(ctx context.Context) RuleFilesystemReadWriteOutput {
	return o
}

// Defines if the path matches or not with the provided list. Default is true.
func (o RuleFilesystemReadWriteOutput) Matching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleFilesystemReadWrite) *bool { return v.Matching }).(pulumi.BoolPtrOutput)
}

// List of paths to match.
func (o RuleFilesystemReadWriteOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleFilesystemReadWrite) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type RuleFilesystemReadWriteArrayOutput struct{ *pulumi.OutputState }

func (RuleFilesystemReadWriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleFilesystemReadWrite)(nil)).Elem()
}

func (o RuleFilesystemReadWriteArrayOutput) ToRuleFilesystemReadWriteArrayOutput() RuleFilesystemReadWriteArrayOutput {
	return o
}

func (o RuleFilesystemReadWriteArrayOutput) ToRuleFilesystemReadWriteArrayOutputWithContext(ctx context.Context) RuleFilesystemReadWriteArrayOutput {
	return o
}

func (o RuleFilesystemReadWriteArrayOutput) Index(i pulumi.IntInput) RuleFilesystemReadWriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleFilesystemReadWrite {
		return vs[0].([]RuleFilesystemReadWrite)[vs[1].(int)]
	}).(RuleFilesystemReadWriteOutput)
}

type RuleNetworkTcp struct {
	// Defines if the port matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// List of ports to match.
	Ports []int `pulumi:"ports"`
}

// RuleNetworkTcpInput is an input type that accepts RuleNetworkTcpArgs and RuleNetworkTcpOutput values.
// You can construct a concrete instance of `RuleNetworkTcpInput` via:
//
//          RuleNetworkTcpArgs{...}
type RuleNetworkTcpInput interface {
	pulumi.Input

	ToRuleNetworkTcpOutput() RuleNetworkTcpOutput
	ToRuleNetworkTcpOutputWithContext(context.Context) RuleNetworkTcpOutput
}

type RuleNetworkTcpArgs struct {
	// Defines if the port matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput `pulumi:"matching"`
	// List of ports to match.
	Ports pulumi.IntArrayInput `pulumi:"ports"`
}

func (RuleNetworkTcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleNetworkTcp)(nil)).Elem()
}

func (i RuleNetworkTcpArgs) ToRuleNetworkTcpOutput() RuleNetworkTcpOutput {
	return i.ToRuleNetworkTcpOutputWithContext(context.Background())
}

func (i RuleNetworkTcpArgs) ToRuleNetworkTcpOutputWithContext(ctx context.Context) RuleNetworkTcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleNetworkTcpOutput)
}

// RuleNetworkTcpArrayInput is an input type that accepts RuleNetworkTcpArray and RuleNetworkTcpArrayOutput values.
// You can construct a concrete instance of `RuleNetworkTcpArrayInput` via:
//
//          RuleNetworkTcpArray{ RuleNetworkTcpArgs{...} }
type RuleNetworkTcpArrayInput interface {
	pulumi.Input

	ToRuleNetworkTcpArrayOutput() RuleNetworkTcpArrayOutput
	ToRuleNetworkTcpArrayOutputWithContext(context.Context) RuleNetworkTcpArrayOutput
}

type RuleNetworkTcpArray []RuleNetworkTcpInput

func (RuleNetworkTcpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleNetworkTcp)(nil)).Elem()
}

func (i RuleNetworkTcpArray) ToRuleNetworkTcpArrayOutput() RuleNetworkTcpArrayOutput {
	return i.ToRuleNetworkTcpArrayOutputWithContext(context.Background())
}

func (i RuleNetworkTcpArray) ToRuleNetworkTcpArrayOutputWithContext(ctx context.Context) RuleNetworkTcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleNetworkTcpArrayOutput)
}

type RuleNetworkTcpOutput struct{ *pulumi.OutputState }

func (RuleNetworkTcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleNetworkTcp)(nil)).Elem()
}

func (o RuleNetworkTcpOutput) ToRuleNetworkTcpOutput() RuleNetworkTcpOutput {
	return o
}

func (o RuleNetworkTcpOutput) ToRuleNetworkTcpOutputWithContext(ctx context.Context) RuleNetworkTcpOutput {
	return o
}

// Defines if the port matches or not with the provided list. Default is true.
func (o RuleNetworkTcpOutput) Matching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleNetworkTcp) *bool { return v.Matching }).(pulumi.BoolPtrOutput)
}

// List of ports to match.
func (o RuleNetworkTcpOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RuleNetworkTcp) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

type RuleNetworkTcpArrayOutput struct{ *pulumi.OutputState }

func (RuleNetworkTcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleNetworkTcp)(nil)).Elem()
}

func (o RuleNetworkTcpArrayOutput) ToRuleNetworkTcpArrayOutput() RuleNetworkTcpArrayOutput {
	return o
}

func (o RuleNetworkTcpArrayOutput) ToRuleNetworkTcpArrayOutputWithContext(ctx context.Context) RuleNetworkTcpArrayOutput {
	return o
}

func (o RuleNetworkTcpArrayOutput) Index(i pulumi.IntInput) RuleNetworkTcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleNetworkTcp {
		return vs[0].([]RuleNetworkTcp)[vs[1].(int)]
	}).(RuleNetworkTcpOutput)
}

type RuleNetworkUdp struct {
	// Defines if the port matches or not with the provided list. Default is true.
	Matching *bool `pulumi:"matching"`
	// List of ports to match.
	Ports []int `pulumi:"ports"`
}

// RuleNetworkUdpInput is an input type that accepts RuleNetworkUdpArgs and RuleNetworkUdpOutput values.
// You can construct a concrete instance of `RuleNetworkUdpInput` via:
//
//          RuleNetworkUdpArgs{...}
type RuleNetworkUdpInput interface {
	pulumi.Input

	ToRuleNetworkUdpOutput() RuleNetworkUdpOutput
	ToRuleNetworkUdpOutputWithContext(context.Context) RuleNetworkUdpOutput
}

type RuleNetworkUdpArgs struct {
	// Defines if the port matches or not with the provided list. Default is true.
	Matching pulumi.BoolPtrInput `pulumi:"matching"`
	// List of ports to match.
	Ports pulumi.IntArrayInput `pulumi:"ports"`
}

func (RuleNetworkUdpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleNetworkUdp)(nil)).Elem()
}

func (i RuleNetworkUdpArgs) ToRuleNetworkUdpOutput() RuleNetworkUdpOutput {
	return i.ToRuleNetworkUdpOutputWithContext(context.Background())
}

func (i RuleNetworkUdpArgs) ToRuleNetworkUdpOutputWithContext(ctx context.Context) RuleNetworkUdpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleNetworkUdpOutput)
}

// RuleNetworkUdpArrayInput is an input type that accepts RuleNetworkUdpArray and RuleNetworkUdpArrayOutput values.
// You can construct a concrete instance of `RuleNetworkUdpArrayInput` via:
//
//          RuleNetworkUdpArray{ RuleNetworkUdpArgs{...} }
type RuleNetworkUdpArrayInput interface {
	pulumi.Input

	ToRuleNetworkUdpArrayOutput() RuleNetworkUdpArrayOutput
	ToRuleNetworkUdpArrayOutputWithContext(context.Context) RuleNetworkUdpArrayOutput
}

type RuleNetworkUdpArray []RuleNetworkUdpInput

func (RuleNetworkUdpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleNetworkUdp)(nil)).Elem()
}

func (i RuleNetworkUdpArray) ToRuleNetworkUdpArrayOutput() RuleNetworkUdpArrayOutput {
	return i.ToRuleNetworkUdpArrayOutputWithContext(context.Background())
}

func (i RuleNetworkUdpArray) ToRuleNetworkUdpArrayOutputWithContext(ctx context.Context) RuleNetworkUdpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleNetworkUdpArrayOutput)
}

type RuleNetworkUdpOutput struct{ *pulumi.OutputState }

func (RuleNetworkUdpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleNetworkUdp)(nil)).Elem()
}

func (o RuleNetworkUdpOutput) ToRuleNetworkUdpOutput() RuleNetworkUdpOutput {
	return o
}

func (o RuleNetworkUdpOutput) ToRuleNetworkUdpOutputWithContext(ctx context.Context) RuleNetworkUdpOutput {
	return o
}

// Defines if the port matches or not with the provided list. Default is true.
func (o RuleNetworkUdpOutput) Matching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleNetworkUdp) *bool { return v.Matching }).(pulumi.BoolPtrOutput)
}

// List of ports to match.
func (o RuleNetworkUdpOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RuleNetworkUdp) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

type RuleNetworkUdpArrayOutput struct{ *pulumi.OutputState }

func (RuleNetworkUdpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleNetworkUdp)(nil)).Elem()
}

func (o RuleNetworkUdpArrayOutput) ToRuleNetworkUdpArrayOutput() RuleNetworkUdpArrayOutput {
	return o
}

func (o RuleNetworkUdpArrayOutput) ToRuleNetworkUdpArrayOutputWithContext(ctx context.Context) RuleNetworkUdpArrayOutput {
	return o
}

func (o RuleNetworkUdpArrayOutput) Index(i pulumi.IntInput) RuleNetworkUdpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleNetworkUdp {
		return vs[0].([]RuleNetworkUdp)[vs[1].(int)]
	}).(RuleNetworkUdpOutput)
}

type TeamUserRole struct {
	// The email of the user in the group.
	Email string `pulumi:"email"`
	// The role for the user in this group.
	// Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
	// Default: ROLE_TEAM_STANDARD.
	Role *string `pulumi:"role"`
}

// TeamUserRoleInput is an input type that accepts TeamUserRoleArgs and TeamUserRoleOutput values.
// You can construct a concrete instance of `TeamUserRoleInput` via:
//
//          TeamUserRoleArgs{...}
type TeamUserRoleInput interface {
	pulumi.Input

	ToTeamUserRoleOutput() TeamUserRoleOutput
	ToTeamUserRoleOutputWithContext(context.Context) TeamUserRoleOutput
}

type TeamUserRoleArgs struct {
	// The email of the user in the group.
	Email pulumi.StringInput `pulumi:"email"`
	// The role for the user in this group.
	// Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
	// Default: ROLE_TEAM_STANDARD.
	Role pulumi.StringPtrInput `pulumi:"role"`
}

func (TeamUserRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamUserRole)(nil)).Elem()
}

func (i TeamUserRoleArgs) ToTeamUserRoleOutput() TeamUserRoleOutput {
	return i.ToTeamUserRoleOutputWithContext(context.Background())
}

func (i TeamUserRoleArgs) ToTeamUserRoleOutputWithContext(ctx context.Context) TeamUserRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamUserRoleOutput)
}

// TeamUserRoleArrayInput is an input type that accepts TeamUserRoleArray and TeamUserRoleArrayOutput values.
// You can construct a concrete instance of `TeamUserRoleArrayInput` via:
//
//          TeamUserRoleArray{ TeamUserRoleArgs{...} }
type TeamUserRoleArrayInput interface {
	pulumi.Input

	ToTeamUserRoleArrayOutput() TeamUserRoleArrayOutput
	ToTeamUserRoleArrayOutputWithContext(context.Context) TeamUserRoleArrayOutput
}

type TeamUserRoleArray []TeamUserRoleInput

func (TeamUserRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TeamUserRole)(nil)).Elem()
}

func (i TeamUserRoleArray) ToTeamUserRoleArrayOutput() TeamUserRoleArrayOutput {
	return i.ToTeamUserRoleArrayOutputWithContext(context.Background())
}

func (i TeamUserRoleArray) ToTeamUserRoleArrayOutputWithContext(ctx context.Context) TeamUserRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamUserRoleArrayOutput)
}

type TeamUserRoleOutput struct{ *pulumi.OutputState }

func (TeamUserRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamUserRole)(nil)).Elem()
}

func (o TeamUserRoleOutput) ToTeamUserRoleOutput() TeamUserRoleOutput {
	return o
}

func (o TeamUserRoleOutput) ToTeamUserRoleOutputWithContext(ctx context.Context) TeamUserRoleOutput {
	return o
}

// The email of the user in the group.
func (o TeamUserRoleOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v TeamUserRole) string { return v.Email }).(pulumi.StringOutput)
}

// The role for the user in this group.
// Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
// Default: ROLE_TEAM_STANDARD.
func (o TeamUserRoleOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TeamUserRole) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type TeamUserRoleArrayOutput struct{ *pulumi.OutputState }

func (TeamUserRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TeamUserRole)(nil)).Elem()
}

func (o TeamUserRoleArrayOutput) ToTeamUserRoleArrayOutput() TeamUserRoleArrayOutput {
	return o
}

func (o TeamUserRoleArrayOutput) ToTeamUserRoleArrayOutputWithContext(ctx context.Context) TeamUserRoleArrayOutput {
	return o
}

func (o TeamUserRoleArrayOutput) Index(i pulumi.IntInput) TeamUserRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TeamUserRole {
		return vs[0].([]TeamUserRole)[vs[1].(int)]
	}).(TeamUserRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyActionInput)(nil)).Elem(), PolicyActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyActionArrayInput)(nil)).Elem(), PolicyActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyActionCaptureInput)(nil)).Elem(), PolicyActionCaptureArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyActionCaptureArrayInput)(nil)).Elem(), PolicyActionCaptureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFalcoExceptionInput)(nil)).Elem(), RuleFalcoExceptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFalcoExceptionArrayInput)(nil)).Elem(), RuleFalcoExceptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFilesystemReadOnlyInput)(nil)).Elem(), RuleFilesystemReadOnlyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFilesystemReadOnlyArrayInput)(nil)).Elem(), RuleFilesystemReadOnlyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFilesystemReadWriteInput)(nil)).Elem(), RuleFilesystemReadWriteArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleFilesystemReadWriteArrayInput)(nil)).Elem(), RuleFilesystemReadWriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleNetworkTcpInput)(nil)).Elem(), RuleNetworkTcpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleNetworkTcpArrayInput)(nil)).Elem(), RuleNetworkTcpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleNetworkUdpInput)(nil)).Elem(), RuleNetworkUdpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleNetworkUdpArrayInput)(nil)).Elem(), RuleNetworkUdpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamUserRoleInput)(nil)).Elem(), TeamUserRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamUserRoleArrayInput)(nil)).Elem(), TeamUserRoleArray{})
	pulumi.RegisterOutputType(PolicyActionOutput{})
	pulumi.RegisterOutputType(PolicyActionArrayOutput{})
	pulumi.RegisterOutputType(PolicyActionCaptureOutput{})
	pulumi.RegisterOutputType(PolicyActionCaptureArrayOutput{})
	pulumi.RegisterOutputType(RuleFalcoExceptionOutput{})
	pulumi.RegisterOutputType(RuleFalcoExceptionArrayOutput{})
	pulumi.RegisterOutputType(RuleFilesystemReadOnlyOutput{})
	pulumi.RegisterOutputType(RuleFilesystemReadOnlyArrayOutput{})
	pulumi.RegisterOutputType(RuleFilesystemReadWriteOutput{})
	pulumi.RegisterOutputType(RuleFilesystemReadWriteArrayOutput{})
	pulumi.RegisterOutputType(RuleNetworkTcpOutput{})
	pulumi.RegisterOutputType(RuleNetworkTcpArrayOutput{})
	pulumi.RegisterOutputType(RuleNetworkUdpOutput{})
	pulumi.RegisterOutputType(RuleNetworkUdpArrayOutput{})
	pulumi.RegisterOutputType(TeamUserRoleOutput{})
	pulumi.RegisterOutputType(TeamUserRoleArrayOutput{})
}
