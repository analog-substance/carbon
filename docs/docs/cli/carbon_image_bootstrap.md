---
title: Image Bootstrap
description: create image build configs
---

## carbon image bootstrap

create image build configs

### Synopsis

create image build configs.
Example

	carbon image bootstrap -n operator-desktop-aws -s aws -t ubuntu-desktop



```
carbon image bootstrap [flags]
```

### Options

```
  -h, --help              help for bootstrap
  -n, --name string       Name of image build
  -s, --service string    Service provider (aws, virtualbox, qemu, multipass)
  -t, --template string   Template to use (default "ubuntu-24.04")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.carbon.yaml)
      --debug           Debug mode
  -j, --json            Output in JSON
```

### SEE ALSO

* [carbon image](carbon_image.md)	 - manage images and image builds

###### Auto generated by spf13/cobra on 9-Oct-2024