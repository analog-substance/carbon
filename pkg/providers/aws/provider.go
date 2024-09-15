package aws

import (
	"github.com/analog-substance/carbon/pkg/providers/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
	"log"
	"slices"
	"strings"
)

type provider struct {
	profiles []string
}

func New() types.Provider {
	return &provider{}
}

func (p *provider) Profiles() []string {
	if len(p.profiles) > 0 {
		return p.profiles
	}
	p.profiles = []string{}
	sections, err := ini.Load(config.DefaultSharedConfigFilename())
	if err != nil {
		return nil
	}

	for _, s := range sections.SectionStrings() {
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

	return p.profiles
}

func (p *provider) IsAvailable() bool {
	return len(p.Profiles()) > 0
}

func (p *provider) Platforms(validNames ...string) []types.Platform {
	platforms := []types.Platform{}

	for _, s := range p.Profiles() {
		// we have filters, check if we are wanted
		if len(validNames) == 0 || slices.Contains(validNames, strings.ToLower(s)) {
			platforms = append(platforms, &platform{s, p})
		}
	}

	return platforms
}

func (p *provider) Name() string {
	return "AWS"
}
