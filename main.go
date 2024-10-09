package main

import (
	"github.com/analog-substance/carbon/pkg/cmd"
	"github.com/analog-substance/util/cli/completion"
	"github.com/analog-substance/util/cli/docs"
	"github.com/analog-substance/util/cli/updater/cobra_updater"
	ver "github.com/analog-substance/util/cli/version"
)

var version = "v0.0.0"
var commit = "replace"

func main() {
	cmd.RootCmd.Version = ver.GetVersionInfo(version, commit)
	cmd.RootCmd.AddCommand(docs.CobraDocsCmd)
	cobra_updater.AddToRootCmd(cmd.RootCmd)
	completion.AddToRootCmd(cmd.RootCmd)
	cmd.Execute()
}
