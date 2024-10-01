package qemu

import (
	"encoding/hex"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
	"os"
	"path"
	"time"
)

type environment struct {
	name    string
	profile types.Profile
	conn    *libvirt.Libvirt
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Profile() types.Profile {
	return e.profile
}

func (e environment) VMs() []types.VM {
	var vms []types.VM
	flags := libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
	domains, _, err := e.conn.ConnectListAllDomains(1, flags)
	if err != nil {
		log.Debug("failed to retrieve domains", "profile", e.Profile().Name(), "err", err)
		os.Exit(3)
	}

	allNets, _, err := e.conn.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive)

	leaseMap := map[string][]string{}
	for _, net := range allNets {
		leases, _, err := e.conn.NetworkGetDhcpLeases(net, libvirt.OptString{}, 1, 0)
		if err != nil {
			log.Debug("error getting leases domain info", "profile", e.Profile().Name(), "err", err)
			continue
		}
		for _, lease := range leases {
			for _, hostname := range lease.Hostname {
				_, ok := leaseMap[hostname]
				if !ok {
					leaseMap[hostname] = []string{}
				}
				leaseMap[hostname] = append(leaseMap[hostname], lease.Ipaddr)
			}
		}
	}

	for _, dom := range domains {
		//state, maxMem, mem, virtCPUs, cpuTime, err := e.conn.DomainGetInfo(dom)
		state, _, _, _, cpuTime, err := e.conn.DomainGetInfo(dom)
		if err != nil {
			log.Debug("error getting libvirt domain info", "profile", e.Profile().Name(), "err", err)
			continue
		}

		domainState := libvirt.DomainState(state)
		name := dom.Name
		id := fmt.Sprintf("%x", dom.UUID)

		publicIPs := []string{}
		privateIPs := []string{}

		if domainState == libvirt.DomainRunning {
			ipAddresses, err := e.conn.DomainInterfaceAddresses(dom, 0, 0)
			if err != nil {
				log.Debug("error getting librt domain interfaces", "profile", e.Profile().Name(), "err", err)
				continue
			}

			if len(ipAddresses) > 0 {
				for _, ifaceAddr := range ipAddresses {
					for _, addr := range ifaceAddr.Addrs {
						publicIPs = append(publicIPs, addr.Addr)
					}
				}
			} else {
				// fallback to look up in default dhcp lease
				_, ok := leaseMap[name]
				if ok {
					publicIPs = leaseMap[name]
				}
			}
		}
		vms = append(vms, &models.Machine{
			InstanceName:       name,
			CurrentState:       stateFromVboxInfo(domainState),
			InstanceID:         fmt.Sprintf("%s", id),
			Env:                e,
			PublicIPAddresses:  publicIPs,
			PrivateIPAddresses: privateIPs,
			CurrentUpTime:      time.Duration(cpuTime),
		})
	}

	return vms
}

func (e environment) StartVM(id string) error {
	dom, err := e.domainFromUUID(id)
	if err != nil {
		return err
	}
	err = e.conn.DomainResume(dom)
	if err != nil {
		err = e.conn.DomainCreate(dom)
	}
	return err
}

func (e environment) StopVM(id string) error {
	dom, err := e.domainFromUUID(id)
	if err != nil {
		return err
	}
	err = e.conn.DomainSuspend(dom)
	return err
}

func (e environment) RestartVM(id string) error {
	dom, err := e.domainFromUUID(id)
	if err != nil {
		return err
	}
	err = e.conn.DomainReboot(dom, libvirt.DomainRebootDefault)
	return err
}

func (e environment) DestroyVM(id string) error {
	dom, err := e.domainFromUUID(id)
	if err != nil {
		return err
	}
	err = e.conn.DomainDestroy(dom)
	return err
}

func (e environment) CreateVM(options types.MachineLaunchOptions) error {
	return nil
}

func (e environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e environment) Images() ([]types.Image, error) {
	ret := []types.Image{}
	listing, _ := os.ReadDir("deployments/images/qemu")
	for _, dirEntry := range listing {
		ret = append(ret, models.NewImage(path.Join("deployments/images/qemu", dirEntry.Name()), dirEntry.Name(), time.Now(), e))

	}
	return ret, nil
}
func stateFromVboxInfo(state libvirt.DomainState) types.MachineState {
	if state == libvirt.DomainPaused {
		return types.StateSleeping
	}
	if state == libvirt.DomainShutoff {
		return types.StateStopped
	}
	if state == libvirt.DomainRunning {
		return types.StateRunning
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}

func (e environment) domainFromUUID(id string) (libvirt.Domain, error) {
	bytes, err := hex.DecodeString(id)
	if err != nil {
		return libvirt.Domain{}, err
	}
	uuid := libvirt.UUID(bytes)
	dom, err := e.conn.DomainLookupByUUID(uuid)
	if err != nil {
		return libvirt.Domain{}, err
	}
	return dom, err
}
