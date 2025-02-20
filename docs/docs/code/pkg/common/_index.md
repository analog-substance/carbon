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
- [func ImagesDir\(\) string](<#ImagesDir>)
- [func LogLevel\(level slog.Level\)](<#LogLevel>)
- [func Logger\(\) \*slog.Logger](<#Logger>)
- [func PackerDir\(\) string](<#PackerDir>)
- [func ProjectsDir\(\) string](<#ProjectsDir>)
- [func WithGroup\(groupName string\) \*slog.Logger](<#WithGroup>)
- [type CarbonConfig](<#CarbonConfig>)
- [type CarbonConfigFile](<#CarbonConfigFile>)
- [type ProfileConfig](<#ProfileConfig>)
- [type ProviderConfig](<#ProviderConfig>)


## Constants

<a name="DefaultDeploymentsDirName"></a>

```go
const DefaultDeploymentsDirName = "deployments"
```

<a name="DefaultImagesDirName"></a>

```go
const DefaultImagesDirName = "images"
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

<a name="ViperDefaultInstanceDir"></a>

```go
const ViperDefaultInstanceDir = ViperPrefix + "default.dir"
```

<a name="ViperDeploymentsDir"></a>

```go
const ViperDeploymentsDir = ViperPrefix + "deployments.dir"
```

<a name="ViperImagesDir"></a>

```go
const ViperImagesDir = ViperPrefix + "images.dir"
```

<a name="ViperPackerDir"></a>

```go
const ViperPackerDir = ViperPrefix + "packer.dir"
```

<a name="ViperPrefix"></a>

```go
const ViperPrefix = "carbon."
```

<a name="ViperTerraformDir"></a>

```go
const ViperTerraformDir = ViperPrefix + "terraform.dir"
```

<a name="ViperTerraformProjectDir"></a>

```go
const ViperTerraformProjectDir = ViperPrefix + "projects.dir"
```

<a name="ImagesDir"></a>
## func [ImagesDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L31>)

```go
func ImagesDir() string
```



<a name="LogLevel"></a>
## func [LogLevel](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L23>)

```go
func LogLevel(level slog.Level)
```



<a name="Logger"></a>
## func [Logger](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L15>)

```go
func Logger() *slog.Logger
```



<a name="PackerDir"></a>
## func [PackerDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L27>)

```go
func PackerDir() string
```



<a name="ProjectsDir"></a>
## func [ProjectsDir](<https://github.com/analog-substance/carbon/blob/main/pkg/common/helpers.go#L35>)

```go
func ProjectsDir() string
```



<a name="WithGroup"></a>
## func [WithGroup](<https://github.com/analog-substance/carbon/blob/main/pkg/common/logging.go#L19>)

```go
func WithGroup(groupName string) *slog.Logger
```



<a name="CarbonConfig"></a>
## type [CarbonConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L17-L19>)



```go
type CarbonConfig struct {
    Providers map[string]ProviderConfig `yaml:"providers"`
}
```

<a name="CarbonConfigFile"></a>
## type [CarbonConfigFile](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L20-L22>)



```go
type CarbonConfigFile struct {
    Carbon CarbonConfig `yaml:"carbon"`
}
```

<a name="ProfileConfig"></a>
## type [ProfileConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L3-L9>)



```go
type ProfileConfig struct {
    Enabled      bool            `yaml:"enabled"`
    Username     string          `yaml:"username"`
    Password     string          `yaml:"password"`
    URL          string          `yaml:"url"`
    Environments map[string]bool `yaml:"environments"`
}
```

<a name="ProviderConfig"></a>
## type [ProviderConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/common/config.go#L11-L15>)



```go
type ProviderConfig struct {
    Enabled      bool                     `yaml:"enabled"`
    AutoDiscover bool                     `yaml:"auto_discover" mapstructure:"auto_discover"`
    Profiles     map[string]ProfileConfig `yaml:"profiles"`
}
```

