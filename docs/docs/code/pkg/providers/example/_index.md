---
title: example
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/example"
```

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
  - [func \(p \*Provider\) IsAvailable\(\) bool](<#Provider.IsAvailable>)
  - [func \(p \*Provider\) Profiles\(\) \[\]types.Profile](<#Provider.Profiles>)


<a name="New"></a>
## func [New](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/provider.go#L18>)

```go
func New() types.Provider
```



<a name="Environment"></a>
## type [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L11-L14>)



```go
type Environment struct {
    // contains filtered or unexported fields
}
```

<a name="Environment.CreateVM"></a>
### func \(\*Environment\) [CreateVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L66>)

```go
func (e *Environment) CreateVM(options types.MachineLaunchOptions) error
```



<a name="Environment.DestroyImage"></a>
### func \(\*Environment\) [DestroyImage](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L62>)

```go
func (e *Environment) DestroyImage(imageID string) error
```



<a name="Environment.DestroyVM"></a>
### func \(\*Environment\) [DestroyVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L58>)

```go
func (e *Environment) DestroyVM(id string) error
```



<a name="Environment.ImageBuilds"></a>
### func \(\*Environment\) [ImageBuilds](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L70>)

```go
func (e *Environment) ImageBuilds() ([]types.ImageBuild, error)
```



<a name="Environment.Images"></a>
### func \(\*Environment\) [Images](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L74>)

```go
func (e *Environment) Images() ([]types.Image, error)
```



<a name="Environment.Name"></a>
### func \(\*Environment\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L16>)

```go
func (e *Environment) Name() string
```



<a name="Environment.Profile"></a>
### func \(\*Environment\) [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L20>)

```go
func (e *Environment) Profile() types.Profile
```



<a name="Environment.RestartVM"></a>
### func \(\*Environment\) [RestartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L54>)

```go
func (e *Environment) RestartVM(id string) error
```



<a name="Environment.StartVM"></a>
### func \(\*Environment\) [StartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L46>)

```go
func (e *Environment) StartVM(id string) error
```



<a name="Environment.StopVM"></a>
### func \(\*Environment\) [StopVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L50>)

```go
func (e *Environment) StopVM(id string) error
```



<a name="Environment.VMs"></a>
### func \(\*Environment\) [VMs](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/environment.go#L24>)

```go
func (e *Environment) VMs() []types.VM
```



<a name="Profile"></a>
## type [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/profile.go#L9-L11>)



```go
type Profile struct {
    types.Profile
}
```

<a name="NewProfile"></a>
### func [NewProfile](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/profile.go#L13>)

```go
func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile
```



<a name="Profile.Environments"></a>
### func \(\*Profile\) [Environments](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/profile.go#L19>)

```go
func (p *Profile) Environments() []types.Environment
```



<a name="Provider"></a>
## type [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/provider.go#L12-L16>)



```go
type Provider struct {
    types.Provider
    // contains filtered or unexported fields
}
```

<a name="Provider.IsAvailable"></a>
### func \(\*Provider\) [IsAvailable](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/provider.go#L24>)

```go
func (p *Provider) IsAvailable() bool
```



<a name="Provider.Profiles"></a>
### func \(\*Provider\) [Profiles](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/example/provider.go#L28>)

```go
func (p *Provider) Profiles() []types.Profile
```



