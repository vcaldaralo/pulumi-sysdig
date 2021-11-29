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

__all__ = ['RuleNetworkArgs', 'RuleNetwork']

@pulumi.input_type
class RuleNetworkArgs:
    def __init__(__self__, *,
                 block_inbound: pulumi.Input[bool],
                 block_outbound: pulumi.Input[bool],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcps: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]] = None,
                 udps: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]] = None):
        """
        The set of arguments for constructing a RuleNetwork resource.
        """
        pulumi.set(__self__, "block_inbound", block_inbound)
        pulumi.set(__self__, "block_outbound", block_outbound)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tcps is not None:
            pulumi.set(__self__, "tcps", tcps)
        if udps is not None:
            pulumi.set(__self__, "udps", udps)

    @property
    @pulumi.getter(name="blockInbound")
    def block_inbound(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "block_inbound")

    @block_inbound.setter
    def block_inbound(self, value: pulumi.Input[bool]):
        pulumi.set(self, "block_inbound", value)

    @property
    @pulumi.getter(name="blockOutbound")
    def block_outbound(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "block_outbound")

    @block_outbound.setter
    def block_outbound(self, value: pulumi.Input[bool]):
        pulumi.set(self, "block_outbound", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def tcps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]]:
        return pulumi.get(self, "tcps")

    @tcps.setter
    def tcps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]]):
        pulumi.set(self, "tcps", value)

    @property
    @pulumi.getter
    def udps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]]:
        return pulumi.get(self, "udps")

    @udps.setter
    def udps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]]):
        pulumi.set(self, "udps", value)


@pulumi.input_type
class _RuleNetworkState:
    def __init__(__self__, *,
                 block_inbound: Optional[pulumi.Input[bool]] = None,
                 block_outbound: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcps: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]] = None,
                 udps: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RuleNetwork resources.
        """
        if block_inbound is not None:
            pulumi.set(__self__, "block_inbound", block_inbound)
        if block_outbound is not None:
            pulumi.set(__self__, "block_outbound", block_outbound)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tcps is not None:
            pulumi.set(__self__, "tcps", tcps)
        if udps is not None:
            pulumi.set(__self__, "udps", udps)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="blockInbound")
    def block_inbound(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "block_inbound")

    @block_inbound.setter
    def block_inbound(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_inbound", value)

    @property
    @pulumi.getter(name="blockOutbound")
    def block_outbound(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "block_outbound")

    @block_outbound.setter
    def block_outbound(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_outbound", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def tcps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]]:
        return pulumi.get(self, "tcps")

    @tcps.setter
    def tcps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkTcpArgs']]]]):
        pulumi.set(self, "tcps", value)

    @property
    @pulumi.getter
    def udps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]]:
        return pulumi.get(self, "udps")

    @udps.setter
    def udps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleNetworkUdpArgs']]]]):
        pulumi.set(self, "udps", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class RuleNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_inbound: Optional[pulumi.Input[bool]] = None,
                 block_outbound: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkTcpArgs']]]]] = None,
                 udps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkUdpArgs']]]]] = None,
                 __props__=None):
        """
        Create a RuleNetwork resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RuleNetwork resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RuleNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_inbound: Optional[pulumi.Input[bool]] = None,
                 block_outbound: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkTcpArgs']]]]] = None,
                 udps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkUdpArgs']]]]] = None,
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
            __props__ = RuleNetworkArgs.__new__(RuleNetworkArgs)

            if block_inbound is None and not opts.urn:
                raise TypeError("Missing required property 'block_inbound'")
            __props__.__dict__["block_inbound"] = block_inbound
            if block_outbound is None and not opts.urn:
                raise TypeError("Missing required property 'block_outbound'")
            __props__.__dict__["block_outbound"] = block_outbound
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tcps"] = tcps
            __props__.__dict__["udps"] = udps
            __props__.__dict__["version"] = None
        super(RuleNetwork, __self__).__init__(
            'sysdig:Secure/ruleNetwork:RuleNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block_inbound: Optional[pulumi.Input[bool]] = None,
            block_outbound: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tcps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkTcpArgs']]]]] = None,
            udps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleNetworkUdpArgs']]]]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'RuleNetwork':
        """
        Get an existing RuleNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleNetworkState.__new__(_RuleNetworkState)

        __props__.__dict__["block_inbound"] = block_inbound
        __props__.__dict__["block_outbound"] = block_outbound
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tcps"] = tcps
        __props__.__dict__["udps"] = udps
        __props__.__dict__["version"] = version
        return RuleNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockInbound")
    def block_inbound(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "block_inbound")

    @property
    @pulumi.getter(name="blockOutbound")
    def block_outbound(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "block_outbound")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tcps(self) -> pulumi.Output[Optional[Sequence['outputs.RuleNetworkTcp']]]:
        return pulumi.get(self, "tcps")

    @property
    @pulumi.getter
    def udps(self) -> pulumi.Output[Optional[Sequence['outputs.RuleNetworkUdp']]]:
        return pulumi.get(self, "udps")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

