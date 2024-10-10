package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// projectListCmd represents the config command
var projectListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Manage and interact with projects",
	Long:    `List projects.`,
	Example: `carbon project list`,
	Run: func(cmd *cobra.Command, args []string) {
		deployments, err := carbonObj.GetProjects()
		if err != nil {
			fmt.Printf("Error getting deployments: %v\n", err)
			os.Exit(1)
		}

		for _, d := range deployments {
			fmt.Printf("%s\n", d.Name())
		}
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

}
