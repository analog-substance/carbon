package digitalocean

import (
	"context"
	"errors"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
)

type Environment struct {
	name     string
	profile  types.Profile
	doClient *godo.Client
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	ctx := context.Background()
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := e.doClient.Droplets.List(ctx, opt)
		if err != nil {
			log.Debug("Error listing Droplets", "error", err)
			break
		}

		for _, d := range droplets {
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

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			log.Debug("Error getting current page", "error", err)
			break
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}
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
