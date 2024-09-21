package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
	types2 "github.com/analog-substance/carbon/pkg/types"
)

type environment struct {
	name     string
	platform types2.Platform
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Platform() types2.Platform {
	return e.platform
}

func (e environment) VMs() []types2.VM {
	var vms []types2.VM
	for _, vboxVM := range api.ListVMs() {

		vms = append(vms, types2.Machine{
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

func stateFromVboxInfo(state string) types2.MachineState {
	if state == "poweroff" {
		return types2.StateOff
	}
	if state == "poweron" {
		return types2.StateRunning
	}
	return types2.StateUnknown
}
