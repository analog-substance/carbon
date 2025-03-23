package qemu

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
)

type Provider struct {
	types.Provider
	path string
}

const providerName = "QEMU"
const profileName = "default"
const environmentName = "local"

func New() types.Provider {
	return &Provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *Provider) IsAvailable() bool {
	return true
}

func (p *Provider) Profiles() []types.Profile {
	var profiles []types.Profile

	if p.IsAvailable() {
		config, ok := p.Provider.GetConfig().Profiles[profileName]
		if !ok {
			log.Debug("no config found", "environmentName", profileName, "config", config)

			config = common.ProfileConfig{
				Enabled: true,
				URL:     string(libvirt.QEMUSystem),
			}
		} else {
			if config.URL == "" {
				config.URL = string(libvirt.QEMUSystem)
			}
		}
		if config.Enabled {
			log.Debug("adding default profile", "environmentName", environmentName, "config", config)
			profiles = append(profiles, NewProfile(environmentName, p, config))
		}
	}
	return profiles

}
