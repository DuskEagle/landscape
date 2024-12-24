package resource

import "context"

type Resource interface {
	Await(context.Context) error
}
