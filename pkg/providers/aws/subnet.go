package aws

import (
	"context"
	"sync"

	"github.com/DuskEagle/landscape/pkg/providers"
	"github.com/DuskEagle/landscape/pkg/resource"
	"github.com/DuskEagle/landscape/pkg/types"
)

type SubnetArgs struct {
	Name      types.StringInput
	VPC       types.StringInput
	CIDRRange types.StringInput
}

type SubnetOutput struct {
	ResourceID types.ResourceID
	wg         *sync.WaitGroup
	ID         types.StringOutput
	Name       types.StringOutput
	VPC        types.StringOutput
	CIDRRange  types.StringOutput
}

type subnetInternal struct {
	ID        string
	Name      string
	VPC       string
	CIDRRange string
}

var _ resource.Resource = &SubnetOutput{}

func (a *AWSProvider) MakeSubnet(
	ctx context.Context,
	id types.ResourceID,
	subnetArgs *SubnetArgs,
	options ...providers.MakeOption,
) (*SubnetOutput, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	internal := &subnetInternal{}
	go func() {
		defer wg.Done()
		// Make AWS call to create GetSubnet here. Populate result into SunbetInternal.
		// The code below is a fake of that process.
		*internal = subnetInternal{
			ID:        "subnet-abcde",
			Name:      subnetArgs.Name.Await(),
			VPC:       subnetArgs.VPC.Await(),
			CIDRRange: subnetArgs.CIDRRange.Await(),
		}
	}()

	return &SubnetOutput{
		wg: &wg,
		ID: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.ID
		}),
		Name: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.Name
		}),
		VPC: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.VPC
		}),
		CIDRRange: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.CIDRRange
		}),
	}, nil
}

func (a *AWSProvider) GetSubnet(ctx context.Context, id types.ResourceID) (*SubnetOutput, error) {
	// TODO(joel): Fetch from provider.
	return &SubnetOutput{}, nil
}

func (a *AWSProvider) DeleteSubnet(ctx context.Context, id types.ResourceID) error {
	// TODO(joel): Implement
	return nil
}

// TODO(joel): Have a way to signal error.
func (subnet *SubnetOutput) Await(ctx context.Context) error {
	subnet.wg.Wait()
	return nil
}
