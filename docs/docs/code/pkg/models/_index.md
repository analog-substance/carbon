---
title: models
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/models"
```

## Index

- [func GetImageBuildsForProvider\(provider string\) \(\[\]types.ImageBuild, error\)](<#GetImageBuildsForProvider>)
- [func NewImage\(imageID string, imageName string, createdAt time.Time, env types.Environment\) types.Image](<#NewImage>)
- [type BuildBlock](<#BuildBlock>)
- [type Image](<#Image>)
  - [func \(i \*Image\) CreatedAt\(\) string](<#Image.CreatedAt>)
  - [func \(i \*Image\) Destroy\(\) error](<#Image.Destroy>)
  - [func \(i \*Image\) Environment\(\) types.Environment](<#Image.Environment>)
  - [func \(i \*Image\) ID\(\) string](<#Image.ID>)
  - [func \(i \*Image\) Launch\(imageLaunchOptions types.ImageLaunchOptions\) error](<#Image.Launch>)
  - [func \(i \*Image\) MarshalJSON\(\) \(\[\]byte, error\)](<#Image.MarshalJSON>)
  - [func \(i \*Image\) Name\(\) string](<#Image.Name>)
  - [func \(i \*Image\) Profile\(\) types.Profile](<#Image.Profile>)
  - [func \(i \*Image\) Provider\(\) types.Provider](<#Image.Provider>)
- [type ImageBuild](<#ImageBuild>)
  - [func NewImageBuild\(buildPath, provider, provisioner string\) \*ImageBuild](<#NewImageBuild>)
  - [func \(b \*ImageBuild\) Build\(\) error](<#ImageBuild.Build>)
  - [func \(b \*ImageBuild\) MarshalJSON\(\) \(\[\]byte, error\)](<#ImageBuild.MarshalJSON>)
  - [func \(b \*ImageBuild\) Name\(\) string](<#ImageBuild.Name>)
  - [func \(b \*ImageBuild\) ProviderType\(\) string](<#ImageBuild.ProviderType>)
  - [func \(b \*ImageBuild\) Provisioner\(\) string](<#ImageBuild.Provisioner>)
- [type Machine](<#Machine>)
  - [func \(m \*Machine\) Cmd\(user string, privateIP bool, cmdArgs ...string\) \(string, error\)](<#Machine.Cmd>)
  - [func \(m \*Machine\) Destroy\(\) error](<#Machine.Destroy>)
  - [func \(m \*Machine\) Environment\(\) types.Environment](<#Machine.Environment>)
  - [func \(m \*Machine\) ExecSSH\(user string, privateIP bool, cmdArgs ...string\) error](<#Machine.ExecSSH>)
  - [func \(m \*Machine\) ID\(\) string](<#Machine.ID>)
  - [func \(m \*Machine\) IPAddress\(\) string](<#Machine.IPAddress>)
  - [func \(m \*Machine\) Name\(\) string](<#Machine.Name>)
  - [func \(m \*Machine\) NewSSHSession\(user string, privateIP bool\) \(\*ssh\_util.Session, error\)](<#Machine.NewSSHSession>)
  - [func \(m \*Machine\) PrivateIPAddress\(\) string](<#Machine.PrivateIPAddress>)
  - [func \(m \*Machine\) Profile\(\) types.Profile](<#Machine.Profile>)
  - [func \(m \*Machine\) Provider\(\) types.Provider](<#Machine.Provider>)
  - [func \(m \*Machine\) Restart\(\) error](<#Machine.Restart>)
  - [func \(m \*Machine\) Start\(\) error](<#Machine.Start>)
  - [func \(m \*Machine\) StartRDPClient\(user string, privateIP bool\) error](<#Machine.StartRDPClient>)
  - [func \(m \*Machine\) StartVNC\(user string, privateIP bool, killVNC bool\) error](<#Machine.StartVNC>)
  - [func \(m \*Machine\) State\(\) string](<#Machine.State>)
  - [func \(m \*Machine\) Stop\(\) error](<#Machine.Stop>)
  - [func \(m \*Machine\) Type\(\) string](<#Machine.Type>)
  - [func \(m \*Machine\) UpTime\(\) time.Duration](<#Machine.UpTime>)
- [type PackerConfig](<#PackerConfig>)
- [type Project](<#Project>)
  - [func NewProject\(buildPath string\) \*Project](<#NewProject>)
  - [func \(d \*Project\) AddMachine\(machine \*types.ProjectMachine, noApply bool\) error](<#Project.AddMachine>)
  - [func \(d \*Project\) GetConfig\(\) \(\*types.ProjectConfig, error\)](<#Project.GetConfig>)
  - [func \(d \*Project\) MarshalJSON\(\) \(\[\]byte, error\)](<#Project.MarshalJSON>)
  - [func \(d \*Project\) Name\(\) string](<#Project.Name>)
  - [func \(d \*Project\) SaveConfig\(\) error](<#Project.SaveConfig>)
  - [func \(d \*Project\) TerraformApply\(\) error](<#Project.TerraformApply>)
- [type SourceBlock](<#SourceBlock>)


<a name="GetImageBuildsForProvider"></a>
## func [GetImageBuildsForProvider](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L99>)

```go
func GetImageBuildsForProvider(provider string) ([]types.ImageBuild, error)
```



<a name="NewImage"></a>
## func [NewImage](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L16>)

```go
func NewImage(imageID string, imageName string, createdAt time.Time, env types.Environment) types.Image
```



<a name="BuildBlock"></a>
## type [BuildBlock](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L87-L92>)



```go
type BuildBlock struct {
    Name        string   `hcl:"name,optional"`
    Description string   `hcl:"description,optional"`
    FromSources []string `hcl:"sources,optional"`
    Config      hcl.Body `hcl:",remain"`
}
```

<a name="Image"></a>
## type [Image](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L9-L14>)



```go
type Image struct {
    // contains filtered or unexported fields
}
```

<a name="Image.CreatedAt"></a>
### func \(\*Image\) [CreatedAt](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L45>)

```go
func (i *Image) CreatedAt() string
```



<a name="Image.Destroy"></a>
### func \(\*Image\) [Destroy](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L59>)

```go
func (i *Image) Destroy() error
```



<a name="Image.Environment"></a>
### func \(\*Image\) [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L63>)

```go
func (i *Image) Environment() types.Environment
```



<a name="Image.ID"></a>
### func \(\*Image\) [ID](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L37>)

```go
func (i *Image) ID() string
```



<a name="Image.Launch"></a>
### func \(\*Image\) [Launch](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L49>)

```go
func (i *Image) Launch(imageLaunchOptions types.ImageLaunchOptions) error
```



<a name="Image.MarshalJSON"></a>
### func \(\*Image\) [MarshalJSON](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L25>)

```go
func (i *Image) MarshalJSON() ([]byte, error)
```



<a name="Image.Name"></a>
### func \(\*Image\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L41>)

```go
func (i *Image) Name() string
```



<a name="Image.Profile"></a>
### func \(\*Image\) [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L71>)

```go
func (i *Image) Profile() types.Profile
```



<a name="Image.Provider"></a>
### func \(\*Image\) [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image.go#L67>)

```go
func (i *Image) Provider() types.Provider
```



<a name="ImageBuild"></a>
## type [ImageBuild](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L19-L23>)



```go
type ImageBuild struct {
    // contains filtered or unexported fields
}
```

<a name="NewImageBuild"></a>
### func [NewImageBuild](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L25>)

```go
func NewImageBuild(buildPath, provider, provisioner string) *ImageBuild
```



<a name="ImageBuild.Build"></a>
### func \(\*ImageBuild\) [Build](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L57>)

```go
func (b *ImageBuild) Build() error
```



<a name="ImageBuild.MarshalJSON"></a>
### func \(\*ImageBuild\) [MarshalJSON](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L33>)

```go
func (b *ImageBuild) MarshalJSON() ([]byte, error)
```



<a name="ImageBuild.Name"></a>
### func \(\*ImageBuild\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L45>)

```go
func (b *ImageBuild) Name() string
```



<a name="ImageBuild.ProviderType"></a>
### func \(\*ImageBuild\) [ProviderType](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L49>)

```go
func (b *ImageBuild) ProviderType() string
```



<a name="ImageBuild.Provisioner"></a>
### func \(\*ImageBuild\) [Provisioner](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L53>)

```go
func (b *ImageBuild) Provisioner() string
```



<a name="Machine"></a>
## type [Machine](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L23-L32>)



```go
type Machine struct {
    InstanceName       string             `json:"name"`
    InstanceID         string             `json:"id"`
    CurrentUpTime      time.Duration      `json:"up_time"`
    InstanceType       string             `json:"type"`
    PublicIPAddresses  []string           `json:"public_ip_addresses"`
    PrivateIPAddresses []string           `json:"private_ip_addresses"`
    CurrentState       types.MachineState `json:"current_state"`
    Env                types.Environment  `json:"-"`
}
```

<a name="Machine.Cmd"></a>
### func \(\*Machine\) [Cmd](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L253>)

```go
func (m *Machine) Cmd(user string, privateIP bool, cmdArgs ...string) (string, error)
```



<a name="Machine.Destroy"></a>
### func \(\*Machine\) [Destroy](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L80>)

```go
func (m *Machine) Destroy() error
```



<a name="Machine.Environment"></a>
### func \(\*Machine\) [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L34>)

```go
func (m *Machine) Environment() types.Environment
```



<a name="Machine.ExecSSH"></a>
### func \(\*Machine\) [ExecSSH](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L96>)

```go
func (m *Machine) ExecSSH(user string, privateIP bool, cmdArgs ...string) error
```



<a name="Machine.ID"></a>
### func \(\*Machine\) [ID](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L50>)

```go
func (m *Machine) ID() string
```



<a name="Machine.IPAddress"></a>
### func \(\*Machine\) [IPAddress](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L54>)

```go
func (m *Machine) IPAddress() string
```



<a name="Machine.Name"></a>
### func \(\*Machine\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L46>)

```go
func (m *Machine) Name() string
```



<a name="Machine.NewSSHSession"></a>
### func \(\*Machine\) [NewSSHSession](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L213>)

```go
func (m *Machine) NewSSHSession(user string, privateIP bool) (*ssh_util.Session, error)
```



<a name="Machine.PrivateIPAddress"></a>
### func \(\*Machine\) [PrivateIPAddress](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L61>)

```go
func (m *Machine) PrivateIPAddress() string
```



<a name="Machine.Profile"></a>
### func \(\*Machine\) [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L38>)

```go
func (m *Machine) Profile() types.Profile
```



<a name="Machine.Provider"></a>
### func \(\*Machine\) [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L42>)

```go
func (m *Machine) Provider() types.Provider
```



<a name="Machine.Restart"></a>
### func \(\*Machine\) [Restart](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L92>)

```go
func (m *Machine) Restart() error
```



<a name="Machine.Start"></a>
### func \(\*Machine\) [Start](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L84>)

```go
func (m *Machine) Start() error
```



<a name="Machine.StartRDPClient"></a>
### func \(\*Machine\) [StartRDPClient](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L191>)

```go
func (m *Machine) StartRDPClient(user string, privateIP bool) error
```



<a name="Machine.StartVNC"></a>
### func \(\*Machine\) [StartVNC](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L128>)

```go
func (m *Machine) StartVNC(user string, privateIP bool, killVNC bool) error
```

StartVNC will create a VNC session on the virtual machine It accomplishes this by:

- SSH to the VM.
- Start VNC if it is not already running.
- Forward a port through the SSH session.
- VNC to the forwarded port.

Requires TigerVNC to be installed.

<a name="Machine.State"></a>
### func \(\*Machine\) [State](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L72>)

```go
func (m *Machine) State() string
```



<a name="Machine.Stop"></a>
### func \(\*Machine\) [Stop](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L88>)

```go
func (m *Machine) Stop() error
```



<a name="Machine.Type"></a>
### func \(\*Machine\) [Type](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L76>)

```go
func (m *Machine) Type() string
```



<a name="Machine.UpTime"></a>
### func \(\*Machine\) [UpTime](<https://github.com/analog-substance/carbon/blob/main/pkg/models/machine.go#L68>)

```go
func (m *Machine) UpTime() time.Duration
```



<a name="PackerConfig"></a>
## type [PackerConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L94-L97>)



```go
type PackerConfig struct {
    Source SourceBlock `hcl:"source,block"`
    Build  BuildBlock  `hcl:"build,block"`
}
```

<a name="Project"></a>
## type [Project](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L15-L19>)



```go
type Project struct {
    // contains filtered or unexported fields
}
```

<a name="NewProject"></a>
### func [NewProject](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L21>)

```go
func NewProject(buildPath string) *Project
```



<a name="Project.AddMachine"></a>
### func \(\*Project\) [AddMachine](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L52>)

```go
func (d *Project) AddMachine(machine *types.ProjectMachine, noApply bool) error
```



<a name="Project.GetConfig"></a>
### func \(\*Project\) [GetConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L105>)

```go
func (d *Project) GetConfig() (*types.ProjectConfig, error)
```



<a name="Project.MarshalJSON"></a>
### func \(\*Project\) [MarshalJSON](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L27>)

```go
func (d *Project) MarshalJSON() ([]byte, error)
```



<a name="Project.Name"></a>
### func \(\*Project\) [Name](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L35>)

```go
func (d *Project) Name() string
```



<a name="Project.SaveConfig"></a>
### func \(\*Project\) [SaveConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L122>)

```go
func (d *Project) SaveConfig() error
```



<a name="Project.TerraformApply"></a>
### func \(\*Project\) [TerraformApply](<https://github.com/analog-substance/carbon/blob/main/pkg/models/project.go#L39>)

```go
func (d *Project) TerraformApply() error
```



<a name="SourceBlock"></a>
## type [SourceBlock](<https://github.com/analog-substance/carbon/blob/main/pkg/models/image_build.go#L81-L85>)



```go
type SourceBlock struct {
    Type   string   `hcl:"type,label"`
    Name   string   `hcl:"name,label"`
    Config hcl.Body `hcl:",remain"`
}
```

