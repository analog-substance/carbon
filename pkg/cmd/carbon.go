package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/carbon"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

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

		carbonConfigFile := &common.CarbonConfigFile{}
		err := viper.Unmarshal(carbonConfigFile)
		if err != nil {
			log.Debug("failed to unmarshal viper config to carbon config struct")
		}

		carbonObj = carbon.New(carbonConfigFile.Carbon)
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

	cobra.OnInitialize(initConfig)
	CarbonCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/carbon.yaml)")
	// need to rethink how I want this to work
	//CarbonCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	//CarbonCmd.PersistentFlags().StringSliceVarP(&profiles, "profile", "p", []string{}, "Profile to use. Like an instance of a provider. Used to specify aws profiles")
	//CarbonCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some providers/profiles support many environments.")
	CarbonCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")
	CarbonCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault(common.ViperDefaultInstanceDir, common.DefaultInstanceDir)
	viper.SetDefault(common.ViperDeploymentsDir, common.DefaultDeploymentsDirName)
	viper.SetDefault(common.ViperPackerDir, filepath.Join(common.DefaultDeploymentsDirName, common.DefaultPackerDirName))
	viper.SetDefault(common.ViperImagesDir, filepath.Join(common.DefaultDeploymentsDirName, common.DefaultImagesDirName))
	viper.SetDefault(common.ViperTerraformDir, filepath.Join(common.DefaultDeploymentsDirName, common.DefaultTerraformDirName))
	viper.SetDefault(common.ViperTerraformProjectDir, filepath.Join(common.DefaultDeploymentsDirName, common.DefaultProjectsDirName))

	viper.SetDefault("carbon.providers.aws.auto_discover", true)
	viper.SetDefault("carbon.providers.aws.enabled", true)
	viper.SetDefault("carbon.providers.aws.profiles.default.enabled", true)

	viper.SetDefault("carbon.providers.virtualbox.enabled", true)
	viper.SetDefault("carbon.providers.virtualbox.auto_discover", true)
	viper.SetDefault("carbon.providers.virtualbox.profiles.default.enabled", true)

	viper.SetDefault("carbon.providers.qemu.enabled", true)
	viper.SetDefault("carbon.providers.qemu.auto_discover", true)
	viper.SetDefault("carbon.providers.qemu.profiles.default.enabled", true)

	viper.SetDefault("carbon.providers.multipass.enabled", true)
	viper.SetDefault("carbon.providers.multipass.auto_discover", true)
	viper.SetDefault("carbon.providers.multipass.profiles.default.enabled", true)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")

		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName("carbon")
	}

	// do not need to handle err here. since an error will occur if the config file doesn't exist
	_ = viper.MergeInConfig()
	log.Debug("carbon config loaded", "config_file", viper.ConfigFileUsed())
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
