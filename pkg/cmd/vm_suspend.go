package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmSuspend represents the config command
var vmSuspend = &cobra.Command{
	Use:     "suspend",
	Short:   "suspend or sleep VMs",
	Long:    `suspend or sleep VMs.`,
	Example: `carbon vm start -n vm-name`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			vmTable(vms, false)
			if AskIfSure(fmt.Sprintf("Do you want to suspend %d machines?", len(vms))) {
				for _, vm := range vms {
					err := vm.Suspend()
					if err != nil {
						log.Error("Error suspending VM", "name", vm.Name(), "err", err)
					}
				}
			}
		} else {
			fmt.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmSuspend)
}
