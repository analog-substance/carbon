package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
	"time"
)

// vmSSH represents the config command
var vmSSH = &cobra.Command{
	Use:   "ssh",
	Short: "SSH to a VM",
	Long: `SSH to a VM.
Carbon will call exec on the ssh binary. This means the SSH process takes
over the carbon process. So SSH agents should just work. 
`,
	Example: `# SSH to a VM
carbon vm ssh -n vm-name


# execute one off command on a VM
carbon vm ssh -n vm-name -- cat /etc/passwd


# proxy through a bastion
carbon vm ssh -n vm-name -- -oProxyCommand="carbon vm ssh -n bastion -- -W %h:%p"


# forward ssh agent
carbon vm ssh -n vm-name -- -A


# open socks proxy
carbon vm ssh -n vm-name -- -D 1080
`,
	Run: func(cmd *cobra.Command, args []string) {
		privateIP, _ := cmd.Flags().GetBool("private-ip")
		user, _ := cmd.Flags().GetString("user")

		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 1 {
			fmt.Println("Too many vms specified.")
		} else if len(vms) == 1 {
			if vms[0].State() != types.StateRunning.Name {
				if AskIfSure("VM is stopped. Would you like to start it?") {
					err := vms[0].Start()
					if err != nil {
						fmt.Printf("Error starting VM: %s\n", err)
					}
					time.Sleep(2 * time.Second)
				}
			}

			err := vms[0].ExecSSH(user, privateIP, args...)
			if err != nil {
				log.Error("failed to ssh to vm", "name", vms[0].Name(), "err", err)
			}
		} else {
			fmt.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmSSH)
	vmSSH.Flags().BoolP("private-ip", "p", false, "Use private IP address")
}
