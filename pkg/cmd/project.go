package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// projectCmd represents the config command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage and interact with deployments",
	Long: `Basically lazy wrappers around tedious things.

So you type less and be more productive!`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(projectCmd)

	projectCmd.PersistentFlags().StringP("project", "p", "", "Name of the project.")

	err := projectCmd.RegisterFlagCompletionFunc("project", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getProjectNames()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getProjectNames() (projectNames []string) {
	projects, err := carbonObj.GetProjects()
	if err == nil {
		for _, project := range projects {
			projectNames = append(projectNames, project.Name())
		}
	}
	return projectNames
}
