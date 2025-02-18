//go:build linux

package rdp_client

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

func StartRDPClient(options Options) error {
	log.Debug("attempting to start RDP Client", "os", "linux")

	rdpClientPath := getRDPClientPath()
	if rdpClientPath != "" {
		log.Debug("rdp client found in applications", "os", "linux", "rdpClientPath", rdpClientPath)
		return builder.Cmd(rdpClientPath, options.Host).Start()
	}

	return fmt.Errorf("unable to find vncviewer")
}

func getRDPClientPath() string {
	executablePath, err := exec.LookPath("rdesktop")
	if err == nil {
		return executablePath
	}

	return ""
}
