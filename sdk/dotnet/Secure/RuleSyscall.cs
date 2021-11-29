// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure
{
    [SysdigResourceType("sysdig:Secure/ruleSyscall:RuleSyscall")]
    public partial class RuleSyscall : Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("matching")]
        public Output<bool?> Matching { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("syscalls")]
        public Output<ImmutableArray<string>> Syscalls { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a RuleSyscall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleSyscall(string name, RuleSyscallArgs? args = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/ruleSyscall:RuleSyscall", name, args ?? new RuleSyscallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RuleSyscall(string name, Input<string> id, RuleSyscallState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/ruleSyscall:RuleSyscall", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RuleSyscall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleSyscall Get(string name, Input<string> id, RuleSyscallState? state = null, CustomResourceOptions? options = null)
        {
            return new RuleSyscall(name, id, state, options);
        }
    }

    public sealed class RuleSyscallArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("matching")]
        public Input<bool>? Matching { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("syscalls")]
        private InputList<string>? _syscalls;
        public InputList<string> Syscalls
        {
            get => _syscalls ?? (_syscalls = new InputList<string>());
            set => _syscalls = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public RuleSyscallArgs()
        {
        }
    }

    public sealed class RuleSyscallState : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("matching")]
        public Input<bool>? Matching { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("syscalls")]
        private InputList<string>? _syscalls;
        public InputList<string> Syscalls
        {
            get => _syscalls ?? (_syscalls = new InputList<string>());
            set => _syscalls = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public RuleSyscallState()
        {
        }
    }
}
