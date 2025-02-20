---
title: api
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/providers/qemu/api"
```

## Index

- [type Domain](<#Domain>)
  - [func \(d \*Domain\) Destroy\(\) error](<#Domain.Destroy>)
  - [func \(d \*Domain\) Reboot\(\) error](<#Domain.Reboot>)
  - [func \(d \*Domain\) Start\(\) error](<#Domain.Start>)
  - [func \(d \*Domain\) Suspend\(\) error](<#Domain.Suspend>)
- [type QEMU](<#QEMU>)
  - [func Connect\(libVirtURL string\) \(\*QEMU, error\)](<#Connect>)
  - [func \(q \*QEMU\) AllNetworks\(\) \(\[\]libvirt.Network, error\)](<#QEMU.AllNetworks>)
  - [func \(q \*QEMU\) Close\(\) error](<#QEMU.Close>)
  - [func \(q \*QEMU\) CreateDomain\(name string, storageVol \*StorageVolume\) \(\*Domain, error\)](<#QEMU.CreateDomain>)
  - [func \(q \*QEMU\) GetDomain\(id string\) \(\*Domain, error\)](<#QEMU.GetDomain>)
  - [func \(q \*QEMU\) GetDomains\(\) \(\[\]\*Domain, error\)](<#QEMU.GetDomains>)
  - [func \(q \*QEMU\) GetStoragePool\(name string\) \(\*StoragePool, error\)](<#QEMU.GetStoragePool>)
  - [func \(q \*QEMU\) GetStoragePools\(\) \(\[\]\*StoragePool, error\)](<#QEMU.GetStoragePools>)
- [type StoragePool](<#StoragePool>)
  - [func \(s \*StoragePool\) GetVolumes\(\) \(\[\]\*StorageVolume, error\)](<#StoragePool.GetVolumes>)
  - [func \(s \*StoragePool\) ImportImage\(name string, imageFile string\) \(\*StorageVolume, error\)](<#StoragePool.ImportImage>)
- [type StorageVolume](<#StorageVolume>)


<a name="Domain"></a>
## type [Domain](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/domain.go#L8-L19>)



```go
type Domain struct {
    LVDomain *libvirt.Domain

    LVDomainState *libvirt.DomainState
    ID            string
    Name          string

    PublicIPAddresses  []string
    PrivateIPAddresses []string
    CurrentUpTime      time.Duration
    // contains filtered or unexported fields
}
```

<a name="Domain.Destroy"></a>
### func \(\*Domain\) [Destroy](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/domain.go#L37>)

```go
func (d *Domain) Destroy() error
```



<a name="Domain.Reboot"></a>
### func \(\*Domain\) [Reboot](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/domain.go#L29>)

```go
func (d *Domain) Reboot() error
```



<a name="Domain.Start"></a>
### func \(\*Domain\) [Start](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/domain.go#L21>)

```go
func (d *Domain) Start() error
```



<a name="Domain.Suspend"></a>
### func \(\*Domain\) [Suspend](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/domain.go#L33>)

```go
func (d *Domain) Suspend() error
```



<a name="QEMU"></a>
## type [QEMU](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L12-L19>)



```go
type QEMU struct {
    // contains filtered or unexported fields
}
```

<a name="Connect"></a>
### func [Connect](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L21>)

```go
func Connect(libVirtURL string) (*QEMU, error)
```



<a name="QEMU.AllNetworks"></a>
### func \(\*QEMU\) [AllNetworks](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L50>)

```go
func (q *QEMU) AllNetworks() ([]libvirt.Network, error)
```



<a name="QEMU.Close"></a>
### func \(\*QEMU\) [Close](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L42>)

```go
func (q *QEMU) Close() error
```



<a name="QEMU.CreateDomain"></a>
### func \(\*QEMU\) [CreateDomain](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L255>)

```go
func (q *QEMU) CreateDomain(name string, storageVol *StorageVolume) (*Domain, error)
```



<a name="QEMU.GetDomain"></a>
### func \(\*QEMU\) [GetDomain](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L124>)

```go
func (q *QEMU) GetDomain(id string) (*Domain, error)
```



<a name="QEMU.GetDomains"></a>
### func \(\*QEMU\) [GetDomains](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L65>)

```go
func (q *QEMU) GetDomains() ([]*Domain, error)
```



<a name="QEMU.GetStoragePool"></a>
### func \(\*QEMU\) [GetStoragePool](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L159>)

```go
func (q *QEMU) GetStoragePool(name string) (*StoragePool, error)
```



<a name="QEMU.GetStoragePools"></a>
### func \(\*QEMU\) [GetStoragePools](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/qemu.go#L140>)

```go
func (q *QEMU) GetStoragePools() ([]*StoragePool, error)
```



<a name="StoragePool"></a>
## type [StoragePool](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/storage_pool.go#L14-L18>)



```go
type StoragePool struct {
    LVStoragePool *libvirt.StoragePool
    Volumes       []*StorageVolume
    // contains filtered or unexported fields
}
```

<a name="StoragePool.GetVolumes"></a>
### func \(\*StoragePool\) [GetVolumes](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/storage_pool.go#L20>)

```go
func (s *StoragePool) GetVolumes() ([]*StorageVolume, error)
```



<a name="StoragePool.ImportImage"></a>
### func \(\*StoragePool\) [ImportImage](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/storage_pool.go#L39>)

```go
func (s *StoragePool) ImportImage(name string, imageFile string) (*StorageVolume, error)
```



<a name="StorageVolume"></a>
## type [StorageVolume](<https://github.com/analog-substance/carbon/blob/main/pkg/providers/qemu/api/storage_volume.go#L5-L8>)



```go
type StorageVolume struct {
    LVStorageVolume *libvirt.StorageVol
    // contains filtered or unexported fields
}
```

