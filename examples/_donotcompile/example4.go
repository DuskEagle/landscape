package main

import (
	"context"
	"log"

	landscape "github.com/DuskEagle/landscape/pkg"
	"github.com/DuskEagle/landscape/pkg/backend/gcs"
	"github.com/DuskEagle/landscape/pkg/providers/aws"
	"github.com/DuskEagle/landscape/pkg/types"
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
	provider, err := aws.NewProvider(project)
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

	// With this "Upsert" method, we must pattern-match the arg to the resource.
	// Create a MakeVPC. Upsert() begins creating the MakeVPC and returns a promise.
	// Fields on the promise will cause other
	vpc, err := provider.MakeVPC(ctx, "myvpc", &aws.VPCArgs{
		Name:      types.String("myvpc"),
		CIDRRange: types.String("10.0.0.0/16"),
	})

	subnet, err := provider.GetSubnet(ctx, "mysubnet", &aws.SubnetArgs{
		Name:      types.String("mysubnet"),
		VPC:       vpc.ID,
		CIDRRange: types.String("10.0.0.0/24"),
	})

	// Explicitly wait for the subnet to finish before continuing. This isn't
	// needed for landscape to work, but could be useful for interacting with
	// other parts of code.
	if err := subnet.Await(ctx); err != nil {
		return err
	}

	// Upsert() might take a variadic list of options, such as .Protect().
	// .Protect would prevent against accidental deletion.
	igw, err := provider.InternetGateway(ctx, "myinternetgateway", &aws.InternetGatewayArgs{}, landscape.Protect())

	routeTable, err := provider.RouteTable(ctx, "myroutetable", &aws.RouteTableArgs{
		Name: "myroutetable",
	})

	route, err := provider.Route(ctx, "route1", &aws.RouteArgs{
		RouteTable:  routeTable,
		Destination: "0.0.0.0/0",
		NextHop:     igw,
	})

	// Now delete the route.
	// Thought: Maybe we'd want to be able to import resources if they exist,
	// not create them if they don't exist, so they could be deleted by
	// landscape without being created by it? Anyway, this is a long way off.
	if err := provider.Delete("route1", provider); err != nil {
		return err
	}
}
