package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// projectApplyCmd represents the config command
var projectApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Run terraform apply",
	Long: `Apply terraform.

Example:

	carbon project apply -p project-name

`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("project")

		project, err := carbonObj.GetProject(name)
		if err != nil {
			fmt.Printf("Failed to get project: %s\n", err)
			os.Exit(1)
		}
		err = project.TerraformApply()
		if err != nil {
			fmt.Printf("Failed to apply project: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectApplyCmd)

}
