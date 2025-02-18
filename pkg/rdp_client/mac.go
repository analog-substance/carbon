//go:build darwin

package rdp_client

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"io/ioutil"
	"os"
	"path/filepath"
)

var fileContentTmpl = "full address:s:%s\n"

func StartRDPClient(options Options) error {
	log.Debug("attempting to start RDP Client", "os", "mac")
	tmpDir, err := ioutil.TempDir("", "carbon")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	rdpFilePath := filepath.Join(tmpDir, "rdp_client.rb")
	rdpFile, err := os.Open(rdpFilePath)
	if err != nil {
		return err
	}
	fmt.Fprintf(rdpFile, fileContentTmpl, options.Host)

	log.Debug("attempting to open rdp file", "rdpFilePath", rdpFilePath)
	return builder.Cmd("open", rdpFilePath).Stderr(nil).Start()

}
