package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/image_build"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
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
	for _, vboxVM := range api.ListVMs() {

		vms = append(vms, models.Machine{
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

func (e environment) ImageBuilds() []types.ImageBuild {
	imageBuilds, err := image_build.GetImageBuildsForProvider(e.platform.Provider().Name())
	if err != nil {
		log.Printf("Error getting image builds for %s: %s", e.Name(), err)
	}
	imageBuildStructs := []types.ImageBuild{}
	for _, imageBuild := range imageBuilds {
		imageBuildStructs = append(imageBuildStructs, &models.ImageBuild{
			Path: imageBuild,
		})
	}
	return imageBuildStructs
}

func (e environment) Images() []types.Image {
	return []types.Image{}
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

	log.Println("Unknown state for VirtualBox VM:", state)
	return types.StateUnknown
}
