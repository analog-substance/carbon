package common

import (
	"github.com/spf13/viper"
	"os/user"
	"path/filepath"
	"strings"
)

var instancePath string

func getInstancePath() string {
	if instancePath == "" {
		instancePath = viper.GetString(ViperDefaultInstanceDir)
		usr, _ := user.Current()
		dir := usr.HomeDir
		if instancePath == "~" {
			instancePath = dir
		} else if strings.HasPrefix(instancePath, "~/") {
			instancePath = filepath.Join(dir, instancePath[2:])
		}
	}

	return instancePath
}

func PackerDir() string {
	return filepath.Join(getInstancePath(), viper.GetString(ViperPackerDir))
}

func ImagesDir() string {
	return filepath.Join(getInstancePath(), viper.GetString(ViperImagesDir))
}

func ProjectsDir() string {
	return filepath.Join(getInstancePath(), viper.GetString(ViperTerraformProjectDir))
}
