package simple

import (
	"errors"
	"github.com/analog-substance/carbon/pkg/types"
)

type Environment struct {
	name     string
	platform *Platform
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Platform() types.Platform {
	return e.platform
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

func (e *Environment) RestartVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) DestroyVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return errors.New("not implemented")
}

func (e *Environment) Images() []types.Image {
	return []types.Image{}
}
