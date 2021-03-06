// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Secure container runtime rules can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Secure/ruleContainer:RuleContainer example 12345
 * ```
 */
export class RuleContainer extends pulumi.CustomResource {
    /**
     * Get an existing RuleContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleContainerState, opts?: pulumi.CustomResourceOptions): RuleContainer {
        return new RuleContainer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/ruleContainer:RuleContainer';

    /**
     * Returns true if the given object is an instance of RuleContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleContainer.__pulumiType;
    }

    /**
     * List of containers to match.
     */
    public readonly containers!: pulumi.Output<string[] | undefined>;
    /**
     * The description of Secure rule. By default is empty.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defines if the image name matches or not with the provided list. Default is true.
     */
    public readonly matching!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of tags for this rule.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Current version of the resource in Sysdig Secure.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a RuleContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleContainerArgs | RuleContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleContainerState | undefined;
            inputs["containers"] = state ? state.containers : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["matching"] = state ? state.matching : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RuleContainerArgs | undefined;
            inputs["containers"] = args ? args.containers : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["matching"] = args ? args.matching : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RuleContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleContainer resources.
 */
export interface RuleContainerState {
    /**
     * List of containers to match.
     */
    containers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of Secure rule. By default is empty.
     */
    description?: pulumi.Input<string>;
    /**
     * Defines if the image name matches or not with the provided list. Default is true.
     */
    matching?: pulumi.Input<boolean>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags for this rule.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Current version of the resource in Sysdig Secure.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleContainer resource.
 */
export interface RuleContainerArgs {
    /**
     * List of containers to match.
     */
    containers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of Secure rule. By default is empty.
     */
    description?: pulumi.Input<string>;
    /**
     * Defines if the image name matches or not with the provided list. Default is true.
     */
    matching?: pulumi.Input<boolean>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags for this rule.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
