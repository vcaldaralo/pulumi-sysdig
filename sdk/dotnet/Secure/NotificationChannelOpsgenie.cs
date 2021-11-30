// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure
{
    /// <summary>
    /// ## Import
    /// 
    /// Opsgenie notification channels for Secure can be imported using the ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie example 12345
    /// ```
    /// </summary>
    [SysdigResourceType("sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie")]
    public partial class NotificationChannelOpsgenie : Pulumi.CustomResource
    {
        /// <summary>
        /// Key for the API.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// If false, the channel will not emit notifications. Default is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The name of the Notification Channel. Must be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Send a new notification when the alert condition is 
        /// no longer triggered. Default is false.
        /// </summary>
        [Output("notifyWhenOk")]
        public Output<bool> NotifyWhenOk { get; private set; } = null!;

        /// <summary>
        /// Send a new notification when the alert is manually 
        /// acknowledged by a user. Default is false.
        /// </summary>
        [Output("notifyWhenResolved")]
        public Output<bool> NotifyWhenResolved { get; private set; } = null!;

        /// <summary>
        /// Send an initial test notification to check
        /// if the notification channel is working. Default is false.
        /// </summary>
        [Output("sendTestNotification")]
        public Output<bool?> SendTestNotification { get; private set; } = null!;

        /// <summary>
        /// (Computed) The current version of the Notification Channel.
        /// </summary>
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
        /// <summary>
        /// Key for the API.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// If false, the channel will not emit notifications. Default is true.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The name of the Notification Channel. Must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Send a new notification when the alert condition is 
        /// no longer triggered. Default is false.
        /// </summary>
        [Input("notifyWhenOk", required: true)]
        public Input<bool> NotifyWhenOk { get; set; } = null!;

        /// <summary>
        /// Send a new notification when the alert is manually 
        /// acknowledged by a user. Default is false.
        /// </summary>
        [Input("notifyWhenResolved", required: true)]
        public Input<bool> NotifyWhenResolved { get; set; } = null!;

        /// <summary>
        /// Send an initial test notification to check
        /// if the notification channel is working. Default is false.
        /// </summary>
        [Input("sendTestNotification")]
        public Input<bool>? SendTestNotification { get; set; }

        public NotificationChannelOpsgenieArgs()
        {
        }
    }

    public sealed class NotificationChannelOpsgenieState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key for the API.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// If false, the channel will not emit notifications. Default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the Notification Channel. Must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Send a new notification when the alert condition is 
        /// no longer triggered. Default is false.
        /// </summary>
        [Input("notifyWhenOk")]
        public Input<bool>? NotifyWhenOk { get; set; }

        /// <summary>
        /// Send a new notification when the alert is manually 
        /// acknowledged by a user. Default is false.
        /// </summary>
        [Input("notifyWhenResolved")]
        public Input<bool>? NotifyWhenResolved { get; set; }

        /// <summary>
        /// Send an initial test notification to check
        /// if the notification channel is working. Default is false.
        /// </summary>
        [Input("sendTestNotification")]
        public Input<bool>? SendTestNotification { get; set; }

        /// <summary>
        /// (Computed) The current version of the Notification Channel.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public NotificationChannelOpsgenieState()
        {
        }
    }
}
