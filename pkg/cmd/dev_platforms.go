package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// profilesCmd represents the new command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Query supported providers.",
	Long:  `Query supported providers`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, profiles := range carbonObj.Profiles() {
			fmt.Println(profiles.Name())
			for _, env := range profiles.Environments() {
				fmt.Println(env.Name())
				for _, vm := range env.VMs() {
					fmt.Printf("%s (%s)\n", vm.Name(), vm.State())
				}
			}
		}
	},
}

func init() {
	devCmd.AddCommand(profilesCmd)
}
