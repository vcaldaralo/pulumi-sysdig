// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Monitor dashboards can be imported using the dashboard ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Monitor/dashboard:Dashboard example 12345
 * ```
 *
 *  Only dashboards that contain supported panels can be imported. Currently supported panel types are- PromQL timecharts - PromQL numbers - Text Only dashboards that contain supported query types can be imported. Currently supported query types- Percent - Data - Data rate - Number - Number rate - Time
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Monitor/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * Description of the panel.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the panel.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * At least 1 panel is required to define a Dashboard.
     */
    public readonly panels!: pulumi.Output<outputs.Monitor.DashboardPanel[]>;
    /**
     * Define if the dashboard can be accessible without requiring the user to be logged in.
     */
    public readonly public!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) Token defined when the dashboard is set Public.
     */
    public /*out*/ readonly publicToken!: pulumi.Output<string>;
    /**
     * Define the scope of the dashboard and variables for these metrics.
     */
    public readonly scopes!: pulumi.Output<outputs.Monitor.DashboardScope[] | undefined>;
    /**
     * (Computed)  The current version of the Dashboard.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["panels"] = state ? state.panels : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["publicToken"] = state ? state.publicToken : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
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
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * Description of the panel.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the panel.
     */
    name?: pulumi.Input<string>;
    /**
     * At least 1 panel is required to define a Dashboard.
     */
    panels?: pulumi.Input<pulumi.Input<inputs.Monitor.DashboardPanel>[]>;
    /**
     * Define if the dashboard can be accessible without requiring the user to be logged in.
     */
    public?: pulumi.Input<boolean>;
    /**
     * (Computed) Token defined when the dashboard is set Public.
     */
    publicToken?: pulumi.Input<string>;
    /**
     * Define the scope of the dashboard and variables for these metrics.
     */
    scopes?: pulumi.Input<pulumi.Input<inputs.Monitor.DashboardScope>[]>;
    /**
     * (Computed)  The current version of the Dashboard.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * Description of the panel.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the panel.
     */
    name?: pulumi.Input<string>;
    /**
     * At least 1 panel is required to define a Dashboard.
     */
    panels: pulumi.Input<pulumi.Input<inputs.Monitor.DashboardPanel>[]>;
    /**
     * Define if the dashboard can be accessible without requiring the user to be logged in.
     */
    public?: pulumi.Input<boolean>;
    /**
     * Define the scope of the dashboard and variables for these metrics.
     */
    scopes?: pulumi.Input<pulumi.Input<inputs.Monitor.DashboardScope>[]>;
}
