package api

import (
	builder "github.com/NoF0rte/cmd-builder"
	"log"
	"strings"
)

type VBoxVM struct {
	Name    string
	ID      string
	State   string
	GuestOS string

	vmInfo map[string]string
}

func (v *VBoxVM) loadInfo() error {
	vmInfo, err := builder.
		Cmd("vboxmanage", "showvminfo", "--machinereadable", v.ID).
		Output()

	if err != nil {
		return err
	}
	v.vmInfo = map[string]string{}
	for _, line := range strings.Split(vmInfo, "\n") {
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			v.vmInfo[strings.ToLower(strings.Trim(parts[0], "\""))] = strings.Trim(parts[1], "\"")
		}
	}
	v.GuestOS = v.vmInfo["ostype"]
	v.Name = v.vmInfo["name"]
	v.State = v.vmInfo["vmstate"]

	return nil
}

func ListVMs() []VBoxVM {
	output, err := builder.
		Cmd("vboxmanage", "list", "vms").
		Output()

	if err != nil {
		log.Println("Error listing VMs:", err)
	}

	vms := []VBoxVM{}
	for _, vmLine := range strings.Split(string(output), "\n") {
		vmInfo := strings.Split(vmLine, " ")
		vmID := strings.Trim(vmInfo[1], "{}")
		vm := VBoxVM{ID: vmID}
		err = vm.loadInfo()
		if err != nil {
			log.Println("Error loading VM:", err)
		}
		vms = append(vms, vm)
	}

	return vms
}

func StartVM(id string) error {
	_, err := builder.
		Cmd("vboxmanage", "startvm", id).
		Output()
	return err
}

func RestartVM(id string) error {
	_, err := builder.
		Cmd("vboxmanage", "controlvm", id, "reboot").
		Output()
	return err
}

func SleepVM(id string) error {
	_, err := builder.
		Cmd("vboxmanage", "controlvm", id, "savestate").
		Output()
	return err
}
