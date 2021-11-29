// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor
{
    [SysdigResourceType("sysdig:Monitor/alertDowntime:AlertDowntime")]
    public partial class AlertDowntime : Pulumi.CustomResource
    {
        [Output("capture")]
        public Output<Outputs.AlertDowntimeCapture?> Capture { get; private set; } = null!;

        [Output("customNotification")]
        public Output<Outputs.AlertDowntimeCustomNotification?> CustomNotification { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("entitiesToMonitors")]
        public Output<ImmutableArray<string>> EntitiesToMonitors { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notificationChannels")]
        public Output<ImmutableArray<int>> NotificationChannels { get; private set; } = null!;

        [Output("renotificationMinutes")]
        public Output<int?> RenotificationMinutes { get; private set; } = null!;

        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        [Output("severity")]
        public Output<int?> Severity { get; private set; } = null!;

        [Output("team")]
        public Output<int> Team { get; private set; } = null!;

        [Output("triggerAfterMinutes")]
        public Output<int> TriggerAfterMinutes { get; private set; } = null!;

        [Output("triggerAfterPct")]
        public Output<int?> TriggerAfterPct { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AlertDowntime resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertDowntime(string name, AlertDowntimeArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Monitor/alertDowntime:AlertDowntime", name, args ?? new AlertDowntimeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertDowntime(string name, Input<string> id, AlertDowntimeState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Monitor/alertDowntime:AlertDowntime", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertDowntime resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertDowntime Get(string name, Input<string> id, AlertDowntimeState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertDowntime(name, id, state, options);
        }
    }

    public sealed class AlertDowntimeArgs : Pulumi.ResourceArgs
    {
        [Input("capture")]
        public Input<Inputs.AlertDowntimeCaptureArgs>? Capture { get; set; }

        [Input("customNotification")]
        public Input<Inputs.AlertDowntimeCustomNotificationArgs>? CustomNotification { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("entitiesToMonitors", required: true)]
        private InputList<string>? _entitiesToMonitors;
        public InputList<string> EntitiesToMonitors
        {
            get => _entitiesToMonitors ?? (_entitiesToMonitors = new InputList<string>());
            set => _entitiesToMonitors = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("renotificationMinutes")]
        public Input<int>? RenotificationMinutes { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("severity")]
        public Input<int>? Severity { get; set; }

        [Input("triggerAfterMinutes", required: true)]
        public Input<int> TriggerAfterMinutes { get; set; } = null!;

        [Input("triggerAfterPct")]
        public Input<int>? TriggerAfterPct { get; set; }

        public AlertDowntimeArgs()
        {
        }
    }

    public sealed class AlertDowntimeState : Pulumi.ResourceArgs
    {
        [Input("capture")]
        public Input<Inputs.AlertDowntimeCaptureGetArgs>? Capture { get; set; }

        [Input("customNotification")]
        public Input<Inputs.AlertDowntimeCustomNotificationGetArgs>? CustomNotification { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("entitiesToMonitors")]
        private InputList<string>? _entitiesToMonitors;
        public InputList<string> EntitiesToMonitors
        {
            get => _entitiesToMonitors ?? (_entitiesToMonitors = new InputList<string>());
            set => _entitiesToMonitors = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("renotificationMinutes")]
        public Input<int>? RenotificationMinutes { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("severity")]
        public Input<int>? Severity { get; set; }

        [Input("team")]
        public Input<int>? Team { get; set; }

        [Input("triggerAfterMinutes")]
        public Input<int>? TriggerAfterMinutes { get; set; }

        [Input("triggerAfterPct")]
        public Input<int>? TriggerAfterPct { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public AlertDowntimeState()
        {
        }
    }
}
