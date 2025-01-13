package api

import (
	builder "github.com/NoF0rte/cmd-builder"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type VBoxVM struct {
	Name    string
	ID      string
	State   string
	GuestOS string

	vmInfo map[string]string
}

var appPath string

func AppPath() string {
	if appPath == "" {

		vboxExecName := "vboxmanage"
		if //goland:noinspection GoBoolExpressions
		runtime.GOOS == "windows" {
			vboxExecName = "VBoxManage.exe"
		}

		virtualBox, err := exec.LookPath(vboxExecName)
		if err == nil {
			appPath, err = filepath.Abs(virtualBox)
			if err != nil {
				log.Debug("err getting absolute path", "virtualBox", virtualBox, "err", err)
			}
		} else {
			if //goland:noinspection GoBoolExpressions
			runtime.GOOS == "windows" {
				// not in path, lets look in program files
				vBoxManage := `Oracle\VirtualBox\VBoxManage.exe`
				vboxManagePaths := []string{
					filepath.Join(os.Getenv("programfiles"), vBoxManage),
					filepath.Join(os.Getenv("programfiles(x86)"), vBoxManage),
				}

				for _, vboxManagePath := range vboxManagePaths {
					_, err := os.Stat(vboxManagePath)
					if err == nil {
						appPath = vboxManagePath
						break
					}
				}
			}
		}
	}

	return appPath
}

func (v *VBoxVM) loadInfo() error {
	vmInfo, err := builder.
		Cmd(AppPath(), "showvminfo", "--machinereadable", v.ID).
		Output()

	if err != nil {
		return err
	}
	v.vmInfo = map[string]string{}
	for _, line := range strings.Split(vmInfo, "\n") {
		line = strings.TrimSpace(line)
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
		Cmd(AppPath(), "list", "vms").
		Output()

	if err != nil {
		log.Debug("error listing VMs:", "err", err)
	}

	vms := []VBoxVM{}
	for _, vmLine := range strings.Split(string(output), "\n") {
		vmLine = strings.TrimSpace(vmLine)
		if strings.Contains(vmLine, " ") {
			vmInfo := strings.Split(vmLine, " ")
			vmID := strings.Trim(vmInfo[1], "{}")
			vm := VBoxVM{ID: vmID}
			err = vm.loadInfo()
			if err != nil {
				log.Debug("error loading VM", "err", err)
			}
			vms = append(vms, vm)
		}
	}

	return vms
}

func StartVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "startvm", id).
		Output()
	return err
}

func RestartVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "controlvm", id, "reboot").
		Output()
	return err
}

func SleepVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "controlvm", id, "savestate").
		Output()
	return err
}
