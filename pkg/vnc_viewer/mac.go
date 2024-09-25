//go:build darwin

package vnc_viewer

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

func StartViewer(options Options) error {
	vncViewerPath := findVncApp()

	return builder.Cmd("open", "-a", vncViewerPath, "--args", options.PasswordFile, options.Host).Stderr(nil).Start()
}

const lsRegPath string = "/System/Library/Frameworks/CoreServices.framework/Versions/A/Frameworks/LaunchServices.framework/Versions/A/Support/lsregister"

func findVncApp() string {
	output, err := builder.Shell(fmt.Sprintf(`%s -dump | grep -o '/Applications/.*Tiger.*VNC.*\.app' | head -n 1`, lsRegPath)).Output()
	if err != nil {
		panic(err)
	}
	return output
}
