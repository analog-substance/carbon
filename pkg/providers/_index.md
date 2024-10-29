---
title: Providers
description: Providers help carbon work with various services.
weight: 20
---

Providers allow Carbon to interact with external applications to retrieve
information about your operating environment.

By default, providers will automatically discover configuration profiles to use. You can disable the auto discovery and force enabled specific profiles.

```yaml
carbon:
  providers:
    aws:
      auto_discover: false
      profiles:
        default:
          enabled: true
```