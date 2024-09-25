---
title: Carbon
linkTitle: Docs
menu: {main: {weight: 20}}
---
> Infrastructure automation for offensive operations.
> https://analog-substance.github.io/carbon/

## Purpose
***

Carbon's primary purpose is to provide a consistent execution environment to
facilitate offensive security assessments.

It uses:
- Packer to build images.
- Terraform to provision infrastructure.
- Golang project structure.

## Features
***

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


## Install
***
You can download a prebuilt release from our [GitHub Releases](https://github.com/analog-substance/carbon/releases) page.
Or use `go install`.

```sh
go install github.com/analog-substance/carbon@latest
```

Carbon expects Packer, Terraform, and an SSH client to be installed and accessible in your `$PATH`.

## Usage
***

```
Manage and use infrastructure with a consistent interface, regardless of where it lives.

Usage:
  carbon [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Display config information
  dev         Unstable sub-commands for testing random ideas
  help        Help about any command
  image       manage images and image builds
  vm          Manage and interact with VMs

Flags:
      --config string         config file (default is $HOME/.carbon.yaml)
  -e, --environment strings   Environment to use. Some platforms support many environments.
  -h, --help                  help for carbon
  -j, --json                  Output in JSON
  -p, --platform strings      Platform to use. Like an instance of a provider. Used to specify aws profiles
  -P, --provider strings      Provider to use vbox, aws
  -v, --version               version for carbon

Use "carbon [command] --help" for more information about a command.

```

### Images
***

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

### Infrastructure
***

#### Create New Infrastructure

#### Modify Infrastructure

#### Teardown Infrastructure

### Operating
***

#### Starting
```bash
carbon vm start -i i-afde123ae43
```

#### Stopping
```bash
carbon vm stop -i i-afde123ae43
```
#### Connecting to VMs

```bash
carbon vm ssh -i i-afde123ae43
```

## Things to do
***

- docs
  - vhs example gifs
- create/destroy one off VMs
- vm search argument
- provision aws env (create files, call terraform)
- Self Test to ensure dependencies are met
- Cloud init from templates (Base, Operator, Operator Desktop, Implant VM)
- vSphere provider
- vnc to vm
- DNS management
    - point a domain
    - list domains
- Jobs / Distributed execution
- Simple deploy/config of services (Pwndoc, Gophish, modlishka, Guacamole, Sliver, Mythic)
- GCP Provider
- Azure Provider
- LXD Provider
- Different OS (CentOS, Arch)
- Chat Bots
    - Slack Bot
    - Discord Bot
- Web GUI
- tests (lol, this should not be last)
