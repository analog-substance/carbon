package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmStart represents the config command
var vmStart = &cobra.Command{
	Use:     "start",
	Short:   "Start VMs",
	Long:    `start VMs.`,
	Example: `carbon vm start -n vm-name`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			vmTable(vms)
			if AskIfSure(fmt.Sprintf("Do you want to start %d machines?", len(vms))) {
				for _, vm := range vms {
					err := vm.Start()
					if err != nil {
						log.Error("Error starting VM", "name", vm.Name(), "err", err)
					}
				}
			}
		} else {
			fmt.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmStart)
}
