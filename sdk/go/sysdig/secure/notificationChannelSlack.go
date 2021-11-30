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
// Slack notification channels for Secure can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import sysdig:Secure/notificationChannelSlack:NotificationChannelSlack example 12345
// ```
type NotificationChannelSlack struct {
	pulumi.CustomResourceState

	// Channel name from this Slack.
	Channel pulumi.StringOutput `pulumi:"channel"`
	// If false, the channel will not emit notifications. Default is true.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The name of the Notification Channel. Must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk pulumi.BoolOutput `pulumi:"notifyWhenOk"`
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved pulumi.BoolOutput `pulumi:"notifyWhenResolved"`
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification pulumi.BoolPtrOutput `pulumi:"sendTestNotification"`
	// URL of the Slack.
	Url pulumi.StringOutput `pulumi:"url"`
	// (Computed) The current version of the Notification Channel.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNotificationChannelSlack registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannelSlack(ctx *pulumi.Context,
	name string, args *NotificationChannelSlackArgs, opts ...pulumi.ResourceOption) (*NotificationChannelSlack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Channel == nil {
		return nil, errors.New("invalid value for required argument 'Channel'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.NotifyWhenOk == nil {
		return nil, errors.New("invalid value for required argument 'NotifyWhenOk'")
	}
	if args.NotifyWhenResolved == nil {
		return nil, errors.New("invalid value for required argument 'NotifyWhenResolved'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource NotificationChannelSlack
	err := ctx.RegisterResource("sysdig:Secure/notificationChannelSlack:NotificationChannelSlack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationChannelSlack gets an existing NotificationChannelSlack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannelSlack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelSlackState, opts ...pulumi.ResourceOption) (*NotificationChannelSlack, error) {
	var resource NotificationChannelSlack
	err := ctx.ReadResource("sysdig:Secure/notificationChannelSlack:NotificationChannelSlack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationChannelSlack resources.
type notificationChannelSlackState struct {
	// Channel name from this Slack.
	Channel *string `pulumi:"channel"`
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
	// URL of the Slack.
	Url *string `pulumi:"url"`
	// (Computed) The current version of the Notification Channel.
	Version *int `pulumi:"version"`
}

type NotificationChannelSlackState struct {
	// Channel name from this Slack.
	Channel pulumi.StringPtrInput
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
	// URL of the Slack.
	Url pulumi.StringPtrInput
	// (Computed) The current version of the Notification Channel.
	Version pulumi.IntPtrInput
}

func (NotificationChannelSlackState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelSlackState)(nil)).Elem()
}

type notificationChannelSlackArgs struct {
	// Channel name from this Slack.
	Channel string `pulumi:"channel"`
	// If false, the channel will not emit notifications. Default is true.
	Enabled bool `pulumi:"enabled"`
	// The name of the Notification Channel. Must be unique.
	Name *string `pulumi:"name"`
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk bool `pulumi:"notifyWhenOk"`
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved bool `pulumi:"notifyWhenResolved"`
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification *bool `pulumi:"sendTestNotification"`
	// URL of the Slack.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a NotificationChannelSlack resource.
type NotificationChannelSlackArgs struct {
	// Channel name from this Slack.
	Channel pulumi.StringInput
	// If false, the channel will not emit notifications. Default is true.
	Enabled pulumi.BoolInput
	// The name of the Notification Channel. Must be unique.
	Name pulumi.StringPtrInput
	// Send a new notification when the alert condition is
	// no longer triggered. Default is false.
	NotifyWhenOk pulumi.BoolInput
	// Send a new notification when the alert is manually
	// acknowledged by a user. Default is false.
	NotifyWhenResolved pulumi.BoolInput
	// Send an initial test notification to check
	// if the notification channel is working. Default is false.
	SendTestNotification pulumi.BoolPtrInput
	// URL of the Slack.
	Url pulumi.StringInput
}

func (NotificationChannelSlackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelSlackArgs)(nil)).Elem()
}

type NotificationChannelSlackInput interface {
	pulumi.Input

	ToNotificationChannelSlackOutput() NotificationChannelSlackOutput
	ToNotificationChannelSlackOutputWithContext(ctx context.Context) NotificationChannelSlackOutput
}

func (*NotificationChannelSlack) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSlack)(nil))
}

