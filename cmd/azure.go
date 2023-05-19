package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomasscothamilton/seed/internal/azure_seed"
)

var (
	azureSeed azure_seed.AzureSeed
	azureCmd  = &cobra.Command{
		Use: "azure",
		Run: func(cmd *cobra.Command, args []string) {
			// code you wish you had...
		},
	}
)

func init() {
	azureCmd.PersistentFlags().StringVar(&azureSeed.BillingAccountName, "billing-account-name", azureSeed.BillingAccountName, "billing-account-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.BillingProfileName, "billing-profile-name", azureSeed.BillingProfileName, "billing-profile-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.InvoiceSectionName, "invoice-section-name", azureSeed.InvoiceSectionName, "invoice-section-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.SubscriptionName, "subscription-name", azureSeed.SubscriptionName, "subscription-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.ResourceGroupName, "resource-group-name", azureSeed.ResourceGroupName, "resource-group-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.ResourceGroupLocation, "resource-group-location", azureSeed.ResourceGroupLocation, "resource-group-location")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterName, "kubernetes-cluster-name", azureSeed.KubernetesClusterName, "kubernetes-cluster-name")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterLocation, "kubernetes-cluster-location", azureSeed.KubernetesClusterLocation, "kubernetes-cluster-location")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterDnsPrefix, "kubernetes-cluster-dns-prefix", azureSeed.KubernetesClusterDnsPrefix, "kubernetes-cluster-dns-prefix")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesDefaultNodePoolName, "kubernetes-default-node-pool-name", azureSeed.KubernetesDefaultNodePoolName, "kubernetes-default-node-pool-name")
	azureCmd.PersistentFlags().IntVar(&azureSeed.KubernetesDefaultNodePoolNodeCount, "kubernetes-default-node-pool-node-count", 2, "kubernetes-default-node-pool-node-count")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesDefaultNodePoolVmSize, "kubernetes-default-node-pool-vm-size", azureSeed.KubernetesDefaultNodePoolVmSize, "kubernetes-default-node-pool-vm-size")
	azureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterIdentityType, "kubernetes-cluster-identity-type", azureSeed.KubernetesClusterIdentityType, "kubernetes-cluster-identity-type")
	rootCmd.AddCommand(azureCmd)
}
