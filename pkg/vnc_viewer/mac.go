//go:build darwin

package vnc_viewer

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

func StartViewer(options Options) error {

	vncViewerPath := findVncApp()
	if vncViewerPath != "" {
		log.Debug("vncviewer found in applications", "os", "mac", "vncViewerPath", vncViewerPath)
		return builder.Cmd("open", "-a", vncViewerPath, "--args", options.PasswordFile, options.Host).Stderr(nil).Start()
	}

	vncViewerPath = getVNCViewerPath()
	if vncViewerPath != "" {
		log.Debug("vncviewer found in path", "os", "mac", "vncViewerPath", vncViewerPath)
		return builder.Cmd(vncViewerPath, "-SecurityTypes", "VncAuth", "-PasswordFile", options.PasswordFile, options.Host).Start()
	}

	log.Debug("vncviewer not found", "os", "mac", "vncViewerPath", vncViewerPath)
	return fmt.Errorf("unable to find vncviewer")
}

const lsRegPath string = "/System/Library/Frameworks/CoreServices.framework/Versions/A/Frameworks/LaunchServices.framework/Versions/A/Support/lsregister"

func findVncApp() string {
	output, err := builder.Shell(fmt.Sprintf(`%s -dump | grep -o '/Applications/.*Tiger.*VNC.*\.app' | head -n 1`, lsRegPath)).Output()
	if err != nil {
		return ""
	}
	return output
}
