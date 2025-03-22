package carbon

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/digitalocean"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/qemu"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"github.com/analog-substance/carbon/pkg/types"
)

var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	if len(availableProviders) == 0 {
		allProviders := []types.Provider{
			aws.New(),
			qemu.New(),
			virtualbox.New(),
			multipass.New(),
			digitalocean.New(),
		}

		for _, provider := range allProviders {
			isAvailable := provider.IsAvailable()
			log.Debug("assessing provider availability", "provider", provider.Type(), "isAvailable", isAvailable)
			if isAvailable {
				availableProviders = append(availableProviders, provider)
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
