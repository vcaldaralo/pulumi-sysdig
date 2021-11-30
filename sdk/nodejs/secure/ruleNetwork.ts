// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Secure network runtime rules can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Secure/ruleNetwork:RuleNetwork example 12345
 * ```
 */
export class RuleNetwork extends pulumi.CustomResource {
    /**
     * Get an existing RuleNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleNetworkState, opts?: pulumi.CustomResourceOptions): RuleNetwork {
        return new RuleNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/ruleNetwork:RuleNetwork';

    /**
     * Returns true if the given object is an instance of RuleNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleNetwork.__pulumiType;
    }

    /**
     * Detect if there is an inbound connection.
     */
    public readonly blockInbound!: pulumi.Output<boolean>;
    /**
     * Detect if there is an outbound connection.
     */
    public readonly blockOutbound!: pulumi.Output<boolean>;
    /**
     * The description of Secure rule. By default is empty.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of tags for this rule.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tcps!: pulumi.Output<outputs.Secure.RuleNetworkTcp[] | undefined>;
    public readonly udps!: pulumi.Output<outputs.Secure.RuleNetworkUdp[] | undefined>;
    /**
     * Current version of the resource in Sysdig Secure.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a RuleNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleNetworkArgs | RuleNetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleNetworkState | undefined;
            inputs["blockInbound"] = state ? state.blockInbound : undefined;
            inputs["blockOutbound"] = state ? state.blockOutbound : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tcps"] = state ? state.tcps : undefined;
            inputs["udps"] = state ? state.udps : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RuleNetworkArgs | undefined;
            if ((!args || args.blockInbound === undefined) && !opts.urn) {
                throw new Error("Missing required property 'blockInbound'");
            }
            if ((!args || args.blockOutbound === undefined) && !opts.urn) {
                throw new Error("Missing required property 'blockOutbound'");
            }
            inputs["blockInbound"] = args ? args.blockInbound : undefined;
            inputs["blockOutbound"] = args ? args.blockOutbound : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tcps"] = args ? args.tcps : undefined;
            inputs["udps"] = args ? args.udps : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RuleNetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleNetwork resources.
 */
export interface RuleNetworkState {
    /**
     * Detect if there is an inbound connection.
     */
    blockInbound?: pulumi.Input<boolean>;
    /**
     * Detect if there is an outbound connection.
     */
    blockOutbound?: pulumi.Input<boolean>;
    /**
     * The description of Secure rule. By default is empty.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags for this rule.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tcps?: pulumi.Input<pulumi.Input<inputs.Secure.RuleNetworkTcp>[]>;
    udps?: pulumi.Input<pulumi.Input<inputs.Secure.RuleNetworkUdp>[]>;
    /**
     * Current version of the resource in Sysdig Secure.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleNetwork resource.
 */
export interface RuleNetworkArgs {
    /**
     * Detect if there is an inbound connection.
     */
    blockInbound: pulumi.Input<boolean>;
    /**
     * Detect if there is an outbound connection.
     */
    blockOutbound: pulumi.Input<boolean>;
    /**
     * The description of Secure rule. By default is empty.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the Secure rule. It must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags for this rule.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tcps?: pulumi.Input<pulumi.Input<inputs.Secure.RuleNetworkTcp>[]>;
    udps?: pulumi.Input<pulumi.Input<inputs.Secure.RuleNetworkUdp>[]>;
}
