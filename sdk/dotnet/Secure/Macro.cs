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
    /// Secure macros can be imported using the ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import sysdig:Secure/macro:Macro example 12345
    /// ```
    /// </summary>
    [SysdigResourceType("sysdig:Secure/macro:Macro")]
    public partial class Macro : Pulumi.CustomResource
    {
        /// <summary>
        /// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
        /// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
        /// append macro called "foo" but not a second one. By default this is false.
        /// </summary>
        [Output("append")]
        public Output<bool?> Append { get; private set; } = null!;

        /// <summary>
        /// Macro condition. It can contain lists or other macros.
        /// </summary>
        [Output("condition")]
        public Output<string> Condition { get; private set; } = null!;

        /// <summary>
        /// The name of the macro. It must be unique if it's not in append mode.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Macro resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Macro(string name, MacroArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Secure/macro:Macro", name, args ?? new MacroArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Macro(string name, Input<string> id, MacroState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Secure/macro:Macro", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Macro resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Macro Get(string name, Input<string> id, MacroState? state = null, CustomResourceOptions? options = null)
        {
            return new Macro(name, id, state, options);
        }
    }

    public sealed class MacroArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
        /// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
        /// append macro called "foo" but not a second one. By default this is false.
        /// </summary>
        [Input("append")]
        public Input<bool>? Append { get; set; }

        /// <summary>
        /// Macro condition. It can contain lists or other macros.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// The name of the macro. It must be unique if it's not in append mode.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MacroArgs()
        {
        }
    }

    public sealed class MacroState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
        /// The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
        /// append macro called "foo" but not a second one. By default this is false.
        /// </summary>
        [Input("append")]
        public Input<bool>? Append { get; set; }

        /// <summary>
        /// Macro condition. It can contain lists or other macros.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// The name of the macro. It must be unique if it's not in append mode.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public MacroState()
        {
        }
    }
}
