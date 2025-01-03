package aws

import (
	"context"
	"sync"

	"github.com/DuskEagle/landscape/pkg/providers"
	"github.com/DuskEagle/landscape/pkg/resource"
	"github.com/DuskEagle/landscape/pkg/types"
)

type VPCArgs struct {
	Name      types.StringInput
	CIDRRange types.StringInput
}

type VPCOutput struct {
	wg        *sync.WaitGroup
	ID        types.StringOutput
	Name      types.StringOutput
	CIDRRange types.StringOutput
}

type vpcInternal struct {
	ID        string
	Name      string
	CIDRRange string
}

var _ resource.Resource = &VPCOutput{}

func (a *AWSProvider) MakeVPC(
	ctx context.Context,
	id types.ResourceID,
	args *VPCArgs,
	options ...providers.MakeOption,
) (*VPCOutput, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	internal := &vpcInternal{}
	go func() {
		defer wg.Done()
		// Make AWS call to create MakeVPC here. Populate result into VPCInternal.
	}()
	return &VPCOutput{
		ID: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.ID
		}),
		Name: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.Name
		}),
		CIDRRange: types.NewStringOutput(func() string {
			wg.Wait()
			return internal.CIDRRange
		}),
	}, nil
}

// TODO(joel): Have a way to signal error.
func (vpc *VPCOutput) Await(ctx context.Context) error {
	vpc.wg.Wait()
	return nil
}
