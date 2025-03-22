---
title: vnc_viewer
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/vnc_viewer"
```

## Index

- [func Start\(options Options\) error](<#Start>)
- [func StartViewer\(options Options\) error](<#StartViewer>)
- [type Options](<#Options>)


<a name="Start"></a>
## func [Start](<https://github.com/analog-substance/carbon/blob/main/pkg/vnc_viewer/options.go#L14>)

```go
func Start(options Options) error
```



<a name="StartViewer"></a>
## func [StartViewer](<https://github.com/analog-substance/carbon/blob/main/pkg/vnc_viewer/linux.go#L10>)

```go
func StartViewer(options Options) error
```



<a name="Options"></a>
## type [Options](<https://github.com/analog-substance/carbon/blob/main/pkg/vnc_viewer/options.go#L8-L12>)



```go
type Options struct {
    Delay        int
    PasswordFile string
    Host         string
}
```

