//go:build linux

package rdp_client

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

func StartRDPClient(options Options) error {
	rdpClientPath := getRDPClientPath()
	if rdpClientPath != "" {
		log.Debug("rdp client found in applications", "os", "linux", "rdpClientPath", rdpClientPath)
		return builder.Cmd(rdpClientPath, options.Host).Start()
	}

	return fmt.Errorf("unable to find vncviewer")
}
