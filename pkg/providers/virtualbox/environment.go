package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
	"github.com/analog-substance/carbon/pkg/types"
	"os"
	"path"
	"time"
)

type environment struct {
	name    string
	profile types.Profile
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Profile() types.Profile {
	return e.profile
}

func (e environment) VMs() []types.VM {
	var vms []types.VM
	for _, vboxVM := range api.ListVMs() {

		vms = append(vms, &models.Machine{
			InstanceName: vboxVM.Name,
			CurrentState: stateFromVboxInfo(vboxVM.State),
			InstanceID:   vboxVM.ID,
			Env:          e,
		})
	}
	return vms
}

func (e environment) StartVM(id string) error {
	return api.StartVM(id)
}

func (e environment) StopVM(id string) error {
	// sleep
	return api.SleepVM(id)
}

func (e environment) RestartVM(id string) error {
	// sleep
	return api.RestartVM(id)
}

func (e environment) DestroyVM(id string) error {
	return nil
}

func (e environment) CreateVM(options types.MachineLaunchOptions) error {
	return nil
}

func (e environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e environment) Images() ([]types.Image, error) {
	ret := []types.Image{}
	listing, _ := os.ReadDir("deployments/images/virtualbox")
	for _, dirEntry := range listing {
		ret = append(ret, models.NewImage(path.Join("deployments/images/virtualbox", dirEntry.Name()), dirEntry.Name(), time.Now(), e))

	}
	return ret, nil
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

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
