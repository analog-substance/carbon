---
title: api
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/multipass/api"
```

## Index

- [func AppPath\(\) string](<#AppPath>)
- [func RestartVM\(id string\) error](<#RestartVM>)
- [func SleepVM\(id string\) error](<#SleepVM>)
- [func StartVM\(id string\) error](<#StartVM>)
- [type MultipassListOutput](<#MultipassListOutput>)
- [type MultipassVM](<#MultipassVM>)
  - [func ListVMs\(\) \[\]MultipassVM](<#ListVMs>)


<a name="AppPath"></a>
## func [AppPath](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L23>)

```go
func AppPath() string
```



<a name="RestartVM"></a>
## func [RestartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L60>)

```go
func RestartVM(id string) error
```



<a name="SleepVM"></a>
## func [SleepVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L67>)

```go
func SleepVM(id string) error
```



<a name="StartVM"></a>
## func [StartVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L53>)

```go
func StartVM(id string) error
```



<a name="MultipassListOutput"></a>
## type [MultipassListOutput](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L17-L19>)



```go
type MultipassListOutput struct {
    List []MultipassVM `json:"list"`
}
```

<a name="MultipassVM"></a>
## type [MultipassVM](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L10-L15>)



```go
type MultipassVM struct {
    Ipv4    []string `json:"ipv4"`
    Name    string   `json:"name"`
    Release string   `json:"release"`
    State   string   `json:"state"`
}
```

<a name="ListVMs"></a>
### func [ListVMs](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/multipass/api/main.go#L36>)

```go
func ListVMs() []MultipassVM
```



