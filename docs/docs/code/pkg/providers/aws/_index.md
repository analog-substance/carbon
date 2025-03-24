---
title: aws
description: Package aws handles communications with AWS APIs
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/aws"
```

Package aws handles communications with AWS APIs

## Index

- [func New\(\) types.Provider](<#New>)
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
- [type Profile](<#Profile>)
  - [func NewProfile\(name string, providerInstance \*Provider, config common.ProfileConfig\) \*Profile](<#NewProfile>)
  - [func \(p \*Profile\) Environments\(\) \[\]types.Environment](<#Profile.Environments>)
- [type Provider](<#Provider>)
  - [func \(p \*Provider\) AWSProfiles\(\) \[\]string](<#Provider.AWSProfiles>)
  - [func \(p \*Provider\) IsAvailable\(\) bool](<#Provider.IsAvailable>)
  - [func \(p \*Provider\) Profiles\(\) \[\]types.Profile](<#Provider.Profiles>)


<a name="New"></a>
## func [New](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/provider.go#L24>)

```go
func New() types.Provider
```

New creates new instance of an AWS Provider and returns it. Defaults to no awsProfileNames, this forces a query of the AWS config at runtime.

<a name="Environment"></a>
## type [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L14-L20>)



```go
type Environment struct {
    // contains filtered or unexported fields
}
```

<a name="Environment.CreateVM"></a>
### func \(\*Environment\) [CreateVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L86>)

```go
func (e *Environment) CreateVM(options types.MachineLaunchOptions) error
```



<a name="Environment.DestroyImage"></a>
### func \(\*Environment\) [DestroyImage](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L187>)

```go
func (e *Environment) DestroyImage(imageID string) error
```



<a name="Environment.DestroyVM"></a>
### func \(\*Environment\) [DestroyVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L79>)

```go
func (e *Environment) DestroyVM(id string) error
```



<a name="Environment.ImageBuilds"></a>
### func \(\*Environment\) [ImageBuilds](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L95>)

```go
func (e *Environment) ImageBuilds() ([]types.ImageBuild, error)
```



<a name="Environment.Images"></a>
### func \(\*Environment\) [Images](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L98>)

```go
func (e *Environment) Images() ([]types.Image, error)
```



<a name="Environment.Name"></a>
### func \(\*Environment\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L22>)

```go
func (e *Environment) Name() string
```



<a name="Environment.Profile"></a>
### func \(\*Environment\) [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L26>)

```go
func (e *Environment) Profile() types.Profile
```



<a name="Environment.RestartVM"></a>
### func \(\*Environment\) [RestartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L72>)

```go
func (e *Environment) RestartVM(id string) error
```



<a name="Environment.StartVM"></a>
### func \(\*Environment\) [StartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L56>)

```go
func (e *Environment) StartVM(id string) error
```



<a name="Environment.StopVM"></a>
### func \(\*Environment\) [StopVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L63>)

```go
func (e *Environment) StopVM(id string) error
```



<a name="Environment.VMs"></a>
### func \(\*Environment\) [VMs](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/environment.go#L30>)

```go
func (e *Environment) VMs() []types.VM
```



<a name="Profile"></a>
## type [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/profile.go#L13-L15>)



```go
type Profile struct {
    types.Profile
}
```

<a name="NewProfile"></a>
### func [NewProfile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/profile.go#L17>)

```go
func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile
```



<a name="Profile.Environments"></a>
### func \(\*Profile\) [Environments](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/profile.go#L23>)

```go
func (p *Profile) Environments() []types.Environment
```



<a name="Provider"></a>
## type [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/provider.go#L16-L20>)



```go
type Provider struct {
    types.Provider
    // contains filtered or unexported fields
}
```

<a name="Provider.AWSProfiles"></a>
### func \(\*Provider\) [AWSProfiles](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/provider.go#L31>)

```go
func (p *Provider) AWSProfiles() []string
```



<a name="Provider.IsAvailable"></a>
### func \(\*Provider\) [IsAvailable](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/provider.go#L62>)

```go
func (p *Provider) IsAvailable() bool
```



<a name="Provider.Profiles"></a>
### func \(\*Provider\) [Profiles](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/aws/provider.go#L66>)

```go
func (p *Provider) Profiles() []types.Profile
```



