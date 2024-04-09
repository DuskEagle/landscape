package landscape

import "github.com/DuskEagle/landscape/pkg/backend"

type Project interface{}

type project struct{}

var _ Project = &project{}

func NewProject(projectName string, backend backend.Backend) (Project, error) {
	return &project{}, nil
}
