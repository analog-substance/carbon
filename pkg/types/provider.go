package types

import "github.com/analog-substance/carbon/pkg/common"

type Provider interface {
	Profiles() []Profile
	Name() string
	Type() string
	IsAvailable() bool
	SetConfig(config common.ProviderConfig)
	GetConfig() common.ProviderConfig
}
