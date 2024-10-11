package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
	"os"
)

// vmLaunchCmd represents the image command
var vmLaunchCmd = &cobra.Command{
	Use:     "launch",
	Short:   "launch a new vm from an image",
	Long:    `launch a new vm from an image.`,
	Example: `carbon vm launch -I qemu/carbon-ubuntu-desktop-20241007212910 -n vm-name`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		imageID, _ := cmd.Flags().GetString("image-id")
		image, err := carbonObj.GetImage(imageID)
		if err != nil {
			log.Error("failed to find image", "err", err)
			os.Exit(1)
		}

		err = image.Launch(types.ImageLaunchOptions{Name: name})
		if err != nil {
			log.Error("failed to launch vm", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	vmCmd.AddCommand(vmLaunchCmd)
	vmLaunchCmd.Flags().StringP("name", "n", "", "Name of new VM")
	vmLaunchCmd.Flags().StringP("image-id", "I", "", "ID of image")

	err := vmLaunchCmd.RegisterFlagCompletionFunc("image-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageIDs()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}
