// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationChannelWebhook struct {
	pulumi.CustomResourceState

	Enabled              pulumi.BoolOutput    `pulumi:"enabled"`
	Name                 pulumi.StringOutput  `pulumi:"name"`
	NotifyWhenOk         pulumi.BoolOutput    `pulumi:"notifyWhenOk"`
	NotifyWhenResolved   pulumi.BoolOutput    `pulumi:"notifyWhenResolved"`
	SendTestNotification pulumi.BoolPtrOutput `pulumi:"sendTestNotification"`
	Url                  pulumi.StringOutput  `pulumi:"url"`
	Version              pulumi.IntOutput     `pulumi:"version"`
}

// NewNotificationChannelWebhook registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannelWebhook(ctx *pulumi.Context,
	name string, args *NotificationChannelWebhookArgs, opts ...pulumi.ResourceOption) (*NotificationChannelWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource NotificationChannelWebhook
	err := ctx.RegisterResource("sysdig:Secure/notificationChannelWebhook:NotificationChannelWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationChannelWebhook gets an existing NotificationChannelWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannelWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelWebhookState, opts ...pulumi.ResourceOption) (*NotificationChannelWebhook, error) {
	var resource NotificationChannelWebhook
	err := ctx.ReadResource("sysdig:Secure/notificationChannelWebhook:NotificationChannelWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationChannelWebhook resources.
type notificationChannelWebhookState struct {
	Enabled              *bool   `pulumi:"enabled"`
	Name                 *string `pulumi:"name"`
	NotifyWhenOk         *bool   `pulumi:"notifyWhenOk"`
	NotifyWhenResolved   *bool   `pulumi:"notifyWhenResolved"`
	SendTestNotification *bool   `pulumi:"sendTestNotification"`
	Url                  *string `pulumi:"url"`
	Version              *int    `pulumi:"version"`
}

type NotificationChannelWebhookState struct {
	Enabled              pulumi.BoolPtrInput
	Name                 pulumi.StringPtrInput
	NotifyWhenOk         pulumi.BoolPtrInput
	NotifyWhenResolved   pulumi.BoolPtrInput
	SendTestNotification pulumi.BoolPtrInput
	Url                  pulumi.StringPtrInput
	Version              pulumi.IntPtrInput
}

func (NotificationChannelWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelWebhookState)(nil)).Elem()
}

type notificationChannelWebhookArgs struct {
	Enabled              bool    `pulumi:"enabled"`
	Name                 *string `pulumi:"name"`
	NotifyWhenOk         bool    `pulumi:"notifyWhenOk"`
	NotifyWhenResolved   bool    `pulumi:"notifyWhenResolved"`
	SendTestNotification *bool   `pulumi:"sendTestNotification"`
	Url                  string  `pulumi:"url"`
}

// The set of arguments for constructing a NotificationChannelWebhook resource.
type NotificationChannelWebhookArgs struct {
	Enabled              pulumi.BoolInput
	Name                 pulumi.StringPtrInput
	NotifyWhenOk         pulumi.BoolInput
	NotifyWhenResolved   pulumi.BoolInput
	SendTestNotification pulumi.BoolPtrInput
	Url                  pulumi.StringInput
}

func (NotificationChannelWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelWebhookArgs)(nil)).Elem()
}

type NotificationChannelWebhookInput interface {
	pulumi.Input

	ToNotificationChannelWebhookOutput() NotificationChannelWebhookOutput
	ToNotificationChannelWebhookOutputWithContext(ctx context.Context) NotificationChannelWebhookOutput
}

func (*NotificationChannelWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelWebhook)(nil))
}

func (i *NotificationChannelWebhook) ToNotificationChannelWebhookOutput() NotificationChannelWebhookOutput {
	return i.ToNotificationChannelWebhookOutputWithContext(context.Background())
}

func (i *NotificationChannelWebhook) ToNotificationChannelWebhookOutputWithContext(ctx context.Context) NotificationChannelWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelWebhookOutput)
}

func (i *NotificationChannelWebhook) ToNotificationChannelWebhookPtrOutput() NotificationChannelWebhookPtrOutput {
	return i.ToNotificationChannelWebhookPtrOutputWithContext(context.Background())
}

func (i *NotificationChannelWebhook) ToNotificationChannelWebhookPtrOutputWithContext(ctx context.Context) NotificationChannelWebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelWebhookPtrOutput)
}

type NotificationChannelWebhookPtrInput interface {
	pulumi.Input

	ToNotificationChannelWebhookPtrOutput() NotificationChannelWebhookPtrOutput
	ToNotificationChannelWebhookPtrOutputWithContext(ctx context.Context) NotificationChannelWebhookPtrOutput
}

type notificationChannelWebhookPtrType NotificationChannelWebhookArgs

func (*notificationChannelWebhookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelWebhook)(nil))
}

func (i *notificationChannelWebhookPtrType) ToNotificationChannelWebhookPtrOutput() NotificationChannelWebhookPtrOutput {
	return i.ToNotificationChannelWebhookPtrOutputWithContext(context.Background())
}

func (i *notificationChannelWebhookPtrType) ToNotificationChannelWebhookPtrOutputWithContext(ctx context.Context) NotificationChannelWebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelWebhookPtrOutput)
}

// NotificationChannelWebhookArrayInput is an input type that accepts NotificationChannelWebhookArray and NotificationChannelWebhookArrayOutput values.
// You can construct a concrete instance of `NotificationChannelWebhookArrayInput` via:
//
//          NotificationChannelWebhookArray{ NotificationChannelWebhookArgs{...} }
type NotificationChannelWebhookArrayInput interface {
	pulumi.Input

	ToNotificationChannelWebhookArrayOutput() NotificationChannelWebhookArrayOutput
	ToNotificationChannelWebhookArrayOutputWithContext(context.Context) NotificationChannelWebhookArrayOutput
}

