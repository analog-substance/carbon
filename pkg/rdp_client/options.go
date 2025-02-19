package rdp_client

import (
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
