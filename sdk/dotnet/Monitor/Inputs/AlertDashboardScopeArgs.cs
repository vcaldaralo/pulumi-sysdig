// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Inputs
{

    public sealed class AlertDashboardScopeArgs : Pulumi.ResourceArgs
    {
        [Input("comparator")]
        public Input<string>? Comparator { get; set; }

        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public AlertDashboardScopeArgs()
        {
        }
    }
}
