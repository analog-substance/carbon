package common

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var instanceCfg *CarbonConfigFile

type ProfileConfig struct {
	Enabled         bool            `yaml:"enabled"`
	Username        string          `yaml:"username" `
	Password        string          `yaml:"password" `
	PasswordCommand string          `yaml:"password_command"`
	Use1PassCLI     bool            `yaml:"use_1pass_cli" `
	URL             string          `yaml:"url"`
	Environments    map[string]bool `yaml:"environments"`
}

func (pc *ProfileConfig) Get(search []string) any {

	if len(search) == 0 {
		return pc
	}
	first := search[0]
	search = search[1:]

	switch first {
	case "enabled":
		return pc.Enabled
	case "ussername":
		return pc.Username
	case "password":
		return pc.Password
	case "password_command":
		return pc.PasswordCommand
	case "use_1pass_cli":
		return pc.Use1PassCLI
	case "URL":
		return pc.URL
	case "Environments":
		return pc.Environments
	}
	return ""
}

func (pc *ProfileConfig) Set(search []string, val any) *ProfileConfig {

	first := search[0]
	search = search[1:]

	if len(search) == 0 {

		switch first {
		case "enabled":
			pc.Enabled = val.(bool)
		case "username":
			pc.Username = val.(string)
			log.Debug("asd", "pc", pc)
		case "password":
			pc.Password = val.(string)
		case "password_command":
			pc.PasswordCommand = val.(string)
		case "use_1pass_cli":
			pc.Use1PassCLI = val.(bool)
		case "URL":
			pc.URL = val.(string)
		}
	}
	//fmt.Println(pc)
	//fmt.Println(pc)
	return nil
}

func (pc *ProfileConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type defaultProfile ProfileConfig
	raw := defaultProfile{Enabled: true} // Set default values
	if err := unmarshal(&raw); err != nil {
		return err
	}
	*pc = ProfileConfig(raw)
	return nil
}

func (pc *ProfileConfig) GetPassword() string {
	if pc.Use1PassCLI {
		pc.PasswordCommand = fmt.Sprintf("op read \"%s\"", pc.Password)
	}

	if pc.PasswordCommand != "" {
		log.Debug("password command", "command", pc.PasswordCommand)
		p, err := exec.Command("sh", "-c", pc.PasswordCommand).Output()
		if err != nil {
			log.Debug("failed to execute password command", "cmd", pc.PasswordCommand, "err", err)
		}
		return strings.TrimSpace(string(p))
	}

	return pc.Password
}

type ProviderConfig struct {
	Enabled      bool                     `yaml:"enabled"`
	AutoDiscover bool                     `yaml:"auto_discover"`
	Profiles     map[string]ProfileConfig `yaml:"profiles"`
}

func (pc *ProviderConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type defaultProvider ProviderConfig
	raw := defaultProvider{Enabled: true, AutoDiscover: true, Profiles: map[string]ProfileConfig{
		"default": DefaultProfileConfig(),
	}} // Set default values
	if err := unmarshal(&raw); err != nil {
		return err
	}
	*pc = ProviderConfig(raw)
	return nil
}

func (pc *ProviderConfig) Get(search []string) any {

	if len(search) == 0 {
		return pc
	}

	first := search[0]
	search = search[1:]

	switch first {
	case "auto_discover":
		return pc.AutoDiscover
	case "enabled":
		return pc.Enabled
	case "profiles":
		if len(search) == 0 {
			return pc.Profiles
		}
		if p, ok := pc.Profiles[search[0]]; ok {
			search = search[1:]
			return p.Get(search)
		}
	}

	return ""
}

func (pc *ProviderConfig) Set(search []string, val any) *ProviderConfig {
	first := search[0]
	search = search[1:]

	switch first {
	case "auto_discover":
		if len(search) == 0 {
			pc.AutoDiscover = val.(bool)
		}
	case "enabled":
		if len(search) == 0 {
			pc.Enabled = val.(bool)
		}
	case "profiles":
		if len(search) > 0 {
			profileName := search[0]
			if pp, ok := pc.Profiles[profileName]; ok {
				search = search[1:]
				pp.Set(search, val)
				pc.Profiles[profileName] = pp
			}
		}
	}
	return nil
}

