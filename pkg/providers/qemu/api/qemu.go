package api

import (
	"encoding/hex"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"libvirt.org/go/libvirtxml"
	"net/url"
	"time"
)

type QEMU struct {
	libVirtURL string
	conn       *libvirt.Libvirt

	leaseMap     map[string][]string
	allNets      []libvirt.Network
	domains      []*Domain
	storagePools []*StoragePool
}

func Connect(libVirtURL string) (*QEMU, error) {

	if libVirtURL == "" {
		return nil, fmt.Errorf("no libvirt URL found")
	}

	uri, err := url.Parse(libVirtURL)
	if err != nil {
		return nil, err
	}

	conn, err := libvirt.ConnectToURI(uri)
	if err != nil {
		return nil, err
	}
	return &QEMU{
		libVirtURL: libVirtURL,
		conn:       conn,
	}, nil
}

func (q *QEMU) Close() error {
	if q.conn != nil {
		return q.conn.ConnectClose()
	}

	return nil
}

func (q *QEMU) AllNetworks() ([]libvirt.Network, error) {
	if q.allNets == nil {

		allNets, _, err := q.conn.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive)
		if err != nil {
			return nil, err
		}

		q.allNets = allNets
	}

	return q.allNets, nil

}

func (q *QEMU) GetDomains() ([]*Domain, error) {

	if q.domains == nil {

		flags := libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
		domains, _, err := q.conn.ConnectListAllDomains(1, flags)
		if err != nil {
			return nil, err
		}

		q.domains = []*Domain{}
		for _, dom := range domains {
			//state, maxMem, mem, virtCPUs, cpuTime, err := e.conn.DomainGetInfo(dom)
			state, _, _, _, cpuTime, err := q.conn.DomainGetInfo(dom)
			if err != nil {
				log.Debug("error getting libvirt domain info", "err", err)
				continue
			}

			domainState := libvirt.DomainState(state)
			name := dom.Name
			id := fmt.Sprintf("%x", dom.UUID)

			publicIPs := []string{}
			privateIPs := []string{}

			sources := []libvirt.DomainInterfaceAddressesSource{
				libvirt.DomainInterfaceAddressesSrcLease,
				libvirt.DomainInterfaceAddressesSrcAgent,
				libvirt.DomainInterfaceAddressesSrcArp,
			}

			for _, source := range sources {
				ipAddresses, err := q.conn.DomainInterfaceAddresses(dom, uint32(source), 0)
				if err != nil {
					log.Debug("error getting libvirt domain interface addresses", "source", source, "err", err)
				}
				for _, ifaceAddr := range ipAddresses {
					for _, addr := range ifaceAddr.Addrs {
						publicIPs = append(publicIPs, addr.Addr)
					}
				}
			}

			q.domains = append(q.domains, &Domain{
				LVDomain:           &dom,
				LVDomainState:      &domainState,
				Name:               name,
				ID:                 id,
				PublicIPAddresses:  publicIPs,
				PrivateIPAddresses: privateIPs,
				CurrentUpTime:      time.Duration(cpuTime),
			})
		}
	}

	return q.domains, nil
}

func (q *QEMU) GetDomain(id string) (*Domain, error) {
	bytes, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}
	uuid := libvirt.UUID(bytes)
	dom, err := q.conn.DomainLookupByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return &Domain{
		qemu:     q,
		LVDomain: &dom,
	}, nil
}

func (q *QEMU) GetStoragePools() ([]*StoragePool, error) {
	if q.storagePools == nil {
		storagePools, _, err := q.conn.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsActive)
		if err != nil {
			return nil, err
		}

		q.storagePools = []*StoragePool{}
		for _, storagePool := range storagePools {
			q.storagePools = append(q.storagePools, &StoragePool{
				qemu:          q,
				LVStoragePool: &storagePool,
			})
		}
	}

	return q.storagePools, nil
}

func (q *QEMU) GetStoragePool(name string) (*StoragePool, error) {
	storagePools, err := q.GetStoragePools()
	if err != nil {
		return nil, err
	}
	for _, storagePool := range storagePools {
		if storagePool.LVStoragePool.Name == name {
			return storagePool, nil
		}
	}

	return nil, fmt.Errorf("storage pool '%s' not found", name)
}

func newDomainFromVol(vol *StorageVolume) libvirtxml.Domain {

	domainDef := libvirtxml.Domain{
		Type: "kvm",
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type:    "hvm",
				Arch:    "x86_64",
				Machine: "pc",
			},
		},
		Memory: &libvirtxml.DomainMemory{
			Unit:  "MiB",
			Value: 4096,
		},
		VCPU: &libvirtxml.DomainVCPU{
			Placement: "static",
			Value:     1,
		},
		CPU: &libvirtxml.DomainCPU{},
		Devices: &libvirtxml.DomainDeviceList{
			Graphics: []libvirtxml.DomainGraphic{
				{
					Spice: &libvirtxml.DomainGraphicSpice{},
				},
			},
			//Channels: []libvirtxml.DomainChannel{
			//	{
			//		Type: "unix",
			//		Target: &libvirtxml.DomainChannelTarget{
			//			Type: "virtio",
			//			Name: "org.qemu.guest_agent.0",
			//		},
			//	},
			//},
			RNGs: []libvirtxml.DomainRNG{
				{
					Model: "virtio",
					Backend: &libvirtxml.DomainRNGBackend{
						BuiltIn: &libvirtxml.DomainRNGBackendBuiltIn{},
					},
				},
			},
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						Volume: &libvirtxml.DomainDiskSourceVolume{
							Pool:   vol.LVStorageVolume.Pool,
							Volume: vol.LVStorageVolume.Name,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
						Bus: "sata",
					},
					Boot: &libvirtxml.DomainDeviceBoot{
						Order: 1,
					},
				},
			},
			Interfaces: []libvirtxml.DomainInterface{
				{
					//MAC:    &libvirtxml.DomainInterfaceMAC{Address: macAddr()},
					Source: &libvirtxml.DomainInterfaceSource{Network: &libvirtxml.DomainInterfaceSourceNetwork{Network: "default"}},
					Model:  &libvirtxml.DomainInterfaceModel{Type: "virtio"},
				},
			},
		},
		Features: &libvirtxml.DomainFeatureList{
			PAE:  &libvirtxml.DomainFeature{},
			ACPI: &libvirtxml.DomainFeature{},
			APIC: &libvirtxml.DomainFeatureAPIC{},
		},
	}

	return domainDef
}

func (q *QEMU) CreateDomain(name string, storageVol *StorageVolume) (*Domain, error) {

	domXML := newDomainFromVol(storageVol)
	domXML.Name = name

	//domXML.Devices.Disks

	xmldoc, err := domXML.Marshal()
	if err != nil {
		return nil, err
	}

	dom, err := q.conn.DomainCreateXML(xmldoc, libvirt.DomainNone)
	if err != nil {
		return nil, err
	}

	return &Domain{
		qemu:     q,
		LVDomain: &dom,
	}, nil
}
