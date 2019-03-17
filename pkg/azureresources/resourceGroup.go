package azureresources

import (
	"os"
	"text/template"
)

// ResourceGroup is an Azure resource group
type ResourceGroup struct {
	Name     string
	Location string
	Tags     map[string]string
}

const resourceTemplate = `
resource "azurerm_resource_group" "{{.Name}}" {
	name			= "{{.Name}}"
	location		= "{{.Location}}"
}
`

// TestTemplate is a test function for use during development. It will be removed
// TODO: Move to test package
func TestTemplate() {
	testGroup := ResourceGroup{"testName", "testLocation", nil}
	tmpl, err := template.New("test").Parse(resourceTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, testGroup)
	if err != nil {
		panic(err)
	}
}
