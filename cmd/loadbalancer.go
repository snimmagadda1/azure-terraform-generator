package cmd

import (
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureresources"

	"github.com/spf13/cobra"
)

var lbResourceGroupName string
var lbName string

// loadbalancerCmd represents the loadbalancer command
var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Generate a terraform resource for an Azure load balancer",
	Long: `By providing the name of the load balancer and its resource group one can quickly generate the associated terraform resource. For example:

	azure-terraform-generator loadbalancer -n test-loadbalancer -t TestResourceGroupName

	output:

resource "azurerm_lb" "test-loadbalancer" {
  name                   =    "test-loadbalancer"
  resource_group_name    =    "TestResourceGroupName"
  location               =    "eastus"
  sku                    =    "Basic"

  tags    =    {

          environment = "development",
          testTagKey = "testTagValue"
  }

  frontend_ip_configuration {
      name                             =    "LoadBalancerFrontEnd"
      private_ip_address_allocation    =    "Dynamic"
      public_ip_address_id             =    "/subscriptions/{subscription-id}/resourceGroups/TestResourceGroupName/providers/Microsoft.Network/publicIPAddresses/test-public-ip"

  }
}
`,
	Run: func(cmd *cobra.Command, args []string) {
		azureresources.CreateTerraLoadBalancer(lbResourceGroupName, lbName)
	},
}

func init() {
	loadbalancerCmd.Flags().StringVarP(&lbResourceGroupName, "group-name", "g", "", "name of resource group")
	loadbalancerCmd.MarkFlagRequired("group-name")

	loadbalancerCmd.Flags().StringVarP(&lbName, "name", "n", "", "name of load balancer")
	loadbalancerCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(loadbalancerCmd)
}
