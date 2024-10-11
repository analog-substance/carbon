package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
)

// projectAddCmd represents the config command
var projectAddCmd = &cobra.Command{
	Use:     "add-machine",
	Short:   "Add a new machine to a project",
	Long:    `Add a new machine to the project.`,
	Example: `carbon project add-machine -p example-qemu-carbon -n modlishka -P qemu -i carbon-ubuntu-desktop-20241008201758`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")
		vmName, _ := cmd.Flags().GetString("name")
		serviceProvider, _ := cmd.Flags().GetString("service")
		vmImage, _ := cmd.Flags().GetString("image")

		project, err := carbonObj.GetProject(projectName)
		if err != nil {
			fmt.Printf("Failed to get project: %s\n", err)
		}

		err = project.AddMachine(&types.ProjectMachine{
			Name:     vmName,
			Provider: serviceProvider,
			Image:    vmImage,
		})
		if err != nil {
			fmt.Printf("Failed to save project: %s\n", err)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
	projectAddCmd.PersistentFlags().StringP("name", "n", "", "Name of the VM to add.")
	projectAddCmd.PersistentFlags().StringP("image", "i", "", "Name of the image to use.")

	addServiceProviderFlag(projectAddCmd)

	err := projectAddCmd.RegisterFlagCompletionFunc("image", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageIDs()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}
