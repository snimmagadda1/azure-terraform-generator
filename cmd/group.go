package cmd

import (
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureresources"
	"github.com/spf13/cobra"
)

var groupName string

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Generate terraform resource for an Azure ResourceGroup",
	Long: `By providing the name of an Azure resource group one can quickly generate the associated terraform resource:

azure-terraform-generate --name TestResourceGroupName

output: 

resource "azurerm_resource_group" "TestResourceGroupName" {
	name     = "TestResourceGroupName"
	location = "East US"
	tags = {
		key1 : val1, 
		key2 : val2
	}
}`,
	Run: func(cmd *cobra.Command, args []string) {
		azureresources.CreateTerraResourceGroup(groupName)
	},
}

func init() {
	groupCmd.Flags().StringVarP(&groupName, "name", "n", "", "name of resource group")
	groupCmd.MarkFlagRequired("name")

	RootCmd.AddCommand(groupCmd)

}
