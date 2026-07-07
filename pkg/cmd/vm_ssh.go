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

SSM tunneling (--ssm):
  SSH can be tunnelled over AWS Systems Manager Session Manager instead of a
  direct network connection. This is useful when the instance has no public IP
  or inbound SSH is blocked.

  One-time setup required:
    1. Install the AWS CLI: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
    2. Install the Session Manager plugin: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html
    3. Ensure the target instance has:
         - The SSM Agent running (pre-installed on most AWS-provided AMIs)
         - An IAM instance profile with the AmazonSSMManagedInstanceCore policy
    4. Add the following to your ~/.ssh/config to use carbon as the proxy:

         Host i-* mi-*
           ProxyCommand carbon vm ssh -i %h --ssm -- -W %h:%p
           User ubuntu
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


# SSH tunnelled over AWS SSM
carbon vm ssh -n vm-name --ssm


# port forward over SSM
carbon vm ssh -n vm-name --ssm -- -L 8080:localhost:80


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
