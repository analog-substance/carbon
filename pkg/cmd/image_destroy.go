package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// imageDestroyCmd represents the image command
var imageDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy/delete images",
	Long: `destroy/delete images.
Example

	carbon image destroy -i qemu/some-image-123123123

`,
	Run: func(cmd *cobra.Command, args []string) {
		imageID, _ := cmd.Flags().GetString("image-id")

		image, err := carbonObj.GetImage(imageID)
		if err == nil {
			err = image.Destroy()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: ", err)
		}

	},
}

func init() {
	imageCmd.AddCommand(imageDestroyCmd)
	imageDestroyCmd.Flags().StringP("image-id", "i", "", "ID of image to delete")

	err := imageDestroyCmd.RegisterFlagCompletionFunc("image-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageIDs()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getImageIDs() (imageIDs []string) {
	images, err := carbonObj.GetImages()
	if err != nil {
		log.Error(err.Error())
	}

	for _, image := range images {
		imageIDs = append(imageIDs, image.ID())
	}
	return imageIDs
}
