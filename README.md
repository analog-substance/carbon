# Carbon

Infrastructure Automations and Configurations

## Purpose

Carbon's primary purpose is to provide a consistent execution environment to facilitate offensive security assessments.



***

## Ideas

- [ ] Support different linux distributions
- [ ] Minimal Base VM
    - [ ] VirtualBox
    - [ ] ESXi
    - [ ] LXD
    - [ ] ProxMox
    - [ ] QEMU
    - [ ] AWS AMI
    - [ ] GCP
    - [ ] Azure Virtual Image
- [ ] Desktop VM
    - [x] VirtualBox
    - [x] vSphere (ESXi)
    - [ ] Azure Virtual Image
    - [ ] LXD
    - [ ] ProxMox
    - [ ] QEMU
    - [ ] AWS AMI
    - [ ] GCP
- [ ] Guacamole Server
    - [ ] VirtualBox
    - [ ] ESXi
    - [ ] LXD
    - [ ] ProxMox
    - [ ] QEMU
    - [ ] AWS AMI
    - [ ] GCP
    - [ ] Azure Virtual Image
- [ ] Implant VM
    - [ ] VirtualBox
    - [ ] ESXi
    - [ ] LXD
    - [ ] ProxMox
    - [ ] QEMU
    - [ ] AWS AMI
    - [ ] GCP
    - [ ] Azure Virtual Image
- [ ] Terraform
- [ ] Op Templates
- [ ] CLI (probably separate repo)
    - [ ] Self Test Dependencies Met
    - [ ] Self Update
    - [ ] Build VMs
    - [ ] Point Domains
        - [ ] AWS (Route53)
        - [ ] GCP
        - [ ] Azure
    - [ ] Infrastucture Management
    - [ ] Connect to infrastructure
    - [ ] Slack Bot (maybe separate repo)



```
carbon
    vm
        create
        ssh
        start
        stop
        delete
        vnc
        restart
    infra
        create
        destroy
        list
    image
        create
        build
        delete
    dns
        list
        create
        delete
    jobs (maybe not this, probably just use argo or embed the argo cli with params already populated based on the context)
        create
        list
        delete
        run
```

Things to do

- build image (call packer)
- provision aws env (create files, call terraform)
- create a new vm on infrastructure
- ssh to an vm
- start an vm
- stop an vm
- vnc to vm
- create new build config
- point a domain
- list domains

****

What will we use to deploy resources?
- AWS
- QEMU
- VirtualBox
- vSphere

****
