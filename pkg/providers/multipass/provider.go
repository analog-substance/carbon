package multipass

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/multipass/api"
	"github.com/analog-substance/carbon/pkg/types"
	"os/exec"
	"path/filepath"
)

type provider struct {
	types.Provider
	path string
}

const providerName = "Multipass"
const profileName = "default"
const environmentName = "local"

func New() types.Provider {
	return &provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *provider) appPath() string {
	if p.path == "" {
		virtualBox, err := exec.LookPath("multipass")
		if err == nil {
			p.path, err = filepath.Abs(virtualBox)
		}
	}
	return p.path
}

func (p *provider) IsAvailable() bool {
	return api.AppPath() != ""
}

func (p *provider) Profiles() []types.Profile {
	profiles := []types.Profile{}
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
