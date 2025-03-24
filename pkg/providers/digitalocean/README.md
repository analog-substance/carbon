---
title: DigitalOcean WIP
description: DigitalOcean Provider for Carbon
weight: 20
---

**This is a WIP**

### Map

| DigitalOcean | Carbon       | Description                                                                                                 |
|--------------|--------------|-------------------------------------------------------------------------------------------------------------|
| Account      | Profile      | Default token is pulled from the `DIGITALOCEAN_TOKEN` environment variable, or from the `doctl` config file |
| Project      | Environment  |                                                                                                             |
|              |              |                                                                                                             |



- List Droplets as machines.

While not a lot. All functionality on machines (SSH, VNC, RDP) should work assuming the droplet supports them.