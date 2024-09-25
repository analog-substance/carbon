package simple

import (
	"github.com/analog-substance/carbon/pkg/types"
)

type Platform struct {
	profileName string
	provider    *Provider
}

func (p *Platform) Environments(validNames ...string) []types.Environment {
	return []types.Environment{
		&Environment{
			name:     "simple",
			platform: p,
		},
	}
}
func (p *Platform) Name() string {
	return p.profileName
}
func (p *Platform) Provider() types.Provider {
	return p.provider
}
