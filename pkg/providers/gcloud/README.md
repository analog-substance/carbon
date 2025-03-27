---
title: GCloud Provider WIP
description: GCP Compute provider
weight: 20
---

**This is a WIP**

### Map

| GCP            | Carbon       | Description |
|----------------|--------------|-------------|
| Account        | Profile      |             |
| Project & Zone | Environment  |             |
|                |              |             |

### Configuration

use project/zone for the environment name.
```yaml
carbon:
    providers:
        gcloud:
            profiles:
                default:
                    environments:
                      gcp-project/us-east3-c
```