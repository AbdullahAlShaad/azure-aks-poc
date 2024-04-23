package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

func main() {
	GetKubernetesVersions()
}

func GetKubernetesVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewManagedClustersClient(os.Getenv("subscriptionId"), cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	resp, err := clientFactory.ListKubernetesVersions(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to list kubernetes versions: %v", err)
	}
	versions := make([]string, 0)
	for _, ver := range resp.Values {
		for key, _ := range (*ver).PatchVersions {
			versions = append(versions, key)
		}
	}
	fmt.Println("Kubernetes Versions: ")
	fmt.Println(versions)
}
