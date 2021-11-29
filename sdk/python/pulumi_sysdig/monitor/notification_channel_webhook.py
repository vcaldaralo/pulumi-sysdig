# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NotificationChannelWebhookArgs', 'NotificationChannelWebhook']

@pulumi.input_type
class NotificationChannelWebhookArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str],
                 additional_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NotificationChannelWebhook resource.
        """
        pulumi.set(__self__, "url", url)
        if additional_headers is not None:
            pulumi.set(__self__, "additional_headers", additional_headers)
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

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "additional_headers")

    @additional_headers.setter
    def additional_headers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "additional_headers", value)

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
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "notify_when_ok")

    @notify_when_ok.setter
    def notify_when_ok(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_ok", value)

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "notify_when_resolved")

    @notify_when_resolved.setter
    def notify_when_resolved(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_resolved", value)

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "send_test_notification")

    @send_test_notification.setter
    def send_test_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_test_notification", value)


@pulumi.input_type
class _NotificationChannelWebhookState:
    def __init__(__self__, *,
                 additional_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering NotificationChannelWebhook resources.
        """
        if additional_headers is not None:
            pulumi.set(__self__, "additional_headers", additional_headers)
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
        if url is not None:
            pulumi.set(__self__, "url", url)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "additional_headers")

    @additional_headers.setter
    def additional_headers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "additional_headers", value)

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
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "notify_when_ok")

    @notify_when_ok.setter
    def notify_when_ok(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_ok", value)

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "notify_when_resolved")

    @notify_when_resolved.setter
    def notify_when_resolved(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_when_resolved", value)

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "send_test_notification")

    @send_test_notification.setter
    def send_test_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_test_notification", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class NotificationChannelWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NotificationChannelWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationChannelWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NotificationChannelWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NotificationChannelWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationChannelWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_when_ok: Optional[pulumi.Input[bool]] = None,
                 notify_when_resolved: Optional[pulumi.Input[bool]] = None,
                 send_test_notification: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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
            __props__ = NotificationChannelWebhookArgs.__new__(NotificationChannelWebhookArgs)

            __props__.__dict__["additional_headers"] = additional_headers
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["notify_when_ok"] = notify_when_ok
            __props__.__dict__["notify_when_resolved"] = notify_when_resolved
            __props__.__dict__["send_test_notification"] = send_test_notification
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["version"] = None
        super(NotificationChannelWebhook, __self__).__init__(
            'sysdig:Monitor/notificationChannelWebhook:NotificationChannelWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notify_when_ok: Optional[pulumi.Input[bool]] = None,
            notify_when_resolved: Optional[pulumi.Input[bool]] = None,
            send_test_notification: Optional[pulumi.Input[bool]] = None,
            url: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'NotificationChannelWebhook':
        """
        Get an existing NotificationChannelWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationChannelWebhookState.__new__(_NotificationChannelWebhookState)

        __props__.__dict__["additional_headers"] = additional_headers
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["notify_when_ok"] = notify_when_ok
        __props__.__dict__["notify_when_resolved"] = notify_when_resolved
        __props__.__dict__["send_test_notification"] = send_test_notification
        __props__.__dict__["url"] = url
        __props__.__dict__["version"] = version
        return NotificationChannelWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "additional_headers")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyWhenOk")
    def notify_when_ok(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "notify_when_ok")

    @property
    @pulumi.getter(name="notifyWhenResolved")
    def notify_when_resolved(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "notify_when_resolved")

    @property
    @pulumi.getter(name="sendTestNotification")
    def send_test_notification(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "send_test_notification")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

