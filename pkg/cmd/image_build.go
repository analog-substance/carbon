package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// imageBuildCmd represents the image command
var imageBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "build an image",
	Long:  `build an image`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		provider, _ := cmd.Flags().GetString("provider-type")
		provisioner, _ := cmd.Flags().GetString("provisioner")
		err := carbonObj.BuildImage(name, provider, provisioner)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBuildCmd)
	imageBuildCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBuildCmd.Flags().StringP("provider-type", "t", "", "Name of provider to use")
	imageBuildCmd.Flags().StringP("provisioner", "a", "cloud-init", "Name of provisioner to use")
}
