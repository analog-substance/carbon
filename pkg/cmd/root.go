package cmd

import (
	"github.com/analog-substance/carbon/pkg/carbon"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cfgFile string

// var providers []string
// var profiles []string
// var environments []string
var jsonOutput bool
var carbonObj *carbon.Carbon

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "carbon",
	Short: "Infrastructure ops simplified",
	Long:  `Manage and use infrastructure with a consistent interface, regardless of where it lives.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		carbonConfigFile := common.CarbonConfigFile{}
		err := viper.Unmarshal(&carbonConfigFile)
		if err != nil {
			log.Fatalf("unable to decode into config struct, %v", err)
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
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carbon.yaml)")
	// need to rethink how I want this to work
	//RootCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	//RootCmd.PersistentFlags().StringSliceVarP(&profiles, "profile", "p", []string{}, "Profile to use. Like an instance of a provider. Used to specify aws profiles")
	//RootCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some providers/profiles support many environments.")
	RootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault(common.ViperPackerDir, "deployments/packer")

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

func lowerStringSlice(strs []string) []string {
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
	}
	return strs
}
