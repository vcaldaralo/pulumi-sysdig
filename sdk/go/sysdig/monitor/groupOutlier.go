// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GroupOutlier struct {
	pulumi.CustomResourceState

	Capture               GroupOutlierCapturePtrOutput            `pulumi:"capture"`
	CustomNotification    GroupOutlierCustomNotificationPtrOutput `pulumi:"customNotification"`
	Description           pulumi.StringPtrOutput                  `pulumi:"description"`
	Enabled               pulumi.BoolPtrOutput                    `pulumi:"enabled"`
	Monitors              pulumi.StringArrayOutput                `pulumi:"monitors"`
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	NotificationChannels  pulumi.IntArrayOutput                   `pulumi:"notificationChannels"`
	RenotificationMinutes pulumi.IntPtrOutput                     `pulumi:"renotificationMinutes"`
	Scope                 pulumi.StringPtrOutput                  `pulumi:"scope"`
	Severity              pulumi.IntPtrOutput                     `pulumi:"severity"`
	Team                  pulumi.IntOutput                        `pulumi:"team"`
	TriggerAfterMinutes   pulumi.IntOutput                        `pulumi:"triggerAfterMinutes"`
	Version               pulumi.IntOutput                        `pulumi:"version"`
}

// NewGroupOutlier registers a new resource with the given unique name, arguments, and options.
func NewGroupOutlier(ctx *pulumi.Context,
	name string, args *GroupOutlierArgs, opts ...pulumi.ResourceOption) (*GroupOutlier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Monitors == nil {
		return nil, errors.New("invalid value for required argument 'Monitors'")
	}
	if args.TriggerAfterMinutes == nil {
		return nil, errors.New("invalid value for required argument 'TriggerAfterMinutes'")
	}
	var resource GroupOutlier
	err := ctx.RegisterResource("sysdig:Monitor/groupOutlier:GroupOutlier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupOutlier gets an existing GroupOutlier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupOutlier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupOutlierState, opts ...pulumi.ResourceOption) (*GroupOutlier, error) {
	var resource GroupOutlier
	err := ctx.ReadResource("sysdig:Monitor/groupOutlier:GroupOutlier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupOutlier resources.
type groupOutlierState struct {
	Capture               *GroupOutlierCapture            `pulumi:"capture"`
	CustomNotification    *GroupOutlierCustomNotification `pulumi:"customNotification"`
	Description           *string                         `pulumi:"description"`
	Enabled               *bool                           `pulumi:"enabled"`
	Monitors              []string                        `pulumi:"monitors"`
	Name                  *string                         `pulumi:"name"`
	NotificationChannels  []int                           `pulumi:"notificationChannels"`
	RenotificationMinutes *int                            `pulumi:"renotificationMinutes"`
	Scope                 *string                         `pulumi:"scope"`
	Severity              *int                            `pulumi:"severity"`
	Team                  *int                            `pulumi:"team"`
	TriggerAfterMinutes   *int                            `pulumi:"triggerAfterMinutes"`
	Version               *int                            `pulumi:"version"`
}

type GroupOutlierState struct {
	Capture               GroupOutlierCapturePtrInput
	CustomNotification    GroupOutlierCustomNotificationPtrInput
	Description           pulumi.StringPtrInput
	Enabled               pulumi.BoolPtrInput
	Monitors              pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	NotificationChannels  pulumi.IntArrayInput
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	Severity              pulumi.IntPtrInput
	Team                  pulumi.IntPtrInput
	TriggerAfterMinutes   pulumi.IntPtrInput
	Version               pulumi.IntPtrInput
}

func (GroupOutlierState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupOutlierState)(nil)).Elem()
}

type groupOutlierArgs struct {
	Capture               *GroupOutlierCapture            `pulumi:"capture"`
	CustomNotification    *GroupOutlierCustomNotification `pulumi:"customNotification"`
	Description           *string                         `pulumi:"description"`
	Enabled               *bool                           `pulumi:"enabled"`
	Monitors              []string                        `pulumi:"monitors"`
	Name                  *string                         `pulumi:"name"`
	NotificationChannels  []int                           `pulumi:"notificationChannels"`
	RenotificationMinutes *int                            `pulumi:"renotificationMinutes"`
	Scope                 *string                         `pulumi:"scope"`
	Severity              *int                            `pulumi:"severity"`
	TriggerAfterMinutes   int                             `pulumi:"triggerAfterMinutes"`
}

// The set of arguments for constructing a GroupOutlier resource.
type GroupOutlierArgs struct {
	Capture               GroupOutlierCapturePtrInput
	CustomNotification    GroupOutlierCustomNotificationPtrInput
	Description           pulumi.StringPtrInput
	Enabled               pulumi.BoolPtrInput
	Monitors              pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	NotificationChannels  pulumi.IntArrayInput
	RenotificationMinutes pulumi.IntPtrInput
	Scope                 pulumi.StringPtrInput
	Severity              pulumi.IntPtrInput
	TriggerAfterMinutes   pulumi.IntInput
}

func (GroupOutlierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupOutlierArgs)(nil)).Elem()
}

