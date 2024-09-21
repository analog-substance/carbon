package virtualbox

import (
	types2 "github.com/analog-substance/carbon/pkg/types"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type provider struct {
	path string
}

func New() types2.Provider {
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

func (p *provider) Platforms(validNames ...string) []types2.Platform {
	platforms := []types2.Platform{}

	// we have filters, check if we are wanted
	if len(validNames) > 0 && !slices.Contains(validNames, strings.ToLower(p.Name())) {
		return platforms
	}

	if p.IsAvailable() {
		platforms = append(platforms, platform{p.Name(), p})
	}
	return platforms
}
func (p *provider) Name() string {
	return "VirtualBox"
}
