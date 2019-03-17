package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureauth"
	"github.com/snimmagadda1/azure-terraform-generator/pkg/azureresources"

	"github.com/pkg/errors"
)

func getGroups(sess *azureauth.AzureSession) ([]string, error) {
	tab := make([]string, 0)
	var err error

	grClient := resources.NewGroupsClient(sess.SubscriptionID)
	grClient.Authorizer = sess.Authorizer

	for list, err := grClient.ListComplete(context.Background(), "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			return nil, errors.Wrap(err, "error traverising RG list")
		}
		rgName := *list.Value().Name
		tab = append(tab, rgName)
	}
	return tab, err
}

func main() {
	fmt.Println("Development in progress!")
	azureresources.TestTemplate()

	os.Setenv("AZURE_AUTH_LOCATION", "/Users/snimmag6/go/src/github.com/snimmagadda1/azure-terraform-generator/my.auth")

	sess, err := azureauth.NewSessionFromFile()

	groups, err := getGroups(sess)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, group := range groups {
		fmt.Println(group)
	}
}
