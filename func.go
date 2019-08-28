package main

import (
	"github.com/oracle/oci-go-sdk/common"
	"context"
	_ "encoding/json"
	"fmt"
	"os"
	"io"

	fdk "github.com/fnproject/fdk-go"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/common/auth"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
	// reader := os.Stdin
	// writer := os.Stdout
	// myHandler(context.TODO(), reader, writer)
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {

	provider, err := auth.ResourcePrincipalConfigurationProvider()

	if err != nil {
		fmt.Println(err)
		return
	}

	compartmentID := os.Getenv("COMPARTMENT_ID")

	request := core.ListInstancesRequest{
		CompartmentId: &compartmentID,
		LifecycleState: core.InstanceLifecycleStateRunning,
	}

	fmt.Println(request)

	client, err := core.NewComputeClientWithConfigurationProvider(provider)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Override the region, this is an optional step.
	// the InstancePrincipalsConfigurationProvider defaults to the region
	// in which the compute instance is currently running
	client.SetRegion(string(common.RegionAPTokyo1))

	listInstancesResponse, err := client.ListInstances(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range listInstancesResponse.Items {
		fmt.Printf("list of Compute Instance: %s \n", *item.DisplayName)
	}
}
