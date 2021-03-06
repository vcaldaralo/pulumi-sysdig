// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// PromQL Monitor alerts can be imported using the alert ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Monitor/alertPromql:AlertPromql example 12345
// ```
type AlertPromql struct {
	pulumi.CustomResourceState

	Capture AlertPromqlCapturePtrOutput `pulumi:"capture"`
	// Allows to define a custom notification title, prepend and append text.
	CustomNotification AlertPromqlCustomNotificationPtrOutput `pulumi:"customNotification"`
	// The description of Monitor alert.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Boolean that defines if the alert is enabled or not. Defaults to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the Monitor alert. It must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of notification channel IDs where an alert must be sent to once fired.
	NotificationChannels pulumi.IntArrayOutput `pulumi:"notificationChannels"`
	// PromQL-based metric expression to alert on. Example: `histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[5m]) > 0.15` or `predict_linear(sysdig_fs_free_bytes{fstype!~"tmpfs"}[1h], 24*3600) < 10000000000`.
	Promql pulumi.StringOutput `pulumi:"promql"`
	// Number of minutes for the alert to re-notify until the status is solved.
	RenotificationMinutes pulumi.IntPtrOutput    `pulumi:"renotificationMinutes"`
	Scope                 pulumi.StringPtrOutput `pulumi:"scope"`
	// Severity of the Monitor alert. It must be a value between 0 and 7,
	// with 0 being the most critical and 7 the less critical. Defaults to 4.
	Severity pulumi.IntPtrOutput `pulumi:"severity"`
	// Team ID that owns the alert.
	Team pulumi.IntOutput `pulumi:"team"`
	// Threshold of time for the status to stabilize until the alert is fired.
	TriggerAfterMinutes pulumi.IntOutput `pulumi:"triggerAfterMinutes"`
	// Current version of the resource in Sysdig Monitor.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAlertPromql registers a new resource with the given unique name, arguments, and options.
func NewAlertPromql(ctx *pulumi.Context,
	name string, args *AlertPromqlArgs, opts ...pulumi.ResourceOption) (*AlertPromql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Promql == nil {
		return nil, errors.New("invalid value for required argument 'Promql'")
	}
	if args.TriggerAfterMinutes == nil {
		return nil, errors.New("invalid value for required argument 'TriggerAfterMinutes'")
	}
	var resource AlertPromql
	err := ctx.RegisterResource("sysdig:Monitor/alertPromql:AlertPromql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertPromql gets an existing AlertPromql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertPromql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertPromqlState, opts ...pulumi.ResourceOption) (*AlertPromql, error) {
	var resource AlertPromql
	err := ctx.ReadResource("sysdig:Monitor/alertPromql:AlertPromql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertPromql resources.
type alertPromqlState struct {
	Capture *AlertPromqlCapture `pulumi:"capture"`
	// Allows to define a custom notification title, prepend and append text.
	CustomNotification *AlertPromqlCustomNotification `pulumi:"customNotification"`
	// The description of Monitor alert.
	Description *string `pulumi:"description"`
	// Boolean that defines if the alert is enabled or not. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Monitor alert. It must be unique.
	Name *string `pulumi:"name"`
	// List of notification channel IDs where an alert must be sent to once fired.
	NotificationChannels []int `pulumi:"notificationChannels"`
	// PromQL-based metric expression to alert on. Example: `histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[5m]) > 0.15` or `predict_linear(sysdig_fs_free_bytes{fstype!~"tmpfs"}[1h], 24*3600) < 10000000000`.
	Promql *string `pulumi:"promql"`
	// Number of minutes for the alert to re-notify until the status is solved.
	RenotificationMinutes *int    `pulumi:"renotificationMinutes"`
	Scope                 *string `pulumi:"scope"`
	// Severity of the Monitor alert. It must be a value between 0 and 7,
	// with 0 being the most critical and 7 the less critical. Defaults to 4.
	Severity *int `pulumi:"severity"`
	// Team ID that owns the alert.
	Team *int `pulumi:"team"`
	// Threshold of time for the status to stabilize until the alert is fired.
	TriggerAfterMinutes *int `pulumi:"triggerAfterMinutes"`
	// Current version of the resource in Sysdig Monitor.
	Version *int `pulumi:"version"`
}

type AlertPromqlState struct {
	Capture AlertPromqlCapturePtrInput
	// Allows to define a custom notification title, prepend and append text.
	CustomNotification AlertPromqlCustomNotificationPtrInput
	// The description of Monitor alert.
	Description pulumi.StringPtrInput
	// Boolean that defines if the alert is enabled or not. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// The name of the Monitor alert. It must be unique.
	Name pulumi.StringPtrInput
	// List of notification channel IDs where an alert must be sent to once fired.
	NotificationChannels pulumi.IntArrayInput
	// PromQL-based metric expression to alert on. Example: `histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[5m]) > 0.15` or `predict_linear(sysdig_fs_free_bytes{fstype!~"tmpfs"}[1h], 24*3600) < 10000000000`.
	Promql pulumi.StringPtrInput
	// Number of minutes for the alert to re-notify until the status is solved.
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	// Severity of the Monitor alert. It must be a value between 0 and 7,
	// with 0 being the most critical and 7 the less critical. Defaults to 4.
	Severity pulumi.IntPtrInput
	// Team ID that owns the alert.
	Team pulumi.IntPtrInput
	// Threshold of time for the status to stabilize until the alert is fired.
	TriggerAfterMinutes pulumi.IntPtrInput
	// Current version of the resource in Sysdig Monitor.
	Version pulumi.IntPtrInput
}

func (AlertPromqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPromqlState)(nil)).Elem()
}

type alertPromqlArgs struct {
	Capture *AlertPromqlCapture `pulumi:"capture"`
	// Allows to define a custom notification title, prepend and append text.
	CustomNotification *AlertPromqlCustomNotification `pulumi:"customNotification"`
	// The description of Monitor alert.
	Description *string `pulumi:"description"`
	// Boolean that defines if the alert is enabled or not. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Monitor alert. It must be unique.
	Name *string `pulumi:"name"`
	// List of notification channel IDs where an alert must be sent to once fired.
	NotificationChannels []int `pulumi:"notificationChannels"`
	// PromQL-based metric expression to alert on. Example: `histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[5m]) > 0.15` or `predict_linear(sysdig_fs_free_bytes{fstype!~"tmpfs"}[1h], 24*3600) < 10000000000`.
	Promql string `pulumi:"promql"`
	// Number of minutes for the alert to re-notify until the status is solved.
	RenotificationMinutes *int    `pulumi:"renotificationMinutes"`
	Scope                 *string `pulumi:"scope"`
	// Severity of the Monitor alert. It must be a value between 0 and 7,
	// with 0 being the most critical and 7 the less critical. Defaults to 4.
	Severity *int `pulumi:"severity"`
	// Threshold of time for the status to stabilize until the alert is fired.
	TriggerAfterMinutes int `pulumi:"triggerAfterMinutes"`
}

// The set of arguments for constructing a AlertPromql resource.
type AlertPromqlArgs struct {
	Capture AlertPromqlCapturePtrInput
	// Allows to define a custom notification title, prepend and append text.
	CustomNotification AlertPromqlCustomNotificationPtrInput
	// The description of Monitor alert.
	Description pulumi.StringPtrInput
	// Boolean that defines if the alert is enabled or not. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// The name of the Monitor alert. It must be unique.
	Name pulumi.StringPtrInput
	// List of notification channel IDs where an alert must be sent to once fired.
	NotificationChannels pulumi.IntArrayInput
	// PromQL-based metric expression to alert on. Example: `histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[5m]) > 0.15` or `predict_linear(sysdig_fs_free_bytes{fstype!~"tmpfs"}[1h], 24*3600) < 10000000000`.
	Promql pulumi.StringInput
	// Number of minutes for the alert to re-notify until the status is solved.
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	// Severity of the Monitor alert. It must be a value between 0 and 7,
	// with 0 being the most critical and 7 the less critical. Defaults to 4.
	Severity pulumi.IntPtrInput
	// Threshold of time for the status to stabilize until the alert is fired.
	TriggerAfterMinutes pulumi.IntInput
}

func (AlertPromqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPromqlArgs)(nil)).Elem()
}

