package azureresources

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureauth"
)

// VirtualNetwork is an Azure VM
type VirtualNetwork struct {
	Name               string
	ResourceGroupName  string
	AddressSpace       []string
	Location           string
	DDOSProtectionPlan DdosProtectionPlan
	DNSServers         []string
	Subnets            []Subnet
	Tags               map[string]*string
}

// DdosProtectionPlan is a block for an Azure VM
type DdosProtectionPlan struct {
	ID     string
	enable bool
}

// Subnet is a block for an Azure subnet
type Subnet struct {
	Name          string
	AddressPrefix string
	SecurityGroup string
}

func createVnetClient(sess *azureauth.AzureSession) network.VirtualNetworksClient {
	vnetClient := network.NewVirtualNetworksClient(sess.SubscriptionID)
	vnetClient.Authorizer = sess.Authorizer
	return vnetClient
}

func getVnet(ctx context.Context, sess *azureauth.AzureSession, groupName string, vnetName string) (network.VirtualNetwork, error) {
	vnetClient := createVnetClient(sess)
	return vnetClient.Get(ctx, groupName, vnetName, "")
}

func GetVnetStruct(groupName string, vnetName string) VirtualNetwork {

	sess, err := azureauth.NewSessionFromFile()

	if err != nil {
		log.Fatalf("Failed to initialize authorized session: %v\n", err)
	}

	foundVnet, err := getVnet(context.Background(), sess, groupName, vnetName)

	if err != nil {
		log.Fatalf("Failed to get Virtual Network: %v\n", err)
	}

	virtualNetwork := VirtualNetwork{
		Name: *foundVnet.Name,
	}

	return virtualNetwork

}
