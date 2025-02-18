//go:build darwin

package rdp_client

import (
	"fmt"
	"os"
	"time"

	builder "github.com/NoF0rte/cmd-builder"
)

var fileContentTmpl = "full address:s:%s\nusername:s:%s\n"

func StartRDPClient(options Options) error {
	if options.User == "" {
		options.User = "administrator"
	}

	log.Debug("attempting to start RDP Client", "os", "mac")
	tmpFile, err := os.CreateTemp("", "carbon_*.rdp")
	if err != nil {
		return err
	}
	rdpFilePath := tmpFile.Name()
	defer os.Remove(rdpFilePath)
	fmt.Fprintf(tmpFile, fileContentTmpl, options.Host, options.User)
	tmpFile.Close()

	log.Debug("attempting to open rdp file", "rdpFilePath", rdpFilePath)
	cmdRes := builder.Cmd("/usr/bin/open", rdpFilePath).Stderr(nil).Start()
	time.Sleep(500 * time.Millisecond)

	return cmdRes
}
