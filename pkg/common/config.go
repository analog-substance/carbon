package common

type ProfileConfig struct {
	Enabled      bool            `yaml:"enabled"`
	Username     string          `yaml:"username"`
	Password     string          `yaml:"password"`
	URL          string          `yaml:"url"`
	Environments map[string]bool `yaml:"environments"`
}

type ProviderConfig struct {
	Enabled  bool                     `yaml:"enabled"`
	Profiles map[string]ProfileConfig `yaml:"profiles"`
}

type CarbonConfig struct {
	Providers map[string]ProviderConfig `yaml:"providers"`
}
type CarbonConfigFile struct {
	Carbon CarbonConfig `yaml:"carbon"`
}
