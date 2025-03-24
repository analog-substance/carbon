package common

import (
	"os/user"
	"path/filepath"
	"strings"
)

var instancePath string

func getInstancePath() string {
	if instancePath == "" {
		instancePath = GetConfig().Carbon.Dir[DefaultInstanceConfigKey]
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
	return filepath.Join(getInstancePath(), GetConfig().Carbon.Dir[PackerConfigKey])
}

func ImagesDir() string {
	return filepath.Join(getInstancePath(), GetConfig().Carbon.Dir[ImagesConfigKey])
}

func ProjectsDir() string {
	return filepath.Join(getInstancePath(), GetConfig().Carbon.Dir[TerraformProjectConfigKey])
}
