package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmStart represents the config command
var vmStart = &cobra.Command{
	Use:   "start",
	Short: "Start VMs",
	Long:  `start VMs`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			for _, vm := range vms {
				err := vm.Start()
				if err != nil {
					log.Error("Error starting VM", "name", vm.Name(), "err", err)
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
