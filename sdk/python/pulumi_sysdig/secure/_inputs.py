# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'PolicyActionArgs',
    'PolicyActionCaptureArgs',
    'RuleFalcoExceptionArgs',
    'RuleFilesystemReadOnlyArgs',
    'RuleFilesystemReadWriteArgs',
    'RuleNetworkTcpArgs',
    'RuleNetworkUdpArgs',
    'TeamUserRoleArgs',
]

@pulumi.input_type
class PolicyActionArgs:
    def __init__(__self__, *,
                 captures: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyActionCaptureArgs']]]] = None,
                 container: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['PolicyActionCaptureArgs']]] captures: Captures with Sysdig the stream of system calls:
        :param pulumi.Input[str] container: The action applied to container when this Policy is
               triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
               no action will be applied at the container level.
        """
        if captures is not None:
            pulumi.set(__self__, "captures", captures)
        if container is not None:
            pulumi.set(__self__, "container", container)

    @property
    @pulumi.getter
    def captures(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyActionCaptureArgs']]]]:
        """
        Captures with Sysdig the stream of system calls:
        """
        return pulumi.get(self, "captures")

    @captures.setter
    def captures(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyActionCaptureArgs']]]]):
        pulumi.set(self, "captures", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input[str]]:
        """
        The action applied to container when this Policy is
        triggered. Can be *stop*, *pause* or *kill*. If this is not specified,
        no action will be applied at the container level.
        """
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container", value)


@pulumi.input_type
class PolicyActionCaptureArgs:
    def __init__(__self__, *,
                 seconds_after_event: pulumi.Input[int],
                 seconds_before_event: pulumi.Input[int]):
        """
        :param pulumi.Input[int] seconds_after_event: Captures the system calls for the amount
               of seconds after the policy was triggered.
        :param pulumi.Input[int] seconds_before_event: Captures the system calls during the
               amount of seconds before the policy was triggered.
        """
        pulumi.set(__self__, "seconds_after_event", seconds_after_event)
        pulumi.set(__self__, "seconds_before_event", seconds_before_event)

    @property
    @pulumi.getter(name="secondsAfterEvent")
    def seconds_after_event(self) -> pulumi.Input[int]:
        """
        Captures the system calls for the amount
        of seconds after the policy was triggered.
        """
        return pulumi.get(self, "seconds_after_event")

    @seconds_after_event.setter
    def seconds_after_event(self, value: pulumi.Input[int]):
        pulumi.set(self, "seconds_after_event", value)

    @property
    @pulumi.getter(name="secondsBeforeEvent")
    def seconds_before_event(self) -> pulumi.Input[int]:
        """
        Captures the system calls during the
        amount of seconds before the policy was triggered.
        """
        return pulumi.get(self, "seconds_before_event")

    @seconds_before_event.setter
    def seconds_before_event(self, value: pulumi.Input[int]):
        pulumi.set(self, "seconds_before_event", value)


@pulumi.input_type
class RuleFalcoExceptionArgs:
    def __init__(__self__, *,
                 comps: pulumi.Input[Sequence[pulumi.Input[str]]],
                 fields: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: pulumi.Input[str],
                 values: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] comps: Contains comparison operators that align 1-1 with the items in the fields property.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fields: Contains one or more fields that will extract a value from the syscall/k8s_audit events.
        :param pulumi.Input[str] name: The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
        :param pulumi.Input[str] values: Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
        """
        pulumi.set(__self__, "comps", comps)
        pulumi.set(__self__, "fields", fields)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def comps(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Contains comparison operators that align 1-1 with the items in the fields property.
        """
        return pulumi.get(self, "comps")

    @comps.setter
    def comps(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "comps", value)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Contains one or more fields that will extract a value from the syscall/k8s_audit events.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the exception. Only used to provide a handy name, and to potentially link together values in a later rule that has `append = true`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[str]:
        """
        Contains tuples of values. Each item in the tuple should align 1-1 with the corresponding field and comparison operator. Since the value can be a string, a list of strings or a list of a list of strings, the value of this field must be supplied in JSON format. You can use the default `jsonencode` function to provide this value. See the usage example on the top.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class RuleFilesystemReadOnlyArgs:
    def __init__(__self__, *,
                 paths: pulumi.Input[Sequence[pulumi.Input[str]]],
                 matching: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to match.
        :param pulumi.Input[bool] matching: Defines if the path matches or not with the provided list. Default is true.
        """
        pulumi.set(__self__, "paths", paths)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of paths to match.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def matching(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if the path matches or not with the provided list. Default is true.
        """
        return pulumi.get(self, "matching")

    @matching.setter
    def matching(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "matching", value)


@pulumi.input_type
class RuleFilesystemReadWriteArgs:
    def __init__(__self__, *,
                 paths: pulumi.Input[Sequence[pulumi.Input[str]]],
                 matching: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to match.
        :param pulumi.Input[bool] matching: Defines if the path matches or not with the provided list. Default is true.
        """
        pulumi.set(__self__, "paths", paths)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of paths to match.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def matching(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if the path matches or not with the provided list. Default is true.
        """
        return pulumi.get(self, "matching")

    @matching.setter
    def matching(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "matching", value)


@pulumi.input_type
class RuleNetworkTcpArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[Sequence[pulumi.Input[int]]],
                 matching: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ports: List of ports to match.
        :param pulumi.Input[bool] matching: Defines if the port matches or not with the provided list. Default is true.
        """
        pulumi.set(__self__, "ports", ports)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        List of ports to match.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def matching(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if the port matches or not with the provided list. Default is true.
        """
        return pulumi.get(self, "matching")

    @matching.setter
    def matching(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "matching", value)


@pulumi.input_type
class RuleNetworkUdpArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[Sequence[pulumi.Input[int]]],
                 matching: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ports: List of ports to match.
        :param pulumi.Input[bool] matching: Defines if the port matches or not with the provided list. Default is true.
        """
        pulumi.set(__self__, "ports", ports)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        List of ports to match.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def matching(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if the port matches or not with the provided list. Default is true.
        """
        return pulumi.get(self, "matching")

    @matching.setter
    def matching(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "matching", value)


@pulumi.input_type
class TeamUserRoleArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 role: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] email: The email of the user in the group.
        :param pulumi.Input[str] role: The role for the user in this group.
               Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
               Default: ROLE_TEAM_STANDARD.
        """
        pulumi.set(__self__, "email", email)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The email of the user in the group.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role for the user in this group.
        Valid roles are: ROLE_TEAM_STANDARD, ROLE_TEAM_EDIT, ROLE_TEAM_READ, ROLE_TEAM_MANAGER.
        Default: ROLE_TEAM_STANDARD.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


