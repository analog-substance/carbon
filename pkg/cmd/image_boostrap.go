package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// imageBootstrapCmd represents the image command
var imageBootstrapCmd = &cobra.Command{
	Use:     "bootstrap",
	Short:   "Create packer files and other image build configs.",
	Long:    `Create packer files and other image build configs.`,
	Example: `carbon image bootstrap -n operator-desktop-aws -S aws -t ubuntu-desktop`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		osDir, _ := cmd.Flags().GetString("template")
		serviceProvider, _ := cmd.Flags().GetString("service")

		provider, err := carbonObj.GetProvider(serviceProvider)
		if err != nil {
			log.Error("failed to get provider", "provider", serviceProvider, "err", err)
			os.Exit(1)
		}

		imageBuild, err := provider.NewImageBuild(name, osDir)
		//err := carbonObj.CreateImageBuild(name, osDir, serviceProvider)
		if err != nil {
			log.Error("failed to bootstrap packer build", "err", err)
		}

		fmt.Printf("Image build successfully created: %s\n", imageBuild.Name())
	},
}

func init() {
	imageCmd.AddCommand(imageBootstrapCmd)
	imageBootstrapCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBootstrapCmd.Flags().StringP("template", "t", "ubuntu-24.04", "Template to use")
	addServiceProviderFlag(imageBootstrapCmd)

	err := imageBootstrapCmd.RegisterFlagCompletionFunc("template", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := carbonObj.GetImageBuildTemplates()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}

}
