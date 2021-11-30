// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Secure macros can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Secure/macro:Macro example 12345
 * ```
 */
export class Macro extends pulumi.CustomResource {
    /**
     * Get an existing Macro resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MacroState, opts?: pulumi.CustomResourceOptions): Macro {
        return new Macro(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/macro:Macro';

    /**
     * Returns true if the given object is an instance of Macro.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Macro {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Macro.__pulumiType;
    }

    /**
     * Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
     * The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
     * append macro called "foo" but not a second one. By default this is false.
     */
    public readonly append!: pulumi.Output<boolean | undefined>;
    /**
     * Macro condition. It can contain lists or other macros.
     */
    public readonly condition!: pulumi.Output<string>;
    /**
     * The name of the macro. It must be unique if it's not in append mode.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Macro resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MacroArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MacroArgs | MacroState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MacroState | undefined;
            inputs["append"] = state ? state.append : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as MacroArgs | undefined;
            if ((!args || args.condition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'condition'");
            }
            inputs["append"] = args ? args.append : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Macro.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Macro resources.
 */
export interface MacroState {
    /**
     * Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
     * The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
     * append macro called "foo" but not a second one. By default this is false.
     */
    append?: pulumi.Input<boolean>;
    /**
     * Macro condition. It can contain lists or other macros.
     */
    condition?: pulumi.Input<string>;
    /**
     * The name of the macro. It must be unique if it's not in append mode.
     */
    name?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Macro resource.
 */
export interface MacroArgs {
    /**
     * Adds these elements to an existing macro. Used to extend existing macros provided by Sysdig.
     * The macros can only be extended once, for example if there is an existing macro called "foo", one can have another
     * append macro called "foo" but not a second one. By default this is false.
     */
    append?: pulumi.Input<boolean>;
    /**
     * Macro condition. It can contain lists or other macros.
     */
    condition: pulumi.Input<string>;
    /**
     * The name of the macro. It must be unique if it's not in append mode.
     */
    name?: pulumi.Input<string>;
}
