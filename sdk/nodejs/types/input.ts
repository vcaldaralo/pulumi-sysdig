// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export namespace Monitor {
    export interface AlertAnomalyCapture {
        /**
         * Time frame in seconds of the capture.
         */
        duration: pulumi.Input<number>;
        /**
         * Defines the name of the capture file.
         */
        filename: pulumi.Input<string>;
        /**
         * Additional filter to apply to the capture. For example: `proc.name contains nginx`.
         */
        filter?: pulumi.Input<string>;
    }

    export interface AlertAnomalyCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface AlertDowntimeCapture {
        /**
         * Time frame in seconds of the capture.
         */
        duration: pulumi.Input<number>;
        /**
         * Defines the name of the capture file.
         */
        filename: pulumi.Input<string>;
        /**
         * Additional filter to apply to the capture. For example: `proc.name contains nginx`.
         */
        filter?: pulumi.Input<string>;
    }

    export interface AlertDowntimeCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface AlertEventCapture {
        /**
         * Time frame in seconds of the capture.
         */
        duration: pulumi.Input<number>;
        /**
         * Defines the name of the capture file.
         */
        filename: pulumi.Input<string>;
        /**
         * Additional filter to apply to the capture. For example: `proc.name contains nginx`.
         */
        filter?: pulumi.Input<string>;
    }

    export interface AlertEventCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface AlertGroupOutlierCapture {
        /**
         * Time frame in seconds of the capture.
         */
        duration: pulumi.Input<number>;
        /**
         * Defines the name of the capture file.
         */
        filename: pulumi.Input<string>;
        /**
         * Additional filter to apply to the capture. For example: `proc.name contains nginx`.
         */
        filter?: pulumi.Input<string>;
    }

    export interface AlertGroupOutlierCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface AlertMetricCapture {
        /**
         * Time frame in seconds of the capture.
         */
        duration: pulumi.Input<number>;
        /**
         * Defines the name of the capture file.
         */
        filename: pulumi.Input<string>;
        /**
         * Additional filter to apply to the capture. For example: `proc.name contains nginx`.
         */
        filter?: pulumi.Input<string>;
    }

    export interface AlertMetricCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface AlertPromqlCapture {
        duration: pulumi.Input<number>;
        filename: pulumi.Input<string>;
        filter?: pulumi.Input<string>;
    }

    export interface AlertPromqlCustomNotification {
        /**
         * Text to add after the alert template.
         */
        append?: pulumi.Input<string>;
        /**
         * Text to add before the alert template.
         */
        prepend?: pulumi.Input<string>;
        /**
         * Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
         */
        title: pulumi.Input<string>;
    }

    export interface DashboardPanel {
        /**
         * If true, the text will be autosized in the panel.
         * This field is ignored for all panel types except `text`.
         */
        autosizeText?: pulumi.Input<boolean>;
        /**
         * This field is required if the panel type is `text`. It represents the 
         * text that will be displayed in the panel.
         */
        content?: pulumi.Input<string>;
        /**
         * Description of the panel.
         */
        description?: pulumi.Input<string>;
        /**
         * Height of the panel. Min value: 1.
         */
        height: pulumi.Input<number>;
        /**
         * Name of the panel.
         */
        name: pulumi.Input<string>;
        /**
         * Position of the panel in the X axis. Min value: 0, max value: 23.
         */
        posX: pulumi.Input<number>;
        /**
         * Position of the panel in the Y axis. Min value: 0.
         */
        posY: pulumi.Input<number>;
        /**
         * The PromQL query that will show information in the panel. 
         * If the type of the panel is `timechart`, then it can be specified multiple
         * times, to have multiple metrics in the same graph.
         * If the type of the panel is `number` then only one can be specified.
         * This field is required if the panel type is `timechart` or `number`.
         */
        queries?: pulumi.Input<pulumi.Input<inputs.Monitor.DashboardPanelQuery>[]>;
        /**
         * If true, the panel will have a transparent background.
         * This field is ignored for all panel types except `text`.
         */
        transparentBackground?: pulumi.Input<boolean>;
        /**
         * Kind of panel, must be either `timechart`, `number` or `text`.
         */
        type: pulumi.Input<string>;
        /**
         * If true, the title of the panel will be displayed. Default: false.
         * This field is ignored for all panel types except `text`.
         */
        visibleTitle?: pulumi.Input<boolean>;
        /**
         * Width of the panel. Min value: 1, max value: 24.
         */
        width: pulumi.Input<number>;
    }

    export interface DashboardPanelQuery {
        /**
         * The PromQL query. Must be a valid PromQL query with existing
         * metrics in Sysdig Monitor.
         */
        promql: pulumi.Input<string>;
        /**
         * The type of metric for this query. Can be one of: `percent`, `data`, `data rate`, 
         * `number`, `number rate`, `time`.
         */
        unit: pulumi.Input<string>;
    }

    export interface DashboardScope {
        /**
         * Operator to relate the metric with some value. It is only required if the value to filter by is set, or the variable field is not set. Valid values are: `in`, `notIn`, `equals`, `notEquals`, `contains`, `notContains` and `startsWith`.
         */
        comparator?: pulumi.Input<string>;
        /**
         * Metric to scope by, common examples are `host.hostName`, `kubernetes.namespace.name` or `kubernetes.cluster.name`, but you can use all the Sysdig-supported values shown in the UI. Note that kubernetes-related values only appear when Sysdig detects Kubernetes metadata.
         */
        metric: pulumi.Input<string>;
        /**
         * List of values to filter by, if comparator is set. If the comparator is not `in` or `notIn` the list must contain only 1 value.
         */
        values?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Assigns this metric to a value name and allows PromQL to reference it.
         */
        variable?: pulumi.Input<string>;
    }

    export interface TeamEntrypoint {
        /**
         * Sets up the defined Dashboard name as entrypoint.
         * Warning: This field must only be added if the `type` is "Dashboards".
         */
        selection?: pulumi.Input<string>;
        /**
         * Main entrypoint for the team.
         * Valid options are: Explore, Dashboards, Events, Alerts, Settings.
         */
        type: pulumi.Input<string>;
    }

    export interface TeamUserRole {
        /**
         * The email of the user in the group.
         */
        email: pulumi.Input<string>;
        /**
         * The role for the user in this group.
         * Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
         * Default: ROLE_TEAM_STANDARD.
         */
        role?: pulumi.Input<string>;
    }
}

