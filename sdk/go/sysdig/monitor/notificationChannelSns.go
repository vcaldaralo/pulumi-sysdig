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
// Amazon SNS notification channels for Monitor can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Monitor/notificationChannelSns:NotificationChannelSns example 12345
// ```
type NotificationChannelSns struct {
	pulumi.CustomResourceState

	// If false, the channel will not emit notifications. Default is true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the Notification Channel. Must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk pulumi.BoolPtrOutput `pulumi:"notifyWhenOk"`
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved pulumi.BoolPtrOutput `pulumi:"notifyWhenResolved"`
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification pulumi.BoolPtrOutput `pulumi:"sendTestNotification"`
	// List of ARNs from the SNS topics.
	Topics pulumi.StringArrayOutput `pulumi:"topics"`
	// (Computed) The current version of the Notification Channel.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNotificationChannelSns registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannelSns(ctx *pulumi.Context,
	name string, args *NotificationChannelSnsArgs, opts ...pulumi.ResourceOption) (*NotificationChannelSns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Topics == nil {
		return nil, errors.New("invalid value for required argument 'Topics'")
	}
	var resource NotificationChannelSns
	err := ctx.RegisterResource("sysdig:Monitor/notificationChannelSns:NotificationChannelSns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationChannelSns gets an existing NotificationChannelSns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannelSns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelSnsState, opts ...pulumi.ResourceOption) (*NotificationChannelSns, error) {
	var resource NotificationChannelSns
	err := ctx.ReadResource("sysdig:Monitor/notificationChannelSns:NotificationChannelSns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationChannelSns resources.
type notificationChannelSnsState struct {
	// If false, the channel will not emit notifications. Default is true.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Notification Channel. Must be unique.
	Name *string `pulumi:"name"`
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk *bool `pulumi:"notifyWhenOk"`
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved *bool `pulumi:"notifyWhenResolved"`
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification *bool `pulumi:"sendTestNotification"`
	// List of ARNs from the SNS topics.
	Topics []string `pulumi:"topics"`
	// (Computed) The current version of the Notification Channel.
	Version *int `pulumi:"version"`
}

type NotificationChannelSnsState struct {
	// If false, the channel will not emit notifications. Default is true.
	Enabled pulumi.BoolPtrInput
	// The name of the Notification Channel. Must be unique.
	Name pulumi.StringPtrInput
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk pulumi.BoolPtrInput
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved pulumi.BoolPtrInput
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification pulumi.BoolPtrInput
	// List of ARNs from the SNS topics.
	Topics pulumi.StringArrayInput
	// (Computed) The current version of the Notification Channel.
	Version pulumi.IntPtrInput
}

func (NotificationChannelSnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelSnsState)(nil)).Elem()
}

type notificationChannelSnsArgs struct {
	// If false, the channel will not emit notifications. Default is true.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Notification Channel. Must be unique.
	Name *string `pulumi:"name"`
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk *bool `pulumi:"notifyWhenOk"`
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved *bool `pulumi:"notifyWhenResolved"`
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification *bool `pulumi:"sendTestNotification"`
	// List of ARNs from the SNS topics.
	Topics []string `pulumi:"topics"`
}

// The set of arguments for constructing a NotificationChannelSns resource.
type NotificationChannelSnsArgs struct {
	// If false, the channel will not emit notifications. Default is true.
	Enabled pulumi.BoolPtrInput
	// The name of the Notification Channel. Must be unique.
	Name pulumi.StringPtrInput
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk pulumi.BoolPtrInput
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved pulumi.BoolPtrInput
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification pulumi.BoolPtrInput
	// List of ARNs from the SNS topics.
	Topics pulumi.StringArrayInput
}

func (NotificationChannelSnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelSnsArgs)(nil)).Elem()
}

type NotificationChannelSnsInput interface {
	pulumi.Input

	ToNotificationChannelSnsOutput() NotificationChannelSnsOutput
	ToNotificationChannelSnsOutputWithContext(ctx context.Context) NotificationChannelSnsOutput
}

func (*NotificationChannelSns) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSns)(nil))
}

func (i *NotificationChannelSns) ToNotificationChannelSnsOutput() NotificationChannelSnsOutput {
	return i.ToNotificationChannelSnsOutputWithContext(context.Background())
}

func (i *NotificationChannelSns) ToNotificationChannelSnsOutputWithContext(ctx context.Context) NotificationChannelSnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsOutput)
}

func (i *NotificationChannelSns) ToNotificationChannelSnsPtrOutput() NotificationChannelSnsPtrOutput {
	return i.ToNotificationChannelSnsPtrOutputWithContext(context.Background())
}

func (i *NotificationChannelSns) ToNotificationChannelSnsPtrOutputWithContext(ctx context.Context) NotificationChannelSnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsPtrOutput)
}

type NotificationChannelSnsPtrInput interface {
	pulumi.Input

	ToNotificationChannelSnsPtrOutput() NotificationChannelSnsPtrOutput
	ToNotificationChannelSnsPtrOutputWithContext(ctx context.Context) NotificationChannelSnsPtrOutput
}

type notificationChannelSnsPtrType NotificationChannelSnsArgs

func (*notificationChannelSnsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSns)(nil))
}

func (i *notificationChannelSnsPtrType) ToNotificationChannelSnsPtrOutput() NotificationChannelSnsPtrOutput {
	return i.ToNotificationChannelSnsPtrOutputWithContext(context.Background())
}

func (i *notificationChannelSnsPtrType) ToNotificationChannelSnsPtrOutputWithContext(ctx context.Context) NotificationChannelSnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsPtrOutput)
}

