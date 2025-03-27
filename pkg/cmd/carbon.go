package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/carbon"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"path/filepath"
)

const cfgFileName = "carbon.yaml"

var cfgFile string

// var providers []string
// var profiles []string
// var environments []string
var jsonOutput bool
var debug bool
var carbonObj *carbon.Carbon
var log *slog.Logger

// CarbonCmd represents the base command when called without any subcommands
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
- DigitalOcean
- GCloud
- Multipass (Local)
- QEMU (Local)
- VirtualBox (Local)
- vSphere

There are plans to bring support to the following:

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

		carbonObj = carbon.New(cmd.Root().Version, carbonConfigFile.Carbon)
		updateConfigHelp()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := CarbonCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log = common.WithGroup("cmd")

	CarbonCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/carbon.yaml)")
	// need to rethink how I want this to work
	//CarbonCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	//CarbonCmd.PersistentFlags().StringSliceVarP(&profiles, "profile", "p", []string{}, "Profile to use. Like an instance of a provider. Used to specify aws profiles")
	//CarbonCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some providers/profiles support many environments.")
	CarbonCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")
	CarbonCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debug mode")
}

func addServiceProviderFlag(c *cobra.Command) {
	c.Flags().StringP("service", "S", "", "Service provider (aws, virtualbox, qemu, multipass)")

	err := c.RegisterFlagCompletionFunc("service", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getServiceProviders()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getServiceProviders() []string {
	var serviceProviderNames []string
	serviceProviders := carbonObj.Providers()
	for _, provider := range serviceProviders {
		serviceProviderNames = append(serviceProviderNames, provider.Type())
	}
	return serviceProviderNames
}
