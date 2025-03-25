package vsphere

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
)

const defaultProfileName = "default"
const providerName = "vSphere"

type Provider struct {
	types.Provider
	profilesLoaded bool
	profiles       []types.Profile
}

func New() types.Provider {
	return &Provider{
		Provider: base.NewWithName(providerName),
	}
}

func (p *Provider) IsAvailable() bool {
	return true
}

func (p *Provider) Profiles() []types.Profile {
	if p.profilesLoaded {
		return p.profiles
	}
	p.profilesLoaded = true

	if len(p.Provider.GetConfig().Profiles) == 0 {
		defaultConfig := common.ProfileConfig{
			Enabled: true,
		}
		p.profiles = append(p.profiles, NewProfile(defaultProfileName, p, defaultConfig))
	}

	for profileName, profileConfig := range p.Provider.GetConfig().Profiles {
		if profileConfig.Enabled {
			p.profiles = append(p.profiles, NewProfile(profileName, p, profileConfig))
		}
	}

	return p.profiles
}
