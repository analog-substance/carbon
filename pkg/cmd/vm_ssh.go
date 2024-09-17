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
	"github.com/analog-substance/carbon/pkg/providers/types"
	"github.com/spf13/cobra"
	"log"
)

// vmSSH represents the config command
var vmSSH = &cobra.Command{
	Use:   "ssh",
	Short: "ssh to a vm",
	Long:  `ssh to a vm`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		user, _ := cmd.Flags().GetString("user")
		var vm types.VM
		if name != "" {
			vm = carbonObj.FindVMByName(name)
		} else if id != "" {
			vm = carbonObj.FindVMByID(id)
		}
		if vm != nil {
			err := vm.ExecSSH(user)
			if err != nil {
				log.Println("Error SSHing to VM:", err)
			}
		} else {
			log.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmSSH)
	vmSSH.Flags().StringP("name", "name", "", "Name of the VM")
	vmSSH.Flags().StringP("id", "i", "", "ID of machine to start")
	vmSSH.Flags().StringP("user", "u", "ubuntu", "SSH Username")
}