type AlertPromqlInput interface {
	pulumi.Input

	ToAlertPromqlOutput() AlertPromqlOutput
	ToAlertPromqlOutputWithContext(ctx context.Context) AlertPromqlOutput
}

func (*AlertPromql) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertPromql)(nil))
}

func (i *AlertPromql) ToAlertPromqlOutput() AlertPromqlOutput {
	return i.ToAlertPromqlOutputWithContext(context.Background())
}

func (i *AlertPromql) ToAlertPromqlOutputWithContext(ctx context.Context) AlertPromqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPromqlOutput)
}

func (i *AlertPromql) ToAlertPromqlPtrOutput() AlertPromqlPtrOutput {
	return i.ToAlertPromqlPtrOutputWithContext(context.Background())
}

func (i *AlertPromql) ToAlertPromqlPtrOutputWithContext(ctx context.Context) AlertPromqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPromqlPtrOutput)
}

type AlertPromqlPtrInput interface {
	pulumi.Input

	ToAlertPromqlPtrOutput() AlertPromqlPtrOutput
	ToAlertPromqlPtrOutputWithContext(ctx context.Context) AlertPromqlPtrOutput
}

type alertPromqlPtrType AlertPromqlArgs

func (*alertPromqlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertPromql)(nil))
}

func (i *alertPromqlPtrType) ToAlertPromqlPtrOutput() AlertPromqlPtrOutput {
	return i.ToAlertPromqlPtrOutputWithContext(context.Background())
}

func (i *alertPromqlPtrType) ToAlertPromqlPtrOutputWithContext(ctx context.Context) AlertPromqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPromqlPtrOutput)
}

