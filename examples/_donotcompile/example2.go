package main

import (
	"context"
	"errors"
	"log"

	landscape "github.com/DuskEagle/landscape/pkg"
	"github.com/DuskEagle/landscape/pkg/backend/gcs"
	"github.com/DuskEagle/landscape/pkg/providers/aws"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	backend, err := gcs.NewGCSBackend("gs://bucket")
	if err != nil {
		return err
	}
	project, err := landscape.NewProject("myproject", backend)
	if err != nil {
		return err
	}

	// Provider is initiated with project.
	provider, err := aws.NewProvider(project)
	if err != nil {
		return err
	}

	vpcIface, err := provider.Get("myvpc")
	if err != nil {
		return err
	}
	vpc, ok := vpcIface.(*ec2.Vpc)
	if !ok {
		return errors.New("invalid cast")
	}

	subnet, err := provider.NewSubnet(ctx, "mysubnet", &ec2.SubnetArgs{
		Vpc:              vpc.Id(),
		AvailabilityZone: "us-east-1a",
	})
	if err != nil {
		return err
	}

}
