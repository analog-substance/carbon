package example

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
)

type Profile struct {
	types.Profile
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {

	log.Debug("password", "password", config.GetPassword())
	return &Profile{
		Profile: base.NewProfile(name, providerInstance, config),
	}
}

func (p *Profile) Environments() []types.Environment {
	var environments []types.Environment

	environments = append(environments, &Environment{
		name:    "example",
		profile: p,
	})

	return environments
}
