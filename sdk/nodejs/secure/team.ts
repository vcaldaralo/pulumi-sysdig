// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class Team extends pulumi.CustomResource {
    /**
     * Get an existing Team resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamState, opts?: pulumi.CustomResourceOptions): Team {
        return new Team(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/team:Team';

    /**
     * Returns true if the given object is an instance of Team.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Team {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Team.__pulumiType;
    }

    public readonly defaultTeam!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly filter!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly scopeBy!: pulumi.Output<string | undefined>;
    public readonly theme!: pulumi.Output<string | undefined>;
    public readonly useSysdigCapture!: pulumi.Output<boolean | undefined>;
    public readonly userRoles!: pulumi.Output<outputs.Secure.TeamUserRole[] | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Team resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TeamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamArgs | TeamState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamState | undefined;
            inputs["defaultTeam"] = state ? state.defaultTeam : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scopeBy"] = state ? state.scopeBy : undefined;
            inputs["theme"] = state ? state.theme : undefined;
            inputs["useSysdigCapture"] = state ? state.useSysdigCapture : undefined;
            inputs["userRoles"] = state ? state.userRoles : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as TeamArgs | undefined;
            inputs["defaultTeam"] = args ? args.defaultTeam : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scopeBy"] = args ? args.scopeBy : undefined;
            inputs["theme"] = args ? args.theme : undefined;
            inputs["useSysdigCapture"] = args ? args.useSysdigCapture : undefined;
            inputs["userRoles"] = args ? args.userRoles : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Team.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Team resources.
 */
export interface TeamState {
    defaultTeam?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    filter?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    scopeBy?: pulumi.Input<string>;
    theme?: pulumi.Input<string>;
    useSysdigCapture?: pulumi.Input<boolean>;
    userRoles?: pulumi.Input<pulumi.Input<inputs.Secure.TeamUserRole>[]>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Team resource.
 */
export interface TeamArgs {
    defaultTeam?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    filter?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    scopeBy?: pulumi.Input<string>;
    theme?: pulumi.Input<string>;
    useSysdigCapture?: pulumi.Input<boolean>;
    userRoles?: pulumi.Input<pulumi.Input<inputs.Secure.TeamUserRole>[]>;
}