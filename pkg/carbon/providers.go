package carbon

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/digitalocean"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/qemu"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"github.com/analog-substance/carbon/pkg/providers/vsphere"
	"github.com/analog-substance/carbon/pkg/types"
)

var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	defer (common.Time("available providers"))()

	if len(availableProviders) == 0 {

		type providerAvailability struct {
			provider  types.Provider
			available bool
		}
		c := make(chan providerAvailability)
		for _, provider := range AllProviders {
			go func() {
				//defer (common.Time(fmt.Sprintf("(%s)provider.isAvailable", provider.Name())))()
				c <- providerAvailability{
					provider:  provider,
					available: provider.IsAvailable(),
				}
			}()
		}

		result := make([]providerAvailability, len(AllProviders))
		for i, _ := range result {
			result[i] = <-c
			log.Debug("assessing provider availability", "provider", result[i].provider.Type(), "isAvailable", result[i].available)
			if result[i].available {
				availableProviders = append(availableProviders, result[i].provider)
			}
		}
	}

	return availableProviders
}

func (c *Carbon) Providers() []types.Provider {
	return c.providers
}

func (c *Carbon) GetProvider(providerType string) (types.Provider, error) {
	for _, provider := range c.Providers() {
		if provider.Type() == providerType {
			return provider, nil
		}
	}

	return nil, fmt.Errorf("provider '%s' not found", providerType)
}

var AllProviders = []types.Provider{
	aws.New(),
	qemu.New(),
	virtualbox.New(),
	multipass.New(),
	digitalocean.New(),
	vsphere.New(),
}
