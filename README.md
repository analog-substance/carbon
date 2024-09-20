# Carbon
> Infrastructure automation and configuration management
___

## Purpose

Carbon's primary purpose is to provide a consistent execution environment to
facilitate offensive security assessments.

It uses:
- Packer to build images.
- Terraform to provision infrastructure.
- Golang project structure.

It is aimed at supporting various services to ensure it can scale with you as
your operations become more complex.

| Feature                 | AWS ✅ | Azure ❌ | GCP ❌ | QEMU ✅ | Multipass ✅ | VirtualBox ✅ | vSphere ❌ |
|-------------------------|-------|---------|-------|--------|-------------|--------------|-----------|
| Image build templates   | ✅     | ❌       | ❌     | ✅      | ❌ N/A       | ✅            | ✅         |
| Create Images           | ✅     | ❌       | ❌     | ✅      | ❌           | ✅            | ✅         |
| Infrastructure Creation | ❌     | ❌       | ❌     | ❌      | ❌ N/A       | ❌            | ❌         |
| VM Management           | ✅     | ❌       | ❌     | ✅      | ✅           | ✅            | ❌         |

❌ = Not Right Now, but soon  
✅ = Supported

_____

## Install

```bash
go install github.com/analog-substance/carbon@latest
```

_____
## Usage
_____

### Images

#### Bootstrap Image Build Configuration

```bash
carbon image bootstrap -n my-image -t ubuntu-24.04 -s aws
```

#### Build Images


```bash
carbon image build -n my-image
```

#### Manage Images and Image Builds

List image build configs.

```bash
carbon image list -b
```
_____

### Infrastructure

#### Create New Infrastructure

#### Modify Infrastructure

#### Teardown Infrastructure
_____

### Operating

#### Starting and Stopping 

```bash
carbon vm start -i i-afde123ae43
carbon vm stop -i i-afde123ae43
```
#### Connecting to VMs

```bash
carbon vm ssh -i i-afde123ae43
```

***

Things to do

- provision aws env (create files, call terraform)
- create a new vm on infrastructure
- vnc to vm
- point a domain
- list domains
- Jobs / Distributed execution
- Cloud init templates (Base, Operator, Operator Desktop, Implant VM)
- Simple deploy/config of services (Pwndoc, Gophish, modlishka, Guacamole, Sliver, Mythic)
- vSphere provider
- GCP Provider
- Azure Provider
- LXD Provider
- Different OS (CentOS, Arch)
- Self Test to ensure dependencies are met
- Slack Bot
- Discord Bot
- Web GUI
- DNS management
- docsy config
- vhs example gifs
- docs
- tests
