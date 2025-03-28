---
title: Vm Stop
description: Stop VM(s)
---

## Synopsis

Stop VM(s).

By default, carbon will attempt to hibernate the machine rather than power it off.


```
carbon vm stop [flags]
```

## Examples

```bash
carbon vm stop -n vm-name
```

## Options

```
  -h, --help   help for stop
```

## Options inherited from parent commands

```
      --config string   config file (default is $HOME/carbon.yaml)
      --debug           Debug mode
      --host strings    Hostname or IP Address.
  -i, --id string       ID of machine to start.
  -j, --json            Output in JSON
  -n, --name string     Name of the VM.
  -u, --user string     SSH Username. (default "ubuntu")
```

## SEE ALSO

* [carbon vm](carbon_vm.md)	 - Manage and interact with VMs.

###### Auto generated by spf13/cobra on 25-Mar-2025
