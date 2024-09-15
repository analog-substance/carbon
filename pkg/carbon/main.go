package carbon

import (
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/libvirt"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/types"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"strings"
)

type Options struct {
	Providers    []string
	Platforms    []string
	Environments []string
}

type Carbon struct {
	options      Options
	providers    []types.Provider
	platforms    []types.Platform
	environments []types.Environment
	machines     []types.VM
}

func New(options Options) *Carbon {

	carbon := &Carbon{options: options}

	if options.Providers == nil || len(options.Providers) == 0 {
		carbon.providers = AvailableProviders()
	} else {
		provs := []types.Provider{}
		for _, provider := range AvailableProviders() {
			for _, providerStr := range options.Providers {
				if strings.ToLower(providerStr) == strings.ToLower(provider.Name()) {
					provs = append(provs, provider)
				}
			}
		}
		carbon.providers = provs
	}

	return carbon
}

func (c *Carbon) Providers() []types.Provider {
	return c.providers
}

func (c *Carbon) Platforms() []types.Platform {
	if len(c.platforms) == 0 {
		c.platforms = []types.Platform{}
		for _, provider := range c.Providers() {
			c.platforms = append(c.platforms, provider.Platforms(c.options.Platforms...)...)
		}
	}

	return c.platforms
}

func (c *Carbon) GetVMs() []types.VM {
	if len(c.machines) == 0 {
		c.machines = []types.VM{}
		for _, platform := range c.Platforms() {
			for _, env := range platform.Environments(c.options.Environments...) {
				c.machines = append(c.machines, env.VMs()...)
			}
		}

	}
	return c.machines
}

func (c *Carbon) FindVMByID(id string) types.VM {
	for _, vm := range c.GetVMs() {
		if vm.ID() == id {
			return vm
		}
	}
	return nil
}

var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	if len(availableProviders) == 0 {
		allProviders := []types.Provider{
			aws.New(),
			libvirt.New(),
			virtualbox.New(),
			multipass.New(),
		}

		for _, provider := range allProviders {
			if provider.IsAvailable() {
				availableProviders = append(availableProviders, provider)
			}
		}
	}

	return availableProviders
}
