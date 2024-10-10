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

// imageListCmd represents the image command
var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list images",
	Long:  `list images and image builds.`,
	Example: `# List all images
carbon image list


#list image builds
carbon image list -b
`,
	Run: func(cmd *cobra.Command, args []string) {
		listBuilds, _ := cmd.Flags().GetBool("builds")

		if jsonOutput {
			if listBuilds {
				imagesBuilds, err := carbonObj.GetImageBuilds()
				if err != nil {
					log.Info("failed to list image builds", "err", err)
					os.Exit(2)
				}

				out, err := json.MarshalIndent(imagesBuilds, "", "  ")
				if err != nil {
					log.Error("error marshalling JSON", "err", err)
				}
				fmt.Println(string(out))

			} else {

				images, err := carbonObj.GetImages()
				if err != nil {
					log.Info("failed to list images", "err", err)
					os.Exit(2)
				}
				out, err := json.MarshalIndent(images, "", "  ")
				if err != nil {
					log.Error("error marshalling JSON", "err", err)
				}
				fmt.Println(string(out))

			}

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

			var headers []string

			if listBuilds {
				imagesBuilds, err := carbonObj.GetImageBuilds()
				if err != nil {
					log.Info("failed to list image builds", "err", err)
					os.Exit(2)
				}
				headers = []string{"Name", "Provisioner", "Provider"}
				for _, imageBuild := range imagesBuilds {
					data = append(data, []string{
						imageBuild.Name(),
						imageBuild.Provisioner(),
						imageBuild.ProviderType(),
					})
				}
			} else {

				images, err := carbonObj.GetImages()
				if err != nil {
					log.Info("failed to list images", "err", err)
					os.Exit(2)
				}

				headers = []string{"Name", "Created", "Env", "Profile", "Provider", "ID"}

				for _, image := range images {

					data = append(data, []string{
						image.Name(),
						image.CreatedAt(),
						image.Environment().Name(),
						image.Environment().Profile().Name(),
						image.Environment().Profile().Provider().Name(),
						image.ID(),
					})
				}
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
	imageCmd.AddCommand(imageListCmd)
	imageListCmd.Flags().BoolP("builds", "b", false, "List build configs")
}
