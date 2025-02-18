package rdp_client

import (
	"os/exec"
	"time"
)

type Options struct {
	Delay int
	User  string
	Host  string
}

func Start(options Options) error {
	if options.Delay > 0 {
		time.Sleep(time.Duration(options.Delay) * time.Second)
	}

	return StartRDPClient(options)
}

func getRDPClientPath() string {
	executablePath, err := exec.LookPath("rdesktop")
	if err == nil {
		return executablePath
	}

	return ""
}
