package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmRestart represents the config command
var vmRestart = &cobra.Command{
	Use:   "restart",
	Short: "Restart VM(s)",
	Long:  `Restart VM(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			for _, vm := range vms {
				err := vm.Restart()
				if err != nil {
					log.Error("Error restarting VM", vm.Name(), err)
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
