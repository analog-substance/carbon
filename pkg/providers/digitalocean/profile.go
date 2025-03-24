package digitalocean

import (
	"context"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
)

type Profile struct {
	types.Profile
	doClient *godo.Client
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {

	log.Debug("password", "password", config.GetPassword())
	doClient := godo.NewFromToken(config.GetPassword())
	return &Profile{
		Profile:  base.NewProfile(name, providerInstance, config),
		doClient: doClient,
	}
}

func (p *Profile) Environments() []types.Environment {
	var environments []types.Environment

	ctx := context.Background()
	opt := &godo.ListOptions{}
	for {
		projects, resp, err := p.doClient.Projects.List(ctx, opt)
		if err != nil {
			log.Debug("Error listing projects", "error", err)
		}

		for _, project := range projects {
			shouldInclude := p.ShouldIncludeEnvironment(project.ID)
			log.Debug("validating Environment visibility", "Environment", project.Name, "shouldInclude", shouldInclude)
			if shouldInclude {
				environments = append(environments, &Environment{
					name:      project.Name,
					profile:   p,
					doClient:  p.doClient,
					doProject: &project,
				})
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

	return environments
}
