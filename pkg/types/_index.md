---
title: Types
description: Carbon has several types defined for providers to use
weight: 10
---

Carbon organizes your machines in a tree like:

- Providers
  - Platforms
    - Environments

An example could look like this:

- Provider: AWS
  - Platform: default aws profile
    - Environment: VPC 001
    - Environment: VPC 002
  - Platform: red team aws profile
      - Environment: VPC 001
      - Environment: VPC 002
- Provider: VirtualBox
    - Platform: local
        - Environment: local
- Provider: Multipass
    - Platform: local
        - Environment: local
- Provider: vSphere
    - Platform: whatever.vsphere.local
        - Environment: Datacenter 01
        - Environment: Datacenter 02