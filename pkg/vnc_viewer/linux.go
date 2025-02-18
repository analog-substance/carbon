//go:build linux

package vnc_viewer

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"os/exec"
)

func StartViewer(options Options) error {
	vncViewerPath := getVNCViewerPath()
	if vncViewerPath != "" {
		log.Debug("vncviewer found in applications", "os", "linux", "vncViewerPath", vncViewerPath)
		return builder.Cmd(vncViewerPath, "-SecurityTypes", "VncAuth", "-PasswordFile", options.PasswordFile, options.Host).Start()
	}
	return fmt.Errorf("unable to find vncviewer")
}
