package azureresources

import (
	"context"
	"log"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureauth"
)

// LoadBalancer is an Azure Load Balancer
type LoadBalancer struct {
	Name              string
	ResourceGroupName string
	Location          string
	Sku               network.LoadBalancerSkuName
	Tags              map[string]*string
	// Frontend IP Configuration properties
	FIPName                       string
	FIPSubnetID                   string
	FIPPrivateIPAddress           string
	FIPPrivateIPAddressAllocation network.IPAllocationMethod
	FIPPublicIPAddressID          string
	// TODO is this available via api for frontend ip config block?
	// Zones                         *[]string
}

const loadBalancerTemplate = `
resource "azurerm_lb" "{{.Name}}" {
	name	=	"{{.Name}}"
	resource_group_name	=	"{{.ResourceGroupName}}"
	location	=	"{{.Location}}"
	sku	=	"{{.Sku}}"
	{{if not .Tags}}{{else}}
	tags	=	{
		{{$first := true}}{{range $key, $value := .Tags}}{{if $first}}{{$first = false}}{{else}},{{end}}
		{{$key}} = "{{$value}}"{{end}}
	}
	{{end}}
	frontend_ip_configuration {
		name	=	"{{.FIPName}}"
		private_ip_address_allocation	=	"{{.FIPPrivateIPAddressAllocation}}"
		public_ip_address_id	=	"{{.FIPPublicIPAddressID}}"
		{{if not .FIPSubnetID}}{{else}}
		subnet_id	=	"{{.FIPSubnetID}}"
	}{{end}}
	}
}
`

func createLBClient(sess *azureauth.AzureSession) network.LoadBalancersClient {
	lbClient := network.NewLoadBalancersClient(sess.SubscriptionID)
	lbClient.Authorizer = sess.Authorizer
	return lbClient
}

func getLoadBalancer(ctx context.Context, sess *azureauth.AzureSession, groupName string, lbName string) (network.LoadBalancer, error) {
	lbClient := createLBClient(sess)
	return lbClient.Get(ctx, groupName, lbName, "")
}

func getLBStruct(groupName string, lbName string) LoadBalancer {
	sess, err := azureauth.NewSessionFromFile()

	if err != nil {
		log.Fatalf("Failed to initialize authorized session: %v\n", err)
	}

	foundLb, err := getLoadBalancer(context.Background(), sess, groupName, lbName)

	if err != nil {
		log.Fatalf("Failed to get Load Balancer: %v\n", err)
	}

	loadBalancer := LoadBalancer{
		Name:              *foundLb.Name,
		ResourceGroupName: groupName,
		Location:          *foundLb.Location,
		Sku:               foundLb.Sku.Name,
		Tags:              foundLb.Tags,
	}

	// Get properties object of foundLb
	lbProperties := foundLb.LoadBalancerPropertiesFormat

	// Get FrontendIPConfigurations
	frontendIPConfigurations := lbProperties.FrontendIPConfigurations

	if len(*frontendIPConfigurations) > 1 {
		log.Fatal("Only a single frontend_ip_configuration block currently supported")
	}

	loadBalancer.FIPName = *(*frontendIPConfigurations)[0].Name
	frontendIPSubnet := (*frontendIPConfigurations)[0].FrontendIPConfigurationPropertiesFormat.Subnet
	if frontendIPSubnet != nil {
		loadBalancer.FIPSubnetID = *frontendIPSubnet.ID
	}
	frontendIPProperties := (*frontendIPConfigurations)[0].FrontendIPConfigurationPropertiesFormat
	if frontendIPProperties.PrivateIPAddress != nil {
		loadBalancer.FIPPrivateIPAddress = *frontendIPProperties.PrivateIPAddress
	}
	loadBalancer.FIPPrivateIPAddressAllocation = frontendIPProperties.PrivateIPAllocationMethod
	loadBalancer.FIPPublicIPAddressID = *frontendIPProperties.PublicIPAddress.ID

	return loadBalancer
}

// CreateTerraLoadBalancer creates a terraform resource defining an Azure LoadBalancer object
func CreateTerraLoadBalancer(resourceGroup string, lbName string) {
	loadBalancer := getLBStruct(resourceGroup, lbName)
	tmpl := template.Must(template.New("loadbalancer").Parse(loadBalancerTemplate))
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	if err := tmpl.Execute(w, loadBalancer); err != nil {
		log.Fatalf("Failed to parse PublicIPAddress to Terraform resource: %v\n", err)
	}
	w.Flush()
}
