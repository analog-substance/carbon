package cmd

import (
	"github.com/spf13/cobra"
)

// deploymentCmd represents the config command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage and interact with deployments",
	Long: `Basically lazy wrappers around tedious things.

So you type less and be more productive!`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	devCmd.AddCommand(deploymentCmd)

	deploymentCmd.PersistentFlags().StringP("name", "n", "", "Name of the VM.")
}
