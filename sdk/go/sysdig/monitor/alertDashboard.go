// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertDashboard struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput         `pulumi:"description"`
	Name        pulumi.StringOutput            `pulumi:"name"`
	Panels      AlertDashboardPanelArrayOutput `pulumi:"panels"`
	Public      pulumi.BoolPtrOutput           `pulumi:"public"`
	PublicToken pulumi.StringOutput            `pulumi:"publicToken"`
	Scopes      AlertDashboardScopeArrayOutput `pulumi:"scopes"`
	Version     pulumi.IntOutput               `pulumi:"version"`
}

// NewAlertDashboard registers a new resource with the given unique name, arguments, and options.
func NewAlertDashboard(ctx *pulumi.Context,
	name string, args *AlertDashboardArgs, opts ...pulumi.ResourceOption) (*AlertDashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Panels == nil {
		return nil, errors.New("invalid value for required argument 'Panels'")
	}
	var resource AlertDashboard
	err := ctx.RegisterResource("sysdig:Monitor/alertDashboard:AlertDashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertDashboard gets an existing AlertDashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertDashboardState, opts ...pulumi.ResourceOption) (*AlertDashboard, error) {
	var resource AlertDashboard
	err := ctx.ReadResource("sysdig:Monitor/alertDashboard:AlertDashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertDashboard resources.
type alertDashboardState struct {
	Description *string               `pulumi:"description"`
	Name        *string               `pulumi:"name"`
	Panels      []AlertDashboardPanel `pulumi:"panels"`
	Public      *bool                 `pulumi:"public"`
	PublicToken *string               `pulumi:"publicToken"`
	Scopes      []AlertDashboardScope `pulumi:"scopes"`
	Version     *int                  `pulumi:"version"`
}

type AlertDashboardState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Panels      AlertDashboardPanelArrayInput
	Public      pulumi.BoolPtrInput
	PublicToken pulumi.StringPtrInput
	Scopes      AlertDashboardScopeArrayInput
	Version     pulumi.IntPtrInput
}

func (AlertDashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertDashboardState)(nil)).Elem()
}

type alertDashboardArgs struct {
	Description *string               `pulumi:"description"`
	Name        *string               `pulumi:"name"`
	Panels      []AlertDashboardPanel `pulumi:"panels"`
	Public      *bool                 `pulumi:"public"`
	Scopes      []AlertDashboardScope `pulumi:"scopes"`
}

// The set of arguments for constructing a AlertDashboard resource.
type AlertDashboardArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Panels      AlertDashboardPanelArrayInput
	Public      pulumi.BoolPtrInput
	Scopes      AlertDashboardScopeArrayInput
}

func (AlertDashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertDashboardArgs)(nil)).Elem()
}

type AlertDashboardInput interface {
	pulumi.Input

	ToAlertDashboardOutput() AlertDashboardOutput
	ToAlertDashboardOutputWithContext(ctx context.Context) AlertDashboardOutput
}

func (*AlertDashboard) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDashboard)(nil))
}

func (i *AlertDashboard) ToAlertDashboardOutput() AlertDashboardOutput {
	return i.ToAlertDashboardOutputWithContext(context.Background())
}

func (i *AlertDashboard) ToAlertDashboardOutputWithContext(ctx context.Context) AlertDashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDashboardOutput)
}

func (i *AlertDashboard) ToAlertDashboardPtrOutput() AlertDashboardPtrOutput {
	return i.ToAlertDashboardPtrOutputWithContext(context.Background())
}

func (i *AlertDashboard) ToAlertDashboardPtrOutputWithContext(ctx context.Context) AlertDashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDashboardPtrOutput)
}

type AlertDashboardPtrInput interface {
	pulumi.Input

	ToAlertDashboardPtrOutput() AlertDashboardPtrOutput
	ToAlertDashboardPtrOutputWithContext(ctx context.Context) AlertDashboardPtrOutput
}

type alertDashboardPtrType AlertDashboardArgs

func (*alertDashboardPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDashboard)(nil))
}

func (i *alertDashboardPtrType) ToAlertDashboardPtrOutput() AlertDashboardPtrOutput {
	return i.ToAlertDashboardPtrOutputWithContext(context.Background())
}

func (i *alertDashboardPtrType) ToAlertDashboardPtrOutputWithContext(ctx context.Context) AlertDashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDashboardPtrOutput)
}

// AlertDashboardArrayInput is an input type that accepts AlertDashboardArray and AlertDashboardArrayOutput values.
// You can construct a concrete instance of `AlertDashboardArrayInput` via:
//
//          AlertDashboardArray{ AlertDashboardArgs{...} }
type AlertDashboardArrayInput interface {
	pulumi.Input

	ToAlertDashboardArrayOutput() AlertDashboardArrayOutput
	ToAlertDashboardArrayOutputWithContext(context.Context) AlertDashboardArrayOutput
}

