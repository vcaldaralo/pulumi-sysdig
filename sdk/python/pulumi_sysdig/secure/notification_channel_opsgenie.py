# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NotificationChannelOpsgenieArgs', 'NotificationChannelOpsgenie']

@pulumi.input_type
class NotificationChannelOpsgenieArgs:
    def __init__(__self__, *,
                 api_key: pulumi.Input[str],
                 enabled: pulumi.Input[bool],
                 notify_when_ok: pulumi.Input[bool],
                 notify_when_resolved: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NotificationChannelOpsgenie resource.
        :param pulumi.Input[str] api_key: Key for the API.
        :param pulumi.Input[bool] enabled: If false, the channel will not emit notifications. Default is true.
        :param pulumi.Input[bool] notify_when_ok: Send a new notification when the alert condition is 
               no longer triggered. Default is false.
        :param pulumi.Input[bool] notify_when_resolved: Send a new notification when the alert is manually 
               acknowledged by a user. Default is false.
        :param pulumi.Input[str] name: The name of the Notification Channel. Must be unique.
        :param pulumi.Input[bool] send_test_notification: Send an initial test notification to check
               if the notification channel is working. Default is false.
        """
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "notify_when_ok", notify_when_ok)
        pulumi.set(__self__, "notify_when_resolved", notify_when_resolved)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if send_test_notification is not None:
            pulumi.set(__self__, "send_test_notification", send_test_notification)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        """
        Key for the API.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        If false, the channel will not emit notifications. Default is true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> pulumi.Input[bool]:
        """
        Send a new notification when the alert condition is 
        no longer triggered. Default is false.
        """
        return pulumi.get(self, "notify_when_ok")

    @notify_when_ok.setter
    def notify_when_ok(self, value: pulumi.Input[bool]):
        pulumi.set(self, "notify_when_ok", value)

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> pulumi.Input[bool]:
        """
        Send a new notification when the alert is manually 
        acknowledged by a user. Default is false.
        """
        return pulumi.get(self, "notify_when_resolved")

    @notify_when_resolved.setter
    def notify_when_resolved(self, value: pulumi.Input[bool]):
        pulumi.set(self, "notify_when_resolved", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Notification Channel. Must be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Send an initial test notification to check
        if the notification channel is working. Default is false.
        """
        return pulumi.get(self, "send_test_notification")

    @send_test_notification.setter
    def send_test_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_test_notification", value)


@pulumi.input_type
class _NotificationChannelOpsgenieState:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering NotificationChannelOpsgenie resources.
        :param pulumi.Input[str] api_key: Key for the API.
        :param pulumi.Input[bool] enabled: If false, the channel will not emit notifications. Default is true.
        :param pulumi.Input[str] name: The name of the Notification Channel. Must be unique.
        :param pulumi.Input[bool] notify_when_ok: Send a new notification when the alert condition is 
               no longer triggered. Default is false.
        :param pulumi.Input[bool] notify_when_resolved: Send a new notification when the alert is manually 
               acknowledged by a user. Default is false.
        :param pulumi.Input[bool] send_test_notification: Send an initial test notification to check
               if the notification channel is working. Default is false.
        :param pulumi.Input[int] version: (Computed) The current version of the Notification Channel.
        """
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notify_when_ok is not None:
            pulumi.set(__self__, "notify_when_ok", notify_when_ok)
        if notify_when_resolved is not None:
            pulumi.set(__self__, "notify_when_resolved", notify_when_resolved)
        if send_test_notification is not None:
            pulumi.set(__self__, "send_test_notification", send_test_notification)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key for the API.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, the channel will not emit notifications. Default is true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Notification Channel. Must be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> Optional[pulumi.Input[bool]]:
        """
        Send a new notification when the alert condition is 
        no longer triggered. Default is false.
        """
        return pulumi.get(self, "notify_when_ok")

    @notify_when_ok.setter
    def notify_when_ok(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_ok", value)

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> Optional[pulumi.Input[bool]]:
        """
        Send a new notification when the alert is manually 
        acknowledged by a user. Default is false.
        """
        return pulumi.get(self, "notify_when_resolved")

    @notify_when_resolved.setter
    def notify_when_resolved(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_resolved", value)

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Send an initial test notification to check
        if the notification channel is working. Default is false.
        """
        return pulumi.get(self, "send_test_notification")

    @send_test_notification.setter
    def send_test_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_test_notification", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        (Computed) The current version of the Notification Channel.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class NotificationChannelOpsgenie(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Import

        Opsgenie notification channels for Secure can be imported using the ID, e.g.

        ```sh
         $ pulumi import sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie example 12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: Key for the API.
        :param pulumi.Input[bool] enabled: If false, the channel will not emit notifications. Default is true.
        :param pulumi.Input[str] name: The name of the Notification Channel. Must be unique.
        :param pulumi.Input[bool] notify_when_ok: Send a new notification when the alert condition is 
               no longer triggered. Default is false.
        :param pulumi.Input[bool] notify_when_resolved: Send a new notification when the alert is manually 
               acknowledged by a user. Default is false.
        :param pulumi.Input[bool] send_test_notification: Send an initial test notification to check
               if the notification channel is working. Default is false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationChannelOpsgenieArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Opsgenie notification channels for Secure can be imported using the ID, e.g.

        ```sh
         $ pulumi import sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie example 12345
        ```

        :param str resource_name: The name of the resource.
        :param NotificationChannelOpsgenieArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationChannelOpsgenieArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
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
            __props__ = NotificationChannelOpsgenieArgs.__new__(NotificationChannelOpsgenieArgs)

            if api_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_key'")
            __props__.__dict__["api_key"] = api_key
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            if notify_when_ok is None and not opts.urn:
                raise TypeError("Missing required property 'notify_when_ok'")
            __props__.__dict__["notify_when_ok"] = notify_when_ok
            if notify_when_resolved is None and not opts.urn:
                raise TypeError("Missing required property 'notify_when_resolved'")
            __props__.__dict__["notify_when_resolved"] = notify_when_resolved
            __props__.__dict__["send_test_notification"] = send_test_notification
            __props__.__dict__["version"] = None
        super(NotificationChannelOpsgenie, __self__).__init__(
            'sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notify_when_ok: Optional[pulumi.Input[bool]] = None,
            notify_when_resolved: Optional[pulumi.Input[bool]] = None,
            send_test_notification: Optional[pulumi.Input[bool]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'NotificationChannelOpsgenie':
        """
        Get an existing NotificationChannelOpsgenie resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: Key for the API.
        :param pulumi.Input[bool] enabled: If false, the channel will not emit notifications. Default is true.
        :param pulumi.Input[str] name: The name of the Notification Channel. Must be unique.
        :param pulumi.Input[bool] notify_when_ok: Send a new notification when the alert condition is 
               no longer triggered. Default is false.
        :param pulumi.Input[bool] notify_when_resolved: Send a new notification when the alert is manually 
               acknowledged by a user. Default is false.
        :param pulumi.Input[bool] send_test_notification: Send an initial test notification to check
               if the notification channel is working. Default is false.
        :param pulumi.Input[int] version: (Computed) The current version of the Notification Channel.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationChannelOpsgenieState.__new__(_NotificationChannelOpsgenieState)

        __props__.__dict__["api_key"] = api_key
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["notify_when_ok"] = notify_when_ok
        __props__.__dict__["notify_when_resolved"] = notify_when_resolved
        __props__.__dict__["send_test_notification"] = send_test_notification
        __props__.__dict__["version"] = version
        return NotificationChannelOpsgenie(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        Key for the API.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        If false, the channel will not emit notifications. Default is true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Notification Channel. Must be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> pulumi.Output[bool]:
        """
        Send a new notification when the alert condition is 
        no longer triggered. Default is false.
        """
        return pulumi.get(self, "notify_when_ok")

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> pulumi.Output[bool]:
        """
        Send a new notification when the alert is manually 
        acknowledged by a user. Default is false.
        """
        return pulumi.get(self, "notify_when_resolved")

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> pulumi.Output[Optional[bool]]:
        """
        Send an initial test notification to check
        if the notification channel is working. Default is false.
        """
        return pulumi.get(self, "send_test_notification")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        (Computed) The current version of the Notification Channel.
        """
        return pulumi.get(self, "version")

