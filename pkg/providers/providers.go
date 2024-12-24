package providers

import (
	"context"

	"github.com/DuskEagle/landscape/pkg/future"
)

// Not defined yet.
type UpsertOption interface{}

type Provider interface {
	Upsert(context.Context, string, interface{}, ...UpsertOption) (future.Future, error)
}
