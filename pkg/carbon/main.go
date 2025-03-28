// Package carbon provides core application functionality and constants
package carbon

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/types"
	"log/slog"
)

type Options struct {
	Providers    []string
	Profiles     []string
	Environments []string
}

type Carbon struct {
	version      string
	config       common.CarbonConfig
	providers    []types.Provider
	profiles     []types.Profile
	environments []types.Environment
	machines     []types.VM
	imageBuilds  []types.ImageBuild
	images       []types.Image
	projects     []types.Project
}

var log *slog.Logger

func init() {
	log = common.WithGroup("carbon")
	providerTypes := []string{}

	for _, provider := range AllProviders {
		providerTypes = append(providerTypes, provider.Type())
	}
	common.SetProvidersTypes(providerTypes)
}

func New(version string, config common.CarbonConfig) *Carbon {
	carbon := &Carbon{
		version:      version,
		config:       config,
		providers:    []types.Provider{},
		profiles:     []types.Profile{},
		environments: []types.Environment{},
	}
	for _, provider := range AvailableProviders() {
		providerConfig, ok := config.Providers[provider.Type()]
		if !ok {
			providerConfig = common.ProviderConfig{
				Enabled:      true,
				AutoDiscover: true,
			}
		}

		log.Debug("process provider", "enabled", providerConfig.Enabled, "provider", provider.Type(), "config_exists", ok)
		if providerConfig.Enabled {
			// no config, or explicitly enabled
			carbon.providers = append(carbon.providers, provider)
			provider.SetConfig(providerConfig)
		}
	}

	return carbon
}

func (c *Carbon) GetConfig() common.CarbonConfig {
	return c.config
}
