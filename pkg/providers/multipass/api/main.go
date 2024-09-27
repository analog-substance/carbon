package api

import (
	"encoding/json"
	builder "github.com/NoF0rte/cmd-builder"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "[Mutlipass]", 0)
}

type MultipassVM struct {
	Ipv4    []string `json:"ipv4"`
	Name    string   `json:"name"`
	Release string   `json:"release"`
	State   string   `json:"state"`
}

type MultipassListOutput struct {
	List []MultipassVM `json:"list"`
}

var appPath string

func AppPath() string {
	if appPath == "" {
		appPath, err := exec.LookPath("multipass")
		if err == nil {
			absPath, err := filepath.Abs(appPath)
			if err == nil {
				return absPath
			}
		}
	}
	return appPath
}

func ListVMs() []MultipassVM {
	output, err := builder.
		Cmd(AppPath(), "list", "--format", "json").
		Output()

	if err != nil {
		log.Println("Error listing Multipass VMs:", err)
	}

	var listOutput MultipassListOutput
	err = json.Unmarshal([]byte(output), &listOutput)
	if err == nil {
		return listOutput.List
	}
	return nil
}

func StartVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "start", id).
		Output()
	return err
}

func RestartVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "restart", id).
		Output()
	return err
}

func SleepVM(id string) error {
	_, err := builder.
		Cmd(AppPath(), "suspend", id).
		Output()
	return err
}