type GroupOutlierInput interface {
	pulumi.Input

	ToGroupOutlierOutput() GroupOutlierOutput
	ToGroupOutlierOutputWithContext(ctx context.Context) GroupOutlierOutput
}

func (*GroupOutlier) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutlier)(nil))
}

func (i *GroupOutlier) ToGroupOutlierOutput() GroupOutlierOutput {
	return i.ToGroupOutlierOutputWithContext(context.Background())
}

func (i *GroupOutlier) ToGroupOutlierOutputWithContext(ctx context.Context) GroupOutlierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutlierOutput)
}

func (i *GroupOutlier) ToGroupOutlierPtrOutput() GroupOutlierPtrOutput {
	return i.ToGroupOutlierPtrOutputWithContext(context.Background())
}

func (i *GroupOutlier) ToGroupOutlierPtrOutputWithContext(ctx context.Context) GroupOutlierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutlierPtrOutput)
}

type GroupOutlierPtrInput interface {
	pulumi.Input

	ToGroupOutlierPtrOutput() GroupOutlierPtrOutput
	ToGroupOutlierPtrOutputWithContext(ctx context.Context) GroupOutlierPtrOutput
}

type groupOutlierPtrType GroupOutlierArgs

func (*groupOutlierPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupOutlier)(nil))
}

func (i *groupOutlierPtrType) ToGroupOutlierPtrOutput() GroupOutlierPtrOutput {
	return i.ToGroupOutlierPtrOutputWithContext(context.Background())
}

func (i *groupOutlierPtrType) ToGroupOutlierPtrOutputWithContext(ctx context.Context) GroupOutlierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutlierPtrOutput)
}

// GroupOutlierArrayInput is an input type that accepts GroupOutlierArray and GroupOutlierArrayOutput values.
// You can construct a concrete instance of `GroupOutlierArrayInput` via:
//
//          GroupOutlierArray{ GroupOutlierArgs{...} }
type GroupOutlierArrayInput interface {
	pulumi.Input

	ToGroupOutlierArrayOutput() GroupOutlierArrayOutput
	ToGroupOutlierArrayOutputWithContext(context.Context) GroupOutlierArrayOutput
}

type GroupOutlierArray []GroupOutlierInput

func (GroupOutlierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupOutlier)(nil)).Elem()
}

func (i GroupOutlierArray) ToGroupOutlierArrayOutput() GroupOutlierArrayOutput {
	return i.ToGroupOutlierArrayOutputWithContext(context.Background())
}

func (i GroupOutlierArray) ToGroupOutlierArrayOutputWithContext(ctx context.Context) GroupOutlierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutlierArrayOutput)
}

