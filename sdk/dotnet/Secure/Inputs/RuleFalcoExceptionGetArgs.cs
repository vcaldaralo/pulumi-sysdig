// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure.Inputs
{

    public sealed class RuleFalcoExceptionGetArgs : Pulumi.ResourceArgs
    {
        [Input("comps", required: true)]
        private InputList<string>? _comps;

        /// <summary>
        /// Contains comparison operators that align 1-1 with the items in the fields property.
        /// </summary>
        public InputList<string> Comps
        {
            get => _comps ?? (_comps = new InputList<string>());
            set => _comps = value;
        }

        [Input("fields", required: true)]
        private InputList<string>? _fields;

        /// <summary>
        /// Contains one or more fields that will extract a value from the syscall/k8s_audit events.
        /// </summary>
        public InputList<string> Fields
        {
            get => _fields ?? (_fields = new InputList<string>());
            set => _fields = value;
        }

        /// <summary>
        /// The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
        /// </summary>
        [Input("values", required: true)]
        public Input<string> Values { get; set; } = null!;

        public RuleFalcoExceptionGetArgs()
        {
        }
    }
}
