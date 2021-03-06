// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Outputs
{

    [OutputType]
    public sealed class AlertAnomalyCustomNotification
    {
        /// <summary>
        /// Text to add after the alert template.
        /// </summary>
        public readonly string? Append;
        /// <summary>
        /// Text to add before the alert template.
        /// </summary>
        public readonly string? Prepend;
        /// <summary>
        /// Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private AlertAnomalyCustomNotification(
            string? append,

            string? prepend,

            string title)
        {
            Append = append;
            Prepend = prepend;
            Title = title;
        }
    }
}
