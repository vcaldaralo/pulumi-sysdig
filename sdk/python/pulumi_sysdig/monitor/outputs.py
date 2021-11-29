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
    'AlertAnomalyCapture',
    'AlertAnomalyCustomNotification',
    'AlertDashboardPanel',
    'AlertDashboardPanelQuery',
    'AlertDashboardScope',
    'AlertDowntimeCapture',
    'AlertDowntimeCustomNotification',
    'AlertEventCapture',
    'AlertEventCustomNotification',
    'GroupOutlierCapture',
    'GroupOutlierCustomNotification',
    'MetricCapture',
    'MetricCustomNotification',
    'PromqlCapture',
    'PromqlCustomNotification',
    'TeamEntrypoint',
    'TeamUserRole',
]

@pulumi.output_type
class AlertAnomalyCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class AlertAnomalyCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class AlertDashboardPanel(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "posX":
            suggest = "pos_x"
        elif key == "posY":
            suggest = "pos_y"
        elif key == "autosizeText":
            suggest = "autosize_text"
        elif key == "transparentBackground":
            suggest = "transparent_background"
        elif key == "visibleTitle":
            suggest = "visible_title"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AlertDashboardPanel. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AlertDashboardPanel.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AlertDashboardPanel.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 height: int,
                 name: str,
                 pos_x: int,
                 pos_y: int,
                 type: str,
                 width: int,
                 autosize_text: Optional[bool] = None,
                 content: Optional[str] = None,
                 description: Optional[str] = None,
                 queries: Optional[Sequence['outputs.AlertDashboardPanelQuery']] = None,
                 transparent_background: Optional[bool] = None,
                 visible_title: Optional[bool] = None):
        pulumi.set(__self__, "height", height)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "pos_x", pos_x)
        pulumi.set(__self__, "pos_y", pos_y)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "width", width)
        if autosize_text is not None:
            pulumi.set(__self__, "autosize_text", autosize_text)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if queries is not None:
            pulumi.set(__self__, "queries", queries)
        if transparent_background is not None:
            pulumi.set(__self__, "transparent_background", transparent_background)
        if visible_title is not None:
            pulumi.set(__self__, "visible_title", visible_title)

    @property
    @pulumi.getter
    def height(self) -> int:
        return pulumi.get(self, "height")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="posX")
    def pos_x(self) -> int:
        return pulumi.get(self, "pos_x")

    @property
    @pulumi.getter(name="posY")
    def pos_y(self) -> int:
        return pulumi.get(self, "pos_y")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def width(self) -> int:
        return pulumi.get(self, "width")

    @property
    @pulumi.getter(name="autosizeText")
    def autosize_text(self) -> Optional[bool]:
        return pulumi.get(self, "autosize_text")

    @property
    @pulumi.getter
    def content(self) -> Optional[str]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def queries(self) -> Optional[Sequence['outputs.AlertDashboardPanelQuery']]:
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter(name="transparentBackground")
    def transparent_background(self) -> Optional[bool]:
        return pulumi.get(self, "transparent_background")

    @property
    @pulumi.getter(name="visibleTitle")
    def visible_title(self) -> Optional[bool]:
        return pulumi.get(self, "visible_title")


@pulumi.output_type
class AlertDashboardPanelQuery(dict):
    def __init__(__self__, *,
                 promql: str,
                 unit: str):
        pulumi.set(__self__, "promql", promql)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter
    def promql(self) -> str:
        return pulumi.get(self, "promql")

    @property
    @pulumi.getter
    def unit(self) -> str:
        return pulumi.get(self, "unit")


@pulumi.output_type
class AlertDashboardScope(dict):
    def __init__(__self__, *,
                 metric: str,
                 comparator: Optional[str] = None,
                 values: Optional[Sequence[str]] = None,
                 variable: Optional[str] = None):
        pulumi.set(__self__, "metric", metric)
        if comparator is not None:
            pulumi.set(__self__, "comparator", comparator)
        if values is not None:
            pulumi.set(__self__, "values", values)
        if variable is not None:
            pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def metric(self) -> str:
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def comparator(self) -> Optional[str]:
        return pulumi.get(self, "comparator")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def variable(self) -> Optional[str]:
        return pulumi.get(self, "variable")


@pulumi.output_type
class AlertDowntimeCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class AlertDowntimeCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class AlertEventCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class AlertEventCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class GroupOutlierCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class GroupOutlierCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class MetricCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class MetricCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class PromqlCapture(dict):
    def __init__(__self__, *,
                 duration: int,
                 filename: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> int:
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")


@pulumi.output_type
class PromqlCustomNotification(dict):
    def __init__(__self__, *,
                 title: str,
                 append: Optional[str] = None,
                 prepend: Optional[str] = None):
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def append(self) -> Optional[str]:
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def prepend(self) -> Optional[str]:
        return pulumi.get(self, "prepend")


@pulumi.output_type
class TeamEntrypoint(dict):
    def __init__(__self__, *,
                 type: str,
                 selection: Optional[str] = None):
        pulumi.set(__self__, "type", type)
        if selection is not None:
            pulumi.set(__self__, "selection", selection)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def selection(self) -> Optional[str]:
        return pulumi.get(self, "selection")


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

