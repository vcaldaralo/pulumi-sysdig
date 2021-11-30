# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AlertAnomalyCaptureArgs',
    'AlertAnomalyCustomNotificationArgs',
    'AlertDowntimeCaptureArgs',
    'AlertDowntimeCustomNotificationArgs',
    'AlertEventCaptureArgs',
    'AlertEventCustomNotificationArgs',
    'AlertGroupOutlierCaptureArgs',
    'AlertGroupOutlierCustomNotificationArgs',
    'AlertMetricCaptureArgs',
    'AlertMetricCustomNotificationArgs',
    'AlertPromqlCaptureArgs',
    'AlertPromqlCustomNotificationArgs',
    'DashboardPanelArgs',
    'DashboardPanelQueryArgs',
    'DashboardScopeArgs',
    'TeamEntrypointArgs',
    'TeamUserRoleArgs',
]

@pulumi.input_type
class AlertAnomalyCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] duration: Time frame in seconds of the capture.
        :param pulumi.Input[str] filename: Defines the name of the capture file.
        :param pulumi.Input[str] filter: Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        """
        Time frame in seconds of the capture.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        Defines the name of the capture file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertAnomalyCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class AlertDowntimeCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] duration: Time frame in seconds of the capture.
        :param pulumi.Input[str] filename: Defines the name of the capture file.
        :param pulumi.Input[str] filter: Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        """
        Time frame in seconds of the capture.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        Defines the name of the capture file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertDowntimeCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class AlertEventCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] duration: Time frame in seconds of the capture.
        :param pulumi.Input[str] filename: Defines the name of the capture file.
        :param pulumi.Input[str] filter: Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        """
        Time frame in seconds of the capture.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        Defines the name of the capture file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertEventCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class AlertGroupOutlierCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] duration: Time frame in seconds of the capture.
        :param pulumi.Input[str] filename: Defines the name of the capture file.
        :param pulumi.Input[str] filter: Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        """
        Time frame in seconds of the capture.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        Defines the name of the capture file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertGroupOutlierCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class AlertMetricCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] duration: Time frame in seconds of the capture.
        :param pulumi.Input[str] filename: Defines the name of the capture file.
        :param pulumi.Input[str] filter: Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        """
        Time frame in seconds of the capture.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        Defines the name of the capture file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Additional filter to apply to the capture. For example: `proc.name contains nginx`.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertMetricCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class AlertPromqlCaptureArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[int],
                 filename: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "filename", filename)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[int]:
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class AlertPromqlCustomNotificationArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 append: Optional[pulumi.Input[str]] = None,
                 prepend: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] title: Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        :param pulumi.Input[str] append: Text to add after the alert template.
        :param pulumi.Input[str] prepend: Text to add before the alert template.
        """
        pulumi.set(__self__, "title", title)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if prepend is not None:
            pulumi.set(__self__, "prepend", prepend)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Sets the title of the alert. It is commonly defined as `{{__alert_name__}} is {{__alert_status__}}`.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add after the alert template.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def prepend(self) -> Optional[pulumi.Input[str]]:
        """
        Text to add before the alert template.
        """
        return pulumi.get(self, "prepend")

    @prepend.setter
    def prepend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepend", value)


