package libvirt

import (
	"fmt"
	types2 "github.com/analog-substance/carbon/pkg/types"
	"libvirt.org/go/libvirt"
	"log"
	"strconv"
)

type environment struct {
	name     string
	platform types2.Platform
	conn     *libvirt.Connect
}

func (e environment) Name() string {
	return e.name
}

func (e environment) Platform() types2.Platform {
	return e.platform
}

func (e environment) VMs() []types2.VM {
	var vms []types2.VM
	doms, err := e.conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Println("error getting librt domains", e.Platform().Name(), err)
		return vms
	}

	allNets, err := e.conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE)

	leaseMap := map[string][]string{}
	for _, net := range allNets {
		lease, err := net.GetDHCPLeases()
		if err != nil {
			log.Println("error getting leases", err)
		}
		for _, lease := range lease {
			_, ok := leaseMap[lease.Hostname]
			if !ok {
				leaseMap[lease.Hostname] = []string{}
			}
			leaseMap[lease.Hostname] = append(leaseMap[lease.Hostname], lease.IPaddr)
		}
	}

	for _, dom := range doms {
		info, err := dom.GetInfo()
		if err != nil {
			log.Println("error getting librt domain details", e.Platform().Name(), err)
			continue
		}

		name, err := dom.GetName()
		if err != nil {
			log.Println("error getting librt domain details", e.Platform().Name(), err)
			continue
		}
		id, err := dom.GetID()
		if err != nil {
			log.Println("error getting librt domain details", e.Platform().Name(), err)
			continue
		}

		publicIPs := []string{}
		privateIPs := []string{}

		ipAddresses, err := dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
		if err != nil {
			log.Println("error getting librt domain details", e.Platform().Name(), err)
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

		vms = append(vms, types2.Machine{
			InstanceName:       name,
			CurrentState:       stateFromVboxInfo(info.State),
			InstanceID:         fmt.Sprintf("%d", id),
			Env:                e,
			PublicIPAddresses:  publicIPs,
			PrivateIPAddresses: privateIPs,
		})
		err = dom.Free()
		if err != nil {
			log.Println("error freeing libvirt domain", e.Platform().Name(), err)
		}
	}

	return vms
}

func (e environment) StartVM(id string) error {
	newInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	dom, err := e.conn.LookupDomainById(uint32(newInt))
	if err != nil {
		return err
	}

	err = dom.Resume()
	return err
}

func (e environment) StopVM(id string) error {
	// sleep
	newInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	dom, err := e.conn.LookupDomainById(uint32(newInt))
	if err != nil {
		return err
	}
	err = dom.Suspend()
	return err
}

func (e environment) RestartVM(id string) error {
	newInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	dom, err := e.conn.LookupDomainById(uint32(newInt))
	if err != nil {
		return err
	}

	err = dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT)
	return err
}

func stateFromVboxInfo(state libvirt.DomainState) types2.MachineState {
	if state == libvirt.DOMAIN_PMSUSPENDED {
		return types2.StateSleeping
	}
	if state == libvirt.DOMAIN_SHUTOFF {
		return types2.StateOff
	}
	if state == libvirt.DOMAIN_RUNNING {
		return types2.StateRunning
	}
	return types2.StateUnknown
}
