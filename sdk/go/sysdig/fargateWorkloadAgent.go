// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sysdig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func FargateWorkloadAgent(ctx *pulumi.Context, args *FargateWorkloadAgentArgs, opts ...pulumi.InvokeOption) (*FargateWorkloadAgentResult, error) {
	var rv FargateWorkloadAgentResult
	err := ctx.Invoke("sysdig:index/fargateWorkloadAgent:FargateWorkloadAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking FargateWorkloadAgent.
type FargateWorkloadAgentArgs struct {
	CollectorHost        *string `pulumi:"collectorHost"`
	CollectorPort        *string `pulumi:"collectorPort"`
	ContainerDefinitions string  `pulumi:"containerDefinitions"`
	ImageAuthSecret      *string `pulumi:"imageAuthSecret"`
	OrchestratorHost     *string `pulumi:"orchestratorHost"`
	OrchestratorPort     *string `pulumi:"orchestratorPort"`
	SysdigAccessKey      string  `pulumi:"sysdigAccessKey"`
	WorkloadAgentImage   string  `pulumi:"workloadAgentImage"`
}

// A collection of values returned by FargateWorkloadAgent.
type FargateWorkloadAgentResult struct {
	CollectorHost        *string `pulumi:"collectorHost"`
	CollectorPort        *string `pulumi:"collectorPort"`
	ContainerDefinitions string  `pulumi:"containerDefinitions"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string  `pulumi:"id"`
	ImageAuthSecret            *string `pulumi:"imageAuthSecret"`
	OrchestratorHost           *string `pulumi:"orchestratorHost"`
	OrchestratorPort           *string `pulumi:"orchestratorPort"`
	OutputContainerDefinitions string  `pulumi:"outputContainerDefinitions"`
	SysdigAccessKey            string  `pulumi:"sysdigAccessKey"`
	WorkloadAgentImage         string  `pulumi:"workloadAgentImage"`
}

func FargateWorkloadAgentOutput(ctx *pulumi.Context, args FargateWorkloadAgentOutputArgs, opts ...pulumi.InvokeOption) FargateWorkloadAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FargateWorkloadAgentResult, error) {
			args := v.(FargateWorkloadAgentArgs)
			r, err := FargateWorkloadAgent(ctx, &args, opts...)
			return *r, err
		}).(FargateWorkloadAgentResultOutput)
}

// A collection of arguments for invoking FargateWorkloadAgent.
type FargateWorkloadAgentOutputArgs struct {
	CollectorHost        pulumi.StringPtrInput `pulumi:"collectorHost"`
	CollectorPort        pulumi.StringPtrInput `pulumi:"collectorPort"`
	ContainerDefinitions pulumi.StringInput    `pulumi:"containerDefinitions"`
	ImageAuthSecret      pulumi.StringPtrInput `pulumi:"imageAuthSecret"`
	OrchestratorHost     pulumi.StringPtrInput `pulumi:"orchestratorHost"`
	OrchestratorPort     pulumi.StringPtrInput `pulumi:"orchestratorPort"`
	SysdigAccessKey      pulumi.StringInput    `pulumi:"sysdigAccessKey"`
	WorkloadAgentImage   pulumi.StringInput    `pulumi:"workloadAgentImage"`
}

func (FargateWorkloadAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateWorkloadAgentArgs)(nil)).Elem()
}

// A collection of values returned by FargateWorkloadAgent.
type FargateWorkloadAgentResultOutput struct{ *pulumi.OutputState }

func (FargateWorkloadAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateWorkloadAgentResult)(nil)).Elem()
}

func (o FargateWorkloadAgentResultOutput) ToFargateWorkloadAgentResultOutput() FargateWorkloadAgentResultOutput {
	return o
}

func (o FargateWorkloadAgentResultOutput) ToFargateWorkloadAgentResultOutputWithContext(ctx context.Context) FargateWorkloadAgentResultOutput {
	return o
}

func (o FargateWorkloadAgentResultOutput) CollectorHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) *string { return v.CollectorHost }).(pulumi.StringPtrOutput)
}

func (o FargateWorkloadAgentResultOutput) CollectorPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) *string { return v.CollectorPort }).(pulumi.StringPtrOutput)
}

func (o FargateWorkloadAgentResultOutput) ContainerDefinitions() pulumi.StringOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) string { return v.ContainerDefinitions }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o FargateWorkloadAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o FargateWorkloadAgentResultOutput) ImageAuthSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) *string { return v.ImageAuthSecret }).(pulumi.StringPtrOutput)
}

func (o FargateWorkloadAgentResultOutput) OrchestratorHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) *string { return v.OrchestratorHost }).(pulumi.StringPtrOutput)
}

func (o FargateWorkloadAgentResultOutput) OrchestratorPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) *string { return v.OrchestratorPort }).(pulumi.StringPtrOutput)
}

func (o FargateWorkloadAgentResultOutput) OutputContainerDefinitions() pulumi.StringOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) string { return v.OutputContainerDefinitions }).(pulumi.StringOutput)
}

func (o FargateWorkloadAgentResultOutput) SysdigAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) string { return v.SysdigAccessKey }).(pulumi.StringOutput)
}

func (o FargateWorkloadAgentResultOutput) WorkloadAgentImage() pulumi.StringOutput {
	return o.ApplyT(func(v FargateWorkloadAgentResult) string { return v.WorkloadAgentImage }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FargateWorkloadAgentResultOutput{})
}
