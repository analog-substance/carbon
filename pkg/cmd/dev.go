package cmd

import (
	"github.com/spf13/cobra"
)

// devCmd represents the image command
var devCmd = &cobra.Command{
	Use:    "dev",
	Short:  "Unstable sub-commands for testing random ideas",
	Long:   `These sub commands are subject to change without warning`,
	Hidden: true,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("this is unstable called")
	//},
}

func init() {
	RootCmd.AddCommand(devCmd)
}
