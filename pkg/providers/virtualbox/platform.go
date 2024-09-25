package virtualbox

import (
	"github.com/analog-substance/carbon/pkg/types"
	"slices"
)

type platform struct {
	profileName string
	provider    *provider
}

const platformName = "local"

func (p platform) Environments(validNames ...string) []types.Environment {
	// we have filters, check if we are wanted
	if len(validNames) == 0 || slices.Contains(validNames, platformName) {
		return []types.Environment{environment{
			platformName,
			p,
		}}
	}
	return []types.Environment{}
}

func (p platform) Name() string {
	return p.profileName
}

func (p platform) Provider() types.Provider {
	return p.provider
}
