# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AlertGroupOutlierArgs', 'AlertGroupOutlier']

@pulumi.input_type
class AlertGroupOutlierArgs:
    def __init__(__self__, *,
                 monitors: pulumi.Input[Sequence[pulumi.Input[str]]],
                 trigger_after_minutes: pulumi.Input[int],
                 capture: Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']] = None,
                 custom_notification: Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AlertGroupOutlier resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        :param pulumi.Input[int] trigger_after_minutes: Threshold of time for the status to stabilize until the alert is fired.
        :param pulumi.Input['AlertGroupOutlierCaptureArgs'] capture: Enables the creation of a capture file of the syscalls during the event.
        :param pulumi.Input['AlertGroupOutlierCustomNotificationArgs'] custom_notification: Allows to define a custom notification title, prepend and append text.
        :param pulumi.Input[str] description: The description of Monitor alert.
        :param pulumi.Input[bool] enabled: Boolean that defines if the alert is enabled or not. Defaults to true.
        :param pulumi.Input[str] name: The name of the Monitor alert. It must be unique.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] notification_channels: List of notification channel IDs where an alert must be sent to once fired.
        :param pulumi.Input[int] renotification_minutes: Number of minutes for the alert to re-notify until the status is solved.
        :param pulumi.Input[str] scope: Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        :param pulumi.Input[int] severity: Severity of the Monitor alert. It must be a value between 0 and 7,
               with 0 being the most critical and 7 the less critical. Defaults to 4.
        """
        pulumi.set(__self__, "monitors", monitors)
        pulumi.set(__self__, "trigger_after_minutes", trigger_after_minutes)
        if capture is not None:
            pulumi.set(__self__, "capture", capture)
        if custom_notification is not None:
            pulumi.set(__self__, "custom_notification", custom_notification)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_channels is not None:
            pulumi.set(__self__, "notification_channels", notification_channels)
        if renotification_minutes is not None:
            pulumi.set(__self__, "renotification_minutes", renotification_minutes)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)

    @property
    @pulumi.getter
    def monitors(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        """
        return pulumi.get(self, "monitors")

    @monitors.setter
    def monitors(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "monitors", value)

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> pulumi.Input[int]:
        """
        Threshold of time for the status to stabilize until the alert is fired.
        """
        return pulumi.get(self, "trigger_after_minutes")

    @trigger_after_minutes.setter
    def trigger_after_minutes(self, value: pulumi.Input[int]):
        pulumi.set(self, "trigger_after_minutes", value)

    @property
    @pulumi.getter
    def capture(self) -> Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']]:
        """
        Enables the creation of a capture file of the syscalls during the event.
        """
        return pulumi.get(self, "capture")

    @capture.setter
    def capture(self, value: Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']]):
        pulumi.set(self, "capture", value)

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']]:
        """
        Allows to define a custom notification title, prepend and append text.
        """
        return pulumi.get(self, "custom_notification")

    @custom_notification.setter
    def custom_notification(self, value: Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']]):
        pulumi.set(self, "custom_notification", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of Monitor alert.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean that defines if the alert is enabled or not. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Monitor alert. It must be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of notification channel IDs where an alert must be sent to once fired.
        """
        return pulumi.get(self, "notification_channels")

    @notification_channels.setter
    def notification_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "notification_channels", value)

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of minutes for the alert to re-notify until the status is solved.
        """
        return pulumi.get(self, "renotification_minutes")

    @renotification_minutes.setter
    def renotification_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renotification_minutes", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[int]]:
        """
        Severity of the Monitor alert. It must be a value between 0 and 7,
        with 0 being the most critical and 7 the less critical. Defaults to 4.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "severity", value)


@pulumi.input_type
class _AlertGroupOutlierState:
    def __init__(__self__, *,
                 capture: Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']] = None,
                 custom_notification: Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None,
                 team: Optional[pulumi.Input[int]] = None,
                 trigger_after_minutes: Optional[pulumi.Input[int]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AlertGroupOutlier resources.
        :param pulumi.Input['AlertGroupOutlierCaptureArgs'] capture: Enables the creation of a capture file of the syscalls during the event.
        :param pulumi.Input['AlertGroupOutlierCustomNotificationArgs'] custom_notification: Allows to define a custom notification title, prepend and append text.
        :param pulumi.Input[str] description: The description of Monitor alert.
        :param pulumi.Input[bool] enabled: Boolean that defines if the alert is enabled or not. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        :param pulumi.Input[str] name: The name of the Monitor alert. It must be unique.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] notification_channels: List of notification channel IDs where an alert must be sent to once fired.
        :param pulumi.Input[int] renotification_minutes: Number of minutes for the alert to re-notify until the status is solved.
        :param pulumi.Input[str] scope: Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        :param pulumi.Input[int] severity: Severity of the Monitor alert. It must be a value between 0 and 7,
               with 0 being the most critical and 7 the less critical. Defaults to 4.
        :param pulumi.Input[int] team: Team ID that owns the alert.
        :param pulumi.Input[int] trigger_after_minutes: Threshold of time for the status to stabilize until the alert is fired.
        :param pulumi.Input[int] version: Current version of the resource in Sysdig Monitor.
        """
        if capture is not None:
            pulumi.set(__self__, "capture", capture)
        if custom_notification is not None:
            pulumi.set(__self__, "custom_notification", custom_notification)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if monitors is not None:
            pulumi.set(__self__, "monitors", monitors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_channels is not None:
            pulumi.set(__self__, "notification_channels", notification_channels)
        if renotification_minutes is not None:
            pulumi.set(__self__, "renotification_minutes", renotification_minutes)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if team is not None:
            pulumi.set(__self__, "team", team)
        if trigger_after_minutes is not None:
            pulumi.set(__self__, "trigger_after_minutes", trigger_after_minutes)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def capture(self) -> Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']]:
        """
        Enables the creation of a capture file of the syscalls during the event.
        """
        return pulumi.get(self, "capture")

    @capture.setter
    def capture(self, value: Optional[pulumi.Input['AlertGroupOutlierCaptureArgs']]):
        pulumi.set(self, "capture", value)

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']]:
        """
        Allows to define a custom notification title, prepend and append text.
        """
        return pulumi.get(self, "custom_notification")

    @custom_notification.setter
    def custom_notification(self, value: Optional[pulumi.Input['AlertGroupOutlierCustomNotificationArgs']]):
        pulumi.set(self, "custom_notification", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of Monitor alert.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean that defines if the alert is enabled or not. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def monitors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        """
        return pulumi.get(self, "monitors")

    @monitors.setter
    def monitors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "monitors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Monitor alert. It must be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of notification channel IDs where an alert must be sent to once fired.
        """
        return pulumi.get(self, "notification_channels")

    @notification_channels.setter
    def notification_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "notification_channels", value)

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of minutes for the alert to re-notify until the status is solved.
        """
        return pulumi.get(self, "renotification_minutes")

    @renotification_minutes.setter
    def renotification_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renotification_minutes", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[int]]:
        """
        Severity of the Monitor alert. It must be a value between 0 and 7,
        with 0 being the most critical and 7 the less critical. Defaults to 4.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def team(self) -> Optional[pulumi.Input[int]]:
        """
        Team ID that owns the alert.
        """
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "team", value)

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Threshold of time for the status to stabilize until the alert is fired.
        """
        return pulumi.get(self, "trigger_after_minutes")

    @trigger_after_minutes.setter
    def trigger_after_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trigger_after_minutes", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Current version of the resource in Sysdig Monitor.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AlertGroupOutlier(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capture: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCaptureArgs']]] = None,
                 custom_notification: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCustomNotificationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None,
                 trigger_after_minutes: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Import

        Group Outlier Monitor alerts can be imported using the alert ID, e.g.

        ```sh
         $ pulumi import sysdig:Monitor/alertGroupOutlier:AlertGroupOutlier example 12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AlertGroupOutlierCaptureArgs']] capture: Enables the creation of a capture file of the syscalls during the event.
        :param pulumi.Input[pulumi.InputType['AlertGroupOutlierCustomNotificationArgs']] custom_notification: Allows to define a custom notification title, prepend and append text.
        :param pulumi.Input[str] description: The description of Monitor alert.
        :param pulumi.Input[bool] enabled: Boolean that defines if the alert is enabled or not. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        :param pulumi.Input[str] name: The name of the Monitor alert. It must be unique.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] notification_channels: List of notification channel IDs where an alert must be sent to once fired.
        :param pulumi.Input[int] renotification_minutes: Number of minutes for the alert to re-notify until the status is solved.
        :param pulumi.Input[str] scope: Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        :param pulumi.Input[int] severity: Severity of the Monitor alert. It must be a value between 0 and 7,
               with 0 being the most critical and 7 the less critical. Defaults to 4.
        :param pulumi.Input[int] trigger_after_minutes: Threshold of time for the status to stabilize until the alert is fired.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlertGroupOutlierArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Group Outlier Monitor alerts can be imported using the alert ID, e.g.

        ```sh
         $ pulumi import sysdig:Monitor/alertGroupOutlier:AlertGroupOutlier example 12345
        ```

        :param str resource_name: The name of the resource.
        :param AlertGroupOutlierArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlertGroupOutlierArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capture: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCaptureArgs']]] = None,
                 custom_notification: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCustomNotificationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None,
                 trigger_after_minutes: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlertGroupOutlierArgs.__new__(AlertGroupOutlierArgs)

            __props__.__dict__["capture"] = capture
            __props__.__dict__["custom_notification"] = custom_notification
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if monitors is None and not opts.urn:
                raise TypeError("Missing required property 'monitors'")
            __props__.__dict__["monitors"] = monitors
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_channels"] = notification_channels
            __props__.__dict__["renotification_minutes"] = renotification_minutes
            __props__.__dict__["scope"] = scope
            __props__.__dict__["severity"] = severity
            if trigger_after_minutes is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_after_minutes'")
            __props__.__dict__["trigger_after_minutes"] = trigger_after_minutes
            __props__.__dict__["team"] = None
            __props__.__dict__["version"] = None
        super(AlertGroupOutlier, __self__).__init__(
            'sysdig:Monitor/alertGroupOutlier:AlertGroupOutlier',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capture: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCaptureArgs']]] = None,
            custom_notification: Optional[pulumi.Input[pulumi.InputType['AlertGroupOutlierCustomNotificationArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            renotification_minutes: Optional[pulumi.Input[int]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[int]] = None,
            team: Optional[pulumi.Input[int]] = None,
            trigger_after_minutes: Optional[pulumi.Input[int]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AlertGroupOutlier':
        """
        Get an existing AlertGroupOutlier resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AlertGroupOutlierCaptureArgs']] capture: Enables the creation of a capture file of the syscalls during the event.
        :param pulumi.Input[pulumi.InputType['AlertGroupOutlierCustomNotificationArgs']] custom_notification: Allows to define a custom notification title, prepend and append text.
        :param pulumi.Input[str] description: The description of Monitor alert.
        :param pulumi.Input[bool] enabled: Boolean that defines if the alert is enabled or not. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        :param pulumi.Input[str] name: The name of the Monitor alert. It must be unique.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] notification_channels: List of notification channel IDs where an alert must be sent to once fired.
        :param pulumi.Input[int] renotification_minutes: Number of minutes for the alert to re-notify until the status is solved.
        :param pulumi.Input[str] scope: Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        :param pulumi.Input[int] severity: Severity of the Monitor alert. It must be a value between 0 and 7,
               with 0 being the most critical and 7 the less critical. Defaults to 4.
        :param pulumi.Input[int] team: Team ID that owns the alert.
        :param pulumi.Input[int] trigger_after_minutes: Threshold of time for the status to stabilize until the alert is fired.
        :param pulumi.Input[int] version: Current version of the resource in Sysdig Monitor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlertGroupOutlierState.__new__(_AlertGroupOutlierState)

        __props__.__dict__["capture"] = capture
        __props__.__dict__["custom_notification"] = custom_notification
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["monitors"] = monitors
        __props__.__dict__["name"] = name
        __props__.__dict__["notification_channels"] = notification_channels
        __props__.__dict__["renotification_minutes"] = renotification_minutes
        __props__.__dict__["scope"] = scope
        __props__.__dict__["severity"] = severity
        __props__.__dict__["team"] = team
        __props__.__dict__["trigger_after_minutes"] = trigger_after_minutes
        __props__.__dict__["version"] = version
        return AlertGroupOutlier(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capture(self) -> pulumi.Output[Optional['outputs.AlertGroupOutlierCapture']]:
        """
        Enables the creation of a capture file of the syscalls during the event.
        """
        return pulumi.get(self, "capture")

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> pulumi.Output[Optional['outputs.AlertGroupOutlierCustomNotification']]:
        """
        Allows to define a custom notification title, prepend and append text.
        """
        return pulumi.get(self, "custom_notification")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of Monitor alert.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean that defines if the alert is enabled or not. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def monitors(self) -> pulumi.Output[Sequence[str]]:
        """
        Array of metrics to monitor and alert on. Example: `["cpu.used.percent", "cpu.cores.used", "memory.bytes.used", "fs.used.percent", "thread.count", "net.request.count.in"]`.
        """
        return pulumi.get(self, "monitors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Monitor alert. It must be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        List of notification channel IDs where an alert must be sent to once fired.
        """
        return pulumi.get(self, "notification_channels")

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> pulumi.Output[Optional[int]]:
        """
        Number of minutes for the alert to re-notify until the status is solved.
        """
        return pulumi.get(self, "renotification_minutes")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        Part of the infrastructure where the alert is valid. Defaults to the entire infrastructure.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[Optional[int]]:
        """
        Severity of the Monitor alert. It must be a value between 0 and 7,
        with 0 being the most critical and 7 the less critical. Defaults to 4.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def team(self) -> pulumi.Output[int]:
        """
        Team ID that owns the alert.
        """
        return pulumi.get(self, "team")

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> pulumi.Output[int]:
        """
        Threshold of time for the status to stabilize until the alert is fired.
        """
        return pulumi.get(self, "trigger_after_minutes")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Current version of the resource in Sysdig Monitor.
        """
        return pulumi.get(self, "version")

