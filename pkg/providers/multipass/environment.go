package multipass

import (
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
	"github.com/analog-substance/carbon/pkg/types"
	"log"
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
		publicIPs := []string{}
		privateIPs := []string{}

		publicIPs = append(publicIPs, mpVM.Ipv4...)

		vms = append(vms, types.Machine{
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

	log.Println("Unknown state for multipass VM:", state)
	return types.StateUnknown
}
