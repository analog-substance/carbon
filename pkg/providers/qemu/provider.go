package qemu

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
)

type provider struct {
	types.Provider
	path string
}

const providerName = "QEMU"
const profileName = "default"
const environmentName = "local"

func New() types.Provider {
	return &provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *provider) IsAvailable() bool {

	return true
}

func (p *provider) Profiles() []types.Profile {
	profiles := []types.Profile{}

	if p.IsAvailable() {
		config, ok := p.Provider.GetConfig().Profiles[profileName]
		if !ok {
			config = common.ProfileConfig{
				Enabled: true,
				URL:     string(libvirt.QEMUSystem),
			}
		}
		if config.Enabled {
			profiles = append(profiles, NewProfile(profileName, p, config))
		}
	}
	return profiles

}
