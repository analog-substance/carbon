package digitalocean

import (
	"context"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
	"os"
)

const providerName = "DigitalOcean"

type Provider struct {
	types.Provider
	apiToken string

	profiles []string
}

func New() types.Provider {
	return &Provider{
		Provider: base.NewWithName(providerName),
		profiles: []string{},
	}
}

func (p *Provider) APIToken() string {
	if p.apiToken == "" {

		type doYamlConfig struct {
			AccessToken string `yaml:"access-token"`
		}

		doAPIKey := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
		if doAPIKey != "" {
			p.apiToken = doAPIKey
		} else {
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
							p.apiToken = cfs.AccessToken
							break
						}
					}
				}
			}
		}
	}

	return p.apiToken
}

func (p *Provider) IsAvailable() bool {
	return p.APIToken() != ""
}

func (p *Provider) Profiles() []types.Profile {
	var profiles []types.Profile
	if p.IsAvailable() {
		doClient := godo.NewFromToken(p.APIToken())

		ctx := context.Background()
		opt := &godo.ListOptions{}
		for {
			projects, resp, err := doClient.Projects.List(ctx, opt)
			if err != nil {
				log.Debug("Error listing projects", "error", err)
			}

			for _, project := range projects {

				config, ok := p.Provider.GetConfig().Profiles[project.Name]

				if !ok {
					config = common.ProfileConfig{
						Enabled: true,
					}
				}
				if config.Enabled {
					profiles = append(profiles, NewProfile(project.Name, p, config, &project))
				}
			}

			// if we are at the last page, break out the for loop
			if resp.Links == nil || resp.Links.IsLastPage() {
				break
			}

			page, err := resp.Links.CurrentPage()
			if err != nil {
				log.Error("Error getting page from Godo API", "error", err)
				break
			}

			// set the page we want for the next request
			opt.Page = page + 1
		}

	}
	return profiles
}
