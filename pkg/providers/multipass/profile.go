package multipass

import (
	"github.com/analog-substance/carbon/pkg/types"
	"slices"
)

type profile struct {
	profileName string
	provider    *provider
}

const profileName = "local"

func (p profile) Environments(validNames ...string) []types.Environment {
	if len(validNames) == 0 || slices.Contains(validNames, profileName) {
		return []types.Environment{environment{
			profileName,
			p,
		}}
	}
	return []types.Environment{}
}

func (p profile) Name() string {
	return p.profileName
}

func (p profile) Provider() types.Provider {
	return p.provider
}
