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

__all__ = ['RuleFilesystemArgs', 'RuleFilesystem']

@pulumi.input_type
class RuleFilesystemArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_onlies: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]] = None,
                 read_writes: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RuleFilesystem resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if read_onlies is not None:
            pulumi.set(__self__, "read_onlies", read_onlies)
        if read_writes is not None:
            pulumi.set(__self__, "read_writes", read_writes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="readOnlies")
    def read_onlies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]]:
        return pulumi.get(self, "read_onlies")

    @read_onlies.setter
    def read_onlies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]]):
        pulumi.set(self, "read_onlies", value)

    @property
    @pulumi.getter(name="readWrites")
    def read_writes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]]:
        return pulumi.get(self, "read_writes")

    @read_writes.setter
    def read_writes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]]):
        pulumi.set(self, "read_writes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RuleFilesystemState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_onlies: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]] = None,
                 read_writes: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RuleFilesystem resources.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if read_onlies is not None:
            pulumi.set(__self__, "read_onlies", read_onlies)
        if read_writes is not None:
            pulumi.set(__self__, "read_writes", read_writes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if version is not None:
            pulumi.set(__self__, "version", version)

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
    @pulumi.getter(name="readOnlies")
    def read_onlies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]]:
        return pulumi.get(self, "read_onlies")

    @read_onlies.setter
    def read_onlies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadOnlyArgs']]]]):
        pulumi.set(self, "read_onlies", value)

    @property
    @pulumi.getter(name="readWrites")
    def read_writes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]]:
        return pulumi.get(self, "read_writes")

    @read_writes.setter
    def read_writes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleFilesystemReadWriteArgs']]]]):
        pulumi.set(self, "read_writes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class RuleFilesystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_onlies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadOnlyArgs']]]]] = None,
                 read_writes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadWriteArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a RuleFilesystem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RuleFilesystemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RuleFilesystem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RuleFilesystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleFilesystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_onlies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadOnlyArgs']]]]] = None,
                 read_writes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadWriteArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = RuleFilesystemArgs.__new__(RuleFilesystemArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["read_onlies"] = read_onlies
            __props__.__dict__["read_writes"] = read_writes
            __props__.__dict__["tags"] = tags
            __props__.__dict__["version"] = None
        super(RuleFilesystem, __self__).__init__(
            'sysdig:Secure/ruleFilesystem:RuleFilesystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            read_onlies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadOnlyArgs']]]]] = None,
            read_writes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleFilesystemReadWriteArgs']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'RuleFilesystem':
        """
        Get an existing RuleFilesystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleFilesystemState.__new__(_RuleFilesystemState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["read_onlies"] = read_onlies
        __props__.__dict__["read_writes"] = read_writes
        __props__.__dict__["tags"] = tags
        __props__.__dict__["version"] = version
        return RuleFilesystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readOnlies")
    def read_onlies(self) -> pulumi.Output[Optional[Sequence['outputs.RuleFilesystemReadOnly']]]:
        return pulumi.get(self, "read_onlies")

    @property
    @pulumi.getter(name="readWrites")
    def read_writes(self) -> pulumi.Output[Optional[Sequence['outputs.RuleFilesystemReadWrite']]]:
        return pulumi.get(self, "read_writes")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")
