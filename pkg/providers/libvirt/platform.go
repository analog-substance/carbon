package libvirt

import (
	"github.com/analog-substance/carbon/pkg/types"
	"libvirt.org/go/libvirt"
	"log"
	"slices"
)

type platform struct {
	profileName string
	provider    *provider
	connectStr  string
}

const platformName = "qemu"

func (p platform) Environments(validNames ...string) []types.Environment {
	if len(validNames) == 0 || slices.Contains(validNames, platformName) {

		conn, err := libvirt.NewConnect(p.connectStr)
		if err == nil {
			return []types.Environment{environment{
				platformName,
				p,
				conn,
			}}
		} else {
			log.Println("Error connecting to libvirt host", err)
		}
	}
	return []types.Environment{}
}

func (p platform) Name() string {
	return p.profileName
}

func (p platform) Provider() types.Provider {
	return p.provider
}
