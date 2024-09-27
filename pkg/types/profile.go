package types

import "github.com/analog-substance/carbon/pkg/common"

type Profile interface {
	Environments() []Environment
	Name() string
	Provider() Provider
	SetConfig(config common.ProfileConfig)
	GetConfig() common.ProfileConfig
}
