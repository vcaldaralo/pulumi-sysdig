// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure
{
    /// <summary>
    /// ## Import
    /// 
    /// Secure runtime policies can be imported using the ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import sysdig:Secure/policy:Policy example 12345
    /// ```
    /// </summary>
    [SysdigResourceType("sysdig:Secure/policy:Policy")]
    public partial class Policy : Pulumi.CustomResource
    {
        [Output("actions")]
        public Output<ImmutableArray<Outputs.PolicyAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// The description of Secure policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Will secure process with this rule?. By default this is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The name of the Secure policy. It must be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IDs of the notification channels to send alerts to
        /// when the policy is fired.
        /// </summary>
        [Output("notificationChannels")]
        public Output<ImmutableArray<int>> NotificationChannels { get; private set; } = null!;

        /// <summary>
        /// Array with the name of the rules to match.
        /// </summary>
        [Output("ruleNames")]
        public Output<ImmutableArray<string>> RuleNames { get; private set; } = null!;

        /// <summary>
        /// Limit appplication scope based in one expresion. For
        /// example: "host.ip.private = \\"10.0.23.1\\"". By default the rule won't be scoped
        /// and will target the entire infrastructure.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// The severity of Secure policy. The accepted values
        /// are: 0, 1, 2, 3 (High), 4, 5 (Medium), 6 (Low) and 7 (Info). The default value is 4 (Medium).
        /// </summary>
        [Output("severity")]
        public Output<int?> Severity { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of the runtime policy. Must be one of: `falco`, `list_matching`, `k8s_audit`, `aws_cloudtrail`. By default it is `falco`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Secure/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.PolicyActionArgs>? _actions;
        public InputList<Inputs.PolicyActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.PolicyActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The description of Secure policy.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Will secure process with this rule?. By default this is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the Secure policy. It must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;

        /// <summary>
        /// IDs of the notification channels to send alerts to
        /// when the policy is fired.
        /// </summary>
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("ruleNames")]
        private InputList<string>? _ruleNames;

        /// <summary>
        /// Array with the name of the rules to match.
        /// </summary>
        public InputList<string> RuleNames
        {
            get => _ruleNames ?? (_ruleNames = new InputList<string>());
            set => _ruleNames = value;
        }

        /// <summary>
        /// Limit appplication scope based in one expresion. For
        /// example: "host.ip.private = \\"10.0.23.1\\"". By default the rule won't be scoped
        /// and will target the entire infrastructure.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The severity of Secure policy. The accepted values
        /// are: 0, 1, 2, 3 (High), 4, 5 (Medium), 6 (Low) and 7 (Info). The default value is 4 (Medium).
        /// </summary>
        [Input("severity")]
        public Input<int>? Severity { get; set; }

        /// <summary>
        /// Specifies the type of the runtime policy. Must be one of: `falco`, `list_matching`, `k8s_audit`, `aws_cloudtrail`. By default it is `falco`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.PolicyActionGetArgs>? _actions;
        public InputList<Inputs.PolicyActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.PolicyActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The description of Secure policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Will secure process with this rule?. By default this is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the Secure policy. It must be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;

        /// <summary>
        /// IDs of the notification channels to send alerts to
        /// when the policy is fired.
        /// </summary>
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("ruleNames")]
        private InputList<string>? _ruleNames;

        /// <summary>
        /// Array with the name of the rules to match.
        /// </summary>
        public InputList<string> RuleNames
        {
            get => _ruleNames ?? (_ruleNames = new InputList<string>());
            set => _ruleNames = value;
        }

        /// <summary>
        /// Limit appplication scope based in one expresion. For
        /// example: "host.ip.private = \\"10.0.23.1\\"". By default the rule won't be scoped
        /// and will target the entire infrastructure.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The severity of Secure policy. The accepted values
        /// are: 0, 1, 2, 3 (High), 4, 5 (Medium), 6 (Low) and 7 (Info). The default value is 4 (Medium).
        /// </summary>
        [Input("severity")]
        public Input<int>? Severity { get; set; }

        /// <summary>
        /// Specifies the type of the runtime policy. Must be one of: `falco`, `list_matching`, `k8s_audit`, `aws_cloudtrail`. By default it is `falco`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public PolicyState()
        {
        }
    }
}