// GroupOutlierMapInput is an input type that accepts GroupOutlierMap and GroupOutlierMapOutput values.
// You can construct a concrete instance of `GroupOutlierMapInput` via:
//
//          GroupOutlierMap{ "key": GroupOutlierArgs{...} }
type GroupOutlierMapInput interface {
	pulumi.Input

	ToGroupOutlierMapOutput() GroupOutlierMapOutput
	ToGroupOutlierMapOutputWithContext(context.Context) GroupOutlierMapOutput
}

type GroupOutlierMap map[string]GroupOutlierInput

func (GroupOutlierMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupOutlier)(nil)).Elem()
}

func (i GroupOutlierMap) ToGroupOutlierMapOutput() GroupOutlierMapOutput {
	return i.ToGroupOutlierMapOutputWithContext(context.Background())
}

func (i GroupOutlierMap) ToGroupOutlierMapOutputWithContext(ctx context.Context) GroupOutlierMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutlierMapOutput)
}

type GroupOutlierOutput struct{ *pulumi.OutputState }

func (GroupOutlierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutlier)(nil))
}

func (o GroupOutlierOutput) ToGroupOutlierOutput() GroupOutlierOutput {
	return o
}

func (o GroupOutlierOutput) ToGroupOutlierOutputWithContext(ctx context.Context) GroupOutlierOutput {
	return o
}

func (o GroupOutlierOutput) ToGroupOutlierPtrOutput() GroupOutlierPtrOutput {
	return o.ToGroupOutlierPtrOutputWithContext(context.Background())
}

func (o GroupOutlierOutput) ToGroupOutlierPtrOutputWithContext(ctx context.Context) GroupOutlierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupOutlier) *GroupOutlier {
		return &v
	}).(GroupOutlierPtrOutput)
}

type GroupOutlierPtrOutput struct{ *pulumi.OutputState }

func (GroupOutlierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupOutlier)(nil))
}

func (o GroupOutlierPtrOutput) ToGroupOutlierPtrOutput() GroupOutlierPtrOutput {
	return o
}

func (o GroupOutlierPtrOutput) ToGroupOutlierPtrOutputWithContext(ctx context.Context) GroupOutlierPtrOutput {
	return o
}

func (o GroupOutlierPtrOutput) Elem() GroupOutlierOutput {
	return o.ApplyT(func(v *GroupOutlier) GroupOutlier {
		if v != nil {
			return *v
		}
		var ret GroupOutlier
		return ret
	}).(GroupOutlierOutput)
}

type GroupOutlierArrayOutput struct{ *pulumi.OutputState }

func (GroupOutlierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupOutlier)(nil))
}

func (o GroupOutlierArrayOutput) ToGroupOutlierArrayOutput() GroupOutlierArrayOutput {
	return o
}

func (o GroupOutlierArrayOutput) ToGroupOutlierArrayOutputWithContext(ctx context.Context) GroupOutlierArrayOutput {
	return o
}

func (o GroupOutlierArrayOutput) Index(i pulumi.IntInput) GroupOutlierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupOutlier {
		return vs[0].([]GroupOutlier)[vs[1].(int)]
	}).(GroupOutlierOutput)
}

type GroupOutlierMapOutput struct{ *pulumi.OutputState }

func (GroupOutlierMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupOutlier)(nil))
}

func (o GroupOutlierMapOutput) ToGroupOutlierMapOutput() GroupOutlierMapOutput {
	return o
}

func (o GroupOutlierMapOutput) ToGroupOutlierMapOutputWithContext(ctx context.Context) GroupOutlierMapOutput {
	return o
}

func (o GroupOutlierMapOutput) MapIndex(k pulumi.StringInput) GroupOutlierOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupOutlier {
		return vs[0].(map[string]GroupOutlier)[vs[1].(string)]
	}).(GroupOutlierOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutlierOutput{})
	pulumi.RegisterOutputType(GroupOutlierPtrOutput{})
	pulumi.RegisterOutputType(GroupOutlierArrayOutput{})
	pulumi.RegisterOutputType(GroupOutlierMapOutput{})
}
