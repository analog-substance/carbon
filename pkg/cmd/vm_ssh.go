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

Example:

	carbon vm ssh -n vm-name

Carbon will call exec on the ssh binary. This means the SSH process takes
over the carbon process. So SSH agents should just work. 
`,
	Run: func(cmd *cobra.Command, args []string) {
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

			err := vms[0].ExecSSH(user, args...)
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
}
