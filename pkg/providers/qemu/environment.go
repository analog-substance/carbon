package qemu

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/qemu/api"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Environment struct {
	name    string
	profile types.Profile
	qemu    *api.QEMU
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM
	domains, err := e.qemu.GetDomains()
	if err == nil {
		for _, domain := range domains {
			vms = append(vms, &models.Machine{
				InstanceName:       domain.Name,
				CurrentState:       stateFromVboxInfo(*domain.LVDomainState),
				InstanceID:         domain.ID,
				Env:                e,
				PublicIPAddresses:  domain.PublicIPAddresses,
				PrivateIPAddresses: domain.PrivateIPAddresses,
				CurrentUpTime:      domain.CurrentUpTime,
			})
		}
	}

	return vms
}

func (e *Environment) StartVM(id string) error {

	domain, err := e.qemu.GetDomain(id)
	if err != nil {
		return err
	}
	return domain.Start()
}

func (e *Environment) StopVM(id string) error {

	domain, err := e.qemu.GetDomain(id)
	if err != nil {
		return err
	}
	return domain.Suspend()
}

func (e *Environment) RestartVM(id string) error {

	domain, err := e.qemu.GetDomain(id)
	if err != nil {
		return err
	}
	return domain.Reboot()
}

func (e *Environment) DestroyVM(id string) error {
	domain, err := e.qemu.GetDomain(id)
	if err != nil {
		return err
	}
	return domain.Destroy()
}

func (e *Environment) DestroyImage(imageID string) error {
	return base.DestroyImageForFileBasedProvider(imageID)
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {

	defaultStoragePool, err := e.qemu.GetStoragePool("default")
	if err != nil {
		return err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	imagePath := filepath.Join(cwd, viper.GetString(common.ViperImagesDir), options.Image.ID(), options.Image.Name())
	newVol, err := defaultStoragePool.ImportImage(options.Name, imagePath)
	if err != nil {
		return err
	}

	_, err = e.qemu.CreateDomain(options.Name, newVol)
	if err != nil {
		return err
	}

	return nil
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e *Environment) Images() ([]types.Image, error) {
	return base.GetImagesForFileBasedProvider(e.Profile().Provider().Type(), e)
}

func stateFromVboxInfo(state libvirt.DomainState) types.MachineState {
	if state == libvirt.DomainPaused {
		return types.StateSleeping
	}
	if state == libvirt.DomainShutoff {
		return types.StateStopped
	}
	if state == libvirt.DomainRunning {
		return types.StateRunning
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
