---
title: rdp_client
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/rdp_client"
```

## Index

- [func Start\(options Options\) error](<#Start>)
- [func StartRDPClient\(options Options\) error](<#StartRDPClient>)
- [type Options](<#Options>)


<a name="Start"></a>
## func [Start](<https://github.com/analog-substance/carbon/blob/main/pkg/rdp_client/options.go#L13>)

```go
func Start(options Options) error
```



<a name="StartRDPClient"></a>
## func [StartRDPClient](<https://github.com/analog-substance/carbon/blob/main/pkg/rdp_client/linux.go#L11>)

```go
func StartRDPClient(options Options) error
```



<a name="Options"></a>
## type [Options](<https://github.com/analog-substance/carbon/blob/main/pkg/rdp_client/options.go#L7-L11>)



```go
type Options struct {
    Delay int
    User  string
    Host  string
}
```

