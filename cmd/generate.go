package cmd

/*
 * Create the kubernetes seed cluster.
 */

// var (
// 	billingConfig           dataazurermbillingmcaaccountscope.DataAzurermBillingMcaAccountScopeConfig
// 	kubernetesClusterConfig kubernetescluster.KubernetesClusterConfig

// 	providerConfig provider.AzurermProviderConfig

// 	resourceConfig     resourcegroup.ResourceGroupConfig
// 	subscriptionConfig subscription.SubscriptionConfig
// 	createCmd          = &cobra.Command{
// 		Use: "generate [name]",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			// name := args[0]

// 			// // create the app
// 			// app := cdktf.NewApp(&cdktf.AppConfig{})

// 			// // create the stack
// 			// stack := cdktf.NewTerraformStack(app, jsii.String(name))

// 			// // create a provider
// 			// _ = provider.NewAzurermProvider(stack, jsii.String("azure"), &providerConfig)

// 			// // create a billing account
// 			// b := dataazurermbillingmcaaccountscope.NewDataAzurermBillingMcaAccountScope(stack, jsii.String(fmt.Sprintf("%s-billing", name)), &billingConfig)

// 			// // add the billing account id to the subscription config
// 			// subscriptionConfig.BillingScopeId = b.Id()

// 			// // create a subscription
// 			// _ = subscription.NewSubscription(stack, jsii.String(fmt.Sprintf("%s-subscription", name)), &subscriptionConfig)

// 			// // create a resource group
// 			// r := resourcegroup.NewResourceGroup(stack, jsii.String(fmt.Sprintf("%s-resource-group", name)), &resourceConfig)

// 			// // add the resource group name to the kubernetes cluster config
// 			// kubernetesClusterConfig.ResourceGroupName = r.Name()

// 			// // create a kubernetes cluster
// 			// cluster := kubernetescluster.NewKubernetesCluster(stack, jsii.String(fmt.Sprintf("%s-kubernetes-cluster")), &kubernetesClusterConfig)

// 			// // output stuff
// 			// cdktf.NewTerraformOutput(stack, jsii.String("names"), &cdktf.TerraformOutputConfig{
// 			// 	Value: &[]*string{cluster.Name(), r.Name()},
// 			// })

// 			// // run the stack
// 			// app.Synth()
// 		},
// 	}
// )

// func init() {
// 	// //+ provider config
// 	// providerConfig = provider.AzurermProviderConfig{
// 	// 	Features: &provider.AzurermProviderFeatures{},
// 	// }
// 	// //- provider config

// 	// //+ billing config
// 	// billingConfig = dataazurermbillingmcaaccountscope.DataAzurermBillingMcaAccountScopeConfig{
// 	// 	BillingAccountName: jsii.String("1f9055e4-5caf-5d19-6103-6fef9bb70785:4c43b9c9-c07b-4877-91f9-f5a87e34a652_2019-05-31"),
// 	// 	BillingProfileName: jsii.String("YPRK-RHGJ-BG7-PGB"),
// 	// 	InvoiceSectionName: jsii.String("ELBV-HGBS-PJA-PGB"),
// 	// }
// 	// createCmd.Flags().StringVarP(billingConfig.BillingAccountName, "billing-account-name", "", *billingConfig.BillingAccountName, "billing-account-name")
// 	// createCmd.Flags().StringVarP(billingConfig.BillingAccountName, "billing-profile-name", "", *billingConfig.BillingAccountName, "billing-profile-name")
// 	// createCmd.Flags().StringVarP(billingConfig.BillingAccountName, "invoice-section-name", "", *billingConfig.BillingAccountName, "invoice-section-name")
// 	// //- billing config

// 	// //+ subscription config
// 	// subscriptionConfig = subscription.SubscriptionConfig{
// 	// 	SubscriptionName: jsii.String("containerly-default-subscription"),
// 	// 	BillingScopeId:   billingConfig.Id,
// 	// }
// 	// createCmd.Flags().StringVarP(subscriptionConfig.SubscriptionName, "subscription-name", "", *subscriptionConfig.SubscriptionName, "subscription-name")
// 	// //createCmd.Flags().StringVarP(subscriptionConfig.BillingScopeId, "billing-scope-id", "", *subscriptionConfig.BillingScopeId, "billing-scope-id")
// 	// //+ subscription config

// 	// //+ resource config
// 	// resourceConfig = resourcegroup.ResourceGroupConfig{
// 	// 	Name:     jsii.String("containerly-default-resource-group"),
// 	// 	Location: jsii.String("eastus"),
// 	// }
// 	// createCmd.Flags().StringVarP(resourceConfig.Name, "resource-group-name", "", *resourceConfig.Name, "resource-group-name")
// 	// createCmd.Flags().StringVarP(resourceConfig.Location, "resource-group-location", "", *resourceConfig.Location, "resource-group-location")
// 	// //- resource config

// 	// //+ kubernetes cluster config
// 	// kubernetesClusterConfig = kubernetescluster.KubernetesClusterConfig{
// 	// 	Name:      jsii.String("containerly-default-kubernetes-cluster"),
// 	// 	Location:  jsii.String("eastus"),
// 	// 	DnsPrefix: jsii.String("containerly"),
// 	// 	DefaultNodePool: &kubernetescluster.KubernetesClusterDefaultNodePool{
// 	// 		Name:      jsii.String("containerly-default-node-pool"),
// 	// 		NodeCount: jsii.Number(1),
// 	// 		VmSize:    jsii.String("Standard_DS2_v2"),
// 	// 	},
// 	// 	Identity: &kubernetescluster.KubernetesClusterIdentity{Type: jsii.String("SystemAssigned")},
// 	// }
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.Name, "kubernetes-cluster-name", "", *kubernetesClusterConfig.Name, "kubernetes-cluster-name")
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.Location, "kubernetes-cluster-location", "", *kubernetesClusterConfig.Location, "kubernetes-cluster-location")
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.DnsPrefix, "kubernetes-cluster-dns-prefix", "", *kubernetesClusterConfig.DnsPrefix, "kubernetes-cluster-dns-prefix")
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.DefaultNodePool.Name, "kubernetes-cluster-default-node-pool-name", "", *kubernetesClusterConfig.DefaultNodePool.Name, "kubernetes-cluster-default-node-pool-name")
// 	// createCmd.Flags().Float64VarP(kubernetesClusterConfig.DefaultNodePool.NodeCount, "kubernetes-cluster-default-node-pool-node-count", "", *kubernetesClusterConfig.DefaultNodePool.NodeCount, "kubernetes-cluster-default-node-pool-node-count")
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.DefaultNodePool.VmSize, "kubernetes-cluster-default-node-pool-vm-size", "", *kubernetesClusterConfig.DefaultNodePool.VmSize, "kubernetes-cluster-default-node-pool-vm-size")
// 	// createCmd.Flags().StringVarP(kubernetesClusterConfig.Identity.Type, "kubernetes-cluster-identity-type", "", *kubernetesClusterConfig.Identity.Type, "kubernetes-cluster-identity-type")
// 	//- kubernetes cluster config

// 	rootCmd.AddCommand(createCmd)
// }
