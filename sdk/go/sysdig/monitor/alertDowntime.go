// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertDowntime struct {
	pulumi.CustomResourceState

	Capture               AlertDowntimeCapturePtrOutput            `pulumi:"capture"`
	CustomNotification    AlertDowntimeCustomNotificationPtrOutput `pulumi:"customNotification"`
	Description           pulumi.StringPtrOutput                   `pulumi:"description"`
	Enabled               pulumi.BoolPtrOutput                     `pulumi:"enabled"`
	EntitiesToMonitors    pulumi.StringArrayOutput                 `pulumi:"entitiesToMonitors"`
	Name                  pulumi.StringOutput                      `pulumi:"name"`
	NotificationChannels  pulumi.IntArrayOutput                    `pulumi:"notificationChannels"`
	RenotificationMinutes pulumi.IntPtrOutput                      `pulumi:"renotificationMinutes"`
	Scope                 pulumi.StringPtrOutput                   `pulumi:"scope"`
	Severity              pulumi.IntPtrOutput                      `pulumi:"severity"`
	Team                  pulumi.IntOutput                         `pulumi:"team"`
	TriggerAfterMinutes   pulumi.IntOutput                         `pulumi:"triggerAfterMinutes"`
	TriggerAfterPct       pulumi.IntPtrOutput                      `pulumi:"triggerAfterPct"`
	Version               pulumi.IntOutput                         `pulumi:"version"`
}

// NewAlertDowntime registers a new resource with the given unique name, arguments, and options.
func NewAlertDowntime(ctx *pulumi.Context,
	name string, args *AlertDowntimeArgs, opts ...pulumi.ResourceOption) (*AlertDowntime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntitiesToMonitors == nil {
		return nil, errors.New("invalid value for required argument 'EntitiesToMonitors'")
	}
	if args.TriggerAfterMinutes == nil {
		return nil, errors.New("invalid value for required argument 'TriggerAfterMinutes'")
	}
	var resource AlertDowntime
	err := ctx.RegisterResource("sysdig:Monitor/alertDowntime:AlertDowntime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertDowntime gets an existing AlertDowntime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertDowntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertDowntimeState, opts ...pulumi.ResourceOption) (*AlertDowntime, error) {
	var resource AlertDowntime
	err := ctx.ReadResource("sysdig:Monitor/alertDowntime:AlertDowntime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertDowntime resources.
type alertDowntimeState struct {
	Capture               *AlertDowntimeCapture            `pulumi:"capture"`
	CustomNotification    *AlertDowntimeCustomNotification `pulumi:"customNotification"`
	Description           *string                          `pulumi:"description"`
	Enabled               *bool                            `pulumi:"enabled"`
	EntitiesToMonitors    []string                         `pulumi:"entitiesToMonitors"`
	Name                  *string                          `pulumi:"name"`
	NotificationChannels  []int                            `pulumi:"notificationChannels"`
	RenotificationMinutes *int                             `pulumi:"renotificationMinutes"`
	Scope                 *string                          `pulumi:"scope"`
	Severity              *int                             `pulumi:"severity"`
	Team                  *int                             `pulumi:"team"`
	TriggerAfterMinutes   *int                             `pulumi:"triggerAfterMinutes"`
	TriggerAfterPct       *int                             `pulumi:"triggerAfterPct"`
	Version               *int                             `pulumi:"version"`
}

type AlertDowntimeState struct {
	Capture               AlertDowntimeCapturePtrInput
	CustomNotification    AlertDowntimeCustomNotificationPtrInput
	Description           pulumi.StringPtrInput
	Enabled               pulumi.BoolPtrInput
	EntitiesToMonitors    pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	NotificationChannels  pulumi.IntArrayInput
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	Severity              pulumi.IntPtrInput
	Team                  pulumi.IntPtrInput
	TriggerAfterMinutes   pulumi.IntPtrInput
	TriggerAfterPct       pulumi.IntPtrInput
	Version               pulumi.IntPtrInput
}

func (AlertDowntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertDowntimeState)(nil)).Elem()
}

type alertDowntimeArgs struct {
	Capture               *AlertDowntimeCapture            `pulumi:"capture"`
	CustomNotification    *AlertDowntimeCustomNotification `pulumi:"customNotification"`
	Description           *string                          `pulumi:"description"`
	Enabled               *bool                            `pulumi:"enabled"`
	EntitiesToMonitors    []string                         `pulumi:"entitiesToMonitors"`
	Name                  *string                          `pulumi:"name"`
	NotificationChannels  []int                            `pulumi:"notificationChannels"`
	RenotificationMinutes *int                             `pulumi:"renotificationMinutes"`
	Scope                 *string                          `pulumi:"scope"`
	Severity              *int                             `pulumi:"severity"`
	TriggerAfterMinutes   int                              `pulumi:"triggerAfterMinutes"`
	TriggerAfterPct       *int                             `pulumi:"triggerAfterPct"`
}

// The set of arguments for constructing a AlertDowntime resource.
type AlertDowntimeArgs struct {
	Capture               AlertDowntimeCapturePtrInput
	CustomNotification    AlertDowntimeCustomNotificationPtrInput
	Description           pulumi.StringPtrInput
	Enabled               pulumi.BoolPtrInput
	EntitiesToMonitors    pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	NotificationChannels  pulumi.IntArrayInput
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	Severity              pulumi.IntPtrInput
	TriggerAfterMinutes   pulumi.IntInput
	TriggerAfterPct       pulumi.IntPtrInput
}

func (AlertDowntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertDowntimeArgs)(nil)).Elem()
}

