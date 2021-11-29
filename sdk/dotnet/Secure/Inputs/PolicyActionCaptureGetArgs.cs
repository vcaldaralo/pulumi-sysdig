// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure.Inputs
{

    public sealed class PolicyActionCaptureGetArgs : Pulumi.ResourceArgs
    {
        [Input("secondsAfterEvent", required: true)]
        public Input<int> SecondsAfterEvent { get; set; } = null!;

        [Input("secondsBeforeEvent", required: true)]
        public Input<int> SecondsBeforeEvent { get; set; } = null!;

        public PolicyActionCaptureGetArgs()
        {
        }
    }
}
