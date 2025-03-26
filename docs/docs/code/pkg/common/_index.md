---
title: common
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/common"
```

## Index

- [Constants](<#constants>)
- [func Get\(s string\) any](<#Get>)
- [func ImagesDir\(\) string](<#ImagesDir>)
- [func Keys\(\) \[\]string](<#Keys>)
- [func LogLevel\(level slog.Level\)](<#LogLevel>)
- [func Logger\(\) \*slog.Logger](<#Logger>)
- [func PackerDir\(\) string](<#PackerDir>)
- [func ProjectsDir\(\) string](<#ProjectsDir>)
- [func Set\(s string, v any\) any](<#Set>)
- [func SetProvidersTypes\(p \[\]string\)](<#SetProvidersTypes>)
- [func Time\(what string\) func\(\)](<#Time>)
- [func WithGroup\(groupName string\) \*slog.Logger](<#WithGroup>)
- [type CarbonConfig](<#CarbonConfig>)
  - [func \(cc \*CarbonConfig\) Get\(search \[\]string\) any](<#CarbonConfig.Get>)
  - [func \(cc \*CarbonConfig\) Keys\(prefix string\) \[\]string](<#CarbonConfig.Keys>)
  - [func \(cc \*CarbonConfig\) Set\(search \[\]string, val any\) \*CarbonConfig](<#CarbonConfig.Set>)
- [type CarbonConfigFile](<#CarbonConfigFile>)
  - [func GetConfig\(\) \*CarbonConfigFile](<#GetConfig>)
  - [func \(cf \*CarbonConfigFile\) Get\(search \[\]string\) any](<#CarbonConfigFile.Get>)
  - [func \(cf \*CarbonConfigFile\) Keys\(\) \[\]string](<#CarbonConfigFile.Keys>)
  - [func \(cf \*CarbonConfigFile\) MergeInConfigFile\(cfgFile string\) error](<#CarbonConfigFile.MergeInConfigFile>)
  - [func \(cf \*CarbonConfigFile\) Set\(search \[\]string, val any\) \*CarbonConfigFile](<#CarbonConfigFile.Set>)
- [type ProfileConfig](<#ProfileConfig>)
  - [func DefaultProfileConfig\(\) ProfileConfig](<#DefaultProfileConfig>)
  - [func \(pc \*ProfileConfig\) Get\(search \[\]string\) any](<#ProfileConfig.Get>)
  - [func \(pc \*ProfileConfig\) GetPassword\(\) string](<#ProfileConfig.GetPassword>)
  - [func \(pc \*ProfileConfig\) Keys\(prefix string\) \[\]string](<#ProfileConfig.Keys>)
  - [func \(pc \*ProfileConfig\) Set\(search \[\]string, val any\) \*ProfileConfig](<#ProfileConfig.Set>)
  - [func \(pc \*ProfileConfig\) UnmarshalYAML\(unmarshal func\(interface\{\}\) error\) error](<#ProfileConfig.UnmarshalYAML>)
- [type ProviderConfig](<#ProviderConfig>)
  - [func DefaultProviderConfig\(\) ProviderConfig](<#DefaultProviderConfig>)
  - [func \(pc \*ProviderConfig\) Get\(search \[\]string\) any](<#ProviderConfig.Get>)
  - [func \(pc \*ProviderConfig\) Keys\(prefix string\) \[\]string](<#ProviderConfig.Keys>)
  - [func \(pc \*ProviderConfig\) Set\(search \[\]string, val any\) \*ProviderConfig](<#ProviderConfig.Set>)
  - [func \(pc \*ProviderConfig\) UnmarshalYAML\(unmarshal func\(interface\{\}\) error\) error](<#ProviderConfig.UnmarshalYAML>)


## Constants

<a name="DefaultDeploymentsDirName"></a>

```go
const DefaultDeploymentsDirName = "deployments"
```

<a name="DefaultImagesDirName"></a>

```go
const DefaultImagesDirName = "images"
```

<a name="DefaultInstanceConfigKey"></a>

```go
const DefaultInstanceConfigKey = "instance"
```

<a name="DefaultInstanceDir"></a>

```go
const DefaultInstanceDir = "."
```

<a name="DefaultPackerDirName"></a>

```go
const DefaultPackerDirName = "packer"
```

<a name="DefaultProjectsDirName"></a>

```go
const DefaultProjectsDirName = "projects"
```

<a name="DefaultTerraformDirName"></a>

```go
const DefaultTerraformDirName = "terraform"
```

<a name="DeploymentsConfigKey"></a>

```go
const DeploymentsConfigKey = "deployments"
```

<a name="ImagesConfigKey"></a>

```go
const ImagesConfigKey = "images"
```

<a name="PackerConfigKey"></a>

```go
const PackerConfigKey = "packer"
```

<a name="TerraformConfigKey"></a>

```go
const TerraformConfigKey = "terraform"
```

<a name="TerraformProjectConfigKey"></a>

```go
const TerraformProjectConfigKey = "projects"
```

<a name="Get"></a>
## func [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L374>)

```go
func Get(s string) any
```



<a name="ImagesDir"></a>
## func [ImagesDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L30>)

```go
func ImagesDir() string
```



<a name="Keys"></a>
## func [Keys](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L370>)

```go
func Keys() []string
```



<a name="LogLevel"></a>
## func [LogLevel](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L27>)

```go
func LogLevel(level slog.Level)
```



<a name="Logger"></a>
## func [Logger](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L19>)

```go
func Logger() *slog.Logger
```



<a name="PackerDir"></a>
## func [PackerDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L26>)

```go
func PackerDir() string
```



<a name="ProjectsDir"></a>
## func [ProjectsDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L34>)

```go
func ProjectsDir() string
```



<a name="Set"></a>
## func [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L378>)

```go
func Set(s string, v any) any
```



<a name="SetProvidersTypes"></a>
## func [SetProvidersTypes](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L385>)

```go
func SetProvidersTypes(p []string)
```



<a name="Time"></a>
## func [Time](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L31>)

```go
func Time(what string) func()
```



<a name="WithGroup"></a>
## func [WithGroup](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L23>)

```go
func WithGroup(groupName string) *slog.Logger
```



<a name="CarbonConfig"></a>
## type [CarbonConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L204-L207>)



```go
type CarbonConfig struct {
    Dir       map[string]string         `yaml:"dir"`
    Providers map[string]ProviderConfig `yaml:"providers"`
}
```

<a name="CarbonConfig.Get"></a>
### func \(\*CarbonConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L223>)

```go
func (cc *CarbonConfig) Get(search []string) any
```



<a name="CarbonConfig.Keys"></a>
### func \(\*CarbonConfig\) [Keys](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L209>)

```go
func (cc *CarbonConfig) Keys(prefix string) []string
```



<a name="CarbonConfig.Set"></a>
### func \(\*CarbonConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L255>)

```go
func (cc *CarbonConfig) Set(search []string, val any) *CarbonConfig
```



<a name="CarbonConfigFile"></a>
## type [CarbonConfigFile](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L279-L281>)



```go
type CarbonConfigFile struct {
    Carbon CarbonConfig `yaml:"carbon"`
}
```

<a name="GetConfig"></a>
### func [GetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L349>)

```go
func GetConfig() *CarbonConfigFile
```



<a name="CarbonConfigFile.Get"></a>
### func \(\*CarbonConfigFile\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L302>)

```go
func (cf *CarbonConfigFile) Get(search []string) any
```



<a name="CarbonConfigFile.Keys"></a>
### func \(\*CarbonConfigFile\) [Keys](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L297>)

```go
func (cf *CarbonConfigFile) Keys() []string
```



<a name="CarbonConfigFile.MergeInConfigFile"></a>
### func \(\*CarbonConfigFile\) [MergeInConfigFile](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L283>)

```go
func (cf *CarbonConfigFile) MergeInConfigFile(cfgFile string) error
```



<a name="CarbonConfigFile.Set"></a>
### func \(\*CarbonConfigFile\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L315>)

```go
func (cf *CarbonConfigFile) Set(search []string, val any) *CarbonConfigFile
```



<a name="ProfileConfig"></a>
## type [ProfileConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L15-L23>)



```go
type ProfileConfig struct {
    Enabled         bool            `yaml:"enabled"`
    Username        string          `yaml:"username" `
    Password        string          `yaml:"password" `
    PasswordCommand string          `yaml:"password_command"`
    Use1PassCLI     bool            `yaml:"use_1pass_cli" `
    URL             string          `yaml:"url"`
    Environments    map[string]bool `yaml:"environments"`
}
```

<a name="DefaultProfileConfig"></a>
### func [DefaultProfileConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L329>)

```go
func DefaultProfileConfig() ProfileConfig
```



<a name="ProfileConfig.Get"></a>
### func \(\*ProfileConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L39>)

```go
func (pc *ProfileConfig) Get(search []string) any
```



<a name="ProfileConfig.GetPassword"></a>
### func \(\*ProfileConfig\) [GetPassword](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L100>)

```go
func (pc *ProfileConfig) GetPassword() string
```



<a name="ProfileConfig.Keys"></a>
### func \(\*ProfileConfig\) [Keys](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L25>)

```go
func (pc *ProfileConfig) Keys(prefix string) []string
```



<a name="ProfileConfig.Set"></a>
### func \(\*ProfileConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L66>)

```go
func (pc *ProfileConfig) Set(search []string, val any) *ProfileConfig
```



<a name="ProfileConfig.UnmarshalYAML"></a>
### func \(\*ProfileConfig\) [UnmarshalYAML](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L90>)

```go
func (pc *ProfileConfig) UnmarshalYAML(unmarshal func(interface{}) error) error
```



<a name="ProviderConfig"></a>
## type [ProviderConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L120-L124>)



```go
type ProviderConfig struct {
    Enabled      bool                     `yaml:"enabled"`
    AutoDiscover bool                     `yaml:"auto_discover"`
    Profiles     map[string]ProfileConfig `yaml:"profiles"`
}
```

<a name="DefaultProviderConfig"></a>
### func [DefaultProviderConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L339>)

```go
func DefaultProviderConfig() ProviderConfig
```



<a name="ProviderConfig.Get"></a>
### func \(\*ProviderConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L152>)

```go
func (pc *ProviderConfig) Get(search []string) any
```



<a name="ProviderConfig.Keys"></a>
### func \(\*ProviderConfig\) [Keys](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L138>)

```go
func (pc *ProviderConfig) Keys(prefix string) []string
```



<a name="ProviderConfig.Set"></a>
### func \(\*ProviderConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L179>)

```go
func (pc *ProviderConfig) Set(search []string, val any) *ProviderConfig
```



<a name="ProviderConfig.UnmarshalYAML"></a>
### func \(\*ProviderConfig\) [UnmarshalYAML](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L126>)

```go
func (pc *ProviderConfig) UnmarshalYAML(unmarshal func(interface{}) error) error
```



