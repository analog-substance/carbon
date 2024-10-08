package api

import (
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"io"
	"libvirt.org/go/libvirtxml"
	"os"
	"strconv"
	"strings"
	"time"
)

type StoragePool struct {
	qemu          *QEMU
	LVStoragePool *libvirt.StoragePool
	Volumes       []*StorageVolume
}

func (s *StoragePool) GetVolumes() ([]*StorageVolume, error) {
	if s.Volumes == nil {
		lspv, _, err := s.qemu.conn.StoragePoolListAllVolumes(*s.LVStoragePool, 1, 0)
		if err != nil {
			return nil, err
		}

		s.Volumes = []*StorageVolume{}
		for _, vol := range lspv {
			s.Volumes = append(s.Volumes, &StorageVolume{
				qemu:            s.qemu,
				LVStorageVolume: &vol,
			})
		}
	}

	return s.Volumes, nil
}

func (s *StoragePool) ImportImage(name string, imageFile string) (*StorageVolume, error) {

	fi, err := os.Stat(imageFile)
	if err != nil {
		return nil, err
	}
	size := uint64(fi.Size())

	storageVolumeXML := &libvirtxml.StorageVolume{
		Type: "file",
		Name: name,
		Target: &libvirtxml.StorageVolumeTarget{
			Path: s.LVStoragePool.Name + name,
			Format: &libvirtxml.StorageVolumeTargetFormat{
				Type: "qcow2",
			},
		},
		Capacity: &libvirtxml.StorageVolumeSize{
			Unit:  "B",
			Value: size,
		},
	}

	xmldoc, err := storageVolumeXML.Marshal()
	if err != nil {
		return nil, err
	}

	volume, err := s.qemu.conn.StorageVolCreateXML(*s.LVStoragePool, xmldoc, 0)
	if err != nil {
		return nil, err
	}

	err = importImage(imageFile, newCopier(s.qemu.conn, volume, size), *storageVolumeXML)
	if err != nil {
		return nil, err
	}
	return &StorageVolume{
		qemu:            s.qemu,
		LVStorageVolume: &volume,
	}, nil
}

func newCopier(virConn *libvirt.Libvirt, volume libvirt.StorageVol, size uint64) func(src io.Reader) error {
	copier := func(src io.Reader) error {
		return virConn.StorageVolUpload(volume, src, 0, size, 0)
	}
	return copier
}

func importImage(imagePath string, copier func(io.Reader) error, vol libvirtxml.StorageVolume) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("error while opening %s: %w", imagePath, err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}
	// we can skip the upload if the modification times are the same
	if vol.Target.Timestamps != nil && vol.Target.Timestamps.Mtime != "" {
		if fi.ModTime() == timeFromEpoch(vol.Target.Timestamps.Mtime) {
			log.Debug("Modification time is the same: skipping image copy")
			return nil
		}
	}

	return copier(file)
}

func timeFromEpoch(str string) time.Time {
	var s, ns int
	var err error

	ts := strings.Split(str, ".")
	if len(ts) == 2 {
		ns, err = strconv.Atoi(ts[1])
		if err != nil {
			ns = 0
		}
	}
	s, err = strconv.Atoi(ts[0])
	if err != nil {
		s = 0
	}

	return time.Unix(int64(s), int64(ns))
}
