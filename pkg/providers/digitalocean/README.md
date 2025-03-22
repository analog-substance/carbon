---
title: DigitalOcean WIP
description: DigitalOcean Provider for Carbon
weight: 20
---

**This is a WIP**

DigitalOcean Projects are mapped to Carbon Profiles, they can only have a single environment.

Current functionality

- Autodetect DigitalOcean config file or use env var `DIGITALOCEAN_ACCESS_TOKEN`
- List Droplets as machines.

While not a lot. All functionality on machines (SSH, VNC, RDP) should work assuming the droplet supports them.