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

// PublicIP is an Azure Public IP
type PublicIP struct {
	Name                 string
	ResourceGroupName    string
	Location             string
	AllocationMethod     network.IPAllocationMethod
	Tags                 map[string]*string
	IPVersion            network.IPVersion
	IdleTimeoutInMinutes *int32
	Sku                  network.PublicIPAddressSkuName
	//TODO  DomainNameLabel      string
	// ReverseFqdn network.PublicIPAddress
	Zones *[]string
}

const publicIPAddressTemplate = `
resource "azurerm_public_ip" "{{.Name}}" {
	name	=	"{{.Name}}"
	resource_group	=	"{{.ResourceGroupName}}"
	location	=	"{{.Location}}"
	sku	=	"{{.Sku}}"
	allocation_method	=	"{{.AllocationMethod}}"
	ip_version	=	"{{.IPVersion}}"
	idle_timeout_in_minutes	=	"{{.IdleTimeoutInMinutes}}"
	{{if not .Tags}}{{else}}
	tags	=	{
		{{$first := true}}{{range $key, $value := .Tags}}{{if $first}}{{$first = false}}{{else}},{{end}}
		{{$key}}	:	{{$value}}{{end}}
	}
	{{end}}
	{{if not .Zones }}{{else}}
	zones	=	"{{range $zone := .Zones}} {{$zone}} {{end}}"{{end}}
}
`

func createIPClient(sess *azureauth.AzureSession) network.PublicIPAddressesClient {
	ipClient := network.NewPublicIPAddressesClient(sess.SubscriptionID)
	ipClient.Authorizer = sess.Authorizer
	return ipClient
}

func getPublicIP(ctx context.Context, sess *azureauth.AzureSession, resourceGroupName string, ipName string) (network.PublicIPAddress, error) {
	ipClient := createIPClient(sess)
	return ipClient.Get(ctx, resourceGroupName, ipName, "")
}

func getPublicIPStruct(resourceGroupName string, ipName string) PublicIP {
	sess, err := azureauth.NewSessionFromFile()

	if err != nil {
		log.Fatalf("Failed to initialize authorized session: %v\n", err)
	}

	foundIP, err := getPublicIP(context.Background(), sess, resourceGroupName, ipName)

	if err != nil {
		log.Fatalf("Failed to get Public IP: %v\n", err)
	}

	publicIP := PublicIP{
		Name:                 *foundIP.Name,
		ResourceGroupName:    resourceGroupName,
		Location:             *foundIP.Location,
		Sku:                  foundIP.Sku.Name,
		AllocationMethod:     foundIP.PublicIPAllocationMethod,
		IPVersion:            foundIP.PublicIPAddressVersion,
		IdleTimeoutInMinutes: foundIP.IdleTimeoutInMinutes,
		Tags:                 foundIP.Tags,
		// ReverseFqdn:          foundIP.ReverseFqdn,
	}

	return publicIP

}

// CreateTerraPublicIPAddress creates a terraform resource defining an Azure PublicIPAddress object
func CreateTerraPublicIPAddress(resourceGroupName string, ipName string) {
	publicIP := getPublicIPStruct(resourceGroupName, ipName)

	tmpl := template.Must(template.New("public-ip").Parse(publicIPAddressTemplate))
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	if err := tmpl.Execute(w, publicIP); err != nil {
		log.Fatalf("Failed to parse PublicIPAddress to Terraform resource: %v\n", err)
	}
	w.Flush()
}
