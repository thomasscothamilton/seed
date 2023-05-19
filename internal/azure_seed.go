package azure_seed

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