func (i *NotificationChannelSlack) ToNotificationChannelSlackOutput() NotificationChannelSlackOutput {
	return i.ToNotificationChannelSlackOutputWithContext(context.Background())
}

func (i *NotificationChannelSlack) ToNotificationChannelSlackOutputWithContext(ctx context.Context) NotificationChannelSlackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSlackOutput)
}

func (i *NotificationChannelSlack) ToNotificationChannelSlackPtrOutput() NotificationChannelSlackPtrOutput {
	return i.ToNotificationChannelSlackPtrOutputWithContext(context.Background())
}

func (i *NotificationChannelSlack) ToNotificationChannelSlackPtrOutputWithContext(ctx context.Context) NotificationChannelSlackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSlackPtrOutput)
}

type NotificationChannelSlackPtrInput interface {
	pulumi.Input

	ToNotificationChannelSlackPtrOutput() NotificationChannelSlackPtrOutput
	ToNotificationChannelSlackPtrOutputWithContext(ctx context.Context) NotificationChannelSlackPtrOutput
}

type notificationChannelSlackPtrType NotificationChannelSlackArgs

func (*notificationChannelSlackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSlack)(nil))
}

func (i *notificationChannelSlackPtrType) ToNotificationChannelSlackPtrOutput() NotificationChannelSlackPtrOutput {
	return i.ToNotificationChannelSlackPtrOutputWithContext(context.Background())
}

func (i *notificationChannelSlackPtrType) ToNotificationChannelSlackPtrOutputWithContext(ctx context.Context) NotificationChannelSlackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSlackPtrOutput)
}

// NotificationChannelSlackArrayInput is an input type that accepts NotificationChannelSlackArray and NotificationChannelSlackArrayOutput values.
// You can construct a concrete instance of `NotificationChannelSlackArrayInput` via:
//
//          NotificationChannelSlackArray{ NotificationChannelSlackArgs{...} }
type NotificationChannelSlackArrayInput interface {
	pulumi.Input

	ToNotificationChannelSlackArrayOutput() NotificationChannelSlackArrayOutput
	ToNotificationChannelSlackArrayOutputWithContext(context.Context) NotificationChannelSlackArrayOutput
}

type NotificationChannelSlackArray []NotificationChannelSlackInput

func (NotificationChannelSlackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationChannelSlack)(nil)).Elem()
}

func (i NotificationChannelSlackArray) ToNotificationChannelSlackArrayOutput() NotificationChannelSlackArrayOutput {
	return i.ToNotificationChannelSlackArrayOutputWithContext(context.Background())
}

func (i NotificationChannelSlackArray) ToNotificationChannelSlackArrayOutputWithContext(ctx context.Context) NotificationChannelSlackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSlackArrayOutput)
}

// NotificationChannelSlackMapInput is an input type that accepts NotificationChannelSlackMap and NotificationChannelSlackMapOutput values.
// You can construct a concrete instance of `NotificationChannelSlackMapInput` via:
//
//          NotificationChannelSlackMap{ "key": NotificationChannelSlackArgs{...} }
type NotificationChannelSlackMapInput interface {
	pulumi.Input

	ToNotificationChannelSlackMapOutput() NotificationChannelSlackMapOutput
	ToNotificationChannelSlackMapOutputWithContext(context.Context) NotificationChannelSlackMapOutput
}

type NotificationChannelSlackMap map[string]NotificationChannelSlackInput

func (NotificationChannelSlackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationChannelSlack)(nil)).Elem()
}

