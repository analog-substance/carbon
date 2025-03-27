package digitalocean

import (
	"context"
	"errors"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
	"strconv"
)

type Environment struct {
	name      string
	profile   *Profile
	doClient  *godo.Client
	doProject *godo.Project
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile.ToProfile()
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	for _, d := range e.profile.GetProjectVMs(e.doProject.ID) {
		var publicIPs []string
		var privateIPs []string
		pIPs, err := d.PrivateIPv4()
		if err != nil {
			log.Debug("Error getting private IPv4", "error", err)
		} else {
			privateIPs = append(privateIPs, pIPs)
		}
		pIPs, err = d.PublicIPv4()
		if err != nil {
			log.Debug("Error getting private IPv4", "error", err)
		} else {
			publicIPs = append(publicIPs, pIPs)
		}

		vms = append(vms, &models.Machine{
			InstanceName: d.Name,
			CurrentState: stateFromStatus(d.Status),
			InstanceID:   fmt.Sprintf("%d", d.ID),
			Env:          e,
			//CurrentUpTime:      d,
			InstanceType:       d.Size.Slug,
			PrivateIPAddresses: privateIPs,
			PublicIPAddresses:  publicIPs,
		})
	}
	return vms
}

func (e *Environment) StartVM(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, _, err = e.doClient.DropletActions.PowerOn(context.Background(), intId)
	return err
}

func (e *Environment) StopVM(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, _, err = e.doClient.DropletActions.Shutdown(context.Background(), intId)
	return err
}

func (e *Environment) SuspendVM(id string) error {
	return errors.New("not implemented")
}

func (e *Environment) RestartVM(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, _, err = e.doClient.DropletActions.Reboot(context.Background(), intId)
	return err
}

func (e *Environment) DestroyVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) DestroyImage(imageID string) error {
	return errors.New("not yet implemented")
	return base.DestroyImageForFileBasedProvider(imageID)
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return errors.New("not yet implemented")
	return nil
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
	if state == "active" {
		return types.StateRunning
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
