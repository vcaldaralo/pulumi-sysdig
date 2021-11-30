// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * VictorOPS notification channels for Monitor can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Monitor/notificationChannelVictorops:NotificationChannelVictorops example 12345
 * ```
 */
export class NotificationChannelVictorops extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannelVictorops resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationChannelVictoropsState, opts?: pulumi.CustomResourceOptions): NotificationChannelVictorops {
        return new NotificationChannelVictorops(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Monitor/notificationChannelVictorops:NotificationChannelVictorops';

    /**
     * Returns true if the given object is an instance of NotificationChannelVictorops.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannelVictorops {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannelVictorops.__pulumiType;
    }

    /**
     * Key for the API.
     */
    public readonly apiKey!: pulumi.Output<string>;
    /**
     * If false, the channel will not emit notifications. Default is true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Notification Channel. Must be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Send a new notification when the alert condition is 
     * no longer triggered. Default is false.
     */
    public readonly notifyWhenOk!: pulumi.Output<boolean | undefined>;
    /**
     * Send a new notification when the alert is manually 
     * acknowledged by a user. Default is false.
     */
    public readonly notifyWhenResolved!: pulumi.Output<boolean | undefined>;
    /**
     * Routing key for VictorOps.
     */
    public readonly routingKey!: pulumi.Output<string>;
    /**
     * Send an initial test notification to check
     * if the notification channel is working. Default is false.
     */
    public readonly sendTestNotification!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) The current version of the Notification Channel.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a NotificationChannelVictorops resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelVictoropsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationChannelVictoropsArgs | NotificationChannelVictoropsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationChannelVictoropsState | undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifyWhenOk"] = state ? state.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = state ? state.notifyWhenResolved : undefined;
            inputs["routingKey"] = state ? state.routingKey : undefined;
            inputs["sendTestNotification"] = state ? state.sendTestNotification : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as NotificationChannelVictoropsArgs | undefined;
            if ((!args || args.apiKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKey'");
            }
            if ((!args || args.routingKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routingKey'");
            }
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifyWhenOk"] = args ? args.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = args ? args.notifyWhenResolved : undefined;
            inputs["routingKey"] = args ? args.routingKey : undefined;
            inputs["sendTestNotification"] = args ? args.sendTestNotification : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationChannelVictorops.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationChannelVictorops resources.
 */
export interface NotificationChannelVictoropsState {
    /**
     * Key for the API.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * If false, the channel will not emit notifications. Default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the Notification Channel. Must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Send a new notification when the alert condition is 
     * no longer triggered. Default is false.
     */
    notifyWhenOk?: pulumi.Input<boolean>;
    /**
     * Send a new notification when the alert is manually 
     * acknowledged by a user. Default is false.
     */
    notifyWhenResolved?: pulumi.Input<boolean>;
    /**
     * Routing key for VictorOps.
     */
    routingKey?: pulumi.Input<string>;
    /**
     * Send an initial test notification to check
     * if the notification channel is working. Default is false.
     */
    sendTestNotification?: pulumi.Input<boolean>;
    /**
     * (Computed) The current version of the Notification Channel.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NotificationChannelVictorops resource.
 */
export interface NotificationChannelVictoropsArgs {
    /**
     * Key for the API.
     */
    apiKey: pulumi.Input<string>;
    /**
     * If false, the channel will not emit notifications. Default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the Notification Channel. Must be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Send a new notification when the alert condition is 
     * no longer triggered. Default is false.
     */
    notifyWhenOk?: pulumi.Input<boolean>;
    /**
     * Send a new notification when the alert is manually 
     * acknowledged by a user. Default is false.
     */
    notifyWhenResolved?: pulumi.Input<boolean>;
    /**
     * Routing key for VictorOps.
     */
    routingKey: pulumi.Input<string>;
    /**
     * Send an initial test notification to check
     * if the notification channel is working. Default is false.
     */
    sendTestNotification?: pulumi.Input<boolean>;
}
