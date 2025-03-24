package digitalocean

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
	"os"
)

const apiKeyEnvVar = "DIGITALOCEAN_TOKEN"
const defaultProfileName = "default"
const providerName = "DigitalOcean"

type Provider struct {
	types.Provider
	profilesLoaded bool
	profiles       []types.Profile
}

func New() types.Provider {
	return &Provider{
		Provider: base.NewWithName(providerName),
	}
}

func (p *Provider) IsAvailable() bool {
	return true
}

func (p *Provider) Profiles() []types.Profile {
	if p.profilesLoaded {
		return p.profiles
	}
	p.profilesLoaded = true

	if len(p.Provider.GetConfig().Profiles) == 0 {
		apiKey := os.Getenv(apiKeyEnvVar)
		if apiKey != "" {
			defaultConfig := common.ProfileConfig{
				Enabled:  true,
				Password: apiKey,
			}
			p.profiles = append(p.profiles, NewProfile(defaultProfileName, p, defaultConfig))
		} else {

			type doYamlConfig struct {
				AccessToken string `yaml:"access-token"`
			}

			// check config file
			configFiles := []string{}

			path, err := homedir.Expand("~/.config/doctl/config.yaml")
			if err == nil {
				configFiles = append(configFiles, path)
			}

			for _, file := range configFiles {
				if _, err := os.Stat(file); err == nil {
					cfBytes, err := os.ReadFile(file)
					if err == nil {
						var cfs doYamlConfig
						err = yaml.Unmarshal(cfBytes, &cfs)
						if err == nil {
							if cfs.AccessToken != "" {

								defaultConfig := common.ProfileConfig{
									Enabled:  true,
									Password: cfs.AccessToken,
								}
								p.profiles = append(p.profiles, NewProfile(defaultProfileName, p, defaultConfig))
								break
							}
						}
					}
				}
			}
		}
	}
	log.Debug("provider config", "config", p.Provider.GetConfig())

	for profileName, profileConfig := range p.Provider.GetConfig().Profiles {
		log.Debug("profile config", "profile", profileName, "config", profileConfig)
		if profileConfig.Enabled {
			p.profiles = append(p.profiles, NewProfile(defaultProfileName, p, profileConfig))
		}
	}

	return p.profiles
}
