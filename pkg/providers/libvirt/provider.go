package libvirt

import (
	types2 "github.com/analog-substance/carbon/pkg/types"
	"slices"
	"strings"
)

type provider struct {
	path string
}

func New() types2.Provider {
	return &provider{}
}

func (p *provider) IsAvailable() bool {

	return true
}

func (p *provider) Platforms(validNames ...string) []types2.Platform {
	platforms := []types2.Platform{}
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
