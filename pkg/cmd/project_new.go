package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// projectNewCmd represents the config command
var projectNewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new project.",
	Long:    `Create a new project.`,
	Example: `carbon project new -n project-name`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")
		serviceProvider, _ := cmd.Flags().GetString("service")
		force, _ := cmd.Flags().GetBool("force")

		project, err := carbonObj.NewProject(projectName, serviceProvider, force)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Project created: %s\n", project.Name())
		//
		//provider, err := carbonObj.GetProvider(serviceProvider)
		//if err != nil {
		//	log.Error("Failed to get provider", "err", err)
		//	os.Exit(1)
		//}
		//
		//project, err := carbonObj.GetProject(projectName)
		//if !force && (err == nil || project != nil) {
		//	if !AskIfSure("Project already exists. would you like to overwrite the files?") {
		//		os.Exit(1)
		//	}
		//	force = true
		//}
		//
		//project, err = provider.NewProject(projectName, force)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Printf("Project created: %s\n", project.Name())
	},
}

func init() {
	projectCmd.AddCommand(projectNewCmd)
	projectNewCmd.PersistentFlags().StringP("name", "n", "", "Name of the new project.")
	projectNewCmd.PersistentFlags().BoolP("force", "f", false, "Force over writing files.")
	addServiceProviderFlag(projectNewCmd)
	projectNewCmd.MarkFlagRequired("name")
}
