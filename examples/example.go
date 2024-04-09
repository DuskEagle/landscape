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
	vpc, err := provider.VPC(ctx, "myvpc", &aws.VPCArgs{
		Name:      "myvpc",
		CIDRRange: "10.0.0.0/16",
	}).Upsert()

	subnet, err := provider.Subnet(ctx, "mysubnet", &aws.SubnetArgs{
		Name:      "mysubnet",
		VPC:       vpc,
		CIDRRange: "10.0.0.0/24",
	}).Upsert()

	// Explicitly wait for the subnet to finish before continuing. This isn't
	// needed for landscape to work, but could be useful for interacting with
	// other parts of code.
	if err := subnet.Await(ctx); err != nil {
		return err
	}

	// Upsert() might take a variadic list of options, such as .Protect().
	// .Protect would prevent against accidental deletion.
	igw, err := provider.InternetGateway(ctx, "myinternetgateway", &aws.InternetGatewayArgs{}).
		Upsert(landscape.Protect())

	// We don't call Upsert here - that's a mistake. RouteTable will actually
	// be a resource that we can potentially call other methods on.
	// But if we end up not wanting other methods beyond Upsert, then we might
	// drop the Upsert call from all other resources and use this approach.
	routeTable, err := provider.RouteTable(ctx, "myroutetable", &aws.RouteTableArgs{
		Name: "myroutetable",
	})

	route, err := provider.Route(ctx, "route1", &aws.RouteArgs{
		RouteTable:  routeTable,
		Destination: "0.0.0.0/0",
		NextHop:     igw,
	}).Upsert()

	// Now delete the route.
	// Unless provider has some delete method, we need to do this on the project
	// level. Having a GetProject method avoids having to pass both a provider
	// and a Project around. But having to pass the provider as an argument to
	// .Delete is annoying.
	if err := provider.GetProject().Delete("route1", provider); err != nil {
		return err
	}

}
