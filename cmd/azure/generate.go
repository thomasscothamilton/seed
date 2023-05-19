package cmd

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/kubernetescluster"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/provider"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/resourcegroup"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/spf13/cobra"
)

var (
	AzureGenerateCmd = &cobra.Command{
		Use:   "generate [name]",
		Short: "https://github.com/thomasscothamilton/seed#generate",
		Long:  "https://github.com/thomasscothamilton/seed#generate",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			// create the app
			app := cdktf.NewApp(&cdktf.AppConfig{})

			// create the stack
			stack := cdktf.NewTerraformStack(app, jsii.String(name))

			// create a provider
			_ = provider.NewAzurermProvider(stack, jsii.String(fmt.Sprintf("%s-azure", name)), &provider.AzurermProviderConfig{
				Features: &provider.AzurermProviderFeatures{},
			})

			// create a resource group
			r := resourcegroup.NewResourceGroup(stack, jsii.String(azureSeed.ResourceGroupName), &resourcegroup.ResourceGroupConfig{
				Name:     &azureSeed.ResourceGroupName,
				Location: &azureSeed.ResourceGroupLocation,
			})

			// create a kubernetes cluster
			_ = kubernetescluster.NewKubernetesCluster(stack, jsii.String(azureSeed.KubernetesClusterName), &kubernetescluster.KubernetesClusterConfig{
				ResourceGroupName: r.Name(),
				Name:              jsii.String(azureSeed.KubernetesClusterName),
				Location:          jsii.String(azureSeed.KubernetesClusterLocation),
				DnsPrefix:         jsii.String(azureSeed.KubernetesClusterDnsPrefix),
				Identity: &kubernetescluster.KubernetesClusterIdentity{
					Type: jsii.String(azureSeed.KubernetesClusterIdentityType),
				},
				DefaultNodePool: &kubernetescluster.KubernetesClusterDefaultNodePool{
					Name:      jsii.String(azureSeed.KubernetesDefaultNodePoolName),
					NodeCount: jsii.Number(azureSeed.KubernetesDefaultNodePoolNodeCount),
					VmSize:    jsii.String(azureSeed.KubernetesDefaultNodePoolVmSize),
				},
			})

			// synthesize the stack
			app.Synth()
		},
	}
)

func init() {
	AzureCmd.AddCommand(AzureGenerateCmd)
}
