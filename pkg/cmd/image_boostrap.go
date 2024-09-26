package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// imageBootstrapCmd represents the image command
var imageBootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "create image build configs",
	Long:  `create image build configs`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		osDir, _ := cmd.Flags().GetString("template")
		serviceProvider, _ := cmd.Flags().GetString("service")
		err := carbonObj.CreateImageBuild(name, osDir, serviceProvider)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBootstrapCmd)
	imageBootstrapCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBootstrapCmd.Flags().StringP("template", "t", "ubuntu-24.04", "Template to use")
	imageBootstrapCmd.Flags().StringP("service", "s", "", "Service provider (aws, virtualbox, qemu, multipass)")
}
