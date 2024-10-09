package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmStop represents the config command
var vmStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop VM(s)",
	Long: `Stop VM(s).

Example:

	carbon vm stop -n vm-name

By default, carbon will attempt to hibernate the machine rather than power it off.
`,
	Run: func(cmd *cobra.Command, args []string) {
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 0 {
			vmTable(vms)
			if AskIfSure(fmt.Sprintf("Do you want to stop %d machines?", len(vms))) {
				for _, vm := range vms {
					err := vm.Stop()
					if err != nil {
						log.Error("Error stopping VM", "name", vm.Name(), "err", err)
					}
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
