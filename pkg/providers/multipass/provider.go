package multipass

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
	"github.com/analog-substance/carbon/pkg/types"
	"os/exec"
	"path/filepath"
)

type Provider struct {
	types.Provider
	path string
}

const providerName = "Multipass"
const profileName = "default"
const environmentName = "local"

func New() types.Provider {
	return &Provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *Provider) appPath() string {
	if p.path == "" {
		multipassPath, err := exec.LookPath("multipass")
		if err == nil {
			p.path, err = filepath.Abs(multipassPath)
		}
	}
	return p.path
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
