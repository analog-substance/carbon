package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// devVMCmd represents the new command
var devVMCmd = &cobra.Command{
	Use:   "platform",
	Short: "Query supported platforms.",
	Long:  `Query supported platform`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, platform := range carbonObj.Platforms() {
			for _, env := range platform.Environments() {
				for _, vm := range env.VMs() {
					fmt.Printf("%s / %s: %s (%s)\n", platform.Name(), env.Name(), vm.Name(), vm.State())
				}
			}
		}
	},
}

func init() {
	devCmd.AddCommand(devVMCmd)
}
