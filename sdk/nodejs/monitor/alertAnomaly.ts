// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class AlertAnomaly extends pulumi.CustomResource {
    /**
     * Get an existing AlertAnomaly resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertAnomalyState, opts?: pulumi.CustomResourceOptions): AlertAnomaly {
        return new AlertAnomaly(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Monitor/alertAnomaly:AlertAnomaly';

    /**
     * Returns true if the given object is an instance of AlertAnomaly.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertAnomaly {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertAnomaly.__pulumiType;
    }

    public readonly capture!: pulumi.Output<outputs.Monitor.AlertAnomalyCapture | undefined>;
    public readonly customNotification!: pulumi.Output<outputs.Monitor.AlertAnomalyCustomNotification | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly monitors!: pulumi.Output<string[]>;
    public readonly multipleAlertsBies!: pulumi.Output<string[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly notificationChannels!: pulumi.Output<number[] | undefined>;
    public readonly renotificationMinutes!: pulumi.Output<number | undefined>;
    public readonly scope!: pulumi.Output<string | undefined>;
    public readonly severity!: pulumi.Output<number | undefined>;
    public /*out*/ readonly team!: pulumi.Output<number>;
    public readonly triggerAfterMinutes!: pulumi.Output<number>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a AlertAnomaly resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertAnomalyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertAnomalyArgs | AlertAnomalyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertAnomalyState | undefined;
            inputs["capture"] = state ? state.capture : undefined;
            inputs["customNotification"] = state ? state.customNotification : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["monitors"] = state ? state.monitors : undefined;
            inputs["multipleAlertsBies"] = state ? state.multipleAlertsBies : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationChannels"] = state ? state.notificationChannels : undefined;
            inputs["renotificationMinutes"] = state ? state.renotificationMinutes : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["severity"] = state ? state.severity : undefined;
            inputs["team"] = state ? state.team : undefined;
            inputs["triggerAfterMinutes"] = state ? state.triggerAfterMinutes : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AlertAnomalyArgs | undefined;
            if ((!args || args.monitors === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitors'");
            }
            if ((!args || args.triggerAfterMinutes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggerAfterMinutes'");
            }
            inputs["capture"] = args ? args.capture : undefined;
            inputs["customNotification"] = args ? args.customNotification : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["monitors"] = args ? args.monitors : undefined;
            inputs["multipleAlertsBies"] = args ? args.multipleAlertsBies : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationChannels"] = args ? args.notificationChannels : undefined;
            inputs["renotificationMinutes"] = args ? args.renotificationMinutes : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["severity"] = args ? args.severity : undefined;
            inputs["triggerAfterMinutes"] = args ? args.triggerAfterMinutes : undefined;
            inputs["team"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlertAnomaly.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertAnomaly resources.
 */
export interface AlertAnomalyState {
    capture?: pulumi.Input<inputs.Monitor.AlertAnomalyCapture>;
    customNotification?: pulumi.Input<inputs.Monitor.AlertAnomalyCustomNotification>;
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    monitors?: pulumi.Input<pulumi.Input<string>[]>;
    multipleAlertsBies?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    notificationChannels?: pulumi.Input<pulumi.Input<number>[]>;
    renotificationMinutes?: pulumi.Input<number>;
    scope?: pulumi.Input<string>;
    severity?: pulumi.Input<number>;
    team?: pulumi.Input<number>;
    triggerAfterMinutes?: pulumi.Input<number>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AlertAnomaly resource.
 */
export interface AlertAnomalyArgs {
    capture?: pulumi.Input<inputs.Monitor.AlertAnomalyCapture>;
    customNotification?: pulumi.Input<inputs.Monitor.AlertAnomalyCustomNotification>;
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    monitors: pulumi.Input<pulumi.Input<string>[]>;
    multipleAlertsBies?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    notificationChannels?: pulumi.Input<pulumi.Input<number>[]>;
    renotificationMinutes?: pulumi.Input<number>;
    scope?: pulumi.Input<string>;
    severity?: pulumi.Input<number>;
    triggerAfterMinutes: pulumi.Input<number>;
}