# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'PolicyAction',
    'PolicyActionCapture',
    'RuleFalcoException',
    'RuleFilesystemReadOnly',
    'RuleFilesystemReadWrite',
    'RuleNetworkTcp',
    'RuleNetworkUdp',
    'TeamUserRole',
]

@pulumi.output_type
class PolicyAction(dict):
    def __init__(__self__, *,
                 captures: Optional[Sequence['outputs.PolicyActionCapture']] = None,
                 container: Optional[str] = None):
        if captures is not None:
            pulumi.set(__self__, "captures", captures)
        if container is not None:
            pulumi.set(__self__, "container", container)

    @property
    @pulumi.getter
    def captures(self) -> Optional[Sequence['outputs.PolicyActionCapture']]:
        return pulumi.get(self, "captures")

    @property
    @pulumi.getter
    def container(self) -> Optional[str]:
        return pulumi.get(self, "container")


@pulumi.output_type
class PolicyActionCapture(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secondsAfterEvent":
            suggest = "seconds_after_event"
        elif key == "secondsBeforeEvent":
            suggest = "seconds_before_event"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyActionCapture. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyActionCapture.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyActionCapture.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 seconds_after_event: int,
                 seconds_before_event: int):
        pulumi.set(__self__, "seconds_after_event", seconds_after_event)
        pulumi.set(__self__, "seconds_before_event", seconds_before_event)

    @property
    @pulumi.getter(name="secondsAfterEvent")
    def seconds_after_event(self) -> int:
        return pulumi.get(self, "seconds_after_event")

    @property
    @pulumi.getter(name="secondsBeforeEvent")
    def seconds_before_event(self) -> int:
        return pulumi.get(self, "seconds_before_event")


@pulumi.output_type
class RuleFalcoException(dict):
    def __init__(__self__, *,
                 comps: Sequence[str],
                 fields: Sequence[str],
                 name: str,
                 values: str):
        pulumi.set(__self__, "comps", comps)
        pulumi.set(__self__, "fields", fields)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def comps(self) -> Sequence[str]:
        return pulumi.get(self, "comps")

    @property
    @pulumi.getter
    def fields(self) -> Sequence[str]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> str:
        return pulumi.get(self, "values")


@pulumi.output_type
class RuleFilesystemReadOnly(dict):
    def __init__(__self__, *,
                 paths: Sequence[str],
                 matching: Optional[bool] = None):
        pulumi.set(__self__, "paths", paths)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def paths(self) -> Sequence[str]:
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def matching(self) -> Optional[bool]:
        return pulumi.get(self, "matching")


@pulumi.output_type
class RuleFilesystemReadWrite(dict):
    def __init__(__self__, *,
                 paths: Sequence[str],
                 matching: Optional[bool] = None):
        pulumi.set(__self__, "paths", paths)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def paths(self) -> Sequence[str]:
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def matching(self) -> Optional[bool]:
        return pulumi.get(self, "matching")


@pulumi.output_type
class RuleNetworkTcp(dict):
    def __init__(__self__, *,
                 ports: Sequence[int],
                 matching: Optional[bool] = None):
        pulumi.set(__self__, "ports", ports)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def ports(self) -> Sequence[int]:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def matching(self) -> Optional[bool]:
        return pulumi.get(self, "matching")


@pulumi.output_type
class RuleNetworkUdp(dict):
    def __init__(__self__, *,
                 ports: Sequence[int],
                 matching: Optional[bool] = None):
        pulumi.set(__self__, "ports", ports)
        if matching is not None:
            pulumi.set(__self__, "matching", matching)

    @property
    @pulumi.getter
    def ports(self) -> Sequence[int]:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def matching(self) -> Optional[bool]:
        return pulumi.get(self, "matching")


@pulumi.output_type
class TeamUserRole(dict):
    def __init__(__self__, *,
                 email: str,
                 role: Optional[str] = None):
        pulumi.set(__self__, "email", email)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")


