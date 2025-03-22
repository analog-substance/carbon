package digitalocean

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
)

type Profile struct {
	types.Profile
	doClient  *godo.Client
	doProject *godo.Project
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig, doProject *godo.Project) *Profile {

	doClient := godo.NewFromToken(providerInstance.APIToken())
	return &Profile{
		base.NewProfile(name, providerInstance, config),
		doClient,
		doProject,
	}
}

func (p *Profile) Environments() []types.Environment {

	enabled, ok := p.Profile.GetConfig().Environments[p.doProject.Environment]
	if !ok || enabled {
		return []types.Environment{&Environment{
			p.doProject.Environment,
			p,
			p.doClient,
		}}
	}
	return []types.Environment{}
}
