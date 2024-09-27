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
var profiles []string
var environments []string
var jsonOutput bool
var carbonObj *carbon.Carbon

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "carbon",
	Short: "Infrastructure ops simplified",
	Long:  `Manage and use infrastructure with a consistent interface, regardless of where it lives.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		providers := viper.GetStringSlice("carbon.providers")
		profiles := viper.GetStringSlice("carbon.profiles")
		environments := viper.GetStringSlice("carbon.environments")

		carbonObj = carbon.New(carbon.Options{
			Providers:    lowerStringSlice(providers),
			Profiles:     lowerStringSlice(profiles),
			Environments: lowerStringSlice(environments),
		})
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
	RootCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	RootCmd.PersistentFlags().StringSliceVarP(&profiles, "profile", "p", []string{}, "Profile to use. Like an instance of a provider. Used to specify aws profiles")
	RootCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some providers/profiles support many environments.")
	RootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")

	err := viper.BindPFlag("carbon.providers", RootCmd.PersistentFlags().Lookup("provider"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("carbon.profiles", RootCmd.PersistentFlags().Lookup("profiles"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("carbon.environments", RootCmd.PersistentFlags().Lookup("environment"))
	if err != nil {
		log.Fatal(err)
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault("carbon.providers", []string{})
	viper.SetDefault("carbon.profiles", []string{})
	viper.SetDefault("carbon.environments", []string{})
	viper.SetDefault("carbon.credentials", []map[string]string{})
	viper.SetDefault("carbon.cloud-init.dir", "cloud-init")
	viper.SetDefault(common.ConfigPackerDir, "deployments/packer")

	//const CloudInitDir = "cloud-init"
	//const PackerDir = "deployments/packer"
	//const PackerFileSuffixCloudInit = "-cloud-init.pkr.hcl"
	//const PackerFileSuffixAnsible = "-ansible.pkr.hcl"
	//const PackerFileSuffixVariables = "-variables.pkr.hcl"
	//const PackerFilePrivateVarsExample = "private.auto.pkrvars.hcl.example"
	//const PackerFileIsoVars = "iso-variables.pkr.hcl"
	//const PackerFileLocalVars = "local-variables.pkr.hcl"
	//const PackerFilePacker = "packer.pkr.hcl"
	//const ISOVarUsage = "var.iso_url"

	// Use config file from the flag.
	// Find home directory.
	//cwd, err := os.Getwd()
	//if err != nil {
	//	log.Println(err)
	//}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName("carbon")
	}

	//err := viper.ReadInConfig()
	//cobra.CheckErr(err)

	//settings := viper.AllSettings()
	//bytes, err := json.MarshalIndent(settings, "", "  ")
	//cobra.CheckErr(err)
	//log.Printf(string(bytes))

	viper.AddConfigPath(".")

	// do not need to handle err here. since an error will occur if the config file doesn't exist
	_ = viper.MergeInConfig()
	//cobra.CheckErr(err)

	//settings = viper.AllSettings()
	//bytes, err = json.MarshalIndent(settings, "", "  ")
	//cobra.CheckErr(err)
	//log.Printf(string(bytes))
}

func lowerStringSlice(strs []string) []string {
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
	}
	return strs
}
