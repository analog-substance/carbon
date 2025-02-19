//go:build windows

package rdp_client

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
)

const (
	rdpClient = `mstsc`
)

func StartRDPClient(options Options) error {
	log.Debug("attempting to start RDP Client", "os", "windows")

	return builder.Shell(fmt.Sprintf("& '%s' /v:%s > $null 2>&1", rdpClient, options.Host)).Start()

	return fmt.Errorf("unable to find rdp client")
}
