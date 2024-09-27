package image_build

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func GetImageBuildsForProvider(provider string) ([]string, error) {
	ret := []string{}
	images, err := GetImageBuilds()
	if err != nil {
		return ret, err
	}
	for _, imageBuildDir := range images {
		if strings.HasPrefix(imageBuildDir, provider) {
			ret = append(ret, imageBuildDir)
		}
	}
	return ret, nil
}

func GetImageBuilds() ([]string, error) {
	ret := []string{}
	packerDir := viper.GetString(common.ConfigPackerDir)
	listing, err := os.ReadDir(packerDir)
	if err != nil {
		return ret, err
	}
	for _, file := range listing {
		ret = append(ret, file.Name())
	}
	return ret, nil
}
