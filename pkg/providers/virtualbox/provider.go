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

func (p *Provider) vboxPath() string {
	if p.vboxmanageExecutablePath == "" {
		virtualBox, err := exec.LookPath("vboxmanage")
		if err == nil {
			p.vboxmanageExecutablePath, err = filepath.Abs(virtualBox)
			if err != nil {
				log.Debug("err getting absolute path", "virtualBox", virtualBox, "err", err)
			}
		}
	}
	return p.vboxmanageExecutablePath
}

func (p *Provider) IsAvailable() bool {
	return p.vboxPath() != ""
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
