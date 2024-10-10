package cmd

import (
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "View or manage images and image builds.",
	Long:  `View or manage images and image builds.`,
}

func init() {
	CarbonCmd.AddCommand(imageCmd)
}
