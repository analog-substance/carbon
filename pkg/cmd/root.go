package cmd

import (
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

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "carbon",
	Short: "Infrastructure ops simplified",
	Long:  `Manage and use infrastructure with a consistent interface, regardless of where it lives.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if debug {
			common.LogLevel(slog.LevelDebug)
		}
		log.Debug("debug mode", "debug", debug)

		carbonConfigFile := common.CarbonConfigFile{}
		err := viper.Unmarshal(&carbonConfigFile)
		if err != nil {
			log.Debug("failed to unmarshal viper config to carbon config struct")
		}

		carbonObj = carbon.New(carbonConfigFile.Carbon)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log = common.WithGroup("cmd")

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carbon.yaml)")
	// need to rethink how I want this to work
	//RootCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	//RootCmd.PersistentFlags().StringSliceVarP(&profiles, "profile", "p", []string{}, "Profile to use. Like an instance of a provider. Used to specify aws profiles")
	//RootCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some providers/profiles support many environments.")
	RootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debug mode")
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

	viper.SetDefault("carbon.providers.aws.enabled", true)
	viper.SetDefault("carbon.providers.aws.profiles.default.enabled", true)
	viper.SetDefault("carbon.providers.virtualbox.enabled", true)
	viper.SetDefault("carbon.providers.virtualbox.profiles.default.enabled", true)
	viper.SetDefault("carbon.providers.qemu.enabled", true)
	viper.SetDefault("carbon.providers.qemu.profiles.default.enabled", true)
	viper.SetDefault("carbon.providers.multipass.enabled", true)
	viper.SetDefault("carbon.providers.multipass.profiles.default.enabled", true)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName("carbon")
	}

	viper.AddConfigPath(".")

	// do not need to handle err here. since an error will occur if the config file doesn't exist
	_ = viper.MergeInConfig()
}
