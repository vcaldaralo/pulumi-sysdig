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

__all__ = ['PromqlArgs', 'Promql']

@pulumi.input_type
class PromqlArgs:
    def __init__(__self__, *,
                 promql: pulumi.Input[str],
                 trigger_after_minutes: pulumi.Input[int],
                 capture: Optional[pulumi.Input['PromqlCaptureArgs']] = None,
                 custom_notification: Optional[pulumi.Input['PromqlCustomNotificationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Promql resource.
        """
        pulumi.set(__self__, "promql", promql)
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
    def promql(self) -> pulumi.Input[str]:
        return pulumi.get(self, "promql")

    @promql.setter
    def promql(self, value: pulumi.Input[str]):
        pulumi.set(self, "promql", value)

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> pulumi.Input[int]:
        return pulumi.get(self, "trigger_after_minutes")

    @trigger_after_minutes.setter
    def trigger_after_minutes(self, value: pulumi.Input[int]):
        pulumi.set(self, "trigger_after_minutes", value)

    @property
    @pulumi.getter
    def capture(self) -> Optional[pulumi.Input['PromqlCaptureArgs']]:
        return pulumi.get(self, "capture")

    @capture.setter
    def capture(self, value: Optional[pulumi.Input['PromqlCaptureArgs']]):
        pulumi.set(self, "capture", value)

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> Optional[pulumi.Input['PromqlCustomNotificationArgs']]:
        return pulumi.get(self, "custom_notification")

    @custom_notification.setter
    def custom_notification(self, value: Optional[pulumi.Input['PromqlCustomNotificationArgs']]):
        pulumi.set(self, "custom_notification", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        return pulumi.get(self, "notification_channels")

    @notification_channels.setter
    def notification_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "notification_channels", value)

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "renotification_minutes")

    @renotification_minutes.setter
    def renotification_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renotification_minutes", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "severity", value)


@pulumi.input_type
class _PromqlState:
    def __init__(__self__, *,
                 capture: Optional[pulumi.Input['PromqlCaptureArgs']] = None,
                 custom_notification: Optional[pulumi.Input['PromqlCustomNotificationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 promql: Optional[pulumi.Input[str]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None,
                 team: Optional[pulumi.Input[int]] = None,
                 trigger_after_minutes: Optional[pulumi.Input[int]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Promql resources.
        """
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
        if promql is not None:
            pulumi.set(__self__, "promql", promql)
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
    def capture(self) -> Optional[pulumi.Input['PromqlCaptureArgs']]:
        return pulumi.get(self, "capture")

    @capture.setter
    def capture(self, value: Optional[pulumi.Input['PromqlCaptureArgs']]):
        pulumi.set(self, "capture", value)

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> Optional[pulumi.Input['PromqlCustomNotificationArgs']]:
        return pulumi.get(self, "custom_notification")

    @custom_notification.setter
    def custom_notification(self, value: Optional[pulumi.Input['PromqlCustomNotificationArgs']]):
        pulumi.set(self, "custom_notification", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        return pulumi.get(self, "notification_channels")

    @notification_channels.setter
    def notification_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "notification_channels", value)

    @property
    @pulumi.getter
    def promql(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "promql")

    @promql.setter
    def promql(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "promql", value)

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "renotification_minutes")

    @renotification_minutes.setter
    def renotification_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renotification_minutes", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def team(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "team", value)

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "trigger_after_minutes")

    @trigger_after_minutes.setter
    def trigger_after_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trigger_after_minutes", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class Promql(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capture: Optional[pulumi.Input[pulumi.InputType['PromqlCaptureArgs']]] = None,
                 custom_notification: Optional[pulumi.Input[pulumi.InputType['PromqlCustomNotificationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 promql: Optional[pulumi.Input[str]] = None,
                 renotification_minutes: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[int]] = None,
                 trigger_after_minutes: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Promql resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PromqlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Promql resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PromqlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PromqlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capture: Optional[pulumi.Input[pulumi.InputType['PromqlCaptureArgs']]] = None,
                 custom_notification: Optional[pulumi.Input[pulumi.InputType['PromqlCustomNotificationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 promql: Optional[pulumi.Input[str]] = None,
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
            __props__ = PromqlArgs.__new__(PromqlArgs)

            __props__.__dict__["capture"] = capture
            __props__.__dict__["custom_notification"] = custom_notification
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_channels"] = notification_channels
            if promql is None and not opts.urn:
                raise TypeError("Missing required property 'promql'")
            __props__.__dict__["promql"] = promql
            __props__.__dict__["renotification_minutes"] = renotification_minutes
            __props__.__dict__["scope"] = scope
            __props__.__dict__["severity"] = severity
            if trigger_after_minutes is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_after_minutes'")
            __props__.__dict__["trigger_after_minutes"] = trigger_after_minutes
            __props__.__dict__["team"] = None
            __props__.__dict__["version"] = None
        super(Promql, __self__).__init__(
            'sysdig:Monitor/promql:Promql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capture: Optional[pulumi.Input[pulumi.InputType['PromqlCaptureArgs']]] = None,
            custom_notification: Optional[pulumi.Input[pulumi.InputType['PromqlCustomNotificationArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            promql: Optional[pulumi.Input[str]] = None,
            renotification_minutes: Optional[pulumi.Input[int]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[int]] = None,
            team: Optional[pulumi.Input[int]] = None,
            trigger_after_minutes: Optional[pulumi.Input[int]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'Promql':
        """
        Get an existing Promql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PromqlState.__new__(_PromqlState)

        __props__.__dict__["capture"] = capture
        __props__.__dict__["custom_notification"] = custom_notification
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["notification_channels"] = notification_channels
        __props__.__dict__["promql"] = promql
        __props__.__dict__["renotification_minutes"] = renotification_minutes
        __props__.__dict__["scope"] = scope
        __props__.__dict__["severity"] = severity
        __props__.__dict__["team"] = team
        __props__.__dict__["trigger_after_minutes"] = trigger_after_minutes
        __props__.__dict__["version"] = version
        return Promql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capture(self) -> pulumi.Output[Optional['outputs.PromqlCapture']]:
        return pulumi.get(self, "capture")

    @property
    @pulumi.getter(name="customNotification")
    def custom_notification(self) -> pulumi.Output[Optional['outputs.PromqlCustomNotification']]:
        return pulumi.get(self, "custom_notification")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> pulumi.Output[Optional[Sequence[int]]]:
        return pulumi.get(self, "notification_channels")

    @property
    @pulumi.getter
    def promql(self) -> pulumi.Output[str]:
        return pulumi.get(self, "promql")

    @property
    @pulumi.getter(name="renotificationMinutes")
    def renotification_minutes(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "renotification_minutes")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def team(self) -> pulumi.Output[int]:
        return pulumi.get(self, "team")

    @property
    @pulumi.getter(name="triggerAfterMinutes")
    def trigger_after_minutes(self) -> pulumi.Output[int]:
        return pulumi.get(self, "trigger_after_minutes")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

