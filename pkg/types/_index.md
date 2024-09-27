---
title: Types
description: Carbon has several types defined for providers to use
weight: 10
---

Carbon organizes your machines in a tree like:

- Providers
  - Profiles
    - Environments

An example could look like this:

- Provider: AWS
  - Profile: default aws profile
    - Environment: VPC 001
    - Environment: VPC 002
  - Profile: red team aws profile
      - Environment: VPC 001
      - Environment: VPC 002
- Provider: VirtualBox
    - Profile: local
        - Environment: local
- Provider: Multipass
    - Profile: local
        - Environment: local
- Provider: vSphere
    - Profile: whatever.vsphere.local
        - Environment: Datacenter 01
        - Environment: Datacenter 02