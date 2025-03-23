// Package aws handles communications with AWS APIs
package aws

// Other docs here are cool as well
import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
	"strings"
)

const providerName = "AWS"

type Provider struct {
	types.Provider
	awsProfileNames []string
	profiles        []types.Profile
}

// New creates new instance of an AWS Provider and returns it.
// Defaults to no awsProfileNames, this forces a query of the AWS config at runtime.
func New() types.Provider {
	return &Provider{
		Provider:        base.NewWithName(providerName),
		awsProfileNames: []string{},
	}
}

func (p *Provider) AWSProfiles() []string {
	if len(p.awsProfileNames) > 0 {
		return p.awsProfileNames
	}
	p.awsProfileNames = []string{}
	sections, err := ini.Load(config.DefaultSharedConfigFilename())
	if err != nil {
		return nil
	}

	for _, s := range sections.SectionStrings() {
		sl := strings.ToLower(s)

		log.Debug("checking config section for profile", "section_name", sl)
		if sl == "default" || strings.HasPrefix(sl, "profile") {
			sec, err := sections.GetSection(s)
			if err != nil {
				log.Debug("error getting config section:", "section", s)
				continue
			}

			if len(sec.Keys()) > 1 {
				name, _ := strings.CutPrefix(s, "profile ")
				p.awsProfileNames = append(p.awsProfileNames, name)
			}
		}
	}

	return p.awsProfileNames
}

func (p *Provider) IsAvailable() bool {
	return len(p.AWSProfiles()) > 0
}

func (p *Provider) Profiles() []types.Profile {
	if p.profiles == nil || len(p.profiles) == 0 {
		for _, profileName := range p.AWSProfiles() {
			profileConfig, ok := p.Provider.GetConfig().Profiles[profileName]
			enabled := p.Provider.GetConfig().AutoDiscover
			if !ok {
				profileConfig = common.ProfileConfig{
					Enabled: enabled,
				}
			}

			log.Debug("aws profile", "profile", profileName, "enabled", profileConfig.Enabled, "ok", ok, "auto_discover", p.Provider.GetConfig().AutoDiscover)
			if profileConfig.Enabled {
				p.profiles = append(p.profiles, NewProfile(profileName, p, profileConfig))
			}
		}

	}
	return p.profiles
}
