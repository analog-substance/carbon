package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/cloud_init"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"strings"
)

// devCloud represents the new command
var devCloud = &cobra.Command{
	Use:   "cloud-init",
	Short: "Cloud init stuff",
	Long:  `for testing things`,
	Run: func(cmd *cobra.Command, args []string) {
		baseDir := "var/cloud-init/ubuntu-24.04"
		listing, err := os.ReadDir(baseDir)
		if err != nil {
			log.Fatal(err)
		}

		tpls := map[string]*cloud_init.CloudConfig{}

		endResult := &cloud_init.CloudConfig{}
		for _, d := range listing {
			if strings.HasSuffix(d.Name(), ".yaml") {
				filebytes, err := os.ReadFile(path.Join(baseDir, d.Name()))
				if err != nil {
					log.Fatal(err)
				}

				tpls[d.Name()] = &cloud_init.CloudConfig{}

				err = yaml.Unmarshal(filebytes, tpls[d.Name()])
				if err != nil {
					log.Fatal(err)
				}

				endResult.MergeWith(tpls[d.Name()])
			}
		}
		d, err := yaml.Marshal(&endResult)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println(string(d))
	},
}

func init() {
	devCmd.AddCommand(devCloud)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
