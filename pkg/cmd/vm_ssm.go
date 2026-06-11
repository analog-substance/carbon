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

var vmSSMPortForwardCmd = &cobra.Command{
	Use:   "port-forward",
	Short: "Forward a port via AWS SSM",
	Long: `Forward a local port to a port on an AWS VM (or a remote host reachable from it) via SSM.

Without --remote-host, forwards local-port -> port on the target instance
(AWS-StartPortForwardingSession).

With --remote-host, forwards local-port -> remote-host:remote-port through
the target instance (AWS-StartPortForwardingSessionToRemoteHost).

Carbon will call exec on the aws binary. This means the AWS SSM process
takes over the carbon process.

Requires the AWS CLI and the Session Manager plugin to be installed.
`,
	Example: `# Forward local port 8080 to port 80 on the instance
carbon vm ssm port-forward -n vm-name --local-port 8080 --remote-port 80

# Forward local port 5432 to a private RDS host through the instance
carbon vm ssm port-forward -n vm-name --local-port 5432 --remote-host db.internal --remote-port 5432

# With a specific AWS profile and region
carbon vm ssm port-forward -n vm-name --local-port 8080 --remote-port 80 --aws-profile prod --region us-east-1
`,
	Run: func(cmd *cobra.Command, args []string) {
		awsProfile, _ := cmd.Flags().GetString("aws-profile")
		region, _ := cmd.Flags().GetString("region")
		localPort, _ := cmd.Flags().GetString("local-port")
		remotePort, _ := cmd.Flags().GetString("remote-port")
		remoteHost, _ := cmd.Flags().GetString("remote-host")

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

			err := execSSMPortForward(vms[0].ID(), awsProfile, region, localPort, remotePort, remoteHost)
			if err != nil {
				log.Error("failed to start ssm port-forward session", "name", vms[0].Name(), "err", err)
			}
		} else {
			fmt.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmSSMCmd)
	vmSSMCmd.PersistentFlags().String("aws-profile", "", "AWS profile to use")
	vmSSMCmd.PersistentFlags().String("region", "", "AWS region")

	vmSSMCmd.AddCommand(vmSSMPortForwardCmd)
	vmSSMPortForwardCmd.Flags().String("local-port", "", "Local port to listen on")
	vmSSMPortForwardCmd.Flags().String("remote-port", "", "Remote port to forward to")
	vmSSMPortForwardCmd.Flags().String("remote-host", "", "Remote host to forward to (enables host-forwarding document)")
	_ = vmSSMPortForwardCmd.MarkFlagRequired("local-port")
	_ = vmSSMPortForwardCmd.MarkFlagRequired("remote-port")
}

func buildSSMArgs(instanceID, awsProfile, region string) []string {
	args := []string{"aws", "ssm", "start-session", "--target", instanceID}
	if awsProfile != "" {
		args = append(args, "--profile", awsProfile)
	}
	if region != "" {
		args = append(args, "--region", region)
	}
	return args
}

func buildSSMPortForwardArgs(instanceID, awsProfile, region, localPort, remotePort, remoteHost string) []string {
	args := []string{"aws", "ssm", "start-session", "--target", instanceID}

	if remoteHost != "" {
		args = append(args,
			"--document-name", "AWS-StartPortForwardingSessionToRemoteHost",
			"--parameters", fmt.Sprintf(`{"host":["%s"],"portNumber":["%s"],"localPortNumber":["%s"]}`, remoteHost, remotePort, localPort),
		)
	} else {
		args = append(args,
			"--document-name", "AWS-StartPortForwardingSession",
			"--parameters", fmt.Sprintf(`{"portNumber":["%s"],"localPortNumber":["%s"]}`, remotePort, localPort),
		)
	}

	if awsProfile != "" {
		args = append(args, "--profile", awsProfile)
	}
	if region != "" {
		args = append(args, "--region", region)
	}
	return args
}

func execSSM(instanceID, awsProfile, region string) error {
	awsPath, err := exec.LookPath("aws")
	if err != nil {
		return fmt.Errorf("aws CLI not found: %w", err)
	}

	cmdArgs := buildSSMArgs(instanceID, awsProfile, region)

	if runtime.GOOS == "windows" {
		c := exec.Command(awsPath, cmdArgs[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	}
	return syscall.Exec(awsPath, cmdArgs, os.Environ())
}

func execSSMPortForward(instanceID, awsProfile, region, localPort, remotePort, remoteHost string) error {
	awsPath, err := exec.LookPath("aws")
	if err != nil {
		return fmt.Errorf("aws CLI not found: %w", err)
	}

	cmdArgs := buildSSMPortForwardArgs(instanceID, awsProfile, region, localPort, remotePort, remoteHost)

	if runtime.GOOS == "windows" {
		c := exec.Command(awsPath, cmdArgs[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	}
	return syscall.Exec(awsPath, cmdArgs, os.Environ())
}
