---
title: types
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/types"
```

## Index

- [Variables](<#variables>)
- [type Environment](<#Environment>)
- [type Image](<#Image>)
- [type ImageBuild](<#ImageBuild>)
- [type ImageLaunchOptions](<#ImageLaunchOptions>)
- [type MachineLaunchOptions](<#MachineLaunchOptions>)
- [type MachineState](<#MachineState>)
- [type Profile](<#Profile>)
- [type Project](<#Project>)
- [type ProjectConfig](<#ProjectConfig>)
- [type ProjectMachine](<#ProjectMachine>)
- [type Provider](<#Provider>)
- [type VM](<#VM>)


## Variables

<a name="StateRunning"></a>

```go
var StateRunning = MachineState{"Running"}
```

<a name="StateSleeping"></a>

```go
var StateSleeping = MachineState{"Sleeping"}
```

<a name="StateStarting"></a>

```go
var StateStarting = MachineState{"Starting"}
```

<a name="StateStopped"></a>

```go
var StateStopped = MachineState{"Stopped"}
```

<a name="StateStopping"></a>

```go
var StateStopping = MachineState{"Stopping"}
```

<a name="StateTerminated"></a>

```go
var StateTerminated = MachineState{"Terminated"}
```

<a name="StateTerminating"></a>

```go
var StateTerminating = MachineState{"Terminating"}
```

<a name="StateUnknown"></a>

```go
var StateUnknown = MachineState{"Unknown"}
```

<a name="Environment"></a>
## type [Environment](<https://github.com/analog-substance/carbon/blob/main/pkg/types/environment.go#L9-L21>)



```go
type Environment interface {
    Name() string
    VMs() []VM
    Profile() Profile
    StartVM(string) error
    StopVM(string) error
    RestartVM(string) error
    ImageBuilds() ([]ImageBuild, error)
    Images() ([]Image, error)
    CreateVM(MachineLaunchOptions) error
    DestroyVM(string) error
    DestroyImage(string) error
}
```

<a name="Image"></a>
## type [Image](<https://github.com/analog-substance/carbon/blob/main/pkg/types/image.go#L7-L16>)



```go
type Image interface {
    ID() string
    Name() string
    CreatedAt() string
    Environment() Environment
    Profile() Profile
    Provider() Provider
    Launch(imageLaunchOptions ImageLaunchOptions) error
    Destroy() error
}
```

<a name="ImageBuild"></a>
## type [ImageBuild](<https://github.com/analog-substance/carbon/blob/main/pkg/types/image_build.go#L3-L8>)



```go
type ImageBuild interface {
    Name() string
    ProviderType() string
    Provisioner() string
    Build() error
}
```

<a name="ImageLaunchOptions"></a>
## type [ImageLaunchOptions](<https://github.com/analog-substance/carbon/blob/main/pkg/types/image.go#L3-L5>)



```go
type ImageLaunchOptions struct {
    Name string
}
```

<a name="MachineLaunchOptions"></a>
## type [MachineLaunchOptions](<https://github.com/analog-substance/carbon/blob/main/pkg/types/environment.go#L3-L7>)



```go
type MachineLaunchOptions struct {
    CloudInitTpl string `json:"cloud-init"`
    Image        Image  `json:"image"`
    Name         string `json:"name"`
}
```

<a name="MachineState"></a>
## type [MachineState](<https://github.com/analog-substance/carbon/blob/main/pkg/types/vm.go#L8-L10>)



```go
type MachineState struct {
    Name string `json:"name"`
}
```

<a name="Profile"></a>
## type [Profile](<https://github.com/analog-substance/carbon/blob/main/pkg/types/profile.go#L5-L12>)



```go
type Profile interface {
    Environments() []Environment
    Name() string
    Provider() Provider
    SetConfig(config common.ProfileConfig)
    GetConfig() common.ProfileConfig
    ShouldIncludeEnvironment(envName string) bool
}
```

<a name="Project"></a>
## type [Project](<https://github.com/analog-substance/carbon/blob/main/pkg/types/project.go#L3-L7>)



```go
type Project interface {
    Name() string
    TerraformApply() error
    AddMachine(machine *ProjectMachine, noApply bool) error
}
```

<a name="ProjectConfig"></a>
## type [ProjectConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/types/vm.go#L64-L66>)



```go
type ProjectConfig struct {
    Machines []*ProjectMachine `yaml:"machines"`
}
```

<a name="ProjectMachine"></a>
## type [ProjectMachine](<https://github.com/analog-substance/carbon/blob/main/pkg/types/vm.go#L54-L62>)



```go
type ProjectMachine struct {
    Name       string `yaml:"name"`
    Image      string `yaml:"image,omitempty"`
    Type       string `yaml:"type,omitempty"`
    Profile    string `yaml:"profile,omitempty"`
    Purpose    string `yaml:"purpose,omitempty"`
    VolumeSize int    `yaml:"volume_size,omitempty"`
    Provider   string `yaml:"provider,omitempty"`
}
```

<a name="Provider"></a>
## type [Provider](<https://github.com/analog-substance/carbon/blob/main/pkg/types/provider.go#L5-L14>)



```go
type Provider interface {
    Profiles() []Profile
    Name() string
    Type() string
    IsAvailable() bool
    SetConfig(config common.ProviderConfig)
    GetConfig() common.ProviderConfig
    NewImageBuild(name string, tplDir string) (ImageBuild, error)
    NewProject(name string, force bool) (Project, error)
}
```

<a name="VM"></a>
## type [VM](<https://github.com/analog-substance/carbon/blob/main/pkg/types/vm.go#L22-L52>)

VM interface provides access to useful information and actions related to Virtual Machines

```go
type VM interface {
    // Name returns the name of a virtual machine
    Name() string

    // ID returns the ID of the virtual machine
    ID() string

    // IPAddress returns the public IP address of the virtual machine
    IPAddress() string

    // PrivateIPAddress of the virtual machine
    PrivateIPAddress() string
    UpTime() time.Duration
    State() string
    Type() string

    Environment() Environment
    Profile() Profile
    Provider() Provider

    Destroy() error
    Start() error
    Stop() error
    Restart() error

    ExecSSH(string, bool, ...string) error
    StartVNC(user string, privateIP bool, killVNC bool) error
    StartRDPClient(user string, privateIP bool) error
    Cmd(string, bool, ...string) (string, error)
    NewSSHSession(string, bool) (*ssh_util.Session, error)
}
```

