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

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 entrypoints: pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]],
                 can_see_infrastructure_events: Optional[pulumi.Input[bool]] = None,
                 can_use_aws_data: Optional[pulumi.Input[bool]] = None,
                 can_use_sysdig_capture: Optional[pulumi.Input[bool]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_by: Optional[pulumi.Input[str]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 user_roles: Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]] entrypoints: Main entry point for the current team in the product. 
               See the Entrypoint argument reference section for more information.
        :param pulumi.Input[bool] can_see_infrastructure_events: TODO. Default: false.
        :param pulumi.Input[bool] can_use_aws_data: TODO. Default: false.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[str] filter: If the team can only see some resources, 
               write down a filter of such resources.
        :param pulumi.Input[str] name: The name of the Monitor Team. It must be unique and must not exist in Secure.
        :param pulumi.Input[str] scope_by: Scope for the team. Default: "container".
        :param pulumi.Input[str] theme: Colour of the team. Default: "#73A1F7".
        """
        pulumi.set(__self__, "entrypoints", entrypoints)
        if can_see_infrastructure_events is not None:
            pulumi.set(__self__, "can_see_infrastructure_events", can_see_infrastructure_events)
        if can_use_aws_data is not None:
            pulumi.set(__self__, "can_use_aws_data", can_use_aws_data)
        if can_use_sysdig_capture is not None:
            pulumi.set(__self__, "can_use_sysdig_capture", can_use_sysdig_capture)
        if default_team is not None:
            pulumi.set(__self__, "default_team", default_team)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_by is not None:
            pulumi.set(__self__, "scope_by", scope_by)
        if theme is not None:
            pulumi.set(__self__, "theme", theme)
        if user_roles is not None:
            pulumi.set(__self__, "user_roles", user_roles)

    @property
    @pulumi.getter
    def entrypoints(self) -> pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]]:
        """
        Main entry point for the current team in the product. 
        See the Entrypoint argument reference section for more information.
        """
        return pulumi.get(self, "entrypoints")

    @entrypoints.setter
    def entrypoints(self, value: pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]]):
        pulumi.set(self, "entrypoints", value)

    @property
    @pulumi.getter(name="canSeeInfrastructureEvents")
    def can_see_infrastructure_events(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_see_infrastructure_events")

    @can_see_infrastructure_events.setter
    def can_see_infrastructure_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_see_infrastructure_events", value)

    @property
    @pulumi.getter(name="canUseAwsData")
    def can_use_aws_data(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_use_aws_data")

    @can_use_aws_data.setter
    def can_use_aws_data(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_use_aws_data", value)

    @property
    @pulumi.getter(name="canUseSysdigCapture")
    def can_use_sysdig_capture(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "can_use_sysdig_capture")

    @can_use_sysdig_capture.setter
    def can_use_sysdig_capture(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_use_sysdig_capture", value)

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "default_team")

    @default_team.setter
    def default_team(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_team", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        If the team can only see some resources, 
        write down a filter of such resources.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Monitor Team. It must be unique and must not exist in Secure.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="scopeBy")
    def scope_by(self) -> Optional[pulumi.Input[str]]:
        """
        Scope for the team. Default: "container".
        """
        return pulumi.get(self, "scope_by")

    @scope_by.setter
    def scope_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_by", value)

    @property
    @pulumi.getter
    def theme(self) -> Optional[pulumi.Input[str]]:
        """
        Colour of the team. Default: "#73A1F7".
        """
        return pulumi.get(self, "theme")

    @theme.setter
    def theme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme", value)

    @property
    @pulumi.getter(name="userRoles")
    def user_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]]:
        return pulumi.get(self, "user_roles")

    @user_roles.setter
    def user_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]]):
        pulumi.set(self, "user_roles", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 can_see_infrastructure_events: Optional[pulumi.Input[bool]] = None,
                 can_use_aws_data: Optional[pulumi.Input[bool]] = None,
                 can_use_sysdig_capture: Optional[pulumi.Input[bool]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entrypoints: Optional[pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_by: Optional[pulumi.Input[str]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 user_roles: Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[bool] can_see_infrastructure_events: TODO. Default: false.
        :param pulumi.Input[bool] can_use_aws_data: TODO. Default: false.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]] entrypoints: Main entry point for the current team in the product. 
               See the Entrypoint argument reference section for more information.
        :param pulumi.Input[str] filter: If the team can only see some resources, 
               write down a filter of such resources.
        :param pulumi.Input[str] name: The name of the Monitor Team. It must be unique and must not exist in Secure.
        :param pulumi.Input[str] scope_by: Scope for the team. Default: "container".
        :param pulumi.Input[str] theme: Colour of the team. Default: "#73A1F7".
        """
        if can_see_infrastructure_events is not None:
            pulumi.set(__self__, "can_see_infrastructure_events", can_see_infrastructure_events)
        if can_use_aws_data is not None:
            pulumi.set(__self__, "can_use_aws_data", can_use_aws_data)
        if can_use_sysdig_capture is not None:
            pulumi.set(__self__, "can_use_sysdig_capture", can_use_sysdig_capture)
        if default_team is not None:
            pulumi.set(__self__, "default_team", default_team)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entrypoints is not None:
            pulumi.set(__self__, "entrypoints", entrypoints)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_by is not None:
            pulumi.set(__self__, "scope_by", scope_by)
        if theme is not None:
            pulumi.set(__self__, "theme", theme)
        if user_roles is not None:
            pulumi.set(__self__, "user_roles", user_roles)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="canSeeInfrastructureEvents")
    def can_see_infrastructure_events(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_see_infrastructure_events")

    @can_see_infrastructure_events.setter
    def can_see_infrastructure_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_see_infrastructure_events", value)

    @property
    @pulumi.getter(name="canUseAwsData")
    def can_use_aws_data(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_use_aws_data")

    @can_use_aws_data.setter
    def can_use_aws_data(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_use_aws_data", value)

    @property
    @pulumi.getter(name="canUseSysdigCapture")
    def can_use_sysdig_capture(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "can_use_sysdig_capture")

    @can_use_sysdig_capture.setter
    def can_use_sysdig_capture(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_use_sysdig_capture", value)

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "default_team")

    @default_team.setter
    def default_team(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_team", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def entrypoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]]]:
        """
        Main entry point for the current team in the product. 
        See the Entrypoint argument reference section for more information.
        """
        return pulumi.get(self, "entrypoints")

    @entrypoints.setter
    def entrypoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamEntrypointArgs']]]]):
        pulumi.set(self, "entrypoints", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        If the team can only see some resources, 
        write down a filter of such resources.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Monitor Team. It must be unique and must not exist in Secure.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="scopeBy")
    def scope_by(self) -> Optional[pulumi.Input[str]]:
        """
        Scope for the team. Default: "container".
        """
        return pulumi.get(self, "scope_by")

    @scope_by.setter
    def scope_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_by", value)

    @property
    @pulumi.getter
    def theme(self) -> Optional[pulumi.Input[str]]:
        """
        Colour of the team. Default: "#73A1F7".
        """
        return pulumi.get(self, "theme")

    @theme.setter
    def theme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme", value)

    @property
    @pulumi.getter(name="userRoles")
    def user_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]]:
        return pulumi.get(self, "user_roles")

    @user_roles.setter
    def user_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamUserRoleArgs']]]]):
        pulumi.set(self, "user_roles", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_see_infrastructure_events: Optional[pulumi.Input[bool]] = None,
                 can_use_aws_data: Optional[pulumi.Input[bool]] = None,
                 can_use_sysdig_capture: Optional[pulumi.Input[bool]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entrypoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamEntrypointArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_by: Optional[pulumi.Input[str]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 user_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamUserRoleArgs']]]]] = None,
                 __props__=None):
        """
        ## Import

        Monitor Teams can be imported using the ID, e.g.

        ```sh
         $ pulumi import sysdig:Monitor/team:Team example 12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_see_infrastructure_events: TODO. Default: false.
        :param pulumi.Input[bool] can_use_aws_data: TODO. Default: false.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamEntrypointArgs']]]] entrypoints: Main entry point for the current team in the product. 
               See the Entrypoint argument reference section for more information.
        :param pulumi.Input[str] filter: If the team can only see some resources, 
               write down a filter of such resources.
        :param pulumi.Input[str] name: The name of the Monitor Team. It must be unique and must not exist in Secure.
        :param pulumi.Input[str] scope_by: Scope for the team. Default: "container".
        :param pulumi.Input[str] theme: Colour of the team. Default: "#73A1F7".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Monitor Teams can be imported using the ID, e.g.

        ```sh
         $ pulumi import sysdig:Monitor/team:Team example 12345
        ```

        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_see_infrastructure_events: Optional[pulumi.Input[bool]] = None,
                 can_use_aws_data: Optional[pulumi.Input[bool]] = None,
                 can_use_sysdig_capture: Optional[pulumi.Input[bool]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entrypoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamEntrypointArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_by: Optional[pulumi.Input[str]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 user_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamUserRoleArgs']]]]] = None,
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
            __props__ = TeamArgs.__new__(TeamArgs)

            __props__.__dict__["can_see_infrastructure_events"] = can_see_infrastructure_events
            __props__.__dict__["can_use_aws_data"] = can_use_aws_data
            __props__.__dict__["can_use_sysdig_capture"] = can_use_sysdig_capture
            __props__.__dict__["default_team"] = default_team
            __props__.__dict__["description"] = description
            if entrypoints is None and not opts.urn:
                raise TypeError("Missing required property 'entrypoints'")
            __props__.__dict__["entrypoints"] = entrypoints
            __props__.__dict__["filter"] = filter
            __props__.__dict__["name"] = name
            __props__.__dict__["scope_by"] = scope_by
            __props__.__dict__["theme"] = theme
            __props__.__dict__["user_roles"] = user_roles
            __props__.__dict__["version"] = None
        super(Team, __self__).__init__(
            'sysdig:Monitor/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_see_infrastructure_events: Optional[pulumi.Input[bool]] = None,
            can_use_aws_data: Optional[pulumi.Input[bool]] = None,
            can_use_sysdig_capture: Optional[pulumi.Input[bool]] = None,
            default_team: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            entrypoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamEntrypointArgs']]]]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scope_by: Optional[pulumi.Input[str]] = None,
            theme: Optional[pulumi.Input[str]] = None,
            user_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamUserRoleArgs']]]]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_see_infrastructure_events: TODO. Default: false.
        :param pulumi.Input[bool] can_use_aws_data: TODO. Default: false.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamEntrypointArgs']]]] entrypoints: Main entry point for the current team in the product. 
               See the Entrypoint argument reference section for more information.
        :param pulumi.Input[str] filter: If the team can only see some resources, 
               write down a filter of such resources.
        :param pulumi.Input[str] name: The name of the Monitor Team. It must be unique and must not exist in Secure.
        :param pulumi.Input[str] scope_by: Scope for the team. Default: "container".
        :param pulumi.Input[str] theme: Colour of the team. Default: "#73A1F7".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["can_see_infrastructure_events"] = can_see_infrastructure_events
        __props__.__dict__["can_use_aws_data"] = can_use_aws_data
        __props__.__dict__["can_use_sysdig_capture"] = can_use_sysdig_capture
        __props__.__dict__["default_team"] = default_team
        __props__.__dict__["description"] = description
        __props__.__dict__["entrypoints"] = entrypoints
        __props__.__dict__["filter"] = filter
        __props__.__dict__["name"] = name
        __props__.__dict__["scope_by"] = scope_by
        __props__.__dict__["theme"] = theme
        __props__.__dict__["user_roles"] = user_roles
        __props__.__dict__["version"] = version
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canSeeInfrastructureEvents")
    def can_see_infrastructure_events(self) -> pulumi.Output[Optional[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_see_infrastructure_events")

    @property
    @pulumi.getter(name="canUseAwsData")
    def can_use_aws_data(self) -> pulumi.Output[Optional[bool]]:
        """
        TODO. Default: false.
        """
        return pulumi.get(self, "can_use_aws_data")

    @property
    @pulumi.getter(name="canUseSysdigCapture")
    def can_use_sysdig_capture(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "can_use_sysdig_capture")

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "default_team")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def entrypoints(self) -> pulumi.Output[Sequence['outputs.TeamEntrypoint']]:
        """
        Main entry point for the current team in the product. 
        See the Entrypoint argument reference section for more information.
        """
        return pulumi.get(self, "entrypoints")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional[str]]:
        """
        If the team can only see some resources, 
        write down a filter of such resources.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Monitor Team. It must be unique and must not exist in Secure.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scopeBy")
    def scope_by(self) -> pulumi.Output[Optional[str]]:
        """
        Scope for the team. Default: "container".
        """
        return pulumi.get(self, "scope_by")

    @property
    @pulumi.getter
    def theme(self) -> pulumi.Output[Optional[str]]:
        """
        Colour of the team. Default: "#73A1F7".
        """
        return pulumi.get(self, "theme")

    @property
    @pulumi.getter(name="userRoles")
    def user_roles(self) -> pulumi.Output[Optional[Sequence['outputs.TeamUserRole']]]:
        return pulumi.get(self, "user_roles")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

