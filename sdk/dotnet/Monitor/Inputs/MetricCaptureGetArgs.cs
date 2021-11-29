// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Inputs
{

    public sealed class MetricCaptureGetArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        public MetricCaptureGetArgs()
        {
        }
    }
}
