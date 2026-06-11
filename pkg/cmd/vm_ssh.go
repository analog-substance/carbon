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


# SSH tunneled over AWS SSM (requires AWS-StartSSHSession document enabled on the instance)
carbon vm ssh -n vm-name --ssm

# SSH over SSM with a specific AWS profile and region
carbon vm ssh -n vm-name --ssm --aws-profile my-profile --aws-region us-east-1
`,
	Run: func(cmd *cobra.Command, args []string) {
		privateIP, _ := cmd.Flags().GetBool("private-ip")
		user, _ := cmd.Flags().GetString("user")
		useSSM, _ := cmd.Flags().GetBool("ssm")
		awsProfile, _ := cmd.Flags().GetString("aws-profile")
		awsRegion, _ := cmd.Flags().GetString("aws-region")

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

			var err error
			if useSSM {
				err = vms[0].ExecSSHOverSSM(user, vms[0].ID(), awsProfile, awsRegion, args...)
			} else {
				err = vms[0].ExecSSH(user, privateIP, args...)
			}
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
	vmSSH.Flags().Bool("ssm", false, "Tunnel SSH over AWS SSM (requires AWS-StartSSHSession document enabled on the instance)")
	vmSSH.Flags().String("aws-profile", "", "AWS profile to use when --ssm is set")
	vmSSH.Flags().String("aws-region", "", "AWS region to use when --ssm is set")
}
