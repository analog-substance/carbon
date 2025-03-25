package vsphere

import (
	"context"
	"errors"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	vsphereTypes "github.com/vmware/govmomi/vim25/types"
	"time"
)

type Environment struct {
	name      string
	profile   types.Profile
	apiClient *vim25.Client
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	// Create view of VirtualMachine objects
	m := view.NewManager(e.apiClient)
	ctx := context.Background()

	v, err := m.CreateContainerView(ctx, e.apiClient.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Debug("failed to create VM view", "error", err)
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vsvms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary", "runtime", "network", "guest"}, &vsvms)
	if err != nil {
		log.Debug("failed to retrieve VM summary", "error", err)
	}

	// Print summary per vm (see also: govc/vm/info.go)

	for _, vm := range vsvms {
		//fmt.Printf("%s: %s\n", , vm.Summary.Config.GuestFullName)

		var publicIPs = []string{}
		var privateIPs = []string{vm.Guest.IpAddress}

		uptime := time.Duration(0)
		if vm.Runtime.BootTime != nil {
			uptime = time.Now().Sub(*vm.Runtime.BootTime)
		}

		vms = append(vms, &models.Machine{
			InstanceName:       vm.Summary.Config.Name,
			CurrentState:       stateFromStatus(vm.Runtime.PowerState),
			InstanceID:         vm.Summary.Vm.Value,
			Env:                e,
			CurrentUpTime:      uptime,
			InstanceType:       "large",
			PrivateIPAddresses: privateIPs,
			PublicIPAddresses:  publicIPs,
		})
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

func stateFromStatus(state vsphereTypes.VirtualMachinePowerState) types.MachineState {
	if state == vsphereTypes.VirtualMachinePowerStatePoweredOff {
		return types.StateStopped
	}
	if state == vsphereTypes.VirtualMachinePowerStatePoweredOn {
		return types.StateRunning
	}
	if state == vsphereTypes.VirtualMachinePowerStateSuspended {
		return types.StateStopped
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
