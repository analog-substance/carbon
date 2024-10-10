package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

// imageListCmd represents the image command
var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list images",
	Long: `list images and image builds.
Example

	carbon image list
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

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)

			if listBuilds {
				imagesBuilds, err := carbonObj.GetImageBuilds()
				if err != nil {
					log.Info("failed to list image builds", "err", err)
					os.Exit(2)
				}
				t.AppendHeader(table.Row{"Name", "Provisioner", "Provider"})

				for _, imageBuild := range imagesBuilds {
					t.AppendRows([]table.Row{
						{
							imageBuild.Name(),
							imageBuild.Provisioner(),
							imageBuild.ProviderType(),
						},
					})
				}
			} else {

				images, err := carbonObj.GetImages()
				if err != nil {
					log.Info("failed to list images", "err", err)
					os.Exit(2)
				}
				t.AppendHeader(table.Row{"Name", "Created", "Env", "Profile", "Provider", "ID"})

				for _, image := range images {

					t.AppendRows([]table.Row{
						{
							image.Name(),
							image.CreatedAt(),
							image.Environment().Name(),
							image.Environment().Profile().Name(),
							image.Environment().Profile().Provider().Name(),
							image.ID(),
						},
					})
				}
			}
			t.Render()
		}
	},
}

func init() {
	imageCmd.AddCommand(imageListCmd)
	imageListCmd.Flags().BoolP("builds", "b", false, "List build configs")
}
