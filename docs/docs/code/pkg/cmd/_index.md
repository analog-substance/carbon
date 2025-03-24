---
title: cmd
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/cmd"
```

## Index

- [Variables](<#variables>)
- [func AskIfSure\(msg string\) bool](<#AskIfSure>)
- [func Execute\(\)](<#Execute>)
- [func ListingDir\(dir string\)](<#ListingDir>)


## Variables

<a name="CarbonCmd"></a>CarbonCmd represents the base command when called without any subcommands

```go
var CarbonCmd = &cobra.Command{
    Use:   "carbon",
    Short: "Carbon - Infrastructure automation for offensive operations.",
    Long: `Infrastructure automation for offensive operations.
- ℹ️ Checkout the latest docs [here](https://analog-substance.github.io/carbon/)
- 😢 Have a problem? [Create an Issue](https://github.com/analog-substance/carbon/issues/new?title=Something%20is%20broken)
- ❤️ Enjoying Carbon? [Star the Repo](https://github.com/analog-substance/carbon)

## Purpose

Carbon's primary purpose is to provide a consistent execution environment to
facilitate offensive security assessments.

## Dependencies

- Packer to build images.
- Terraform to provision infrastructure.
- Golang project structure.

## Supported Providers

- AWS
- QEMU (Local)
- VirtualBox (Local)
- vSphere (in progress)
- Multipass (Local)

There are plans to bring support to the following:

- GCP
- Azure
- VMware (Local)
- QEMU (Remote)
`,
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

        if debug {
            common.LogLevel(slog.LevelDebug)
        }
        log.Debug("debug mode", "debug", debug)

        carbonConfigFile := common.GetConfig()
        home, err := homedir.Dir()
        if err != nil {
            log.Debug("error getting home directory", "error", err)
        } else {
            err := carbonConfigFile.MergeInConfigFile(filepath.Join(home, cfgFileName))

            if err != nil {
                log.Debug("error loading carbon config from home", "error", err)
            }
        }

        err = carbonConfigFile.MergeInConfigFile(cfgFileName)
        if err != nil {
            log.Debug("error loading carbon config from home", "error", err)
        }

        carbonObj = carbon.New(carbonConfigFile.Carbon)
        updateConfigHelp()

        return nil
    },
}
```

<a name="AskIfSure"></a>
## func [AskIfSure](<https://github.com/analog-substance/carbon/blob/main/pkg/cmd/vm.go#L63>)

```go
func AskIfSure(msg string) bool
```



<a name="Execute"></a>
## func [Execute](<https://github.com/analog-substance/carbon/blob/main/pkg/cmd/carbon.go#L94>)

```go
func Execute()
```

Execute adds all child commands to the root command and sets flags appropriately. This is called by main.main\(\). It only needs to happen once to the rootCmd.

<a name="ListingDir"></a>
## func [ListingDir](<https://github.com/analog-substance/carbon/blob/main/pkg/cmd/dev_fs.go#L25>)

```go
func ListingDir(dir string)
```



