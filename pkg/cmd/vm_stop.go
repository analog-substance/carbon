package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmStop represents the config command
var vmStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop VM(s)",
	Long:  `Stop VM(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			for _, vm := range vms {
				err := vm.Stop()
				if err != nil {
					log.Error("Error stopping VM", vm.Name(), err)
				}
			}
		} else {
			fmt.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmStop)
}
