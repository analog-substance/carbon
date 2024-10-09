package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vmVNCCmd represents the config command
var vmVNCCmd = &cobra.Command{
	Use:   "vnc",
	Short: "VNC to a VM",
	Long: `VNC to a VM.

Example:

	carbon vm vnc -n vm-name

This will:
- SSH to the target VM.
- Check to see if vncserver is running.
- If not, start vncserver on the remote machine
- If a vnc passwd file does not exist one will be created
- Copy the password file to the local machine
- Setup a tunnel to access vnc
- start vncviewer
`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		killVNC, _ := cmd.Flags().GetBool("kill-vnc")
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 1 {
			fmt.Println("Too many vms specified.")
		} else if len(vms) == 1 {
			err := vms[0].StartVNC(user, killVNC)
			if err != nil {
				log.Error("failed to vnc to vm", "name", vms[0].Name(), "err", err)
			}
		} else {
			fmt.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmVNCCmd)
	vmVNCCmd.Flags().BoolP("kill-vnc", "k", false, "Kill VNC before starting")
}