export namespace Secure {
    export interface PolicyAction {
        /**
         * Captures with Sysdig the stream of system calls:
         */
        captures?: pulumi.Input<pulumi.Input<inputs.Secure.PolicyActionCapture>[]>;
        /**
         * The action applied to container when this Policy is
         * triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
         * no action will be applied at the container level.
         */
        container?: pulumi.Input<string>;
    }

    export interface PolicyActionCapture {
        /**
         * Captures the system calls for the amount
         * of seconds after the policy was triggered.
         */
        secondsAfterEvent: pulumi.Input<number>;
        /**
         * Captures the system calls during the
         * amount of seconds before the policy was triggered.
         */
        secondsBeforeEvent: pulumi.Input<number>;
    }

    export interface RuleFalcoException {
        /**
         * Contains comparison operators that align 1-1 with the items in the fields property.
         */
        comps: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Contains one or more fields that will extract a value from the syscall/k8s_audit events.
         */
        fields: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
         */
        name: pulumi.Input<string>;
        /**
         * Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
         */
        values: pulumi.Input<string>;
    }

    export interface RuleFilesystemReadOnly {
        /**
         * Defines if the path matches or not with the provided list. Default is true.
         */
        matching?: pulumi.Input<boolean>;
        /**
         * List of paths to match.
         */
        paths: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface RuleFilesystemReadWrite {
        /**
         * Defines if the path matches or not with the provided list. Default is true.
         */
        matching?: pulumi.Input<boolean>;
        /**
         * List of paths to match.
         */
        paths: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface RuleNetworkTcp {
        /**
         * Defines if the port matches or not with the provided list. Default is true.
         */
        matching?: pulumi.Input<boolean>;
        /**
         * List of ports to match.
         */
        ports: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface RuleNetworkUdp {
        /**
         * Defines if the port matches or not with the provided list. Default is true.
         */
        matching?: pulumi.Input<boolean>;
        /**
         * List of ports to match.
         */
        ports: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface TeamUserRole {
        /**
         * The email of the user in the group.
         */
        email: pulumi.Input<string>;
        /**
         * The role for the user in this group.
         * Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
         * Default: ROLE_TEAM_STANDARD.
         */
        role?: pulumi.Input<string>;
    }
}
