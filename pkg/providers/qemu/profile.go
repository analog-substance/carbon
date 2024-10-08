package qemu

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/qemu/api"
	"github.com/analog-substance/carbon/pkg/types"
)

type Profile struct {
	types.Profile
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {
	return &Profile{
		base.NewProfile(name, providerInstance, config),
	}
}

func (p *Profile) Environments() []types.Environment {
	enabled, ok := p.Profile.GetConfig().Environments[environmentName]
	if !ok || enabled {
		qemu, err := api.Connect(p.GetConfig().URL)
		if err == nil {
			return []types.Environment{&Environment{
				environmentName,
				p,
				qemu,
			}}
		} else {
			log.Debug("Failed to connect to QEMU", "err", err, "environmentName", environmentName)
		}
	}

	return []types.Environment{}
}
