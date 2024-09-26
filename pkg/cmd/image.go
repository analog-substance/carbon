package cmd

import (
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "manage images and image builds",
	Long:  `manage images and image builds`,
}

func init() {
	RootCmd.AddCommand(imageCmd)
}
