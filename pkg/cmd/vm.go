/*
Copyright © 2024 defektive

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
