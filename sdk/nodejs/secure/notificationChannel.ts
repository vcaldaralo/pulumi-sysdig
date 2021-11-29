// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function notificationChannel(args: NotificationChannelArgs, opts?: pulumi.InvokeOptions): Promise<NotificationChannelResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("sysdig:Secure/notificationChannel:NotificationChannel", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking NotificationChannel.
 */
export interface NotificationChannelArgs {
    name: string;
}

/**
 * A collection of values returned by NotificationChannel.
 */
export interface NotificationChannelResult {
    readonly account: string;
    readonly apiKey: string;
    readonly channel: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly notifyWhenOk: boolean;
    readonly notifyWhenResolved: boolean;
    readonly recipients: string;
    readonly routingKey: string;
    readonly sendTestNotification: boolean;
    readonly serviceKey: string;
    readonly serviceName: string;
    readonly topics: string;
    readonly type: string;
    readonly url: string;
    readonly version: number;
}

export function notificationChannelOutput(args: NotificationChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<NotificationChannelResult> {
    return pulumi.output(args).apply(a => notificationChannel(a, opts))
}

/**
 * A collection of arguments for invoking NotificationChannel.
 */
export interface NotificationChannelOutputArgs {
    name: pulumi.Input<string>;
}