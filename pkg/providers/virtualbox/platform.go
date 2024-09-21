package virtualbox

import (
	types2 "github.com/analog-substance/carbon/pkg/types"
	"slices"
)

type platform struct {
	profileName string
	provider    *provider
}

const platformName = "local"

func (p platform) Environments(validNames ...string) []types2.Environment {
	// we have filters, check if we are wanted
	if len(validNames) == 0 || slices.Contains(validNames, platformName) {
		return []types2.Environment{environment{
			platformName,
			p,
		}}
	}
	return []types2.Environment{}
}

func (p platform) Name() string {
	return p.profileName
}

func (p platform) Provider() types2.Provider {
	return p.provider
}
