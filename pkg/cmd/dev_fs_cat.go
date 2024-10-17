package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/spf13/cobra"
	"os"
	"path"
)

// fsCatCmd represents the new command
var fsCatCmd = &cobra.Command{
	Use:   "cat",
	Short: "cat fs contents",
	Long:  `for testing things`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		filesArgs := fsFiles(".")
		return filesArgs, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			fileBytes, err := deployments.Files.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading file %s: %v\n", file, err)
				continue
			}

			fmt.Println(string(fileBytes))
		}
	},
}

func init() {
	fsCmd.AddCommand(fsCatCmd)

}

func fsFiles(dir string) []string {
	listing, err := deployments.Files.ReadDir(dir)
	if err != nil {
		log.Error("failed to read embedded fs", "err", err)
		os.Exit(2)
	}
	files := []string{}
	for _, file := range listing {
		if file.IsDir() {
			files = append(files, fsFiles(path.Join(dir, file.Name()))...)
		} else {
			files = append(files, path.Join(dir, file.Name()))
		}
	}
	return files
}
