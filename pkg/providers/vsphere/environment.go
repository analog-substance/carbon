package vsphere

import (
	"context"
	"errors"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	vsphereTypes "github.com/vmware/govmomi/vim25/types"
	"time"
)

type Environment struct {
	name       string
	profile    types.Profile
	apiClient  *vim25.Client
	hostSystem *mo.HostSystem
	vsphereVMs []mo.VirtualMachine
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	if len(e.vsphereVMs) == 0 {
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
		err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary", "runtime", "network", "guest"}, &e.vsphereVMs)
		if err != nil {
			log.Debug("failed to retrieve VM summary", "error", err)
		}

	}

	for _, vm := range e.vsphereVMs {
		if vm.Runtime.Host.Value != e.hostSystem.Summary.Host.Value {
			// this is terrible. it seems event the vsphere sdk does something similar. not cool
			continue
		}

		var publicIPs = []string{}
		var privateIPs = []string{vm.Guest.IpAddress}

		uptime := time.Duration(0)
		if vm.Runtime.BootTime != nil {
			uptime = time.Now().Sub(*vm.Runtime.BootTime)
		}

		info := fmt.Sprintf("vx-%dvcpu-%dgb", vm.Summary.Config.NumCpu, vm.Summary.Config.MemorySizeMB/1024)
		vms = append(vms, &models.Machine{
			InstanceName:       vm.Summary.Config.Name,
			CurrentState:       stateFromStatus(vm.Runtime.PowerState),
			InstanceID:         vm.Summary.Vm.Value,
			Env:                e,
			CurrentUpTime:      uptime,
			InstanceType:       info,
			PrivateIPAddresses: privateIPs,
			PublicIPAddresses:  publicIPs,
		})
	}
	return vms
}

func (e *Environment) StartVM(id string) error {
	vsv := e.getVMByID(id)
	if vsv != nil {
		req := vsphereTypes.PowerOnVM_Task{
			This: vsv.Reference(),
		}

		_, err := methods.PowerOnVM_Task(context.Background(), e.apiClient.RoundTripper, &req)
		return err
	}

	return errors.New("no vm found")
}

func (e *Environment) StopVM(id string) error {
	vsv := e.getVMByID(id)
	if vsv != nil {
		req := vsphereTypes.PowerOffVM_Task{
			This: vsv.Reference(),
		}

		_, err := methods.PowerOffVM_Task(context.Background(), e.apiClient.RoundTripper, &req)
		return err
	}

	return errors.New("no vm found")
}

func (e *Environment) SuspendVM(id string) error {
	vsv := e.getVMByID(id)
	if vsv != nil {
		req := vsphereTypes.SuspendVM_Task{
			This: vsv.Reference(),
		}

		_, err := methods.SuspendVM_Task(context.Background(), e.apiClient.RoundTripper, &req)
		return err
	}

	return errors.New("no vm found")
}

func (e *Environment) RestartVM(id string) error {
	vsv := e.getVMByID(id)
	if vsv != nil {
		req := vsphereTypes.RebootGuest{
			This: vsv.Reference(),
		}

		_, err := methods.RebootGuest(context.Background(), e.apiClient.RoundTripper, &req)
		return err
	}

	return errors.New("no vm found")
}

func (e *Environment) DestroyVM(id string) error {
	vsv := e.getVMByID(id)
	if vsv != nil {
		req := vsphereTypes.Destroy_Task{
			This: vsv.Reference(),
		}

		_, err := methods.Destroy_Task(context.Background(), e.apiClient.RoundTripper, &req)
		return err
	}

	return errors.New("no vm found")
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

func (e *Environment) getVMByID(id string) *mo.VirtualMachine {
	for _, vsv := range e.vsphereVMs {
		if id == vsv.Summary.Vm.Value {
			return &vsv
		}
	}
	return nil
}
