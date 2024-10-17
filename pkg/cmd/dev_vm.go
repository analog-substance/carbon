package cmd

import (
	"github.com/analog-substance/carbon/pkg/tui"
	"github.com/spf13/cobra"
)

// devVMCmd represents the new command
var devVMCmd = &cobra.Command{
	Use:   "vm",
	Short: "list vms.",
	Long:  `list vms`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := carbonObj.GetVMs()
		tui.Test(vms)

		//for _, profiles := range carbonObj.Profiles() {
		//	for _, env := range profiles.Environments() {
		//		for _, vm := range env.VMs() {
		//			fmt.Printf("%s / %s: %s (%s)\n", profiles.Name(), env.Name(), vm.Name(), vm.State())
		//		}
		//	}
		//}
	},
}

func init() {
	devCmd.AddCommand(devVMCmd)
}
