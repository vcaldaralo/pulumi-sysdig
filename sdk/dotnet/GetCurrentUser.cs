// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig
{
    public static class GetCurrentUser
    {
        public static Task<GetCurrentUserResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCurrentUserResult>("sysdig:index/getCurrentUser:GetCurrentUser", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetCurrentUserResult
    {
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastName;
        public readonly string Name;
        public readonly string SystemRole;

        [OutputConstructor]
        private GetCurrentUserResult(
            string email,

            string id,

            string lastName,

            string name,

            string systemRole)
        {
            Email = email;
            Id = id;
            LastName = lastName;
            Name = name;
            SystemRole = systemRole;
        }
    }
}