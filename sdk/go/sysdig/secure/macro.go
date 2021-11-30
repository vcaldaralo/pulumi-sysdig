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
// Secure macros can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Secure/macro:Macro example 12345
// ```
type Macro struct {
	pulumi.CustomResourceState

	// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
	// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
	// append macro called "foo" but not a second one. By default this is false.
	Append pulumi.BoolPtrOutput `pulumi:"append"`
	// Macro condition. It can contain lists or other macros.
	Condition pulumi.StringOutput `pulumi:"condition"`
	// The name of the macro. It must be unique if it's not in append mode.
	Name    pulumi.StringOutput `pulumi:"name"`
	Version pulumi.IntOutput    `pulumi:"version"`
}

// NewMacro registers a new resource with the given unique name, arguments, and options.
func NewMacro(ctx *pulumi.Context,
	name string, args *MacroArgs, opts ...pulumi.ResourceOption) (*Macro, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	var resource Macro
	err := ctx.RegisterResource("sysdig:Secure/macro:Macro", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMacro gets an existing Macro resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMacro(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MacroState, opts ...pulumi.ResourceOption) (*Macro, error) {
	var resource Macro
	err := ctx.ReadResource("sysdig:Secure/macro:Macro", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Macro resources.
type macroState struct {
	// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
	// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
	// append macro called "foo" but not a second one. By default this is false.
	Append *bool `pulumi:"append"`
	// Macro condition. It can contain lists or other macros.
	Condition *string `pulumi:"condition"`
	// The name of the macro. It must be unique if it's not in append mode.
	Name    *string `pulumi:"name"`
	Version *int    `pulumi:"version"`
}

type MacroState struct {
	// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
	// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
	// append macro called "foo" but not a second one. By default this is false.
	Append pulumi.BoolPtrInput
	// Macro condition. It can contain lists or other macros.
	Condition pulumi.StringPtrInput
	// The name of the macro. It must be unique if it's not in append mode.
	Name    pulumi.StringPtrInput
	Version pulumi.IntPtrInput
}

func (MacroState) ElementType() reflect.Type {
	return reflect.TypeOf((*macroState)(nil)).Elem()
}

type macroArgs struct {
	// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
	// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
	// append macro called "foo" but not a second one. By default this is false.
	Append *bool `pulumi:"append"`
	// Macro condition. It can contain lists or other macros.
	Condition string `pulumi:"condition"`
	// The name of the macro. It must be unique if it's not in append mode.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Macro resource.
type MacroArgs struct {
	// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
	// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
	// append macro called "foo" but not a second one. By default this is false.
	Append pulumi.BoolPtrInput
	// Macro condition. It can contain lists or other macros.
	Condition pulumi.StringInput
	// The name of the macro. It must be unique if it's not in append mode.
	Name pulumi.StringPtrInput
}

func (MacroArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*macroArgs)(nil)).Elem()
}

type MacroInput interface {
	pulumi.Input

	ToMacroOutput() MacroOutput
	ToMacroOutputWithContext(ctx context.Context) MacroOutput
}

func (*Macro) ElementType() reflect.Type {
	return reflect.TypeOf((*Macro)(nil))
}

func (i *Macro) ToMacroOutput() MacroOutput {
	return i.ToMacroOutputWithContext(context.Background())
}

func (i *Macro) ToMacroOutputWithContext(ctx context.Context) MacroOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroOutput)
}

func (i *Macro) ToMacroPtrOutput() MacroPtrOutput {
	return i.ToMacroPtrOutputWithContext(context.Background())
}

func (i *Macro) ToMacroPtrOutputWithContext(ctx context.Context) MacroPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroPtrOutput)
}

type MacroPtrInput interface {
	pulumi.Input

	ToMacroPtrOutput() MacroPtrOutput
	ToMacroPtrOutputWithContext(ctx context.Context) MacroPtrOutput
}

type macroPtrType MacroArgs

func (*macroPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Macro)(nil))
}

func (i *macroPtrType) ToMacroPtrOutput() MacroPtrOutput {
	return i.ToMacroPtrOutputWithContext(context.Background())
}

