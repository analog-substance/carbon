---
title: cloud_init
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/cloud_init"
```

## Index

- [type AptSource](<#AptSource>)
- [type CloudConfig](<#CloudConfig>)
  - [func \(c \*CloudConfig\) MergeWith\(otherConfig \*CloudConfig\)](<#CloudConfig.MergeWith>)
- [type WriteFile](<#WriteFile>)


<a name="AptSource"></a>
## type [AptSource](<https://github.com/analog-substance/carbon/blob/main/pkg/cloud_init/main.go#L3-L6>)



```go
type AptSource struct {
    Source string `yaml:"source"`
    Keyid  string `yaml:"keyid"`
}
```

<a name="CloudConfig"></a>
## type [CloudConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/cloud_init/main.go#L16-L27>)



```go
type CloudConfig struct {
    Timezone          string   `yaml:"timezone"`
    SSHDeletekeys     bool     `yaml:"ssh_deletekeys"`
    SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys"`
    Apt               struct {
        Sources map[string]AptSource `yaml:"sources"`
    }   `yaml:"apt"`
    WriteFiles     []WriteFile `yaml:"write_files"`
    PackageUpgrade bool        `yaml:"package_upgrade"`
    Packages       []string    `yaml:"packages"`
    Runcmd         [][]string  `yaml:"runcmd"`
}
```

<a name="CloudConfig.MergeWith"></a>
### func \(\*CloudConfig\) [MergeWith](<https://github.com/analog-substance/carbon/blob/main/pkg/cloud_init/main.go#L29>)

```go
func (c *CloudConfig) MergeWith(otherConfig *CloudConfig)
```



<a name="WriteFile"></a>
## type [WriteFile](<https://github.com/analog-substance/carbon/blob/main/pkg/cloud_init/main.go#L8-L14>)



```go
type WriteFile struct {
    Path        string `yaml:"path"`
    Content     string `yaml:"content"`
    Owner       string `yaml:"owner"`
    Permissions string `yaml:"permissions"`
    Encoding    string `yaml:"encoding,omitempty"`
}
```

