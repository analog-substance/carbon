package main

import (
	"github.com/analog-substance/carbon/pkg/cmd"
	"github.com/analog-substance/util/cli/updater/cobra_helper"
	ver "github.com/analog-substance/util/cli/version"
)

var version = "v0.0.0"
var commit = "replace"

func main() {
	cmd.RootCmd.Version = ver.GetVersionInfo(version, commit)
	cmd.RootCmd.AddCommand(cobra_helper.CobraUpdateCmd)
	cmd.Execute()
}
