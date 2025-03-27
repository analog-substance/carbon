package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// projectUpdateCmd represents the config command
var projectUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update terraform module with latest from Carbon.",
	Long:    `Update terraform module with latest from Carbon.`,
	Example: `carbon project update`,
	Run: func(cmd *cobra.Command, args []string) {
		err := carbonObj.CopyTFModule()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)
}
