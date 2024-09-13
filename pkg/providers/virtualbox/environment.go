package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/providers/types"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
)

type environment struct {
	name     string
	platform types.Platform
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Platform() types.Platform {
	return e.platform
}

func (e environment) VMs() []types.VM {
	var vms []types.VM
	for _, vboxVM := range api.ListVMs() {

		vms = append(vms, types.Machine{
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

func stateFromVboxInfo(state string) types.MachineState {
	if state == "poweroff" {
		return types.StateOff
	}
	if state == "poweron" {
		return types.StateRunning
	}
	return types.StateUnknown
}