func (i NotificationChannelSlackMap) ToNotificationChannelSlackMapOutput() NotificationChannelSlackMapOutput {
	return i.ToNotificationChannelSlackMapOutputWithContext(context.Background())
}

func (i NotificationChannelSlackMap) ToNotificationChannelSlackMapOutputWithContext(ctx context.Context) NotificationChannelSlackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSlackMapOutput)
}

type NotificationChannelSlackOutput struct{ *pulumi.OutputState }

func (NotificationChannelSlackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSlack)(nil))
}

func (o NotificationChannelSlackOutput) ToNotificationChannelSlackOutput() NotificationChannelSlackOutput {
	return o
}

func (o NotificationChannelSlackOutput) ToNotificationChannelSlackOutputWithContext(ctx context.Context) NotificationChannelSlackOutput {
	return o
}

func (o NotificationChannelSlackOutput) ToNotificationChannelSlackPtrOutput() NotificationChannelSlackPtrOutput {
	return o.ToNotificationChannelSlackPtrOutputWithContext(context.Background())
}

func (o NotificationChannelSlackOutput) ToNotificationChannelSlackPtrOutputWithContext(ctx context.Context) NotificationChannelSlackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationChannelSlack) *NotificationChannelSlack {
		return &v
	}).(NotificationChannelSlackPtrOutput)
}

type NotificationChannelSlackPtrOutput struct{ *pulumi.OutputState }

func (NotificationChannelSlackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSlack)(nil))
}

func (o NotificationChannelSlackPtrOutput) ToNotificationChannelSlackPtrOutput() NotificationChannelSlackPtrOutput {
	return o
}

func (o NotificationChannelSlackPtrOutput) ToNotificationChannelSlackPtrOutputWithContext(ctx context.Context) NotificationChannelSlackPtrOutput {
	return o
}

func (o NotificationChannelSlackPtrOutput) Elem() NotificationChannelSlackOutput {
	return o.ApplyT(func(v *NotificationChannelSlack) NotificationChannelSlack {
		if v != nil {
			return *v
		}
		var ret NotificationChannelSlack
		return ret
	}).(NotificationChannelSlackOutput)
}

type NotificationChannelSlackArrayOutput struct{ *pulumi.OutputState }

func (NotificationChannelSlackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationChannelSlack)(nil))
}

func (o NotificationChannelSlackArrayOutput) ToNotificationChannelSlackArrayOutput() NotificationChannelSlackArrayOutput {
	return o
}

func (o NotificationChannelSlackArrayOutput) ToNotificationChannelSlackArrayOutputWithContext(ctx context.Context) NotificationChannelSlackArrayOutput {
	return o
}

func (o NotificationChannelSlackArrayOutput) Index(i pulumi.IntInput) NotificationChannelSlackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationChannelSlack {
		return vs[0].([]NotificationChannelSlack)[vs[1].(int)]
	}).(NotificationChannelSlackOutput)
}

type NotificationChannelSlackMapOutput struct{ *pulumi.OutputState }

func (NotificationChannelSlackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationChannelSlack)(nil))
}

func (o NotificationChannelSlackMapOutput) ToNotificationChannelSlackMapOutput() NotificationChannelSlackMapOutput {
	return o
}

func (o NotificationChannelSlackMapOutput) ToNotificationChannelSlackMapOutputWithContext(ctx context.Context) NotificationChannelSlackMapOutput {
	return o
}

func (o NotificationChannelSlackMapOutput) MapIndex(k pulumi.StringInput) NotificationChannelSlackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationChannelSlack {
		return vs[0].(map[string]NotificationChannelSlack)[vs[1].(string)]
	}).(NotificationChannelSlackOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelSlackOutput{})
	pulumi.RegisterOutputType(NotificationChannelSlackPtrOutput{})
	pulumi.RegisterOutputType(NotificationChannelSlackArrayOutput{})
	pulumi.RegisterOutputType(NotificationChannelSlackMapOutput{})
}
