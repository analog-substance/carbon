package vnc_viewer

import (
	"time"
)

type Options struct {
	Delay        int
	PasswordFile string
	Host         string
}

func Start(options Options) error {
	if options.Delay > 0 {
		time.Sleep(time.Duration(options.Delay) * time.Second)
	}

	return StartViewer(options)
}