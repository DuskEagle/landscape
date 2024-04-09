package main

import (
	"context"
	"log"

	landscape "github.com/DuskEagle/landscape/pkg"
	"github.com/DuskEagle/landscape/pkg/backend/gcs"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) (err error) {
	backend, err := gcs.NewGCSBackend("gs://my-bucket")
	if err != nil {
		return err
	}

	project, err := landscape.NewProject("my-example-project", backend)
	if err != nil {
		return err
	}

	// Provider is initiated with project.
	provider, err := landscape.NewAWSProvider(project)
	if err != nil {
		return err
	}

	// If the rest of the function exited cleanly, finish resolving all promises
	// before exiting this code block. Otherwise, cancel any ongoing promises
	// and return the original error.
	defer func() {
		if err != nil {
			provider.Cancel(ctx)
		} else {
			err = provider.AwaitAll(ctx)
		}
	}()

	// Create a VPC. Upsert() begins creating the VPC and returns a promise.
	// Fields on the promise will cause other
	vpc, err := provider.Upsert(ctx, "myvpc", &aws.VPC{
		Name:      "myvpc",
		CIDRRange: "10.0.0.0/16",
	})

	subnet, err := provider.Upsert(ctx, "mysubnet", &aws.Subnet{
		Name:      "mysubnet",
		VPC:       vpc,
		CIDRRange: "10.0.0.0/24",
	})

	// Explicitly wait for the subnet to finish before continuing. This isn't
	// needed for landscape to work, but could be useful for interacting with
	// other parts of code.
	if err := subnet.Await(ctx); err != nil {
		return err
	}

	// Upsert() might take a variadic list of options, such as .Protect().
	// .Protect would prevent against accidental deletion.
	igw, err := provider.Upsert(ctx, "myinternetgateway", &aws.InternetGateway{}, landscape.Protect())

	routeTable, err := provider.Upsert(ctx, "myroutetable", &aws.RouteTable{
		Name: "myroutetable",
	})

	route, err := provider.Upsert(ctx, "route1", &aws.Route{
		RouteTable:  routeTable,
		Destination: "0.0.0.0/0",
		NextHop:     igw,
	})

	// Now delete the route.
	// WIP on how this should be exposed.
	// Unless provider has some delete method, we need to do this on the project
	// level. Having a GetProject method avoids having to pass both a provider
	// and a Project around. But having to pass the provider as an argument to
	// .Delete is annoying.
	if err := provider.GetProject().Delete("route1", provider); err != nil {
		return err
	}

}
