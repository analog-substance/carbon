//go:build windows

package vnc_viewer

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"os"
	"path/filepath"
)

const (
	tigerVNC = `TigerVNC\vncviewer.exe`
)

func StartViewer(options Options) error {
	vncViewerPath := findVNCViewerExecutable()
	if vncViewerPath != "" {
		return builder.Shell(fmt.Sprintf("& '%s' %s %s > $null 2>&1", vncViewerPath, options.PasswordFile, options.Host)).Start()
	}

	return fmt.Errorf("unable to find vncviewer")
}

func findVNCViewerExecutable() string {
	vncPaths := []string{
		filepath.Join(os.Getenv("programfiles"), tigerVNC),
		filepath.Join(os.Getenv("programfiles(x86)"), tigerVNC),
	}

	for _, vncPath := range vncPaths {
		_, err := os.Stat(vncPath)
		if err == nil {
			return vncPath
		}
	}

	panic("TigerVNC not found!")
}
