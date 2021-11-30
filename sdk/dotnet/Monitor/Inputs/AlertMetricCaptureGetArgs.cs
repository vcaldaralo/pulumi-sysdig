// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Inputs
{

    public sealed class AlertMetricCaptureGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time frame in seconds of the capture.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// Defines the name of the capture file.
        /// </summary>
        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        /// <summary>
        /// Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        public AlertMetricCaptureGetArgs()
        {
        }
    }
}
