package qemu

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
	"net/url"
)

type profile struct {
	types.Profile
}

func NewProfile(name string, providerInstance *provider, config common.ProfileConfig) *profile {
	return &profile{
		base.NewProfile(name, providerInstance, config),
	}
}

func (p profile) Environments() []types.Environment {
	enabled, ok := p.Profile.GetConfig().Environments[environmentName]
	if !ok || enabled {
		uri, _ := url.Parse(p.GetConfig().URL)
		conn, err := libvirt.ConnectToURI(uri)
		if err == nil {
			return []types.Environment{environment{
				environmentName,
				p,
				conn,
			}}
		}
	}

	return []types.Environment{}
}
