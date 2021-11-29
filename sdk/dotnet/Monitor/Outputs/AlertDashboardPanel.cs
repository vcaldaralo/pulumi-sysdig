// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor.Outputs
{

    [OutputType]
    public sealed class AlertDashboardPanel
    {
        public readonly bool? AutosizeText;
        public readonly string? Content;
        public readonly string? Description;
        public readonly int Height;
        public readonly string Name;
        public readonly int PosX;
        public readonly int PosY;
        public readonly ImmutableArray<Outputs.AlertDashboardPanelQuery> Queries;
        public readonly bool? TransparentBackground;
        public readonly string Type;
        public readonly bool? VisibleTitle;
        public readonly int Width;

        [OutputConstructor]
        private AlertDashboardPanel(
            bool? autosizeText,

            string? content,

            string? description,

            int height,

            string name,

            int posX,

            int posY,

            ImmutableArray<Outputs.AlertDashboardPanelQuery> queries,

            bool? transparentBackground,

            string type,

            bool? visibleTitle,

            int width)
        {
            AutosizeText = autosizeText;
            Content = content;
            Description = description;
            Height = height;
            Name = name;
            PosX = posX;
            PosY = posY;
            Queries = queries;
            TransparentBackground = transparentBackground;
            Type = type;
            VisibleTitle = visibleTitle;
            Width = width;
        }
    }
}
