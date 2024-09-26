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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
