package simple

import (
	"github.com/analog-substance/carbon/pkg/types"
)

type Profile struct {
	profileName string
	provider    *Provider
}

func (p *Profile) Environments(validNames ...string) []types.Environment {
	return []types.Environment{
		&Environment{
			name:    "simple",
			profile: p,
		},
	}
}
func (p *Profile) Name() string {
	return p.profileName
}
func (p *Profile) Provider() types.Provider {
	return p.provider
}
