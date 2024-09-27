package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/types"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type provider struct {
	path string
}

func New() types.Provider {
	return &provider{}
}

func (p *provider) vboxPath() string {
	if p.path == "" {
		virtualBox, err := exec.LookPath("vboxmanage")
		if err == nil {
			p.path, err = filepath.Abs(virtualBox)
		}
	}
	return p.path
}

func (p *provider) IsAvailable() bool {
	return p.vboxPath() != ""
}

func (p *provider) Profiles(validNames ...string) []types.Profile {
	profiles := []types.Profile{}

	// we have filters, check if we are wanted
	if len(validNames) > 0 && !slices.Contains(validNames, strings.ToLower(p.Name())) {
		return profiles
	}

	if p.IsAvailable() {
		profiles = append(profiles, profile{p.Name(), p})
	}
	return profiles
}
func (p *provider) Name() string {
	return "VirtualBox"
}

func (p *provider) Type() string {
	return strings.ToLower(p.Name())
}
