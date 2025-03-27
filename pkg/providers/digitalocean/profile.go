package digitalocean

import (
	"context"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/godo"
	"strconv"
	"strings"
)

type Profile struct {
	types.Profile
	doClient         *godo.Client
	droplets         map[int]godo.Droplet
	projectResources map[string][]godo.ProjectResource
	projectDroplets  map[string][]*godo.Droplet
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {

	doClient := godo.NewFromToken(config.GetPassword())
	return &Profile{
		Profile:          base.NewProfile(name, providerInstance, config),
		doClient:         doClient,
		droplets:         make(map[int]godo.Droplet),
		projectDroplets:  make(map[string][]*godo.Droplet),
		projectResources: make(map[string][]godo.ProjectResource),
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

func (p *Profile) ToProfile() types.Profile {
	return p
}

func (p *Profile) GetProjectVMs(projectID string) []*godo.Droplet {
	if _, ok := p.projectDroplets[projectID]; !ok {
		p.projectDroplets[projectID] = []*godo.Droplet{}
		droplets := p.GetDroplets()

		for _, res := range p.GetProjectResources(projectID) {
			if strings.HasPrefix(res.URN, "do:droplet:") {
				dIDStr := strings.TrimPrefix(res.URN, "do:droplet:")
				dropletID, err := strconv.Atoi(dIDStr)
				if err != nil {
					log.Error("Error converting droplet ID to int", "error", err)
					continue
				}

				if val, ok := droplets[dropletID]; ok {
					p.projectDroplets[projectID] = append(p.projectDroplets[projectID], &val)
				}

			}
		}
	}
	return p.projectDroplets[projectID]
}

func (p *Profile) GetProjectResources(projectID string) []godo.ProjectResource {
	if _, ok := p.projectResources[projectID]; !ok {

		p.projectResources[projectID] = []godo.ProjectResource{}

		ctx := context.Background()
		opt := &godo.ListOptions{}

		for {
			resources, resp, err := p.doClient.Projects.ListResources(ctx, projectID, opt)
			if err != nil {
				log.Debug("Error listing Droplets", "error", err)
			}

			p.projectResources[projectID] = append(p.projectResources[projectID], resources...)

			if resp.Links == nil || resp.Links.IsLastPage() {
				break
			}

			page, err := resp.Links.CurrentPage()
			if err != nil {
				log.Debug("Error getting current page", "error", err)
				break
			}

			opt.Page = page + 1
		}
	}
	return p.projectResources[projectID]
}

func (p *Profile) GetDroplets() map[int]godo.Droplet {

	if len(p.droplets) == 0 {

		ctx := context.Background()
		opt := &godo.ListOptions{}
		for {
			droplets, resp, err := p.doClient.Droplets.List(ctx, opt)
			if err != nil {
				log.Debug("Error listing Droplets", "error", err)
				break
			}

			for _, d := range droplets {
				p.droplets[d.ID] = d
			}

			// if we are at the last page, break out the for loop
			if resp.Links == nil || resp.Links.IsLastPage() {
				break
			}

			page, err := resp.Links.CurrentPage()
			if err != nil {
				log.Debug("Error getting current page", "error", err)
				break
			}

			// set the page we want for the next request
			opt.Page = page + 1
		}
	}
	return p.droplets

}
