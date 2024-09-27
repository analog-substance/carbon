package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"os/exec"
	"path/filepath"
)

const providerName = "VirtualBox"
const profileName = "default"
const environmentName = "local"

type provider struct {
	types.Provider
	vboxmanageExecutablePath string
}

func New() types.Provider {
	return &provider{
		base.NewWithName(providerName),
		"",
	}
}

func (p *provider) vboxPath() string {
	if p.vboxmanageExecutablePath == "" {
		virtualBox, err := exec.LookPath("vboxmanage")
		if err == nil {
			p.vboxmanageExecutablePath, err = filepath.Abs(virtualBox)
		}
	}
	return p.vboxmanageExecutablePath
}

func (p *provider) IsAvailable() bool {
	return p.vboxPath() != ""
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
