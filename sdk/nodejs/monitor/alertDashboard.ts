// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class AlertDashboard extends pulumi.CustomResource {
    /**
     * Get an existing AlertDashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertDashboardState, opts?: pulumi.CustomResourceOptions): AlertDashboard {
        return new AlertDashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Monitor/alertDashboard:AlertDashboard';

    /**
     * Returns true if the given object is an instance of AlertDashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertDashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertDashboard.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly panels!: pulumi.Output<outputs.Monitor.AlertDashboardPanel[]>;
    public readonly public!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly publicToken!: pulumi.Output<string>;
    public readonly scopes!: pulumi.Output<outputs.Monitor.AlertDashboardScope[] | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a AlertDashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertDashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertDashboardArgs | AlertDashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertDashboardState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["panels"] = state ? state.panels : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["publicToken"] = state ? state.publicToken : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AlertDashboardArgs | undefined;
            if ((!args || args.panels === undefined) && !opts.urn) {
                throw new Error("Missing required property 'panels'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["panels"] = args ? args.panels : undefined;
            inputs["public"] = args ? args.public : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["publicToken"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlertDashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertDashboard resources.
 */
export interface AlertDashboardState {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    panels?: pulumi.Input<pulumi.Input<inputs.Monitor.AlertDashboardPanel>[]>;
    public?: pulumi.Input<boolean>;
    publicToken?: pulumi.Input<string>;
    scopes?: pulumi.Input<pulumi.Input<inputs.Monitor.AlertDashboardScope>[]>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AlertDashboard resource.
 */
export interface AlertDashboardArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    panels: pulumi.Input<pulumi.Input<inputs.Monitor.AlertDashboardPanel>[]>;
    public?: pulumi.Input<boolean>;
    scopes?: pulumi.Input<pulumi.Input<inputs.Monitor.AlertDashboardScope>[]>;
}
