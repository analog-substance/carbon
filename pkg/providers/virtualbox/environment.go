package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
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
	for _, vboxVM := range api.ListVMs() {

		vms = append(vms, &models.Machine{
			InstanceName:       vboxVM.Name,
			CurrentState:       stateFromVboxInfo(vboxVM.State),
			InstanceID:         vboxVM.ID,
			Env:                e,
			PrivateIPAddresses: vboxVM.PrivateIPAddresses,
		})
	}
	return vms
}

func (e *Environment) StartVM(id string) error {
	return api.StartVM(id)
}

func (e *Environment) StopVM(id string) error {
	// sleep
	return api.SleepVM(id)
}

func (e *Environment) RestartVM(id string) error {
	// sleep
	return api.RestartVM(id)
}

func (e *Environment) DestroyVM(id string) error {
	return nil
}

func (e *Environment) DestroyImage(imageID string) error {
	return base.DestroyImageForFileBasedProvider(imageID)
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return nil
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e *Environment) Images() ([]types.Image, error) {
	return base.GetImagesForFileBasedProvider(e.Profile().Provider().Type(), e)
}

func stateFromVboxInfo(state string) types.MachineState {
	if state == "poweroff" {
		return types.StateStopped
	}
	if state == "poweron" {
		return types.StateRunning
	}
	if state == "aborted" {
		return types.StateStopped
	}
	if state == "running" {
		return types.StateRunning
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
