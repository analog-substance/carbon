package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmDestroyCmd represents the image command
var vmDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Long: `Destroy a VM.

Example:

	carbon vm destroy -n vm-name

`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			vmTable(vms)
			if AskIfSure(fmt.Sprintf("Do you want to destroy %d machines?", len(vms))) {
				for _, vm := range vms {
					err := vm.Destroy()
					if err != nil {
						log.Error("Error destroy VM", "name", vm.Name(), "err", err)
					}
				}
			}
		} else {
			fmt.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmDestroyCmd)
}
