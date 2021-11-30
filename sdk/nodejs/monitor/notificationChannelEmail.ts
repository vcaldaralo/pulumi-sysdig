// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Email notification channels for Monitor can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import sysdig:Monitor/notificationChannelEmail:NotificationChannelEmail example 12345
 * ```
 */
export class NotificationChannelEmail extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannelEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationChannelEmailState, opts?: pulumi.CustomResourceOptions): NotificationChannelEmail {
        return new NotificationChannelEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Monitor/notificationChannelEmail:NotificationChannelEmail';

    /**
     * Returns true if the given object is an instance of NotificationChannelEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannelEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannelEmail.__pulumiType;
    }

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
     * List of recipients that will receive 
     * the message.
     */
    public readonly recipients!: pulumi.Output<string[]>;
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
     * Create a NotificationChannelEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationChannelEmailArgs | NotificationChannelEmailState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationChannelEmailState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifyWhenOk"] = state ? state.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = state ? state.notifyWhenResolved : undefined;
            inputs["recipients"] = state ? state.recipients : undefined;
            inputs["sendTestNotification"] = state ? state.sendTestNotification : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as NotificationChannelEmailArgs | undefined;
            if ((!args || args.recipients === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recipients'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifyWhenOk"] = args ? args.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = args ? args.notifyWhenResolved : undefined;
            inputs["recipients"] = args ? args.recipients : undefined;
            inputs["sendTestNotification"] = args ? args.sendTestNotification : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationChannelEmail.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationChannelEmail resources.
 */
export interface NotificationChannelEmailState {
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
     * List of recipients that will receive 
     * the message.
     */
    recipients?: pulumi.Input<pulumi.Input<string>[]>;
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
 * The set of arguments for constructing a NotificationChannelEmail resource.
 */
export interface NotificationChannelEmailArgs {
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
     * List of recipients that will receive 
     * the message.
     */
    recipients: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Send an initial test notification to check
     * if the notification channel is working. Default is false.
     */
    sendTestNotification?: pulumi.Input<boolean>;
}
