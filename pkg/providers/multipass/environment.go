package multipass

import (
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
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
	for _, mpVM := range api.ListVMs() {
		publicIPs := []string{}
		privateIPs := []string{}

		publicIPs = append(publicIPs, mpVM.Ipv4...)

		vms = append(vms, types2.Machine{
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

func stateFromVboxInfo(state string) types2.MachineState {
	if state == "Suspended" {
		return types2.StateSleeping
	}
	if state == "Stopped" {
		return types2.StateOff
	}
	if state == "Running" {
		return types2.StateRunning
	}
	return types2.StateUnknown
}
