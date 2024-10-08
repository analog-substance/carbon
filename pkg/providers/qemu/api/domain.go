package api

import (
	"github.com/digitalocean/go-libvirt"
	"time"
)

type Domain struct {
	qemu     *QEMU
	LVDomain *libvirt.Domain

	LVDomainState *libvirt.DomainState
	ID            string
	Name          string

	PublicIPAddresses  []string
	PrivateIPAddresses []string
	CurrentUpTime      time.Duration
}

func (d *Domain) Start() error {
	err := d.qemu.conn.DomainResume(*d.LVDomain)
	if err != nil {
		err = d.qemu.conn.DomainCreate(*d.LVDomain)
	}
	return err
}

func (d *Domain) Reboot() error {
	return d.qemu.conn.DomainReboot(*d.LVDomain, libvirt.DomainRebootDefault)
}

func (d *Domain) Suspend() error {
	return d.qemu.conn.DomainSuspend(*d.LVDomain)
}

func (d *Domain) Destroy() error {
	return d.qemu.conn.DomainDestroy(*d.LVDomain)
}
