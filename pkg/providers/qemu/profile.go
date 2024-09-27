package qemu

import (
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
	"log"
	"net/url"
	"slices"
)

type profile struct {
	profileName string
	provider    *provider
	connectStr  string
}

const profileName = "qemu"

func (p profile) Environments(validNames ...string) []types.Environment {
	if len(validNames) == 0 || slices.Contains(validNames, profileName) {

		uri, _ := url.Parse(p.connectStr)
		conn, err := libvirt.ConnectToURI(uri)
		if err == nil {
			return []types.Environment{environment{
				profileName,
				p,
				conn,
			}}
		} else {
			log.Println("Error connecting to libvirt host", err)
		}
	}
	return []types.Environment{}
}

func (p profile) Name() string {
	return p.profileName
}

func (p profile) Provider() types.Provider {
	return p.provider
}
