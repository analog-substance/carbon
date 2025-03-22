// Carbon main package
package main

import (
	"github.com/analog-substance/carbon/pkg/cmd"
	"github.com/analog-substance/util/cli/build_info"
	"github.com/analog-substance/util/cli/completion"
	"github.com/analog-substance/util/cli/docs"
	"github.com/analog-substance/util/cli/glamour_help"
	"github.com/analog-substance/util/cli/updater/cobra_updater"
)

var version = "v0.0.0"
var commit = "replace"

func main() {
	versionInfo := build_info.GetVersion(version, commit)
	cmd.CarbonCmd.Version = versionInfo.String()
	cobra_updater.AddToRootCmd(cmd.CarbonCmd, versionInfo)
	completion.AddToRootCmd(cmd.CarbonCmd)
	cmd.CarbonCmd.AddCommand(docs.CobraDocsCmd)
	glamour_help.AddToRootCmd(cmd.CarbonCmd)
	cmd.Execute()
}
