---
title: Vm
description: Manage and interact with VMs
---

## carbon vm

Manage and interact with VMs

### Synopsis

Manage and interact with VMs.

Management functions include

- Starting VMs
- Stopping VMs
- Rebooting VMs
- Launching VMs from images
- Destroying VMs from images



### Options

```
  -h, --help           help for vm
      --host strings   Hostname or IP Address.
  -i, --id string      ID of machine to start.
  -n, --name string    Name of the VM.
  -u, --user string    SSH Username. (default "ubuntu")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.carbon.yaml)
      --debug           Debug mode
  -j, --json            Output in JSON
```

### SEE ALSO

* [carbon](carbon.md)	 - Infrastructure Ops simplified
* [carbon vm destroy](carbon_vm_destroy.md)	 - destroy VM
* [carbon vm launch](carbon_vm_launch.md)	 - launch a new vm from an image
* [carbon vm list](carbon_vm_list.md)	 - List VMs
* [carbon vm restart](carbon_vm_restart.md)	 - Restart VM(s)
* [carbon vm ssh](carbon_vm_ssh.md)	 - SSH to a VM
* [carbon vm start](carbon_vm_start.md)	 - Start VMs
* [carbon vm stop](carbon_vm_stop.md)	 - Stop VM(s)
* [carbon vm vnc](carbon_vm_vnc.md)	 - VNC to a VM

###### Auto generated by spf13/cobra on 9-Oct-2024