package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// projectListCmd represents the config command
var projectListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Manage and interact with projects",
	Long:    `List projects.`,
	Example: `carbon project list`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := carbonObj.GetProjects()
		if err != nil {
			fmt.Printf("Error getting deployments: %v\n", err)
			os.Exit(1)
		}

		if jsonOutput {
			out, err := json.MarshalIndent(projects, "", "  ")
			if err != nil {
				log.Error("error marshalling JSON", "err", err)
			}
			fmt.Println(string(out))
		} else {

			re := lipgloss.NewRenderer(os.Stdout)
			baseStyle := re.NewStyle().Padding(0, 1)
			headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)

			CapitalizeHeaders := func(data []string) []string {
				for i := range data {
					data[i] = strings.ToUpper(data[i])
				}
				return data
			}

			data := [][]string{}
			var headers = []string{"Name"}
			for _, imageBuild := range projects {
				data = append(data, []string{
					imageBuild.Name(),
				})
			}

			ct := table.New().
				Border(lipgloss.NormalBorder()).
				BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
				Headers(CapitalizeHeaders(headers)...).
				Rows(data...).
				StyleFunc(func(row, col int) lipgloss.Style {
					if row == 0 {
						return headerStyle
					}

					even := row%2 == 0

					if even {
						return baseStyle.Foreground(lipgloss.Color("245"))
					}
					return baseStyle.Foreground(lipgloss.Color("252"))
				})
			fmt.Println(ct)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

}
