package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
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
		return []types.Environment{&Environment{
			environmentName,
			p,
		}}
	}
	return []types.Environment{}
}