type AlertDashboardArray []AlertDashboardInput

func (AlertDashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertDashboard)(nil)).Elem()
}

func (i AlertDashboardArray) ToAlertDashboardArrayOutput() AlertDashboardArrayOutput {
	return i.ToAlertDashboardArrayOutputWithContext(context.Background())
}

func (i AlertDashboardArray) ToAlertDashboardArrayOutputWithContext(ctx context.Context) AlertDashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDashboardArrayOutput)
}

// AlertDashboardMapInput is an input type that accepts AlertDashboardMap and AlertDashboardMapOutput values.
// You can construct a concrete instance of `AlertDashboardMapInput` via:
//
//          AlertDashboardMap{ "key": AlertDashboardArgs{...} }
type AlertDashboardMapInput interface {
	pulumi.Input

	ToAlertDashboardMapOutput() AlertDashboardMapOutput
	ToAlertDashboardMapOutputWithContext(context.Context) AlertDashboardMapOutput
}

type AlertDashboardMap map[string]AlertDashboardInput

func (AlertDashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertDashboard)(nil)).Elem()
}

func (i AlertDashboardMap) ToAlertDashboardMapOutput() AlertDashboardMapOutput {
	return i.ToAlertDashboardMapOutputWithContext(context.Background())
}

func (i AlertDashboardMap) ToAlertDashboardMapOutputWithContext(ctx context.Context) AlertDashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDashboardMapOutput)
}

type AlertDashboardOutput struct{ *pulumi.OutputState }

func (AlertDashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDashboard)(nil))
}

func (o AlertDashboardOutput) ToAlertDashboardOutput() AlertDashboardOutput {
	return o
}

func (o AlertDashboardOutput) ToAlertDashboardOutputWithContext(ctx context.Context) AlertDashboardOutput {
	return o
}

func (o AlertDashboardOutput) ToAlertDashboardPtrOutput() AlertDashboardPtrOutput {
	return o.ToAlertDashboardPtrOutputWithContext(context.Background())
}

func (o AlertDashboardOutput) ToAlertDashboardPtrOutputWithContext(ctx context.Context) AlertDashboardPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertDashboard) *AlertDashboard {
		return &v
	}).(AlertDashboardPtrOutput)
}

type AlertDashboardPtrOutput struct{ *pulumi.OutputState }

func (AlertDashboardPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDashboard)(nil))
}

func (o AlertDashboardPtrOutput) ToAlertDashboardPtrOutput() AlertDashboardPtrOutput {
	return o
}

func (o AlertDashboardPtrOutput) ToAlertDashboardPtrOutputWithContext(ctx context.Context) AlertDashboardPtrOutput {
	return o
}

func (o AlertDashboardPtrOutput) Elem() AlertDashboardOutput {
	return o.ApplyT(func(v *AlertDashboard) AlertDashboard {
		if v != nil {
			return *v
		}
		var ret AlertDashboard
		return ret
	}).(AlertDashboardOutput)
}

type AlertDashboardArrayOutput struct{ *pulumi.OutputState }

func (AlertDashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertDashboard)(nil))
}

func (o AlertDashboardArrayOutput) ToAlertDashboardArrayOutput() AlertDashboardArrayOutput {
	return o
}

func (o AlertDashboardArrayOutput) ToAlertDashboardArrayOutputWithContext(ctx context.Context) AlertDashboardArrayOutput {
	return o
}

func (o AlertDashboardArrayOutput) Index(i pulumi.IntInput) AlertDashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertDashboard {
		return vs[0].([]AlertDashboard)[vs[1].(int)]
	}).(AlertDashboardOutput)
}

type AlertDashboardMapOutput struct{ *pulumi.OutputState }

func (AlertDashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AlertDashboard)(nil))
}

func (o AlertDashboardMapOutput) ToAlertDashboardMapOutput() AlertDashboardMapOutput {
	return o
}

func (o AlertDashboardMapOutput) ToAlertDashboardMapOutputWithContext(ctx context.Context) AlertDashboardMapOutput {
	return o
}

func (o AlertDashboardMapOutput) MapIndex(k pulumi.StringInput) AlertDashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AlertDashboard {
		return vs[0].(map[string]AlertDashboard)[vs[1].(string)]
	}).(AlertDashboardOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertDashboardOutput{})
	pulumi.RegisterOutputType(AlertDashboardPtrOutput{})
	pulumi.RegisterOutputType(AlertDashboardArrayOutput{})
	pulumi.RegisterOutputType(AlertDashboardMapOutput{})
}
