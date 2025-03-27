package gcloud

import (
	compute "cloud.google.com/go/compute/apiv1"
	"context"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"strings"
)

type Profile struct {
	types.Profile
	gCloudInstanceClient *compute.InstancesClient
	ctx                  context.Context
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {
	ctx := context.Background()
	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Debug("Error creating GCE client", "error", err)
	}

	return &Profile{
		Profile:              base.NewProfile(name, providerInstance, config),
		gCloudInstanceClient: c,
		ctx:                  ctx,
	}
}

func (p *Profile) Environments() []types.Environment {
	var environments []types.Environment

	for envName, enabled := range p.GetConfig().Environments {

		if enabled {
			ps := strings.Split(envName, "/")
			project := ps[0]
			zone := ps[1]

			environments = append(environments, &Environment{
				name:                 envName,
				gProject:             project,
				zone:                 zone,
				profile:              p,
				gCloudInstanceClient: p.gCloudInstanceClient,
				ctx:                  p.ctx,
			})
		}
	}

	return environments
}
