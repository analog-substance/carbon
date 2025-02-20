---
title: base
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/base"
```

## Index

- [Constants](<#constants>)
- [func DestroyImageForFileBasedProvider\(imageID string\) error](<#DestroyImageForFileBasedProvider>)
- [func GetImagesForFileBasedProvider\(providerType string, e types.Environment\) \(\[\]types.Image, error\)](<#GetImagesForFileBasedProvider>)
- [func New\(\) types.Provider](<#New>)
- [func NewProfile\(name string, providerInstance types.Provider, config common.ProfileConfig\) types.Profile](<#NewProfile>)
- [func NewWithName\(name string\) types.Provider](<#NewWithName>)
- [type Environment](<#Environment>)
  - [func \(e \*Environment\) CreateVM\(options types.MachineLaunchOptions\) error](<#Environment.CreateVM>)
  - [func \(e \*Environment\) DestroyImage\(imageID string\) error](<#Environment.DestroyImage>)
  - [func \(e \*Environment\) DestroyVM\(id string\) error](<#Environment.DestroyVM>)
  - [func \(e \*Environment\) ImageBuilds\(\) \(\[\]types.ImageBuild, error\)](<#Environment.ImageBuilds>)
  - [func \(e \*Environment\) Images\(\) \(\[\]types.Image, error\)](<#Environment.Images>)
  - [func \(e \*Environment\) Name\(\) string](<#Environment.Name>)
  - [func \(e \*Environment\) Profile\(\) types.Profile](<#Environment.Profile>)
  - [func \(e \*Environment\) RestartVM\(id string\) error](<#Environment.RestartVM>)
  - [func \(e \*Environment\) StartVM\(id string\) error](<#Environment.StartVM>)
  - [func \(e \*Environment\) StopVM\(id string\) error](<#Environment.StopVM>)
  - [func \(e \*Environment\) VMs\(\) \[\]types.VM](<#Environment.VMs>)
- [type ImageBuildDate](<#ImageBuildDate>)
- [type Profile](<#Profile>)
  - [func \(p \*Profile\) Environments\(\) \[\]types.Environment](<#Profile.Environments>)
  - [func \(p \*Profile\) GetConfig\(\) common.ProfileConfig](<#Profile.GetConfig>)
  - [func \(p \*Profile\) Name\(\) string](<#Profile.Name>)
  - [func \(p \*Profile\) Provider\(\) types.Provider](<#Profile.Provider>)
  - [func \(p \*Profile\) SetConfig\(config common.ProfileConfig\)](<#Profile.SetConfig>)
  - [func \(p \*Profile\) ShouldIncludeEnvironment\(envName string\) bool](<#Profile.ShouldIncludeEnvironment>)
- [type Provider](<#Provider>)
  - [func \(p \*Provider\) GetConfig\(\) common.ProviderConfig](<#Provider.GetConfig>)
  - [func \(p \*Provider\) IsAvailable\(\) bool](<#Provider.IsAvailable>)
  - [func \(p \*Provider\) Name\(\) string](<#Provider.Name>)
  - [func \(p \*Provider\) NewImageBuild\(name, tplDir string\) \(types.ImageBuild, error\)](<#Provider.NewImageBuild>)
  - [func \(p \*Provider\) NewProject\(name string, force bool\) \(types.Project, error\)](<#Provider.NewProject>)
  - [func \(p \*Provider\) Profiles\(\) \[\]types.Profile](<#Provider.Profiles>)
  - [func \(p \*Provider\) SetConfig\(config common.ProviderConfig\)](<#Provider.SetConfig>)
  - [func \(p \*Provider\) Type\(\) string](<#Provider.Type>)


## Constants

<a name="CloudInitDir"></a>

```go
const CloudInitDir = "cloud-init"
```

<a name="ISOVarUsage"></a>

```go
const ISOVarUsage = "var.iso_url"
```

<a name="PackerFileIsoVars"></a>

```go
const PackerFileIsoVars = "iso-variables.pkr.hcl"
```

<a name="PackerFileLocalVars"></a>

```go
const PackerFileLocalVars = "local-variables.pkr.hcl"
```

<a name="PackerFilePacker"></a>

```go
const PackerFilePacker = "packer.pkr.hcl"
```

<a name="PackerFilePrivateVarsExample"></a>

```go
const PackerFilePrivateVarsExample = "private.auto.pkrvars.hcl.example"
```

<a name="PackerFileSuffixAnsible"></a>

```go
const PackerFileSuffixAnsible = "-ansible.pkr.hcl"
```

<a name="PackerFileSuffixCloudInit"></a>

```go
const PackerFileSuffixCloudInit = "-cloud-init.pkr.hcl"
```

<a name="PackerFileSuffixVariables"></a>

```go
const PackerFileSuffixVariables = "-variables.pkr.hcl"
```

<a name="DestroyImageForFileBasedProvider"></a>
## func [DestroyImageForFileBasedProvider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/util.go#L21>)

```go
func DestroyImageForFileBasedProvider(imageID string) error
```



<a name="GetImagesForFileBasedProvider"></a>
## func [GetImagesForFileBasedProvider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/util.go#L12>)

```go
func GetImagesForFileBasedProvider(providerType string, e types.Environment) ([]types.Image, error)
```



<a name="New"></a>
## func [New](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L37>)

```go
func New() types.Provider
```



<a name="NewProfile"></a>
## func [NewProfile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L15>)

```go
func NewProfile(name string, providerInstance types.Provider, config common.ProfileConfig) types.Profile
```



<a name="NewWithName"></a>
## func [NewWithName](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L43>)

```go
func NewWithName(name string) types.Provider
```



<a name="Environment"></a>
## type [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L9-L12>)



```go
type Environment struct {
    // contains filtered or unexported fields
}
```

<a name="Environment.CreateVM"></a>
### func \(\*Environment\) [CreateVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L44>)

```go
func (e *Environment) CreateVM(options types.MachineLaunchOptions) error
```



<a name="Environment.DestroyImage"></a>
### func \(\*Environment\) [DestroyImage](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L56>)

```go
func (e *Environment) DestroyImage(imageID string) error
```



<a name="Environment.DestroyVM"></a>
### func \(\*Environment\) [DestroyVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L40>)

```go
func (e *Environment) DestroyVM(id string) error
```



<a name="Environment.ImageBuilds"></a>
### func \(\*Environment\) [ImageBuilds](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L48>)

```go
func (e *Environment) ImageBuilds() ([]types.ImageBuild, error)
```



<a name="Environment.Images"></a>
### func \(\*Environment\) [Images](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L52>)

```go
func (e *Environment) Images() ([]types.Image, error)
```



<a name="Environment.Name"></a>
### func \(\*Environment\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L14>)

```go
func (e *Environment) Name() string
```



<a name="Environment.Profile"></a>
### func \(\*Environment\) [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L18>)

```go
func (e *Environment) Profile() types.Profile
```



<a name="Environment.RestartVM"></a>
### func \(\*Environment\) [RestartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L36>)

```go
func (e *Environment) RestartVM(id string) error
```



<a name="Environment.StartVM"></a>
### func \(\*Environment\) [StartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L28>)

```go
func (e *Environment) StartVM(id string) error
```



<a name="Environment.StopVM"></a>
### func \(\*Environment\) [StopVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L32>)

```go
func (e *Environment) StopVM(id string) error
```



<a name="Environment.VMs"></a>
### func \(\*Environment\) [VMs](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/environment.go#L22>)

```go
func (e *Environment) VMs() []types.VM
```



<a name="ImageBuildDate"></a>
## type [ImageBuildDate](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L263-L265>)



```go
type ImageBuildDate struct {
    Name string
}
```

<a name="Profile"></a>
## type [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L9-L13>)



```go
type Profile struct {
    // contains filtered or unexported fields
}
```

<a name="Profile.Environments"></a>
### func \(\*Profile\) [Environments](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L23>)

```go
func (p *Profile) Environments() []types.Environment
```



<a name="Profile.GetConfig"></a>
### func \(\*Profile\) [GetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L42>)

```go
func (p *Profile) GetConfig() common.ProfileConfig
```



<a name="Profile.Name"></a>
### func \(\*Profile\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L31>)

```go
func (p *Profile) Name() string
```



<a name="Profile.Provider"></a>
### func \(\*Profile\) [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L34>)

```go
func (p *Profile) Provider() types.Provider
```



<a name="Profile.SetConfig"></a>
### func \(\*Profile\) [SetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L38>)

```go
func (p *Profile) SetConfig(config common.ProfileConfig)
```



<a name="Profile.ShouldIncludeEnvironment"></a>
### func \(\*Profile\) [ShouldIncludeEnvironment](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/profile.go#L46>)

```go
func (p *Profile) ShouldIncludeEnvironment(envName string) bool
```



<a name="Provider"></a>
## type [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L31-L35>)



```go
type Provider struct {
    // contains filtered or unexported fields
}
```

<a name="Provider.GetConfig"></a>
### func \(\*Provider\) [GetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L74>)

```go
func (p *Provider) GetConfig() common.ProviderConfig
```



<a name="Provider.IsAvailable"></a>
### func \(\*Provider\) [IsAvailable](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L49>)

```go
func (p *Provider) IsAvailable() bool
```



<a name="Provider.Name"></a>
### func \(\*Provider\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L53>)

```go
func (p *Provider) Name() string
```



<a name="Provider.NewImageBuild"></a>
### func \(\*Provider\) [NewImageBuild](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L78>)

```go
func (p *Provider) NewImageBuild(name, tplDir string) (types.ImageBuild, error)
```



<a name="Provider.NewProject"></a>
### func \(\*Provider\) [NewProject](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L176>)

```go
func (p *Provider) NewProject(name string, force bool) (types.Project, error)
```



<a name="Provider.Profiles"></a>
### func \(\*Provider\) [Profiles](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L61>)

```go
func (p *Provider) Profiles() []types.Profile
```



<a name="Provider.SetConfig"></a>
### func \(\*Provider\) [SetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L70>)

```go
func (p *Provider) SetConfig(config common.ProviderConfig)
```



<a name="Provider.Type"></a>
### func \(\*Provider\) [Type](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/base/provider.go#L57>)

```go
func (p *Provider) Type() string
```



