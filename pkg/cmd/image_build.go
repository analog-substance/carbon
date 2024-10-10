package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// imageBuildCmd represents the image command
var imageBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "build an image",
	Long: `build an image.
Example

	carbon image build -m ubuntu-desktop -t qemu

`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		provider, _ := cmd.Flags().GetString("provider-type")
		provisioner, _ := cmd.Flags().GetString("provisioner")
		err := carbonObj.BuildImage(name, provider, provisioner)
		if err != nil {
			log.Error("failed to packer build", "err", err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBuildCmd)
	imageBuildCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBuildCmd.Flags().StringP("provider-type", "t", "", "Name of provider to use")
	imageBuildCmd.Flags().StringP("provisioner", "a", "cloud-init", "Name of provisioner to use")

	err := imageBuildCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageBuildNames()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
	err = imageBuildCmd.RegisterFlagCompletionFunc("provider-type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageBuildProviders()
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

func getImageBuildProviders() (imageBuildProviders []string) {
	imageBuilds, err := carbonObj.GetImageBuilds()
	if err != nil {
		log.Error("failed to get image builds", "err", err)
	}

	providers := make(map[string]bool)
	for _, imageBuild := range imageBuilds {
		providers[imageBuild.ProviderType()] = true
	}
	for provider := range providers {
		imageBuildProviders = append(imageBuildProviders, provider)
	}
	return imageBuildProviders
}
