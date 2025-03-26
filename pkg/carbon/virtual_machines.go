package carbon

import (
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"strings"
	"sync"
	"time"
)

func (c *Carbon) GetVMs() []types.VM {
	if len(c.machines) == 0 {
		started := time.Now()
		defer func() {
			log.Debug("GetVMs timing", "took", time.Since(started))
		}()
		c.machines = []types.VM{}
		mu := sync.Mutex{}
		wait := sync.WaitGroup{}
		for _, profile := range c.Profiles() {
			wait.Add(1)
			go func() {
				for _, env := range profile.Environments() {
					wait.Add(1)
					go func() {
						vmStart := time.Now()
						defer func() {
							log.Debug("env.VMs() timing", "took", time.Since(vmStart), "env", env.Name(), "profile", profile.Name(), "provider", profile.Provider().Name())
						}()
						machines := env.VMs()
						mu.Lock()
						c.machines = append(c.machines, machines...)
						mu.Unlock()
						wait.Done()
					}()
				}
				wait.Done()
			}()
		}
		wait.Wait()
	}
	return c.machines
}

func (c *Carbon) FindVMByID(id string) []types.VM {
	for _, vm := range c.GetVMs() {
		if vm.ID() == id {
			return []types.VM{vm}
		}
	}
	return []types.VM{}
}

func (c *Carbon) FindVMByName(name string) []types.VM {

	vms := []types.VM{}

	for _, vm := range c.GetVMs() {
		lowerName := strings.ToLower(vm.Name())
		name = strings.ToLower(name)
		if strings.Contains(lowerName, name) {
			vms = append(vms, vm)
		}
	}
	return vms
}

func (c *Carbon) VMsFromHosts(hostnames []string) []types.VM {

	simpleProvider := base.New()
	profile := simpleProvider.Profiles()
	envs := profile[0].Environments()

	vms := []types.VM{}
	for _, hostname := range hostnames {
		vms = append(vms, &models.Machine{
			InstanceName:       hostname,
			CurrentState:       types.StateUnknown,
			InstanceID:         hostname,
			Env:                envs[0],
			PublicIPAddresses:  []string{hostname},
			PrivateIPAddresses: []string{hostname},
		})
	}
	return vms
}
