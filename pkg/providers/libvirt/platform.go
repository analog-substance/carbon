package libvirt

import (
	types2 "github.com/analog-substance/carbon/pkg/types"
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

func (p platform) Environments(validNames ...string) []types2.Environment {
	if len(validNames) == 0 || slices.Contains(validNames, platformName) {

		conn, err := libvirt.NewConnect(p.connectStr)
		if err == nil {
			return []types2.Environment{environment{
				platformName,
				p,
				conn,
			}}
		} else {
			log.Println("Error connecting to libvirt host", err)
		}
	}
	return []types2.Environment{}
}

func (p platform) Name() string {
	return p.profileName
}

func (p platform) Provider() types2.Provider {
	return p.provider
}
