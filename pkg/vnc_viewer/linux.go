//go:build linux

package vnc_viewer

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

func StartViewer(options Options) error {
	vncViewerPath := getVNCViewerPath()
	if vncViewerPath != "" {
		return builder.Cmd(vncViewerPath, "-SecurityTypes", "VncAuth", "-PasswordFile", options.PasswordFile, options.Host).Start()
	}
	return fmt.Errorf("unable to find vncviewer")
}
