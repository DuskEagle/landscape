package gcs

import "github.com/DuskEagle/landscape/pkg/backend"

type gcsBackend struct{}

var _ backend.Backend = &gcsBackend{}

func NewGCSBackend(path string) (backend.Backend, error) {
	return &gcsBackend{}, nil
}
