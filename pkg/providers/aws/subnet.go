package aws

import (
	"context"
	"sync"

	"github.com/DuskEagle/landscape/pkg/resource"
	"github.com/DuskEagle/landscape/pkg/types"
)

type SubnetArgs struct {
	Name      types.StringInput
	VPC       types.StringInput
	CIDRRange types.StringInput
}

type SubnetOutput struct {
	wg        *sync.WaitGroup
	ID        types.StringOutput
	Name      types.StringOutput
	VPC       types.StringOutput
	CIDRRange types.StringOutput
}

type subnetInternal struct {
	ID        string
	Name      string
	VPC       string
	CIDRRange string
}

var _ resource.Resource = &SubnetOutput{}

func (a *AWSProvider) Subnet(ctx context.Context, id string, args *SubnetArgs) (*SubnetOutput, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	internal := &subnetInternal{}
	go func() {
		defer wg.Done()
		// Make AWS call to create Subnet here. Populate result into SunbetInternal.
	}()
	return &SubnetOutput{
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

// TODO(joel): Have a way to signal error.
func (subnet *SubnetOutput) Await(ctx context.Context) error {
	subnet.wg.Wait()
	return nil
}
