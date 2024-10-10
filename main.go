package main

import (
	"github.com/analog-substance/carbon/pkg/cmd"
	"github.com/analog-substance/util/cli/completion"
	"github.com/analog-substance/util/cli/docs"
	"github.com/analog-substance/util/cli/glamour_help"
	"github.com/analog-substance/util/cli/updater/cobra_updater"
	ver "github.com/analog-substance/util/cli/version"
)

var version = "v0.0.0"
var commit = "replace"

func main() {
	cmd.CarbonCmd.Version = ver.GetVersionInfo(version, commit)
	cobra_updater.AddToRootCmd(cmd.CarbonCmd)
	completion.AddToRootCmd(cmd.CarbonCmd)
	cmd.CarbonCmd.AddCommand(docs.CobraDocsCmd)
	glamour_help.AddToRootCmd(cmd.CarbonCmd)

	cmd.Execute()
}
