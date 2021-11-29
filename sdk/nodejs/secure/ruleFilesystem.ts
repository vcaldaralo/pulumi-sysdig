// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class RuleFilesystem extends pulumi.CustomResource {
    /**
     * Get an existing RuleFilesystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleFilesystemState, opts?: pulumi.CustomResourceOptions): RuleFilesystem {
        return new RuleFilesystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/ruleFilesystem:RuleFilesystem';

    /**
     * Returns true if the given object is an instance of RuleFilesystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleFilesystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleFilesystem.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly readOnlies!: pulumi.Output<outputs.Secure.RuleFilesystemReadOnly[] | undefined>;
    public readonly readWrites!: pulumi.Output<outputs.Secure.RuleFilesystemReadWrite[] | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a RuleFilesystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleFilesystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleFilesystemArgs | RuleFilesystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleFilesystemState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["readOnlies"] = state ? state.readOnlies : undefined;
            inputs["readWrites"] = state ? state.readWrites : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RuleFilesystemArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["readOnlies"] = args ? args.readOnlies : undefined;
            inputs["readWrites"] = args ? args.readWrites : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RuleFilesystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleFilesystem resources.
 */
export interface RuleFilesystemState {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    readOnlies?: pulumi.Input<pulumi.Input<inputs.Secure.RuleFilesystemReadOnly>[]>;
    readWrites?: pulumi.Input<pulumi.Input<inputs.Secure.RuleFilesystemReadWrite>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleFilesystem resource.
 */
export interface RuleFilesystemArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    readOnlies?: pulumi.Input<pulumi.Input<inputs.Secure.RuleFilesystemReadOnly>[]>;
    readWrites?: pulumi.Input<pulumi.Input<inputs.Secure.RuleFilesystemReadWrite>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