func (i *macroPtrType) ToMacroPtrOutputWithContext(ctx context.Context) MacroPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroPtrOutput)
}

// MacroArrayInput is an input type that accepts MacroArray and MacroArrayOutput values.
// You can construct a concrete instance of `MacroArrayInput` via:
//
//          MacroArray{ MacroArgs{...} }
type MacroArrayInput interface {
	pulumi.Input

	ToMacroArrayOutput() MacroArrayOutput
	ToMacroArrayOutputWithContext(context.Context) MacroArrayOutput
}

type MacroArray []MacroInput

func (MacroArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Macro)(nil)).Elem()
}

func (i MacroArray) ToMacroArrayOutput() MacroArrayOutput {
	return i.ToMacroArrayOutputWithContext(context.Background())
}

func (i MacroArray) ToMacroArrayOutputWithContext(ctx context.Context) MacroArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroArrayOutput)
}

// MacroMapInput is an input type that accepts MacroMap and MacroMapOutput values.
// You can construct a concrete instance of `MacroMapInput` via:
//
//          MacroMap{ "key": MacroArgs{...} }
type MacroMapInput interface {
	pulumi.Input

	ToMacroMapOutput() MacroMapOutput
	ToMacroMapOutputWithContext(context.Context) MacroMapOutput
}

type MacroMap map[string]MacroInput

func (MacroMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Macro)(nil)).Elem()
}

func (i MacroMap) ToMacroMapOutput() MacroMapOutput {
	return i.ToMacroMapOutputWithContext(context.Background())
}

func (i MacroMap) ToMacroMapOutputWithContext(ctx context.Context) MacroMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroMapOutput)
}

type MacroOutput struct{ *pulumi.OutputState }

func (MacroOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Macro)(nil))
}

func (o MacroOutput) ToMacroOutput() MacroOutput {
	return o
}

func (o MacroOutput) ToMacroOutputWithContext(ctx context.Context) MacroOutput {
	return o
}

func (o MacroOutput) ToMacroPtrOutput() MacroPtrOutput {
	return o.ToMacroPtrOutputWithContext(context.Background())
}

func (o MacroOutput) ToMacroPtrOutputWithContext(ctx context.Context) MacroPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Macro) *Macro {
		return &v
	}).(MacroPtrOutput)
}

type MacroPtrOutput struct{ *pulumi.OutputState }

func (MacroPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Macro)(nil))
}

func (o MacroPtrOutput) ToMacroPtrOutput() MacroPtrOutput {
	return o
}

func (o MacroPtrOutput) ToMacroPtrOutputWithContext(ctx context.Context) MacroPtrOutput {
	return o
}

func (o MacroPtrOutput) Elem() MacroOutput {
	return o.ApplyT(func(v *Macro) Macro {
		if v != nil {
			return *v
		}
		var ret Macro
		return ret
	}).(MacroOutput)
}

type MacroArrayOutput struct{ *pulumi.OutputState }

func (MacroArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Macro)(nil))
}

func (o MacroArrayOutput) ToMacroArrayOutput() MacroArrayOutput {
	return o
}

func (o MacroArrayOutput) ToMacroArrayOutputWithContext(ctx context.Context) MacroArrayOutput {
	return o
}

func (o MacroArrayOutput) Index(i pulumi.IntInput) MacroOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Macro {
		return vs[0].([]Macro)[vs[1].(int)]
	}).(MacroOutput)
}

type MacroMapOutput struct{ *pulumi.OutputState }

func (MacroMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Macro)(nil))
}

func (o MacroMapOutput) ToMacroMapOutput() MacroMapOutput {
	return o
}

func (o MacroMapOutput) ToMacroMapOutputWithContext(ctx context.Context) MacroMapOutput {
	return o
}

func (o MacroMapOutput) MapIndex(k pulumi.StringInput) MacroOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Macro {
		return vs[0].(map[string]Macro)[vs[1].(string)]
	}).(MacroOutput)
}

func init() {
	pulumi.RegisterOutputType(MacroOutput{})
	pulumi.RegisterOutputType(MacroPtrOutput{})
	pulumi.RegisterOutputType(MacroArrayOutput{})
	pulumi.RegisterOutputType(MacroMapOutput{})
}
