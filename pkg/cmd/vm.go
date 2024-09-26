package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
)

// vmCmd represents the config command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Manage and interact with VMs",
	Long: `Basically lazy wrappers around tedious things.

So you type less and be more productive!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vm called")
	},
}

func init() {
	RootCmd.AddCommand(vmCmd)

	vmCmd.PersistentFlags().StringP("name", "n", "", "Name of the VM.")
	vmCmd.PersistentFlags().StringP("id", "i", "", "ID of machine to start.")
	vmCmd.PersistentFlags().StringP("user", "u", "ubuntu", "SSH Username.")
	vmCmd.PersistentFlags().StringSlice("host", []string{}, "Hostname or IP Address.")
}

func getVMsFromArgs(cmd *cobra.Command, args []string) []types.VM {
	hosts, _ := cmd.Flags().GetStringSlice("host")
	id, _ := cmd.Flags().GetString("id")
	name, _ := cmd.Flags().GetString("name")

	if len(hosts) > 0 {
		return carbonObj.VMsFromHosts(hosts)
	}

	if name != "" {
		return carbonObj.FindVMByName(name)
	} else if id != "" {
		return carbonObj.FindVMByID(id)
	}
	return carbonObj.GetVMs()
}
