
# Usage


## VirtualBox

### Building with VirtualBox

```bash
packer build --only virtualbox-iso.carbon-vm-ubuntu .
```

### Testing with VirtualBox

A simple script has been provided to create a new VirtualBox machine with the newly created disk image.

```bash
scripts/test-carbon-ubuntu.sh
```

## vSphere

You'll need to copy the `private.auto.pkrvars.hcl.example` to `private.auto.pkrvars.hcl` and populate it with your configuration values.


```bash
packer build --only vsphere-iso.carbon-vm-ubuntu .
```

When finished, you should have a VM powered off in your  server.