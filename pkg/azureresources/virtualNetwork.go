package azureresources

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
