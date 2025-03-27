package multipass

import (
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
	"github.com/analog-substance/carbon/pkg/types"
)

type Environment struct {
	name    string
	profile types.Profile
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM
	for _, mpVM := range api.ListVMs() {
		var publicIPs []string
		var privateIPs []string

		publicIPs = append(publicIPs, mpVM.Ipv4...)

		vms = append(vms, &models.Machine{
			InstanceName:       mpVM.Name,
			CurrentState:       stateFromVboxInfo(mpVM.State),
			InstanceID:         mpVM.Name,
			Env:                e,
			PublicIPAddresses:  publicIPs,
			PrivateIPAddresses: privateIPs,
		})
	}
	return vms
}
func (e *Environment) StartVM(id string) error {
	return api.StartVM(id)
}

func (e *Environment) StopVM(id string) error {
	// sleep
	return api.StopVM(id)
}

func (e *Environment) SuspendVM(id string) error {
	// sleep
	return api.Suspend(id)
}

func (e *Environment) RestartVM(id string) error {
	// sleep
	return api.RestartVM(id)
}

func (e *Environment) DestroyVM(id string) error {
	return nil
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return nil
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Name())
}
func (e *Environment) Images() ([]types.Image, error) {
	return []types.Image{}, nil
}

func stateFromVboxInfo(state string) types.MachineState {
	if state == "Suspended" {
		return types.StateSleeping
	}
	if state == "Starting" {
		return types.StateStarting
	}
	if state == "Stopped" {
		return types.StateStopped
	}
	if state == "Running" {
		return types.StateRunning
	}
	if state == "Deleted" {
		return types.StateTerminated
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}

func (e *Environment) DestroyImage(imageID string) error {
	return base.DestroyImageForFileBasedProvider(imageID)
}
