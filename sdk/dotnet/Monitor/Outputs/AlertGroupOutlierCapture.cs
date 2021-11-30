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
    public sealed class AlertGroupOutlierCapture
    {
        /// <summary>
        /// Time frame in seconds of the capture.
        /// </summary>
        public readonly int Duration;
        /// <summary>
        /// Defines the name of the capture file.
        /// </summary>
        public readonly string Filename;
        /// <summary>
        /// Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        /// </summary>
        public readonly string? Filter;

        [OutputConstructor]
        private AlertGroupOutlierCapture(
            int duration,

            string filename,

            string? filter)
        {
            Duration = duration;
            Filename = filename;
            Filter = filter;
        }
    }
}