type NotificationChannelWebhookArray []NotificationChannelWebhookInput

func (NotificationChannelWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationChannelWebhook)(nil)).Elem()
}

func (i NotificationChannelWebhookArray) ToNotificationChannelWebhookArrayOutput() NotificationChannelWebhookArrayOutput {
	return i.ToNotificationChannelWebhookArrayOutputWithContext(context.Background())
}

func (i NotificationChannelWebhookArray) ToNotificationChannelWebhookArrayOutputWithContext(ctx context.Context) NotificationChannelWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelWebhookArrayOutput)
}

// NotificationChannelWebhookMapInput is an input type that accepts NotificationChannelWebhookMap and NotificationChannelWebhookMapOutput values.
// You can construct a concrete instance of `NotificationChannelWebhookMapInput` via:
//
//          NotificationChannelWebhookMap{ "key": NotificationChannelWebhookArgs{...} }
type NotificationChannelWebhookMapInput interface {
	pulumi.Input

	ToNotificationChannelWebhookMapOutput() NotificationChannelWebhookMapOutput
	ToNotificationChannelWebhookMapOutputWithContext(context.Context) NotificationChannelWebhookMapOutput
}

type NotificationChannelWebhookMap map[string]NotificationChannelWebhookInput

func (NotificationChannelWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationChannelWebhook)(nil)).Elem()
}

func (i NotificationChannelWebhookMap) ToNotificationChannelWebhookMapOutput() NotificationChannelWebhookMapOutput {
	return i.ToNotificationChannelWebhookMapOutputWithContext(context.Background())
}

func (i NotificationChannelWebhookMap) ToNotificationChannelWebhookMapOutputWithContext(ctx context.Context) NotificationChannelWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelWebhookMapOutput)
}

type NotificationChannelWebhookOutput struct{ *pulumi.OutputState }

func (NotificationChannelWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelWebhook)(nil))
}

func (o NotificationChannelWebhookOutput) ToNotificationChannelWebhookOutput() NotificationChannelWebhookOutput {
	return o
}

func (o NotificationChannelWebhookOutput) ToNotificationChannelWebhookOutputWithContext(ctx context.Context) NotificationChannelWebhookOutput {
	return o
}

func (o NotificationChannelWebhookOutput) ToNotificationChannelWebhookPtrOutput() NotificationChannelWebhookPtrOutput {
	return o.ToNotificationChannelWebhookPtrOutputWithContext(context.Background())
}

func (o NotificationChannelWebhookOutput) ToNotificationChannelWebhookPtrOutputWithContext(ctx context.Context) NotificationChannelWebhookPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationChannelWebhook) *NotificationChannelWebhook {
		return &v
	}).(NotificationChannelWebhookPtrOutput)
}

type NotificationChannelWebhookPtrOutput struct{ *pulumi.OutputState }

func (NotificationChannelWebhookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelWebhook)(nil))
}

func (o NotificationChannelWebhookPtrOutput) ToNotificationChannelWebhookPtrOutput() NotificationChannelWebhookPtrOutput {
	return o
}

func (o NotificationChannelWebhookPtrOutput) ToNotificationChannelWebhookPtrOutputWithContext(ctx context.Context) NotificationChannelWebhookPtrOutput {
	return o
}

func (o NotificationChannelWebhookPtrOutput) Elem() NotificationChannelWebhookOutput {
	return o.ApplyT(func(v *NotificationChannelWebhook) NotificationChannelWebhook {
		if v != nil {
			return *v
		}
		var ret NotificationChannelWebhook
		return ret
	}).(NotificationChannelWebhookOutput)
}

type NotificationChannelWebhookArrayOutput struct{ *pulumi.OutputState }

func (NotificationChannelWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationChannelWebhook)(nil))
}

func (o NotificationChannelWebhookArrayOutput) ToNotificationChannelWebhookArrayOutput() NotificationChannelWebhookArrayOutput {
	return o
}

func (o NotificationChannelWebhookArrayOutput) ToNotificationChannelWebhookArrayOutputWithContext(ctx context.Context) NotificationChannelWebhookArrayOutput {
	return o
}

func (o NotificationChannelWebhookArrayOutput) Index(i pulumi.IntInput) NotificationChannelWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationChannelWebhook {
		return vs[0].([]NotificationChannelWebhook)[vs[1].(int)]
	}).(NotificationChannelWebhookOutput)
}

type NotificationChannelWebhookMapOutput struct{ *pulumi.OutputState }

func (NotificationChannelWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationChannelWebhook)(nil))
}

func (o NotificationChannelWebhookMapOutput) ToNotificationChannelWebhookMapOutput() NotificationChannelWebhookMapOutput {
	return o
}

func (o NotificationChannelWebhookMapOutput) ToNotificationChannelWebhookMapOutputWithContext(ctx context.Context) NotificationChannelWebhookMapOutput {
	return o
}

func (o NotificationChannelWebhookMapOutput) MapIndex(k pulumi.StringInput) NotificationChannelWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationChannelWebhook {
		return vs[0].(map[string]NotificationChannelWebhook)[vs[1].(string)]
	}).(NotificationChannelWebhookOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelWebhookOutput{})
	pulumi.RegisterOutputType(NotificationChannelWebhookPtrOutput{})
	pulumi.RegisterOutputType(NotificationChannelWebhookArrayOutput{})
	pulumi.RegisterOutputType(NotificationChannelWebhookMapOutput{})
}
