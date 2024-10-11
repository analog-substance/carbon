package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// imageBuildCmd represents the image command
var imageBuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Build an image.",
	Long:    `build an image.`,
	Example: `carbon image build -t aws -n operator-desktop-aws`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		serviceProvider, _ := cmd.Flags().GetString("service")
		provisioner, _ := cmd.Flags().GetString("provisioner")

		imageBuild, err := carbonObj.GetImageBuild(name, serviceProvider, provisioner)
		if err != nil {
			log.Error("error getting image build", "err", err)
			os.Exit(1)
		}

		err = imageBuild.Build()
		if err != nil {
			log.Error("error building image", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBuildCmd)
	imageBuildCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBuildCmd.Flags().StringP("provisioner", "a", "cloud-init", "Name of provisioner to use")

	addServiceProviderFlag(imageBuildCmd)

	err := imageBuildCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageBuildNames()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getImageBuildNames() (imageBuildNames []string) {
	imageBuilds, err := carbonObj.GetImageBuilds()
	if err != nil {
		log.Error("failed to get image builds", "err", err)
	}

	for _, imageBuild := range imageBuilds {
		imageBuildNames = append(imageBuildNames, imageBuild.Name())
	}
	return imageBuildNames
}