type AlertDowntimeInput interface {
	pulumi.Input

	ToAlertDowntimeOutput() AlertDowntimeOutput
	ToAlertDowntimeOutputWithContext(ctx context.Context) AlertDowntimeOutput
}

func (*AlertDowntime) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDowntime)(nil))
}

func (i *AlertDowntime) ToAlertDowntimeOutput() AlertDowntimeOutput {
	return i.ToAlertDowntimeOutputWithContext(context.Background())
}

func (i *AlertDowntime) ToAlertDowntimeOutputWithContext(ctx context.Context) AlertDowntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDowntimeOutput)
}

func (i *AlertDowntime) ToAlertDowntimePtrOutput() AlertDowntimePtrOutput {
	return i.ToAlertDowntimePtrOutputWithContext(context.Background())
}

func (i *AlertDowntime) ToAlertDowntimePtrOutputWithContext(ctx context.Context) AlertDowntimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDowntimePtrOutput)
}

type AlertDowntimePtrInput interface {
	pulumi.Input

	ToAlertDowntimePtrOutput() AlertDowntimePtrOutput
	ToAlertDowntimePtrOutputWithContext(ctx context.Context) AlertDowntimePtrOutput
}

type alertDowntimePtrType AlertDowntimeArgs

func (*alertDowntimePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDowntime)(nil))
}

func (i *alertDowntimePtrType) ToAlertDowntimePtrOutput() AlertDowntimePtrOutput {
	return i.ToAlertDowntimePtrOutputWithContext(context.Background())
}

func (i *alertDowntimePtrType) ToAlertDowntimePtrOutputWithContext(ctx context.Context) AlertDowntimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDowntimePtrOutput)
}

// AlertDowntimeArrayInput is an input type that accepts AlertDowntimeArray and AlertDowntimeArrayOutput values.
// You can construct a concrete instance of `AlertDowntimeArrayInput` via:
//
//          AlertDowntimeArray{ AlertDowntimeArgs{...} }
type AlertDowntimeArrayInput interface {
	pulumi.Input

	ToAlertDowntimeArrayOutput() AlertDowntimeArrayOutput
	ToAlertDowntimeArrayOutputWithContext(context.Context) AlertDowntimeArrayOutput
}

type AlertDowntimeArray []AlertDowntimeInput

func (AlertDowntimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertDowntime)(nil)).Elem()
}

func (i AlertDowntimeArray) ToAlertDowntimeArrayOutput() AlertDowntimeArrayOutput {
	return i.ToAlertDowntimeArrayOutputWithContext(context.Background())
}

func (i AlertDowntimeArray) ToAlertDowntimeArrayOutputWithContext(ctx context.Context) AlertDowntimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDowntimeArrayOutput)
}

