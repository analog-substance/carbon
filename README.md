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

| Feature                   | AWS | QEMU | VirtualBox | DigitalOcean | vSphere | Multipass | GCP | Azure |
|---------------------------|-----|------|------------|--------------|---------|-----------|-----|-------|
| VM: List                  | ✅   | ✅    | ✅          | ✅            | ✅       | ✅         | ✅   | ❌     |
| VM: Start / Stop / Reboot | ✅   | ✅    | ✅          | ✅            | ✅       | ✅         | ✅   | ❌     |
| VM: SSH / RDP / VNC       | ✅   | ✅    | ✅          | ✅            | ✅       | ✅         | ✅   | ❌     |
| VM: Create / Destroy      | ✅   | ✅    | ✅          | ✅            | ✅       | ✅         | ❌   | ❌     |
| Image Builds              | ✅   | ✅    | ✅          | ❌            | ✅       | ❌ N/A     | ❌   | ❌     |
| Images (Build, Destroy)   | ✅   | ✅    | ✅          | ❌            | ✅       | ❌ N/A     | ❌   | ❌     |
| Infrastructure Creation   | ❌   | ❌    | ❌          | ❌            | ❌       | ❌ N/A     | ❌   | ❌     |

❌ = Not Right Now, but planned  
✅ = Supported


## Install
***
You can download a prebuilt release from our [GitHub Releases](https://github.com/analog-substance/carbon/releases) page.
Or use `go install`.

```sh
go install github.com/analog-substance/carbon@latest
```
Be sure to check out the [providers](pkg/providers) section for additional information on configuring your provider.
## Requirements

Carbon expects the following to be installed and accessible in your `$PATH`.

- Packer
- Terraform
- SSH Client
- vncviewer (TigerVNC)

## Usage
***

```
Manage and use infrastructure with a consistent interface, regardless of where it lives.

## Usage
                                                                                                                                                                                                                                                                                                          
carbon [command]
                                                                                                                                                                                                                                                                                                                  
## Available Commands:
                                                                                                                                                                                                                                                                                                                  
  completion  Generate completion script
  config      View and manage configuration values.
  help        Help about any command
  image       View or manage images and image builds.
  project     Manage and interact with projects
  update      Update carbon to latest version
  vm          Manage and interact with VMs.
                                                                                                                                                                                                                                                                       
## Flags
                                                                                                                                                                                                                                                                                                          
      --config string   config file (default is $HOME/carbon.yaml)
      --debug           Debug mode
  -h, --help            help for carbon
  -j, --json            Output in JSON
  -v, --version         version for carbon
  
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
- provision aws env (create files, call terraform)
- Self Test to ensure dependencies are met
- Cloud init from templates (Base, Operator, Operator Desktop, Implant VM)
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
