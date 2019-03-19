package cmd

import (
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureresources"
	"github.com/spf13/cobra"
)

var resourceGroupName string
var ipName string

// publicipCmd represents the publicip command
var publicipCmd = &cobra.Command{
	Use:   "publicip",
	Short: "Generate terraform resource for an Azure Public IP Address",
	Long: `By providing the name of an Azure resource group one can quickly generate the associated terraform resource:

	azure-terraform-generator publicip --i test-public-ip -g TestResourceGroupName
	
	output: 
	
resource "azurerm_public_ip" "test-public-ip" {
	name                                            = "test-public-ip"
	resource_group                                  = "TestResourceGroupName"
	location                                        = "westeurope"
	allocation_method                               = "Static"
	ip_version                                      = "IPv4"
	idle_timeout_in_minutes                         = "4"


}`,
	Run: func(cmd *cobra.Command, args []string) {
		azureresources.CreateTerraPublicIPAddress(resourceGroupName, ipName)
	},
}

func init() {
	publicipCmd.Flags().StringVarP(&resourceGroupName, "group-name", "g", "", "name of resource group")
	publicipCmd.MarkFlagRequired("group-name")

	publicipCmd.Flags().StringVarP(&ipName, "name", "n", "", "name of public ip address")
	publicipCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(publicipCmd)

}
