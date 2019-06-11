resource "azurerm_resource_group" "rg_test" {
  name     = "testResourceGroup1"
  location = "eastus"

  tags = {
    environment = "Development"
    testTag     = "testTagValue"
  }
}

resource "azurerm_lb" "snn-test-lb" {
	name						= "snn-test-lb"
	resource_group_name					= "testresourcegroup1"
	location					= "eastus"
	sku						= "Basic"

	tags						= {

		another = "one",
		testTagKey = "testTagValue"
	}

	frontend_ip_configuration {
		name= "LoadBalancerFrontEnd"
		private_ip_address_allocation= "Dynamic"
		public_ip_address_id= "/subscriptions/b0879f7b-8c18-4191-a469-f946e8fc4668/resourceGroups/testResourceGroup1/providers/Microsoft.Network/publicIPAddresses/snn-test-ip"

	}
}
