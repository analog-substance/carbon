package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const defaultConfigFile = "carbon.yaml"

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and manage configuration values.",
	Long: `View and manage configuration values.

Carbon loads configuration files from your home directory, then merges it with
a configuration file in the current directory (if it exists). This should allow
you the flexibility you need.
`,
	Example: `# Configure digitalocean credentials
carbon config carbon.providers.digitalocean.profiles.default.use_1pass_cli true
carbon config carbon.providers.digitalocean.profiles.default.password "op://Private/some path/api_key"


# Set a default project directory
carbon config carbon.dir.instance ~/my/path/haxors
`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		keys := common.Keys()
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		count := len(args)
		switch count {
		case 0:
			c := common.GetConfig()
			cb, err := yaml.Marshal(c)
			if err != nil {
				log.Error(err.Error())
			}
			fmt.Println(string(cb))

		case 1: // 1 arg = get/remove/reset config value, 2 args = set config value
			key := args[0]

			c := common.Get(key)

			cb, err := yaml.Marshal(c)
			if err != nil {
				log.Error(err.Error())
			}
			fmt.Println(string(cb))
		case 2:
			key := args[0]
			value := args[1]

			c := common.Set(key, value)
			cb, err := yaml.Marshal(c)
			if err != nil {
				log.Error(err.Error())
			}
			fileH, err := os.OpenFile(defaultConfigFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0o644))
			if err != nil {
				log.Error(err.Error())
			}
			defer fileH.Close()

			_, err = fileH.Write(cb)
			if err != nil {
				log.Debug(err.Error())
			}
		}
	},
}

func init() {
	CarbonCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("save", "s", false, "save the current configuration")
	configCmd.Flags().BoolP("sub-keys", "k", false, "display only the sub-keys")
	configCmd.Flags().BoolP("remove-reset", "r", false, "remove key from the config or reset to default")
}

func updateConfigHelp() {
	configCmd.Long += fmt.Sprintf("## Configuration keys\n- %s", strings.Join(common.Keys(), "\n - "))
}
