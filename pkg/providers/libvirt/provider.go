package libvirt

import (
	"github.com/analog-substance/carbon/pkg/providers/types"
	"slices"
	"strings"
)

type provider struct {
	path string
}

func New() types.Provider {
	return &provider{}
}

func (p *provider) IsAvailable() bool {

	return true
}

func (p *provider) Platforms(validNames ...string) []types.Platform {
	platforms := []types.Platform{}
	// we have filters, check if we are wanted
	if len(validNames) > 0 && !slices.Contains(validNames, strings.ToLower(p.Name())) {
		return platforms
	}

	if p.IsAvailable() {
		platforms = append(platforms, platform{p.Name(), p, "qemu:///system"})
	}
	return platforms
}
func (p *provider) Name() string {
	return "libvirt"
}
