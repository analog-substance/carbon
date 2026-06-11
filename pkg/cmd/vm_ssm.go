package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

var vmSSMCmd = &cobra.Command{
	Use:   "ssm",
	Short: "Start an AWS SSM session to a VM",
	Long: `Start an AWS SSM session to a VM.

Carbon will call exec on the aws binary using ssm start-session. This means
the AWS SSM process takes over the carbon process.

Requires the AWS CLI and the Session Manager plugin to be installed.
`,
	Example: `# SSM to a VM by name
carbon vm ssm -n vm-name

# SSM to a VM by instance ID
carbon vm ssm -i i-1234567890abcdef0

# SSM with a specific AWS profile
carbon vm ssm -n vm-name --aws-profile my-profile

# SSM with a specific AWS region
carbon vm ssm -n vm-name --region us-east-1
`,
	Run: func(cmd *cobra.Command, args []string) {
		awsProfile, _ := cmd.Flags().GetString("aws-profile")
		region, _ := cmd.Flags().GetString("region")

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

			err := execSSM(vms[0].ID(), awsProfile, region)
			if err != nil {
				log.Error("failed to start ssm session", "name", vms[0].Name(), "err", err)
			}
		} else {
			fmt.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmSSMCmd)
	vmSSMCmd.Flags().String("aws-profile", "", "AWS profile to use")
	vmSSMCmd.Flags().String("region", "", "AWS region")
}

func execSSM(instanceID, awsProfile, region string) error {
	awsPath, err := exec.LookPath("aws")
	if err != nil {
		return fmt.Errorf("aws CLI not found: %w", err)
	}

	cmdArgs := []string{"aws", "ssm", "start-session", "--target", instanceID}
	if awsProfile != "" {
		cmdArgs = append(cmdArgs, "--profile", awsProfile)
	}
	if region != "" {
		cmdArgs = append(cmdArgs, "--region", region)
	}

	if runtime.GOOS == "windows" {
		c := exec.Command(awsPath, cmdArgs[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	}
	return syscall.Exec(awsPath, cmdArgs, os.Environ())
}
