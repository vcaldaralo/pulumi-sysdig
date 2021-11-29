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
    public sealed class RuleFalcoException
    {
        public readonly ImmutableArray<string> Comps;
        public readonly ImmutableArray<string> Fields;
        public readonly string Name;
        public readonly string Values;

        [OutputConstructor]
        private RuleFalcoException(
            ImmutableArray<string> comps,

            ImmutableArray<string> fields,

            string name,

            string values)
        {
            Comps = comps;
            Fields = fields;
            Name = name;
            Values = values;
        }
    }
}
