package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
	"github.com/analog-substance/carbon/pkg/types"
)

const providerName = "VirtualBox"
const profileName = "default"
const environmentName = "local"

type Provider struct {
	types.Provider
	vboxmanageExecutablePath string
}

func New() types.Provider {
	return &Provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *Provider) IsAvailable() bool {
	return api.AppPath() != ""
}

func (p *Provider) Profiles() []types.Profile {
	var profiles []types.Profile
	if p.IsAvailable() {
		config, ok := p.Provider.GetConfig().Profiles[profileName]
		if !ok {
			config = common.ProfileConfig{
				Enabled: true,
			}
		}
		if config.Enabled {
			profiles = append(profiles, NewProfile(profileName, p, config))
		}
	}
	return profiles
}
