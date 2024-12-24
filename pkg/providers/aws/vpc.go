package aws

import (
	"context"
	"sync"

	"github.com/DuskEagle/landscape/pkg/resource"
	"github.com/DuskEagle/landscape/pkg/types"
)

type VPCArgs struct {
	Name      types.StringOutput
	CIDRRange types.StringOutput
}

type VPCOutput struct {
	wg *sync.WaitGroup
	ID types.StringOutput
}

var _ resource.Resource = &VPCOutput{}

func (a *AWSProvider) VPC(ctx context.Context, id string, args *VPCArgs) (*VPCOutput, error) {
	//wg := a.GetWaitGroupForResource() // TODO(joel): ?
	wg := sync.WaitGroup{}
	v := &VPCOutput{
		wg: &wg,
	}
	v.ID = types.StringInternal(v, func(r resource.Resource) string {
		_ = v.Await(ctx)
		return v.ID // Need a new type here?
	})
	return v, nil
}

func (vpc *VPCOutput) Await(ctx context.Context) error {
	return nil
}
