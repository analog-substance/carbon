package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// vmList represents the config command
var vmList = &cobra.Command{
	Use:   "list",
	Short: "List VMs",
	Long: `List VMs.

Example:

	carbon vm list

You can also supply a name search:

	carbon vm list -n vm-

This will list VMs with vm- in their name.

`,
	Run: func(cmd *cobra.Command, args []string) {

		vms := getVMsFromArgs(cmd, args)
		if jsonOutput {
			out, err := json.MarshalIndent(vms, "", "  ")
			if err != nil {
				log.Error("error marshalling JSON", "err", err)
			}
			fmt.Println(string(out))
		} else {
			vmTable(vms)
		}
	},
}

func init() {
	vmCmd.AddCommand(vmList)
}
