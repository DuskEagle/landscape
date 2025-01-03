package aws

import (
	landscape "github.com/DuskEagle/landscape/pkg"
)

type AWSProvider struct {
	project landscape.Project
}

func NewProvider(project landscape.Project) (*AWSProvider, error) {
	return &AWSProvider{
		project: project,
	}, nil
}
