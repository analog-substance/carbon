package qemu

import (
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
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

func (p *provider) Profiles(validNames ...string) []types.Profile {
	profiles := []types.Profile{}
	// we have filters, check if we are wanted
	if len(validNames) > 0 && !slices.Contains(validNames, strings.ToLower(p.Name())) {
		return profiles
	}

	if p.IsAvailable() {
		profiles = append(profiles, profile{p.Name(), p, string(libvirt.QEMUSystem)})
	}
	return profiles

}
func (p *provider) Name() string {
	return "QEMU"
}

func (p *provider) Type() string {
	return strings.ToLower(p.Name())
}
