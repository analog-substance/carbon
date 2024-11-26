package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmRestart represents the config command
var vmRestart = &cobra.Command{
	Use:     "restart",
	Short:   "Restart VM(s)",
	Long:    `Restart VM(s).`,
	Example: `carbon vm restart -n vm-name`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			vmTable(vms, false)
			if AskIfSure(fmt.Sprintf("Do you want to restart %d machines?", len(vms))) {
				for _, vm := range vms {
					err := vm.Restart()
					if err != nil {
						log.Error("Error restarting VM", "name", vm.Name(), "err", err)
					}
				}
			}
		} else {
			fmt.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmRestart)
}
