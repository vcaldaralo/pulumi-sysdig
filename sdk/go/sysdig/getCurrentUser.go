// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sysdig

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCurrentUser(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCurrentUserResult, error) {
	var rv GetCurrentUserResult
	err := ctx.Invoke("sysdig:index/getCurrentUser:GetCurrentUser", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by GetCurrentUser.
type GetCurrentUserResult struct {
	// The user email.
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The user's last name.
	LastName string `pulumi:"lastName"`
	// The user's first name.
	Name string `pulumi:"name"`
	// The user's system role.
	SystemRole string `pulumi:"systemRole"`
}
