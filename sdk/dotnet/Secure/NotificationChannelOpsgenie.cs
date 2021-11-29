// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure
{
    [SysdigResourceType("sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie")]
    public partial class NotificationChannelOpsgenie : Pulumi.CustomResource
    {
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

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

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationChannelOpsgenie resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationChannelOpsgenie(string name, NotificationChannelOpsgenieArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie", name, args ?? new NotificationChannelOpsgenieArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationChannelOpsgenie(string name, Input<string> id, NotificationChannelOpsgenieState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotificationChannelOpsgenie resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationChannelOpsgenie Get(string name, Input<string> id, NotificationChannelOpsgenieState? state = null, CustomResourceOptions? options = null)
        {
            return new NotificationChannelOpsgenie(name, id, state, options);
        }
    }

    public sealed class NotificationChannelOpsgenieArgs : Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

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

        public NotificationChannelOpsgenieArgs()
        {
        }
    }

    public sealed class NotificationChannelOpsgenieState : Pulumi.ResourceArgs
    {
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

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

        [Input("version")]
        public Input<int>? Version { get; set; }

        public NotificationChannelOpsgenieState()
        {
        }
    }
}
