package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/cloud_init"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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
			log.Error("failed to get cloud init files", err)
		}

		tpls := map[string]*cloud_init.CloudConfig{}

		endResult := &cloud_init.CloudConfig{}
		for _, d := range listing {
			if strings.HasSuffix(d.Name(), ".yaml") {
				filebytes, err := os.ReadFile(filepath.Join(baseDir, d.Name()))
				if err != nil {
					log.Error("failed to read file", err)
					continue
				}

				tpls[d.Name()] = &cloud_init.CloudConfig{}

				err = yaml.Unmarshal(filebytes, tpls[d.Name()])
				if err != nil {
					log.Error("failed to unmarshal file", err)
					continue
				}

				endResult.MergeWith(tpls[d.Name()])
			}
		}
		d, err := yaml.Marshal(&endResult)
		if err != nil {
			log.Error("failed to marshal data", err)
			os.Exit(2)
		}
		fmt.Println(string(d))
	},
}

func init() {
	devCmd.AddCommand(devCloud)
}
