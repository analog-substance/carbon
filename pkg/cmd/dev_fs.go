package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/spf13/cobra"
	"log"
	"path"
)

// fsCmd represents the new command
var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "List fs contents",
	Long:  `for testing things`,
	Run: func(cmd *cobra.Command, args []string) {
		ListingDir(".")
	},
}

func init() {
	devCmd.AddCommand(fsCmd)
}

func ListingDir(dir string) {
	listing, err := deployments.Files.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range listing {
		fmt.Println(path.Join(dir, file.Name()))
		if file.IsDir() {
			ListingDir(path.Join(dir, file.Name()))
		}
	}
}
