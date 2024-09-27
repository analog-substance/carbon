package simple

import (
	"github.com/analog-substance/carbon/pkg/types"
	"strings"
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

func (p *Provider) Profiles(validNames ...string) []types.Profile {
	return []types.Profile{
		&Profile{
			profileName: "simple",
			provider:    p,
		},
	}
}

func (p *Provider) Name() string {
	return "simple"
}

func (p *Provider) Type() string {
	return strings.ToLower(p.Name())
}
