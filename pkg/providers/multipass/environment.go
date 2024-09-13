package multipass

import (
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
	"github.com/analog-substance/carbon/pkg/providers/types"
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
	for _, mpVM := range api.ListVMs() {

		vms = append(vms, types.Machine{
			InstanceName:       mpVM.Name,
			CurrentState:       stateFromVboxInfo(mpVM.State),
			InstanceID:         mpVM.Name,
			Env:                e,
			PublicIPAddresses:  mpVM.Ipv4[0:1],
			PrivateIPAddresses: mpVM.Ipv4[1:2],
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
	return api.SleepVM(id)
}

func stateFromVboxInfo(state string) types.MachineState {
	if state == "Suspended" {
		return types.StateSleeping
	}
	if state == "Stopped" {
		return types.StateOff
	}
	if state == "Running" {
		return types.StateRunning
	}
	return types.StateUnknown
}
