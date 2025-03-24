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
- [func LogLevel\(level slog.Level\)](<#LogLevel>)
- [func Logger\(\) \*slog.Logger](<#Logger>)
- [func PackerDir\(\) string](<#PackerDir>)
- [func ProjectsDir\(\) string](<#ProjectsDir>)
- [func Set\(s string, v any\) any](<#Set>)
- [func WithGroup\(groupName string\) \*slog.Logger](<#WithGroup>)
- [type CarbonConfig](<#CarbonConfig>)
  - [func \(cc \*CarbonConfig\) Get\(search \[\]string\) any](<#CarbonConfig.Get>)
  - [func \(cc \*CarbonConfig\) Set\(search \[\]string, val any\) \*CarbonConfig](<#CarbonConfig.Set>)
- [type CarbonConfigFile](<#CarbonConfigFile>)
  - [func GetConfig\(\) \*CarbonConfigFile](<#GetConfig>)
  - [func \(cf \*CarbonConfigFile\) Get\(search \[\]string\) any](<#CarbonConfigFile.Get>)
  - [func \(cf \*CarbonConfigFile\) MergeInConfigFile\(cfgFile string\) error](<#CarbonConfigFile.MergeInConfigFile>)
  - [func \(cf \*CarbonConfigFile\) Set\(search \[\]string, val any\) \*CarbonConfigFile](<#CarbonConfigFile.Set>)
- [type ProfileConfig](<#ProfileConfig>)
  - [func DefaultProfileConfig\(\) ProfileConfig](<#DefaultProfileConfig>)
  - [func \(pc \*ProfileConfig\) Get\(search \[\]string\) any](<#ProfileConfig.Get>)
  - [func \(pc \*ProfileConfig\) GetPassword\(\) string](<#ProfileConfig.GetPassword>)
  - [func \(pc \*ProfileConfig\) Set\(search \[\]string, val any\) \*ProfileConfig](<#ProfileConfig.Set>)
  - [func \(pc \*ProfileConfig\) UnmarshalYAML\(unmarshal func\(interface\{\}\) error\) error](<#ProfileConfig.UnmarshalYAML>)
- [type ProviderConfig](<#ProviderConfig>)
  - [func DefaultProviderConfig\(\) ProviderConfig](<#DefaultProviderConfig>)
  - [func \(pc \*ProviderConfig\) Get\(search \[\]string\) any](<#ProviderConfig.Get>)
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
## func [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L330>)

```go
func Get(s string) any
```



<a name="ImagesDir"></a>
## func [ImagesDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L30>)

```go
func ImagesDir() string
```



<a name="LogLevel"></a>
## func [LogLevel](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L25>)

```go
func LogLevel(level slog.Level)
```



<a name="Logger"></a>
## func [Logger](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L17>)

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
## func [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L334>)

```go
func Set(s string, v any) any
```



<a name="WithGroup"></a>
## func [WithGroup](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L21>)

```go
func WithGroup(groupName string) *slog.Logger
```



<a name="CarbonConfig"></a>
## type [CarbonConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L177-L180>)



```go
type CarbonConfig struct {
    Dir       map[string]string         `yaml:"dir"`
    Providers map[string]ProviderConfig `yaml:"providers"`
}
```

<a name="CarbonConfig.Get"></a>
### func \(\*CarbonConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L182>)

```go
func (cc *CarbonConfig) Get(search []string) any
```



<a name="CarbonConfig.Set"></a>
### func \(\*CarbonConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L214>)

```go
func (cc *CarbonConfig) Set(search []string, val any) *CarbonConfig
```



<a name="CarbonConfigFile"></a>
## type [CarbonConfigFile](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L239-L241>)



```go
type CarbonConfigFile struct {
    Carbon CarbonConfig `yaml:"carbon" mapstructure:"carbon"`
}
```

<a name="GetConfig"></a>
### func [GetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L304>)

```go
func GetConfig() *CarbonConfigFile
```



<a name="CarbonConfigFile.Get"></a>
### func \(\*CarbonConfigFile\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L257>)

```go
func (cf *CarbonConfigFile) Get(search []string) any
```



<a name="CarbonConfigFile.MergeInConfigFile"></a>
### func \(\*CarbonConfigFile\) [MergeInConfigFile](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L243>)

```go
func (cf *CarbonConfigFile) MergeInConfigFile(cfgFile string) error
```



<a name="CarbonConfigFile.Set"></a>
### func \(\*CarbonConfigFile\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L270>)

```go
func (cf *CarbonConfigFile) Set(search []string, val any) *CarbonConfigFile
```



<a name="ProfileConfig"></a>
## type [ProfileConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L14-L22>)



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
### func [DefaultProfileConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L284>)

```go
func DefaultProfileConfig() ProfileConfig
```



<a name="ProfileConfig.Get"></a>
### func \(\*ProfileConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L24>)

```go
func (pc *ProfileConfig) Get(search []string) any
```



<a name="ProfileConfig.GetPassword"></a>
### func \(\*ProfileConfig\) [GetPassword](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L89>)

```go
func (pc *ProfileConfig) GetPassword() string
```



<a name="ProfileConfig.Set"></a>
### func \(\*ProfileConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L51>)

```go
func (pc *ProfileConfig) Set(search []string, val any) *ProfileConfig
```



<a name="ProfileConfig.UnmarshalYAML"></a>
### func \(\*ProfileConfig\) [UnmarshalYAML](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L79>)

```go
func (pc *ProfileConfig) UnmarshalYAML(unmarshal func(interface{}) error) error
```



<a name="ProviderConfig"></a>
## type [ProviderConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L106-L110>)



```go
type ProviderConfig struct {
    Enabled      bool                     `yaml:"enabled"`
    AutoDiscover bool                     `yaml:"auto_discover"`
    Profiles     map[string]ProfileConfig `yaml:"profiles"`
}
```

<a name="DefaultProviderConfig"></a>
### func [DefaultProviderConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L294>)

```go
func DefaultProviderConfig() ProviderConfig
```



<a name="ProviderConfig.Get"></a>
### func \(\*ProviderConfig\) [Get](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L124>)

```go
func (pc *ProviderConfig) Get(search []string) any
```



<a name="ProviderConfig.Set"></a>
### func \(\*ProviderConfig\) [Set](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L151>)

```go
func (pc *ProviderConfig) Set(search []string, val any) *ProviderConfig
```



<a name="ProviderConfig.UnmarshalYAML"></a>
### func \(\*ProviderConfig\) [UnmarshalYAML](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L112>)

```go
func (pc *ProviderConfig) UnmarshalYAML(unmarshal func(interface{}) error) error
```



