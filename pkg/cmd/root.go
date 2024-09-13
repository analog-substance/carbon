/*
Copyright © 2024 defektive

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/analog-substance/carbon/pkg/carbon"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cfgFile string

// var providers []string
var platforms []string
var environments []string
var jsonOutput bool
var carbonObj *carbon.Carbon

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "carbon",
	Short: "Infrastructure ops simplified",
	Long:  `Manage and use infrastructure with a consistent interface, regardless of where it lives.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		providers := viper.GetStringSlice("providers")
		platforms := viper.GetStringSlice("platforms")
		environments := viper.GetStringSlice("environments")

		carbonObj = carbon.New(carbon.Options{
			Providers:    lowerStringSlice(providers),
			Platforms:    lowerStringSlice(platforms),
			Environments: lowerStringSlice(environments),
		})
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carbon.yaml)")
	rootCmd.PersistentFlags().StringSliceP("provider", "P", []string{}, "Provider to use vbox, aws")
	rootCmd.PersistentFlags().StringSliceVarP(&platforms, "platform", "p", []string{}, "Platform to use. Like an instance of a provider. Used to specify aws profiles")
	rootCmd.PersistentFlags().StringSliceVarP(&environments, "environment", "e", []string{}, "Environment to use. Some platforms support many environments.")
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON")

	err := viper.BindPFlag("providers", rootCmd.PersistentFlags().Lookup("provider"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("platforms", rootCmd.PersistentFlags().Lookup("platform"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("environments", rootCmd.PersistentFlags().Lookup("environment"))
	if err != nil {
		log.Fatal(err)
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(cwd)
		viper.AddConfigPath(home)
		viper.SetConfigName("carbon")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// do not need to handle err here. since an error will occur if the config file doesn't exist
	_ = viper.ReadInConfig()
}

func lowerStringSlice(strs []string) []string {
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
	}
	return strs
}
