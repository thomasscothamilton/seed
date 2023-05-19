package cmd

import (
	"github.com/spf13/cobra"
)

type AzureSeed struct {
	BillingAccountName                 string `yaml:"billing-account-name"`
	BillingProfileName                 string `yaml:"billing-profile-name"`
	InvoiceSectionName                 string `yaml:"invoice-section-name"`
	SubscriptionName                   string `yaml:"subscription-name"`
	ResourceGroupName                  string `yaml:"resource-group-name"`
	ResourceGroupLocation              string `yaml:"resource-group-location"`
	KubernetesClusterName              string `yaml:"kubernetes-cluster-name"`
	KubernetesClusterLocation          string `yaml:"kubernetes-cluster-location"`
	KubernetesClusterDnsPrefix         string `yaml:"kubernetes-cluster-dns-prefix"`
	KubernetesDefaultNodePoolName      string `yaml:"kubernetes-default-node-pool-name"`
	KubernetesDefaultNodePoolNodeCount int    `yaml:"kubernetes-default-node-pool-node-count"`
	KubernetesDefaultNodePoolVmSize    string `yaml:"kubernetes-default-node-pool-vm-size"`
	KubernetesClusterIdentityType      string `yaml:"kubernetes-cluster-identity-type"`
}

var (
	azureSeed AzureSeed
	AzureCmd  = &cobra.Command{
		Use: "azure",
		Run: func(cmd *cobra.Command, args []string) {
			// code you wish you had...
		},
	}
)

func init() {
	AzureCmd.PersistentFlags().StringVar(&azureSeed.BillingAccountName, "billing-account-name", azureSeed.BillingAccountName, "billing-account-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.BillingProfileName, "billing-profile-name", azureSeed.BillingProfileName, "billing-profile-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.InvoiceSectionName, "invoice-section-name", azureSeed.InvoiceSectionName, "invoice-section-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.SubscriptionName, "subscription-name", azureSeed.SubscriptionName, "subscription-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.ResourceGroupName, "resource-group-name", azureSeed.ResourceGroupName, "resource-group-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.ResourceGroupLocation, "resource-group-location", azureSeed.ResourceGroupLocation, "resource-group-location")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterName, "kubernetes-cluster-name", azureSeed.KubernetesClusterName, "kubernetes-cluster-name")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterLocation, "kubernetes-cluster-location", azureSeed.KubernetesClusterLocation, "kubernetes-cluster-location")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterDnsPrefix, "kubernetes-cluster-dns-prefix", azureSeed.KubernetesClusterDnsPrefix, "kubernetes-cluster-dns-prefix")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesDefaultNodePoolName, "kubernetes-default-node-pool-name", azureSeed.KubernetesDefaultNodePoolName, "kubernetes-default-node-pool-name")
	AzureCmd.PersistentFlags().IntVar(&azureSeed.KubernetesDefaultNodePoolNodeCount, "kubernetes-default-node-pool-node-count", 2, "kubernetes-default-node-pool-node-count")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesDefaultNodePoolVmSize, "kubernetes-default-node-pool-vm-size", azureSeed.KubernetesDefaultNodePoolVmSize, "kubernetes-default-node-pool-vm-size")
	AzureCmd.PersistentFlags().StringVar(&azureSeed.KubernetesClusterIdentityType, "kubernetes-cluster-identity-type", azureSeed.KubernetesClusterIdentityType, "kubernetes-cluster-identity-type")

}
