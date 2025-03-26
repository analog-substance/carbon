package base

import (
	"errors"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
)

type Environment struct {
	name    string
	profile *Profile
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	return vms
}

func (e *Environment) StartVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) StopVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) SuspendVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) RestartVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) DestroyVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return errors.New("not implemented")
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e *Environment) Images() ([]types.Image, error) {
	return []types.Image{}, nil
}

func (e *Environment) DestroyImage(imageID string) error {
	return nil
}
