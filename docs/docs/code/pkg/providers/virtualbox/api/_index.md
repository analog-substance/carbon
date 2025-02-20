---
title: api
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/virtualbox/api"
```

## Index

- [func AppPath\(\) string](<#AppPath>)
- [func RestartVM\(id string\) error](<#RestartVM>)
- [func SleepVM\(id string\) error](<#SleepVM>)
- [func StartVM\(id string\) error](<#StartVM>)
- [type VBoxVM](<#VBoxVM>)
  - [func ListVMs\(\) \[\]VBoxVM](<#ListVMs>)


<a name="AppPath"></a>
## func [AppPath](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L26>)

```go
func AppPath() string
```



<a name="RestartVM"></a>
## func [RestartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L141>)

```go
func RestartVM(id string) error
```



<a name="SleepVM"></a>
## func [SleepVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L148>)

```go
func SleepVM(id string) error
```



<a name="StartVM"></a>
## func [StartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L134>)

```go
func StartVM(id string) error
```



<a name="VBoxVM"></a>
## type [VBoxVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L14-L22>)



```go
type VBoxVM struct {
    Name               string
    ID                 string
    State              string
    GuestOS            string
    UpTime             time.Duration
    PrivateIPAddresses []string
    // contains filtered or unexported fields
}
```

<a name="ListVMs"></a>
### func [ListVMs](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/virtualbox/api/main.go#L107>)

```go
func ListVMs() []VBoxVM
```



