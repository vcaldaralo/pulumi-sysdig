// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Inputs
{

    public sealed class DashboardPanelQueryGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PromQL query. Must be a valid PromQL query with existing
        /// metrics in Sysdig Monitor.
        /// </summary>
        [Input("promql", required: true)]
        public Input<string> Promql { get; set; } = null!;

        /// <summary>
        /// The type of metric for this query. Can be one of: `percent`, `data`, `data rate`, 
        /// `number`, `number rate`, `time`.
        /// </summary>
        [Input("unit", required: true)]
        public Input<string> Unit { get; set; } = null!;

        public DashboardPanelQueryGetArgs()
        {
        }
    }
}
