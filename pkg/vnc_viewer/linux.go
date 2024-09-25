//go:build linux

package vnc_viewer

import (
	builder "github.com/NoF0rte/cmd-builder"
)

func StartViewer(options Options) error {
	return builder.Cmd("vncviewer", "-SecurityTypes", "VncAuth", "-PasswordFile", options.PasswordFile, options.Host).Start()
}
