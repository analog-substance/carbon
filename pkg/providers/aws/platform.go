package aws

import (
	"context"
	types2 "github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
	"slices"
	"strings"
)

type platform struct {
	profileName string
	provider    *provider
}

func (p *platform) Environments(validNames ...string) []types2.Environment {

	var environments []types2.Environment
	var options []func(*config.LoadOptions) error

	options = append(options, config.WithSharedConfigProfile(p.profileName))
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
		name := *vpc.VpcId
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				name = *tag.Value
			}
		}

		if len(validNames) == 0 || slices.Contains(validNames, strings.ToLower(name)) {
			environments = append(environments, &environment{
				name:      name,
				platform:  p,
				ec2Client: ec2Service,
				vpcId:     *vpc.VpcId,
			})
		}
	}
	return environments
}
func (p *platform) Name() string {
	return p.profileName
}
func (p *platform) Provider() types2.Provider {
	return p.provider
}
