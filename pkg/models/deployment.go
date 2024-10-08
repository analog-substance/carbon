package models

import (
	"path/filepath"
)

type Deployment struct {
	buildPath    string
	providerType string
	provisioner  string
}

func NewDeployment(buildPath string) *Deployment {
	return &Deployment{
		buildPath: buildPath,
	}
}

func (d *Deployment) Name() string {
	return filepath.Base(d.buildPath)
}
