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
    public sealed class AlertDashboardScope
    {
        public readonly string? Comparator;
        public readonly string Metric;
        public readonly ImmutableArray<string> Values;
        public readonly string? Variable;

        [OutputConstructor]
        private AlertDashboardScope(
            string? comparator,

            string metric,

            ImmutableArray<string> values,

            string? variable)
        {
            Comparator = comparator;
            Metric = metric;
            Values = values;
            Variable = variable;
        }
    }
}
