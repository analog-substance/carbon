package simple

import (
	"github.com/analog-substance/carbon/pkg/types"
)

type Provider struct {
	profiles []string
}

func New() types.Provider {
	return &Provider{}
}

func (p *Provider) IsAvailable() bool {
	return false
}

func (p *Provider) Platforms(validNames ...string) []types.Platform {
	return []types.Platform{
		&Platform{
			profileName: "simple",
			provider:    p,
		},
	}
}

func (p *Provider) Name() string {
	return "simple"
}