// AlertPromqlArrayInput is an input type that accepts AlertPromqlArray and AlertPromqlArrayOutput values.
// You can construct a concrete instance of `AlertPromqlArrayInput` via:
//
//          AlertPromqlArray{ AlertPromqlArgs{...} }
type AlertPromqlArrayInput interface {
	pulumi.Input

	ToAlertPromqlArrayOutput() AlertPromqlArrayOutput
	ToAlertPromqlArrayOutputWithContext(context.Context) AlertPromqlArrayOutput
}

type AlertPromqlArray []AlertPromqlInput

func (AlertPromqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertPromql)(nil)).Elem()
}

func (i AlertPromqlArray) ToAlertPromqlArrayOutput() AlertPromqlArrayOutput {
	return i.ToAlertPromqlArrayOutputWithContext(context.Background())
}

func (i AlertPromqlArray) ToAlertPromqlArrayOutputWithContext(ctx context.Context) AlertPromqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPromqlArrayOutput)
}

// AlertPromqlMapInput is an input type that accepts AlertPromqlMap and AlertPromqlMapOutput values.
// You can construct a concrete instance of `AlertPromqlMapInput` via:
//
//          AlertPromqlMap{ "key": AlertPromqlArgs{...} }
type AlertPromqlMapInput interface {
	pulumi.Input

	ToAlertPromqlMapOutput() AlertPromqlMapOutput
	ToAlertPromqlMapOutputWithContext(context.Context) AlertPromqlMapOutput
}

type AlertPromqlMap map[string]AlertPromqlInput

func (AlertPromqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertPromql)(nil)).Elem()
}

func (i AlertPromqlMap) ToAlertPromqlMapOutput() AlertPromqlMapOutput {
	return i.ToAlertPromqlMapOutputWithContext(context.Background())
}

func (i AlertPromqlMap) ToAlertPromqlMapOutputWithContext(ctx context.Context) AlertPromqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPromqlMapOutput)
}

type AlertPromqlOutput struct{ *pulumi.OutputState }

func (AlertPromqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertPromql)(nil))
}

func (o AlertPromqlOutput) ToAlertPromqlOutput() AlertPromqlOutput {
	return o
}

func (o AlertPromqlOutput) ToAlertPromqlOutputWithContext(ctx context.Context) AlertPromqlOutput {
	return o
}

func (o AlertPromqlOutput) ToAlertPromqlPtrOutput() AlertPromqlPtrOutput {
	return o.ToAlertPromqlPtrOutputWithContext(context.Background())
}

func (o AlertPromqlOutput) ToAlertPromqlPtrOutputWithContext(ctx context.Context) AlertPromqlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertPromql) *AlertPromql {
		return &v
	}).(AlertPromqlPtrOutput)
}

type AlertPromqlPtrOutput struct{ *pulumi.OutputState }

func (AlertPromqlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertPromql)(nil))
}

func (o AlertPromqlPtrOutput) ToAlertPromqlPtrOutput() AlertPromqlPtrOutput {
	return o
}

func (o AlertPromqlPtrOutput) ToAlertPromqlPtrOutputWithContext(ctx context.Context) AlertPromqlPtrOutput {
	return o
}

func (o AlertPromqlPtrOutput) Elem() AlertPromqlOutput {
	return o.ApplyT(func(v *AlertPromql) AlertPromql {
		if v != nil {
			return *v
		}
		var ret AlertPromql
		return ret
	}).(AlertPromqlOutput)
}

type AlertPromqlArrayOutput struct{ *pulumi.OutputState }

func (AlertPromqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertPromql)(nil))
}

func (o AlertPromqlArrayOutput) ToAlertPromqlArrayOutput() AlertPromqlArrayOutput {
	return o
}

func (o AlertPromqlArrayOutput) ToAlertPromqlArrayOutputWithContext(ctx context.Context) AlertPromqlArrayOutput {
	return o
}

func (o AlertPromqlArrayOutput) Index(i pulumi.IntInput) AlertPromqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertPromql {
		return vs[0].([]AlertPromql)[vs[1].(int)]
	}).(AlertPromqlOutput)
}

type AlertPromqlMapOutput struct{ *pulumi.OutputState }

func (AlertPromqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AlertPromql)(nil))
}

func (o AlertPromqlMapOutput) ToAlertPromqlMapOutput() AlertPromqlMapOutput {
	return o
}

func (o AlertPromqlMapOutput) ToAlertPromqlMapOutputWithContext(ctx context.Context) AlertPromqlMapOutput {
	return o
}

func (o AlertPromqlMapOutput) MapIndex(k pulumi.StringInput) AlertPromqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AlertPromql {
		return vs[0].(map[string]AlertPromql)[vs[1].(string)]
	}).(AlertPromqlOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertPromqlOutput{})
	pulumi.RegisterOutputType(AlertPromqlPtrOutput{})
	pulumi.RegisterOutputType(AlertPromqlArrayOutput{})
	pulumi.RegisterOutputType(AlertPromqlMapOutput{})
}