// NotificationChannelSnsArrayInput is an input type that accepts NotificationChannelSnsArray and NotificationChannelSnsArrayOutput values.
// You can construct a concrete instance of `NotificationChannelSnsArrayInput` via:
//
//          NotificationChannelSnsArray{ NotificationChannelSnsArgs{...} }
type NotificationChannelSnsArrayInput interface {
	pulumi.Input

	ToNotificationChannelSnsArrayOutput() NotificationChannelSnsArrayOutput
	ToNotificationChannelSnsArrayOutputWithContext(context.Context) NotificationChannelSnsArrayOutput
}

type NotificationChannelSnsArray []NotificationChannelSnsInput

func (NotificationChannelSnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationChannelSns)(nil)).Elem()
}

func (i NotificationChannelSnsArray) ToNotificationChannelSnsArrayOutput() NotificationChannelSnsArrayOutput {
	return i.ToNotificationChannelSnsArrayOutputWithContext(context.Background())
}

func (i NotificationChannelSnsArray) ToNotificationChannelSnsArrayOutputWithContext(ctx context.Context) NotificationChannelSnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsArrayOutput)
}

// NotificationChannelSnsMapInput is an input type that accepts NotificationChannelSnsMap and NotificationChannelSnsMapOutput values.
// You can construct a concrete instance of `NotificationChannelSnsMapInput` via:
//
//          NotificationChannelSnsMap{ "key": NotificationChannelSnsArgs{...} }
type NotificationChannelSnsMapInput interface {
	pulumi.Input

	ToNotificationChannelSnsMapOutput() NotificationChannelSnsMapOutput
	ToNotificationChannelSnsMapOutputWithContext(context.Context) NotificationChannelSnsMapOutput
}

type NotificationChannelSnsMap map[string]NotificationChannelSnsInput

func (NotificationChannelSnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationChannelSns)(nil)).Elem()
}

func (i NotificationChannelSnsMap) ToNotificationChannelSnsMapOutput() NotificationChannelSnsMapOutput {
	return i.ToNotificationChannelSnsMapOutputWithContext(context.Background())
}

func (i NotificationChannelSnsMap) ToNotificationChannelSnsMapOutputWithContext(ctx context.Context) NotificationChannelSnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsMapOutput)
}

type NotificationChannelSnsOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSns)(nil))
}

func (o NotificationChannelSnsOutput) ToNotificationChannelSnsOutput() NotificationChannelSnsOutput {
	return o
}

func (o NotificationChannelSnsOutput) ToNotificationChannelSnsOutputWithContext(ctx context.Context) NotificationChannelSnsOutput {
	return o
}

func (o NotificationChannelSnsOutput) ToNotificationChannelSnsPtrOutput() NotificationChannelSnsPtrOutput {
	return o.ToNotificationChannelSnsPtrOutputWithContext(context.Background())
}

func (o NotificationChannelSnsOutput) ToNotificationChannelSnsPtrOutputWithContext(ctx context.Context) NotificationChannelSnsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationChannelSns) *NotificationChannelSns {
		return &v
	}).(NotificationChannelSnsPtrOutput)
}

type NotificationChannelSnsPtrOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSns)(nil))
}

func (o NotificationChannelSnsPtrOutput) ToNotificationChannelSnsPtrOutput() NotificationChannelSnsPtrOutput {
	return o
}

func (o NotificationChannelSnsPtrOutput) ToNotificationChannelSnsPtrOutputWithContext(ctx context.Context) NotificationChannelSnsPtrOutput {
	return o
}

func (o NotificationChannelSnsPtrOutput) Elem() NotificationChannelSnsOutput {
	return o.ApplyT(func(v *NotificationChannelSns) NotificationChannelSns {
		if v != nil {
			return *v
		}
		var ret NotificationChannelSns
		return ret
	}).(NotificationChannelSnsOutput)
}

type NotificationChannelSnsArrayOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationChannelSns)(nil))
}

func (o NotificationChannelSnsArrayOutput) ToNotificationChannelSnsArrayOutput() NotificationChannelSnsArrayOutput {
	return o
}

func (o NotificationChannelSnsArrayOutput) ToNotificationChannelSnsArrayOutputWithContext(ctx context.Context) NotificationChannelSnsArrayOutput {
	return o
}

func (o NotificationChannelSnsArrayOutput) Index(i pulumi.IntInput) NotificationChannelSnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationChannelSns {
		return vs[0].([]NotificationChannelSns)[vs[1].(int)]
	}).(NotificationChannelSnsOutput)
}

type NotificationChannelSnsMapOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationChannelSns)(nil))
}

func (o NotificationChannelSnsMapOutput) ToNotificationChannelSnsMapOutput() NotificationChannelSnsMapOutput {
	return o
}

func (o NotificationChannelSnsMapOutput) ToNotificationChannelSnsMapOutputWithContext(ctx context.Context) NotificationChannelSnsMapOutput {
	return o
}

func (o NotificationChannelSnsMapOutput) MapIndex(k pulumi.StringInput) NotificationChannelSnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationChannelSns {
		return vs[0].(map[string]NotificationChannelSns)[vs[1].(string)]
	}).(NotificationChannelSnsOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelSnsOutput{})
	pulumi.RegisterOutputType(NotificationChannelSnsPtrOutput{})
	pulumi.RegisterOutputType(NotificationChannelSnsArrayOutput{})
	pulumi.RegisterOutputType(NotificationChannelSnsMapOutput{})
}
