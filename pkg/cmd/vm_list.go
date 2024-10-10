package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// vmList represents the config command
var vmList = &cobra.Command{
	Use:   "list",
	Short: "List VMs across all available providers, profiles, and environments.",
	Long:  `List VMs across all available providers, profiles, and environments.`,
	Example: `# list all virtual machines
carbon vm list


# You can also supply a name search, this wil return VMs with names containing 'vm-'
carbon vm list -n vm-
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
