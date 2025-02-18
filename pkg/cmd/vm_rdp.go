package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmRDPCmd represents the config command
var vmRDPCmd = &cobra.Command{
	Use:   "rdp",
	Short: "RDP to a VM",
	Long: `RDP to a VM.
This will:

- RDP to the VM :D
`,
	Example: `carbon vnc rdp -n vm-name`,
	Run: func(cmd *cobra.Command, args []string) {
		privateIP, _ := cmd.Flags().GetBool("private-ip")
		user, _ := cmd.Flags().GetString("user")
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 1 {
			fmt.Println("Too many vms specified.")
		} else if len(vms) == 1 {
			err := vms[0].StartRDPClient(user, privateIP)
			if err != nil {
				log.Error("failed to vnc to vm", "name", vms[0].Name(), "err", err)
			}
		} else {
			fmt.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmRDPCmd)
	vmRDPCmd.Flags().BoolP("private-ip", "p", false, "Use private IP address")
}
