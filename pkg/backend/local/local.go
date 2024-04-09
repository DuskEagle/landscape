package local

import (
	"errors"
	"fmt"
	"os"

	"github.com/DuskEagle/landscape/pkg/backend"
)

type localBackend struct {
	path string
}

var _ backend.Backend = &localBackend{}

func NewLocalBackend(path string) (backend.Backend, error) {
	if err := createFileIfNotExists(path); err != nil {
		return nil, err
	}
	return &localBackend{
		path: path,
	}, nil
}

func createFileIfNotExists(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return os.WriteFile(path, nil, 0644)

	}
	if fileInfo.IsDir() {
		return errors.New(fmt.Sprintf("path %s was a directory, expected a file", path))
	}
	return nil
}
