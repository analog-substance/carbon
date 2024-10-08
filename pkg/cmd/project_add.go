package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/cobra"
)

// projectAddCmd represents the config command
var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add resources to a project",
	Long: `Basically lazy wrappers around tedious things.

So you type less and be more productive!`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project")
		vmName, _ := cmd.Flags().GetString("name")
		vmProvider, _ := cmd.Flags().GetString("provider-type")
		vmImage, _ := cmd.Flags().GetString("image")

		project, err := carbonObj.GetProject(projectName)
		if err != nil {
			fmt.Printf("Failed to get project: %s\n", err)
		}

		err = project.AddMachine(&types.ProjectMachine{
			Name:     vmName,
			Provider: vmProvider,
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
	projectAddCmd.PersistentFlags().StringP("image", "i", "", "Name of the VM to add.")
	projectAddCmd.PersistentFlags().StringP("provider-type", "P", "", "Provider for the new machine")

	err := projectAddCmd.RegisterFlagCompletionFunc("provider-type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageBuildProviders()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}

	err = projectAddCmd.RegisterFlagCompletionFunc("image", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getImageIDs()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}
