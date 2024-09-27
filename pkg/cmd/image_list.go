package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// imageListCmd represents the image command
var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list images",
	Long:  `list images`,
	Run: func(cmd *cobra.Command, args []string) {
		listBuilds, _ := cmd.Flags().GetBool("builds")
		if listBuilds {
			imagesBuilds, err := carbonObj.GetImageBuilds()
			if err != nil {
				log.Info("failed to list image builds", err)
				os.Exit(2)
			}
			for _, imageBuild := range imagesBuilds {
				fmt.Println(imageBuild.Name(), imageBuild.Provisioner(), imageBuild.ProviderType())
			}
		} else {

			imagesBuilds, err := carbonObj.GetImages()
			if err != nil {
				log.Info("failed to list images", err)
				os.Exit(2)
			}
			for _, imageBuild := range imagesBuilds {
				fmt.Println(imageBuild.Name(), imageBuild.Environment().Name())
			}
		}

	},
}

func init() {
	imageCmd.AddCommand(imageListCmd)
	imageListCmd.Flags().BoolP("builds", "b", false, "List build configs")
}
