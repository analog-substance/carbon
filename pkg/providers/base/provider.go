package base

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/types"
	"strings"
)

const providerName = "Base"

type Provider struct {
	name     string
	profiles []string
	config   common.ProviderConfig
}

func New() types.Provider {
	return &Provider{
		name: providerName,
	}
}

func NewWithName(name string) types.Provider {
	return &Provider{
		name: name,
	}
}

func (p *Provider) IsAvailable() bool {
	return false
}

func (p *Provider) Name() string {
	return p.name
}

func (p *Provider) Type() string {
	return strings.ToLower(p.Name())
}

func (p *Provider) Profiles() []types.Profile {
	return []types.Profile{
		&Profile{
			profileName: fmt.Sprintf("%s Profile", p.name),
			provider:    p,
		},
	}
}

func (p *Provider) SetConfig(config common.ProviderConfig) {
	p.config = config
}

func (p *Provider) GetConfig() common.ProviderConfig {
	return p.config
}
