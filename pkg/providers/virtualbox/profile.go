package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
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
		return []types.Environment{environment{
			environmentName,
			p,
		}}
	}
	return []types.Environment{}
}
