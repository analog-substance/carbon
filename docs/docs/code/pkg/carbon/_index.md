---
title: carbon
description: Package carbon provides core application functionality and constants
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/carbon"
```

Package carbon provides core application functionality and constants

## Index

- [Constants](<#constants>)
- [func AvailableProviders\(\) \[\]types.Provider](<#AvailableProviders>)
- [type Carbon](<#Carbon>)
  - [func New\(config common.CarbonConfig\) \*Carbon](<#New>)
  - [func \(c \*Carbon\) FindVMByID\(id string\) \[\]types.VM](<#Carbon.FindVMByID>)
  - [func \(c \*Carbon\) FindVMByName\(name string\) \[\]types.VM](<#Carbon.FindVMByName>)
  - [func \(c \*Carbon\) GetImage\(imageID string\) \(types.Image, error\)](<#Carbon.GetImage>)
  - [func \(c \*Carbon\) GetImageBuild\(name, provider, provisioner string\) \(types.ImageBuild, error\)](<#Carbon.GetImageBuild>)
  - [func \(c \*Carbon\) GetImageBuildTemplates\(\) \[\]string](<#Carbon.GetImageBuildTemplates>)
  - [func \(c \*Carbon\) GetImageBuilds\(\) \(\[\]types.ImageBuild, error\)](<#Carbon.GetImageBuilds>)
  - [func \(c \*Carbon\) GetImages\(\) \(\[\]types.Image, error\)](<#Carbon.GetImages>)
  - [func \(c \*Carbon\) GetProject\(name string\) \(types.Project, error\)](<#Carbon.GetProject>)
  - [func \(c \*Carbon\) GetProjects\(\) \(\[\]types.Project, error\)](<#Carbon.GetProjects>)
  - [func \(c \*Carbon\) GetProvider\(providerType string\) \(types.Provider, error\)](<#Carbon.GetProvider>)
  - [func \(c \*Carbon\) GetVMs\(\) \[\]types.VM](<#Carbon.GetVMs>)
  - [func \(c \*Carbon\) Profiles\(\) \[\]types.Profile](<#Carbon.Profiles>)
  - [func \(c \*Carbon\) Providers\(\) \[\]types.Provider](<#Carbon.Providers>)
  - [func \(c \*Carbon\) VMsFromHosts\(hostnames \[\]string\) \[\]types.VM](<#Carbon.VMsFromHosts>)
- [type Options](<#Options>)


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

<a name="AvailableProviders"></a>
## func [AvailableProviders](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/providers.go#L15>)

```go
func AvailableProviders() []types.Provider
```



<a name="Carbon"></a>
## type [Carbon](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/main.go#L16-L25>)



```go
type Carbon struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/main.go#L33>)

```go
func New(config common.CarbonConfig) *Carbon
```



<a name="Carbon.FindVMByID"></a>
### func \(\*Carbon\) [FindVMByID](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/virtual_machines.go#L23>)

```go
func (c *Carbon) FindVMByID(id string) []types.VM
```



<a name="Carbon.FindVMByName"></a>
### func \(\*Carbon\) [FindVMByName](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/virtual_machines.go#L32>)

```go
func (c *Carbon) FindVMByName(name string) []types.VM
```



<a name="Carbon.GetImage"></a>
### func \(\*Carbon\) [GetImage](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/images.go#L80>)

```go
func (c *Carbon) GetImage(imageID string) (types.Image, error)
```



<a name="Carbon.GetImageBuild"></a>
### func \(\*Carbon\) [GetImageBuild](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/images.go#L49>)

```go
func (c *Carbon) GetImageBuild(name, provider, provisioner string) (types.ImageBuild, error)
```



<a name="Carbon.GetImageBuildTemplates"></a>
### func \(\*Carbon\) [GetImageBuildTemplates](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/images.go#L20>)

```go
func (c *Carbon) GetImageBuildTemplates() []string
```



<a name="Carbon.GetImageBuilds"></a>
### func \(\*Carbon\) [GetImageBuilds](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/images.go#L32>)

```go
func (c *Carbon) GetImageBuilds() ([]types.ImageBuild, error)
```



<a name="Carbon.GetImages"></a>
### func \(\*Carbon\) [GetImages](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/images.go#L63>)

```go
func (c *Carbon) GetImages() ([]types.Image, error)
```



<a name="Carbon.GetProject"></a>
### func \(\*Carbon\) [GetProject](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/projects.go#L30>)

```go
func (c *Carbon) GetProject(name string) (types.Project, error)
```



<a name="Carbon.GetProjects"></a>
### func \(\*Carbon\) [GetProjects](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/projects.go#L12>)

```go
func (c *Carbon) GetProjects() ([]types.Project, error)
```



<a name="Carbon.GetProvider"></a>
### func \(\*Carbon\) [GetProvider](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/providers.go#L41>)

```go
func (c *Carbon) GetProvider(providerType string) (types.Provider, error)
```



<a name="Carbon.GetVMs"></a>
### func \(\*Carbon\) [GetVMs](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/virtual_machines.go#L10>)

```go
func (c *Carbon) GetVMs() []types.VM
```



<a name="Carbon.Profiles"></a>
### func \(\*Carbon\) [Profiles](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/profiles.go#L5>)

```go
func (c *Carbon) Profiles() []types.Profile
```



<a name="Carbon.Providers"></a>
### func \(\*Carbon\) [Providers](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/providers.go#L37>)

```go
func (c *Carbon) Providers() []types.Provider
```



<a name="Carbon.VMsFromHosts"></a>
### func \(\*Carbon\) [VMsFromHosts](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/virtual_machines.go#L46>)

```go
func (c *Carbon) VMsFromHosts(hostnames []string) []types.VM
```



<a name="Options"></a>
## type [Options](<https://github.com/analog-substance/carbon/blob/main/pkg/carbon/main.go#L10-L14>)



```go
type Options struct {
    Providers    []string
    Profiles     []string
    Environments []string
}
```

