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
	"github.com/spf13/cobra"
	"log"
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
					log.Printf("Error restarting VM (%s): %s", vm.Name(), err)
				}
			}
		} else {
			log.Println("No VMs found.")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmRestart)
}
