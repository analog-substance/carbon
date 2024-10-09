package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path"
	"strings"
)

// docsCmd represents the completion command
var docsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "Generate docs ",
	Long:   `Generate documentation markdown files from the source code.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		identity := func(s string) string { return s }
		emptyStr := func(s string) string {

			caser := cases.Title(language.English)

			title := strings.TrimSuffix(strings.ReplaceAll(path.Base(s), "_", " "), ".md")
			var currentCmd *cobra.Command
			var err error
			if title != "carbon" {
				title = strings.TrimPrefix(title, "carbon ")
				args := strings.Split(title, " ")
				currentCmd, _, err = RootCmd.Find(args)
				if err != nil {
					fmt.Println("error getting command", err)
				}
			} else {
				currentCmd = RootCmd
			}

			title = caser.String(title)

			return fmt.Sprintf("---\ntitle: %s\ndescription: %s\n---\n\n", title, currentCmd.Short)
		}

		err := doc.GenMarkdownTreeCustom(RootCmd, "./docs/docs/cli/", emptyStr, identity)
		//err := doc.GenMarkdownTree(RootCmd, "./docs/cli/")
		if err != nil {
			log.Error("error generating docs", "err", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(docsCmd)
}
