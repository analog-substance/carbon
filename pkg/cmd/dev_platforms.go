package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// platformCmd represents the new command
var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Query supported platforms.",
	Long:  `Query supported platform`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, platform := range carbonObj.Platforms() {
			fmt.Println(platform.Name())
			for _, env := range platform.Environments() {
				fmt.Println(env.Name())
				for _, vm := range env.VMs() {
					fmt.Printf("%s (%s)\n", vm.Name(), vm.State())
				}
			}
		}
	},
}

func init() {
	devCmd.AddCommand(platformCmd)
}
