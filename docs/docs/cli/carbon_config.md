---
title: Config
description: Get/Set config information
---

## carbon config

Get/Set config information

### Synopsis

Get/Set config information.

Set vSphere credentials

	carbon config carbon.credentials.vsphere_server.provider vsphere
	carbon config carbon.credentials.vsphere_server.username vsphere_user@vsphere.example
	carbon config carbon.credentials.vsphere_server.password_command 'op read op://Private/vSphere Creds/password'



```
carbon config [flags]
```

### Options

```
  -h, --help           help for config
  -r, --remove-reset   remove key from the config or reset to default
  -s, --save           save the current configuration
  -k, --sub-keys       display only the sub-keys
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.carbon.yaml)
      --debug           Debug mode
  -j, --json            Output in JSON
```

### SEE ALSO

* [carbon](carbon.md)	 - Infrastructure Ops simplified

###### Auto generated by spf13/cobra on 9-Oct-2024