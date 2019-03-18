package azureresources

import (
	"context"
	"os"
	"text/template"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureauth"
)

// ResourceGroup is an Azure resource group
type ResourceGroup struct {
	Name     string
	Location string
	Tags     map[string]*string
}

// TODO: text/template support for tags
const resourceTemplate = `
resource "azurerm_resource_group" "{{.Name}}" {
	name			= "{{.Name}}"
	location		= "{{.Location}}"
}
`

// createGroupsClient returns a new client for interacting w/ Azure resources & resource groups
func createGroupsClient(sess *azureauth.AzureSession) resources.GroupsClient {
	groupsClient := resources.NewGroupsClient(sess.SubscriptionID)
	groupsClient.Authorizer = sess.Authorizer
	return groupsClient
}

// getGroup gets info on the resource group in use
func getGroup(ctx context.Context, sess *azureauth.AzureSession, groupName string) (resources.Group, error) {
	groupsClient := createGroupsClient(sess)
	return groupsClient.Get(ctx, groupName)
}

// getResourceGroup returns the above struct definition of a ResourceGroup
func getResourceGroup(resourceGroup string) ResourceGroup {
	sess, err := azureauth.NewSessionFromFile()

	if err != nil {
		panic(err)
	}

	returnedGroup, err := getGroup(context.Background(), sess, resourceGroup)

	if err != nil {
		panic(err)
	}

	group := ResourceGroup{Name: *returnedGroup.Name, Location: *returnedGroup.Location, Tags: returnedGroup.Tags}
	return group
}

// CreateTerraResourceGroup creates a terraform resource defining the desired azure resource group
// TODO: Handle map[string]*string in templating
func CreateTerraResourceGroup(name string) {

	resourceGroup := getResourceGroup(name)

	tmpl, err := template.New("test").Parse(resourceTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, resourceGroup)
	if err != nil {
		panic(err)
	}

}
