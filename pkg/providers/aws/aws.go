package aws

import (
	landscape "github.com/DuskEagle/landscape/pkg"
	"github.com/DuskEagle/landscape/pkg/providers"
)

type AWSProvider struct {
	project landscape.Project
}

var _ providers.Provider = &AWSProvider{}

func NewProvider(project landscape.Project) (*AWSProvider, error) {
	return &AWSProvider{
		project: project,
	}, nil
}
