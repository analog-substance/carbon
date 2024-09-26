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
		err := carbonObj.BuildImage(name)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBuildCmd)
	imageBuildCmd.Flags().StringP("name", "n", "", "Name of image build")
}
