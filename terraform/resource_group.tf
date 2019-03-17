resource "azurerm_resource_group" "rg_test" {
    name     = "testResourceGroup1"
    location = "East US"
    tags = {
        environment = "Development",
        testTag = "testTagValue"
    }
}