package vsphere

import (
	"errors"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
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

	log.Debug("getting VMs", "env", e.Name())

	var publicIPs = []string{"1.1.1.1"}
	var privateIPs = []string{"10.0.0.1"}

	vms = append(vms, &models.Machine{
		InstanceName: "test name",
		CurrentState: stateFromStatus("d.Status"),
		InstanceID:   fmt.Sprintf("%d", 420),
		Env:          e,
		//CurrentUpTime:      d,
		InstanceType:       "large",
		PrivateIPAddresses: privateIPs,
		PublicIPAddresses:  publicIPs,
	})

	return vms
}

func (e *Environment) StartVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) StopVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) RestartVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) DestroyVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) DestroyImage(imageID string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return errors.New("not yet implemented")
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e *Environment) Images() ([]types.Image, error) {
	return base.GetImagesForFileBasedProvider(e.Profile().Provider().Type(), e)
}

func stateFromStatus(state string) types.MachineState {
	//if state == "poweroff" {
	//	return types.StateStopped
	//}
	//if state == "poweron" {
	//	return types.StateRunning
	//}
	//if state == "aborted" {
	//	return types.StateStopped
	//}
	//if state == "active" {
	//	return types.StateRunning
	//}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
