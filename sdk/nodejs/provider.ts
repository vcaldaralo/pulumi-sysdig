// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the sysdig package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'sysdig';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    public readonly sysdigMonitorApiToken!: pulumi.Output<string | undefined>;
    public readonly sysdigMonitorUrl!: pulumi.Output<string | undefined>;
    public readonly sysdigSecureApiToken!: pulumi.Output<string | undefined>;
    public readonly sysdigSecureUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            inputs["extraHeaders"] = pulumi.output(args ? args.extraHeaders : undefined).apply(JSON.stringify);
            inputs["sysdigMonitorApiToken"] = args ? args.sysdigMonitorApiToken : undefined;
            inputs["sysdigMonitorInsecureTls"] = pulumi.output(args ? args.sysdigMonitorInsecureTls : undefined).apply(JSON.stringify);
            inputs["sysdigMonitorUrl"] = args ? args.sysdigMonitorUrl : undefined;
            inputs["sysdigSecureApiToken"] = args ? args.sysdigSecureApiToken : undefined;
            inputs["sysdigSecureInsecureTls"] = pulumi.output(args ? args.sysdigSecureInsecureTls : undefined).apply(JSON.stringify);
            inputs["sysdigSecureUrl"] = args ? args.sysdigSecureUrl : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    extraHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    sysdigMonitorApiToken?: pulumi.Input<string>;
    sysdigMonitorInsecureTls?: pulumi.Input<boolean>;
    sysdigMonitorUrl?: pulumi.Input<string>;
    sysdigSecureApiToken?: pulumi.Input<string>;
    sysdigSecureInsecureTls?: pulumi.Input<boolean>;
    sysdigSecureUrl?: pulumi.Input<string>;
}
