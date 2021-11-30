// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Inputs
{

    public sealed class AlertPromqlCustomNotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text to add after the alert template.
        /// </summary>
        [Input("append")]
        public Input<string>? Append { get; set; }

        /// <summary>
        /// Text to add before the alert template.
        /// </summary>
        [Input("prepend")]
        public Input<string>? Prepend { get; set; }

        /// <summary>
        /// Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AlertPromqlCustomNotificationArgs()
        {
        }
    }
}