// AlertDowntimeMapInput is an input type that accepts AlertDowntimeMap and AlertDowntimeMapOutput values.
// You can construct a concrete instance of `AlertDowntimeMapInput` via:
//
//          AlertDowntimeMap{ "key": AlertDowntimeArgs{...} }
type AlertDowntimeMapInput interface {
	pulumi.Input

	ToAlertDowntimeMapOutput() AlertDowntimeMapOutput
	ToAlertDowntimeMapOutputWithContext(context.Context) AlertDowntimeMapOutput
}

type AlertDowntimeMap map[string]AlertDowntimeInput

func (AlertDowntimeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertDowntime)(nil)).Elem()
}

func (i AlertDowntimeMap) ToAlertDowntimeMapOutput() AlertDowntimeMapOutput {
	return i.ToAlertDowntimeMapOutputWithContext(context.Background())
}

func (i AlertDowntimeMap) ToAlertDowntimeMapOutputWithContext(ctx context.Context) AlertDowntimeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDowntimeMapOutput)
}

type AlertDowntimeOutput struct{ *pulumi.OutputState }

func (AlertDowntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDowntime)(nil))
}

func (o AlertDowntimeOutput) ToAlertDowntimeOutput() AlertDowntimeOutput {
	return o
}

func (o AlertDowntimeOutput) ToAlertDowntimeOutputWithContext(ctx context.Context) AlertDowntimeOutput {
	return o
}

func (o AlertDowntimeOutput) ToAlertDowntimePtrOutput() AlertDowntimePtrOutput {
	return o.ToAlertDowntimePtrOutputWithContext(context.Background())
}

func (o AlertDowntimeOutput) ToAlertDowntimePtrOutputWithContext(ctx context.Context) AlertDowntimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertDowntime) *AlertDowntime {
		return &v
	}).(AlertDowntimePtrOutput)
}

type AlertDowntimePtrOutput struct{ *pulumi.OutputState }

func (AlertDowntimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDowntime)(nil))
}

func (o AlertDowntimePtrOutput) ToAlertDowntimePtrOutput() AlertDowntimePtrOutput {
	return o
}

func (o AlertDowntimePtrOutput) ToAlertDowntimePtrOutputWithContext(ctx context.Context) AlertDowntimePtrOutput {
	return o
}

func (o AlertDowntimePtrOutput) Elem() AlertDowntimeOutput {
	return o.ApplyT(func(v *AlertDowntime) AlertDowntime {
		if v != nil {
			return *v
		}
		var ret AlertDowntime
		return ret
	}).(AlertDowntimeOutput)
}

type AlertDowntimeArrayOutput struct{ *pulumi.OutputState }

func (AlertDowntimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertDowntime)(nil))
}

func (o AlertDowntimeArrayOutput) ToAlertDowntimeArrayOutput() AlertDowntimeArrayOutput {
	return o
}

func (o AlertDowntimeArrayOutput) ToAlertDowntimeArrayOutputWithContext(ctx context.Context) AlertDowntimeArrayOutput {
	return o
}

func (o AlertDowntimeArrayOutput) Index(i pulumi.IntInput) AlertDowntimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertDowntime {
		return vs[0].([]AlertDowntime)[vs[1].(int)]
	}).(AlertDowntimeOutput)
}

type AlertDowntimeMapOutput struct{ *pulumi.OutputState }

func (AlertDowntimeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AlertDowntime)(nil))
}

func (o AlertDowntimeMapOutput) ToAlertDowntimeMapOutput() AlertDowntimeMapOutput {
	return o
}

func (o AlertDowntimeMapOutput) ToAlertDowntimeMapOutputWithContext(ctx context.Context) AlertDowntimeMapOutput {
	return o
}

func (o AlertDowntimeMapOutput) MapIndex(k pulumi.StringInput) AlertDowntimeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AlertDowntime {
		return vs[0].(map[string]AlertDowntime)[vs[1].(string)]
	}).(AlertDowntimeOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertDowntimeOutput{})
	pulumi.RegisterOutputType(AlertDowntimePtrOutput{})
	pulumi.RegisterOutputType(AlertDowntimeArrayOutput{})
	pulumi.RegisterOutputType(AlertDowntimeMapOutput{})
}
