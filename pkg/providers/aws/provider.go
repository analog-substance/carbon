package aws

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

const providerName = "AWS"

type provider struct {
	types.Provider
	profiles []string
}

func New() types.Provider {
	return &provider{
		base.NewWithName(providerName),
		[]string{},
	}
}

func (p *provider) AWSProfiles() []string {
	if len(p.profiles) > 0 {
		return p.profiles
	}
	p.profiles = []string{}
	sections, err := ini.Load(config.DefaultSharedConfigFilename())
	if err != nil {
		return nil
	}

	for _, s := range sections.SectionStrings() {
		sl := strings.ToLower(s)
		if sl == "default" || strings.HasPrefix(sl, "profile") {
			sec, err := sections.GetSection(s)
			if err != nil {
				log.Println("error getting config section:", s)
				continue
			}

			if len(sec.Keys()) > 1 {
				name, _ := strings.CutPrefix(s, "profile ")
				p.profiles = append(p.profiles, name)
			}
		}
	}

	return p.profiles
}

func (p *provider) IsAvailable() bool {
	return len(p.AWSProfiles()) > 0
}

func (p *provider) Profiles() []types.Profile {
	profiles := []types.Profile{}
	for _, profileName := range p.AWSProfiles() {
		config, ok := p.Provider.GetConfig().Profiles[profileName]
		if !ok {
			config = common.ProfileConfig{
				Enabled: true,
			}
		}
		if config.Enabled {
			profiles = append(profiles, NewProfile(profileName, p, config))
		}
	}

	return profiles
}

//func (p *provider) Name() string {
//	return "AWS"
//}