@pulumi.input_type
class DashboardPanelArgs:
    def __init__(__self__, *,
                 height: pulumi.Input[int],
                 name: pulumi.Input[str],
                 pos_x: pulumi.Input[int],
                 pos_y: pulumi.Input[int],
                 type: pulumi.Input[str],
                 width: pulumi.Input[int],
                 autosize_text: Optional[pulumi.Input[bool]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 queries: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardPanelQueryArgs']]]] = None,
                 transparent_background: Optional[pulumi.Input[bool]] = None,
                 visible_title: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[int] height: Height of the panel. Min value: 1.
        :param pulumi.Input[str] name: Name of the panel.
        :param pulumi.Input[int] pos_x: Position of the panel in the X axis. Min value: 0, max value: 23.
        :param pulumi.Input[int] pos_y: Position of the panel in the Y axis. Min value: 0.
        :param pulumi.Input[str] type: Kind of panel, must be either `timechart`, `number` or `text`.
        :param pulumi.Input[int] width: Width of the panel. Min value: 1, max value: 24.
        :param pulumi.Input[bool] autosize_text: If true, the text will be autosized in the panel.
               This field is ignored for all panel types except `text`.
        :param pulumi.Input[str] content: This field is required if the panel type is `text`. It represents the 
               text that will be displayed in the panel.
        :param pulumi.Input[str] description: Description of the panel.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardPanelQueryArgs']]] queries: The PromQL query that will show information in the panel. 
               If the type of the panel is `timechart`, then it can be specified multiple
               times, to have multiple metrics in the same graph.
               If the type of the panel is `number` then only one can be specified.
               This field is required if the panel type is `timechart` or `number`.
        :param pulumi.Input[bool] transparent_background: If true, the panel will have a transparent background.
               This field is ignored for all panel types except `text`.
        :param pulumi.Input[bool] visible_title: If true, the title of the panel will be displayed. Default: false.
               This field is ignored for all panel types except `text`.
        """
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
    def height(self) -> pulumi.Input[int]:
        """
        Height of the panel. Min value: 1.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: pulumi.Input[int]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the panel.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="posX")
    def pos_x(self) -> pulumi.Input[int]:
        """
        Position of the panel in the X axis. Min value: 0, max value: 23.
        """
        return pulumi.get(self, "pos_x")

    @pos_x.setter
    def pos_x(self, value: pulumi.Input[int]):
        pulumi.set(self, "pos_x", value)

    @property
    @pulumi.getter(name="posY")
    def pos_y(self) -> pulumi.Input[int]:
        """
        Position of the panel in the Y axis. Min value: 0.
        """
        return pulumi.get(self, "pos_y")

    @pos_y.setter
    def pos_y(self, value: pulumi.Input[int]):
        pulumi.set(self, "pos_y", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Kind of panel, must be either `timechart`, `number` or `text`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def width(self) -> pulumi.Input[int]:
        """
        Width of the panel. Min value: 1, max value: 24.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: pulumi.Input[int]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="autosizeText")
    def autosize_text(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the text will be autosized in the panel.
        This field is ignored for all panel types except `text`.
        """
        return pulumi.get(self, "autosize_text")

    @autosize_text.setter
    def autosize_text(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "autosize_text", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        This field is required if the panel type is `text`. It represents the 
        text that will be displayed in the panel.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the panel.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def queries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardPanelQueryArgs']]]]:
        """
        The PromQL query that will show information in the panel. 
        If the type of the panel is `timechart`, then it can be specified multiple
        times, to have multiple metrics in the same graph.
        If the type of the panel is `number` then only one can be specified.
        This field is required if the panel type is `timechart` or `number`.
        """
        return pulumi.get(self, "queries")

    @queries.setter
    def queries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardPanelQueryArgs']]]]):
        pulumi.set(self, "queries", value)

    @property
    @pulumi.getter(name="transparentBackground")
    def transparent_background(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the panel will have a transparent background.
        This field is ignored for all panel types except `text`.
        """
        return pulumi.get(self, "transparent_background")

    @transparent_background.setter
    def transparent_background(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transparent_background", value)

    @property
    @pulumi.getter(name="visibleTitle")
    def visible_title(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the title of the panel will be displayed. Default: false.
        This field is ignored for all panel types except `text`.
        """
        return pulumi.get(self, "visible_title")

    @visible_title.setter
    def visible_title(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "visible_title", value)


@pulumi.input_type
class DashboardPanelQueryArgs:
    def __init__(__self__, *,
                 promql: pulumi.Input[str],
                 unit: pulumi.Input[str]):
        """
        :param pulumi.Input[str] promql: The PromQL query. Must be a valid PromQL query with existing
               metrics in Sysdig Monitor.
        :param pulumi.Input[str] unit: The type of metric for this query. Can be one of: `percent`, `data`, `data rate`, 
               `number`, `number rate`, `time`.
        """
        pulumi.set(__self__, "promql", promql)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter
    def promql(self) -> pulumi.Input[str]:
        """
        The PromQL query. Must be a valid PromQL query with existing
        metrics in Sysdig Monitor.
        """
        return pulumi.get(self, "promql")

    @promql.setter
    def promql(self, value: pulumi.Input[str]):
        pulumi.set(self, "promql", value)

    @property
    @pulumi.getter
    def unit(self) -> pulumi.Input[str]:
        """
        The type of metric for this query. Can be one of: `percent`, `data`, `data rate`, 
        `number`, `number rate`, `time`.
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: pulumi.Input[str]):
        pulumi.set(self, "unit", value)


@pulumi.input_type
class DashboardScopeArgs:
    def __init__(__self__, *,
                 metric: pulumi.Input[str],
                 comparator: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 variable: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] metric: Metric to scope by, common examples are `host.hostName`, `kubernetes.namespace.name` or `kubernetes.cluster.name`, but you can use all the Sysdig-supported values shown in the UI. Note that kubernetes-related values only appear when Sysdig detects Kubernetes metadata.
        :param pulumi.Input[str] comparator: Operator to relate the metric with some value. It is only required if the value to filter by is set, or the variable field is not set. Valid values are: `in`, `notIn`, `equals`, `notEquals`, `contains`, `notContains` and `startsWith`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: List of values to filter by, if comparator is set. If the comparator is not `in` or `notIn` the list must contain only 1 value.
        :param pulumi.Input[str] variable: Assigns this metric to a value name and allows PromQL to reference it.
        """
        pulumi.set(__self__, "metric", metric)
        if comparator is not None:
            pulumi.set(__self__, "comparator", comparator)
        if values is not None:
            pulumi.set(__self__, "values", values)
        if variable is not None:
            pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Input[str]:
        """
        Metric to scope by, common examples are `host.hostName`, `kubernetes.namespace.name` or `kubernetes.cluster.name`, but you can use all the Sysdig-supported values shown in the UI. Note that kubernetes-related values only appear when Sysdig detects Kubernetes metadata.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def comparator(self) -> Optional[pulumi.Input[str]]:
        """
        Operator to relate the metric with some value. It is only required if the value to filter by is set, or the variable field is not set. Valid values are: `in`, `notIn`, `equals`, `notEquals`, `contains`, `notContains` and `startsWith`.
        """
        return pulumi.get(self, "comparator")

    @comparator.setter
    def comparator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comparator", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values to filter by, if comparator is set. If the comparator is not `in` or `notIn` the list must contain only 1 value.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def variable(self) -> Optional[pulumi.Input[str]]:
        """
        Assigns this metric to a value name and allows PromQL to reference it.
        """
        return pulumi.get(self, "variable")

    @variable.setter
    def variable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable", value)


@pulumi.input_type
class TeamEntrypointArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 selection: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Main entrypoint for the team.
               Valid options are: Explore, Dashboards, Events, Alerts, Settings.
        :param pulumi.Input[str] selection: Sets up the defined Dashboard name as entrypoint.
               Warning: This field must only be added if the `type` is "Dashboards".
        """
        pulumi.set(__self__, "type", type)
        if selection is not None:
            pulumi.set(__self__, "selection", selection)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Main entrypoint for the team.
        Valid options are: Explore, Dashboards, Events, Alerts, Settings.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def selection(self) -> Optional[pulumi.Input[str]]:
        """
        Sets up the defined Dashboard name as entrypoint.
        Warning: This field must only be added if the `type` is "Dashboards".
        """
        return pulumi.get(self, "selection")

    @selection.setter
    def selection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "selection", value)


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


