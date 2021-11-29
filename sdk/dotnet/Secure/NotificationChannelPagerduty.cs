// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure
{
    [SysdigResourceType("sysdig:Secure/notificationChannelPagerduty:NotificationChannelPagerduty")]
    public partial class NotificationChannelPagerduty : Pulumi.CustomResource
    {
        [Output("account")]
        public Output<string> Account { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notifyWhenOk")]
        public Output<bool> NotifyWhenOk { get; private set; } = null!;

        [Output("notifyWhenResolved")]
        public Output<bool> NotifyWhenResolved { get; private set; } = null!;

        [Output("sendTestNotification")]
        public Output<bool?> SendTestNotification { get; private set; } = null!;

        [Output("serviceKey")]
        public Output<string> ServiceKey { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationChannelPagerduty resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationChannelPagerduty(string name, NotificationChannelPagerdutyArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Secure/notificationChannelPagerduty:NotificationChannelPagerduty", name, args ?? new NotificationChannelPagerdutyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationChannelPagerduty(string name, Input<string> id, NotificationChannelPagerdutyState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/notificationChannelPagerduty:NotificationChannelPagerduty", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NotificationChannelPagerduty resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationChannelPagerduty Get(string name, Input<string> id, NotificationChannelPagerdutyState? state = null, CustomResourceOptions? options = null)
        {
            return new NotificationChannelPagerduty(name, id, state, options);
        }
    }

    public sealed class NotificationChannelPagerdutyArgs : Pulumi.ResourceArgs
    {
        [Input("account", required: true)]
        public Input<string> Account { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifyWhenOk", required: true)]
        public Input<bool> NotifyWhenOk { get; set; } = null!;

        [Input("notifyWhenResolved", required: true)]
        public Input<bool> NotifyWhenResolved { get; set; } = null!;

        [Input("sendTestNotification")]
        public Input<bool>? SendTestNotification { get; set; }

        [Input("serviceKey", required: true)]
        public Input<string> ServiceKey { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public NotificationChannelPagerdutyArgs()
        {
        }
    }

    public sealed class NotificationChannelPagerdutyState : Pulumi.ResourceArgs
    {
        [Input("account")]
        public Input<string>? Account { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifyWhenOk")]
        public Input<bool>? NotifyWhenOk { get; set; }

        [Input("notifyWhenResolved")]
        public Input<bool>? NotifyWhenResolved { get; set; }

        [Input("sendTestNotification")]
        public Input<bool>? SendTestNotification { get; set; }

        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public NotificationChannelPagerdutyState()
        {
        }
    }
}
