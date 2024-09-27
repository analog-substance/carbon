package aws

import (
	"context"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

//type profile struct {
//	profileName string
//	provider    *provider
//}

type profile struct {
	types.Profile
}

func NewProfile(name string, providerInstance *provider, config common.ProfileConfig) *profile {
	return &profile{
		base.NewProfile(name, providerInstance, config),
	}
}

func (p *profile) Environments() []types.Environment {

	var environments []types.Environment
	var options []func(*config.LoadOptions) error

	options = append(options, config.WithSharedConfigProfile(p.Name()))
	cfg, err := config.LoadDefaultConfig(context.Background(), options...)
	if err != nil {
		log.Println("Error loading  AWS configuration", err)
		return environments
	}
	ec2Service := ec2.NewFromConfig(cfg)
	vpcResults, err := ec2Service.DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	if err != nil {
		log.Println("Error get VPCs", err)
		return environments
	}

	for _, vpc := range vpcResults.Vpcs {
		environmentName := *vpc.VpcId
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				environmentName = *tag.Value
			}
		}

		enabled, ok := p.Profile.GetConfig().Environments[environmentName]
		if !ok || enabled {
			environments = append(environments, &environment{
				name:      environmentName,
				profile:   p,
				ec2Client: ec2Service,
				vpcId:     *vpc.VpcId,
			})
		}
	}
	return environments
}
