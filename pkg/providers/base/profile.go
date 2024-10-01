package base

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/types"
	"strings"
)

type Profile struct {
	profileName string
	provider    types.Provider
	config      common.ProfileConfig
}

func NewProfile(name string, providerInstance types.Provider, config common.ProfileConfig) types.Profile {
	return &Profile{
		profileName: name,
		provider:    providerInstance,
		config:      config,
	}
}

func (p *Profile) Environments() []types.Environment {
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

func (p *Profile) SetConfig(config common.ProfileConfig) {
	p.config = config
}

func (p *Profile) GetConfig() common.ProfileConfig {
	return p.config
}

func (p *Profile) ShouldIncludeEnvironment(envName string) bool {
	envName = strings.ToLower(envName)

	enabled, ok := p.GetConfig().Environments[envName]
	return !ok || enabled
}
