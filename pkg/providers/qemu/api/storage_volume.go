package api

import "github.com/digitalocean/go-libvirt"

type StorageVolume struct {
	qemu            *QEMU
	LVStorageVolume *libvirt.StorageVol
}
