package libvirt

import (
	"encoding/hex"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/digitalocean/go-libvirt"
	"log"
	"time"
)

type environment struct {
	name     string
	platform types.Platform
	conn     *libvirt.Libvirt
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Platform() types.Platform {
	return e.platform
}

func (e environment) VMs() []types.VM {
	var vms []types.VM

	//v, err := e.conn.ConnectGetLibVersion()
	//if err != nil {
	//	log.Fatalf("failed to retrieve libvirt version: %v", err)
	//}
	//fmt.Println("Version:", v)

	flags := libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
	domains, _, err := e.conn.ConnectListAllDomains(1, flags)
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	//if err = l.Disconnect(); err != nil {
	//	log.Fatalf("failed to disconnect: %v", err)
	//}

	allNets, _, err := e.conn.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive)
	//allNets, err := e.conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE)

	leaseMap := map[string][]string{}
	for _, net := range allNets {
		leases, _, err := e.conn.NetworkGetDhcpLeases(net, libvirt.OptString{}, 1, 0)
		if err != nil {
			log.Println("error getting leases", err)
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
			log.Println("error getting libvirt domain info", e.Platform().Name(), err)
			continue
		}

		domainState := libvirt.DomainState(state)
		name := dom.Name
		id := fmt.Sprintf("%x", dom.UUID)

		publicIPs := []string{}
		privateIPs := []string{}

		if domainState == libvirt.DomainRunning {
			ipAddresses, err := e.conn.DomainInterfaceAddresses(dom, 0, 0)
			//ipAddresses, err := dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
			if err != nil {
				log.Println("error getting librt domain interfaces", e.Platform().Name(), err)
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
		vms = append(vms, models.Machine{
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

func (e environment) ImageBuilds() []types.ImageBuild {
	return []types.ImageBuild{}
}
func (e environment) Images() []types.Image {
	return []types.Image{}
}

func stateFromVboxInfo(state libvirt.DomainState) types.MachineState {
	//DomainNostate     DomainState = iota
	//DomainRunning     DomainState = 1
	//DomainBlocked     DomainState = 2
	//DomainPaused      DomainState = 3
	//DomainShutdown    DomainState = 4
	//DomainShutoff     DomainState = 5
	//DomainCrashed     DomainState = 6
	//DomainPmsuspended DomainState = 7

	if state == libvirt.DomainPaused {
		return types.StateSleeping
	}
	if state == libvirt.DomainShutoff {
		return types.StateStopped
	}
	if state == libvirt.DomainRunning {
		return types.StateRunning
	}

	log.Println("unknown libvirt state", state)
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
