// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RuleSyscall extends pulumi.CustomResource {
    /**
     * Get an existing RuleSyscall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleSyscallState, opts?: pulumi.CustomResourceOptions): RuleSyscall {
        return new RuleSyscall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/ruleSyscall:RuleSyscall';

    /**
     * Returns true if the given object is an instance of RuleSyscall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleSyscall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleSyscall.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly matching!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly syscalls!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a RuleSyscall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleSyscallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleSyscallArgs | RuleSyscallState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleSyscallState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["matching"] = state ? state.matching : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["syscalls"] = state ? state.syscalls : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RuleSyscallArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["matching"] = args ? args.matching : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["syscalls"] = args ? args.syscalls : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RuleSyscall.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleSyscall resources.
 */
export interface RuleSyscallState {
    description?: pulumi.Input<string>;
    matching?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    syscalls?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleSyscall resource.
 */
export interface RuleSyscallArgs {
    description?: pulumi.Input<string>;
    matching?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    syscalls?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
