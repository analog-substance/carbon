package cmd

import (
	"github.com/spf13/cobra"
	"log"
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
					log.Printf("Error stopping VM (%s): %s", vm.Name(), err)
				}
			}
		} else {
			log.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmStop)
}
