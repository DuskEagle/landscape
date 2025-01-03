package main

import (
	"context"
	"fmt"
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
		//if err != nil {
		//	provider.Cancel(ctx)
		//} else {
		//	err = provider.AwaitAll(ctx)
		//}
	}()

	vpc, err := provider.MakeVPC(ctx, "myvpc", &aws.VPCArgs{
		Name:      types.String("myvpc"),
		CIDRRange: types.String("10.0.0.0/16"),
	})

	subnet, err := provider.MakeSubnet(ctx, "mysubnet", &aws.SubnetArgs{
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

	subnetFromProvider, err := provider.GetSubnet(ctx, "mysubnet")

	// Now delete the subnet.
	if err := provider.DeleteSubnet(ctx, subnetFromProvider.ResourceID); err != nil {
		return err
	}

	fmt.Println(subnet.ID.Await())

	return nil
}
