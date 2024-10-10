package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// imageBootstrapCmd represents the image command
var imageBootstrapCmd = &cobra.Command{
	Use:     "bootstrap",
	Short:   "Create packer files and other image build configs.",
	Long:    `Create packer files and other image build configs.`,
	Example: `carbon image bootstrap -n operator-desktop-aws -s aws -t ubuntu-desktop`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		osDir, _ := cmd.Flags().GetString("template")
		serviceProvider, _ := cmd.Flags().GetString("service")
		err := carbonObj.CreateImageBuild(name, osDir, serviceProvider)
		if err != nil {
			log.Error("failed to bootstrap packer build", "err", err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBootstrapCmd)
	imageBootstrapCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBootstrapCmd.Flags().StringP("template", "t", "ubuntu-24.04", "Template to use")
	imageBootstrapCmd.Flags().StringP("service", "s", "", "Service provider (aws, virtualbox, qemu, multipass)")

	err := imageBootstrapCmd.RegisterFlagCompletionFunc("template", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := carbonObj.GetImageBuildTemplates()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}

	err = imageBootstrapCmd.RegisterFlagCompletionFunc("service", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getServiceProviders()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getServiceProviders() []string {
	var serviceProviderNames []string
	serviceProviders := carbonObj.Providers()
	for _, provider := range serviceProviders {
		serviceProviderNames = append(serviceProviderNames, provider.Name())
	}
	return serviceProviderNames
}