type CarbonConfig struct {
	Dir       map[string]string         `yaml:"dir"`
	Providers map[string]ProviderConfig `yaml:"providers"`
}

func (cc *CarbonConfig) Get(search []string) any {

	if len(search) == 0 {
		return cc
	}

	first := search[0]
	search = search[1:]

	switch first {
	case "dir":
		if len(search) == 0 {
			return cc.Dir
		}
		if p, ok := cc.Dir[search[0]]; ok {
			return p
		}
	case "providers":
		if len(search) == 0 {
			return cc.Providers
		}
		if p, ok := cc.Providers[search[0]]; ok {
			search = search[1:]
			return p.Get(search)
		} else {
			log.Error("no provider", search[0])
		}
	}

	return cc
}

func (cc *CarbonConfig) Set(search []string, val any) *CarbonConfig {
	first := search[0]
	search = search[1:]

	switch first {
	case "dir":
		if len(search) == 1 {
			cc.Dir[search[0]] = val.(string)
		}
	case "providers":
		if len(search) > 0 {
			providerName := search[0]
			if p, ok := cc.Providers[providerName]; ok {
				search = search[1:]
				p.Set(search, val)
				cc.Providers[providerName] = p
			} else {
				log.Error("no provider", search[0])
			}
		}
	}

	return cc
}

type CarbonConfigFile struct {
	Carbon CarbonConfig `yaml:"carbon" mapstructure:"carbon"`
}

func (cf *CarbonConfigFile) MergeInConfigFile(cfgFile string) error {
	_, err := os.Stat(cfgFile)
	if err == nil || !os.IsNotExist(err) {
		//newCfg := CarbonConfigFile{}
		homeCfgBytes, err := os.ReadFile(cfgFile)
		if err != nil {
			return err
		}

		return yaml.Unmarshal(homeCfgBytes, &cf)
	}
	return err
}

func (cf *CarbonConfigFile) Get(search []string) any {

	first := search[0]
	search = search[1:]

	switch first {
	case "carbon":
		return cf.Carbon.Get(search)
	}

	return ""
}

func (cf *CarbonConfigFile) Set(search []string, val any) *CarbonConfigFile {

	first := search[0]
	search = search[1:]

	switch first {
	case "carbon":
		cf.Carbon.Set(search, val)
	}

	return cf

}

func DefaultProfileConfig() ProfileConfig {
	simpleCfg := "not_used: true"
	var defaultCfg ProfileConfig
	err := yaml.Unmarshal([]byte(simpleCfg), &defaultCfg)
	if err != nil {
		log.Debug("error getting default config", "error", err)
	}
	return defaultCfg
}

func DefaultProviderConfig() ProviderConfig {
	simpleCfg := "not_used: true"
	var defaultCfg ProviderConfig
	err := yaml.Unmarshal([]byte(simpleCfg), &defaultCfg)
	if err != nil {
		log.Debug("error getting default config", "error", err)
	}
	return defaultCfg
}

func GetConfig() *CarbonConfigFile {
	if instanceCfg == nil {
		instanceCfg = &CarbonConfigFile{
			Carbon: CarbonConfig{
				Dir: map[string]string{
					DefaultInstanceConfigKey:  DefaultInstanceDir,
					DeploymentsConfigKey:      DefaultDeploymentsDirName,
					PackerConfigKey:           filepath.Join(DefaultDeploymentsDirName, DefaultPackerDirName),
					ImagesConfigKey:           filepath.Join(DefaultDeploymentsDirName, DefaultImagesDirName),
					TerraformConfigKey:        filepath.Join(DefaultDeploymentsDirName, DefaultTerraformDirName),
					TerraformProjectConfigKey: filepath.Join(DefaultDeploymentsDirName, DefaultProjectsDirName),
				},
				Providers: map[string]ProviderConfig{
					"aws":          DefaultProviderConfig(),
					"digitalocean": DefaultProviderConfig(),
					"multipass":    DefaultProviderConfig(),
					"qemu":         DefaultProviderConfig(),
					"virtualbox":   DefaultProviderConfig(),
				},
			},
		}
	}

	return instanceCfg
}

func Get(s string) any {
	return instanceCfg.Get(strings.Split(s, "."))
}

func Set(s string, v any) any {
	instanceCfg = instanceCfg.Set(strings.Split(s, "."), v)
	return instanceCfg
}
