// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure.Outputs
{

    [OutputType]
    public sealed class PolicyAction
    {
        public readonly ImmutableArray<Outputs.PolicyActionCapture> Captures;
        public readonly string? Container;

        [OutputConstructor]
        private PolicyAction(
            ImmutableArray<Outputs.PolicyActionCapture> captures,

            string? container)
        {
            Captures = captures;
            Container = container;
        }
    }
}
