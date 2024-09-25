//go:build linux

package vnc_viewer

import (
	builder "github.com/NoF0rte/cmd-builder"
	"log"
)

func StartViewer(options Options) error {
	log.Println("starting vnc vierwwer")
	return builder.Cmd("vncviewer", "-SecurityTypes", "VncAuth", "-PasswordFile", options.PasswordFile, options.Host).Start()
}
