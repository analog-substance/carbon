package types

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

type MachineState struct{ Name string }

var StateRunning = MachineState{"Running"}
var StateOff = MachineState{"Off"}
var StateStarting = MachineState{"Starting"}
var StateStopping = MachineState{"Stopping"}
var StateSleeping = MachineState{"Sleeping"}
var StateTerminating = MachineState{"Terminating"}
var StateUnknown = MachineState{"Unknown"}

type VM interface {
	Name() string
	ID() string
	IPAddress() string
	State() string
	Start() error
	Stop() error
	Restart() error
	Environment() Environment
	ExecSSH(string) error
}

type Machine struct {
	InstanceName       string
	InstanceID         string
	CurrentUpTime      time.Duration
	PublicIPAddresses  []string
	PrivateIPAddresses []string
	CurrentState       MachineState
	Env                Environment
}

func (m Machine) Environment() Environment {
	return m.Env
}

func (m Machine) Name() string {
	return m.InstanceName
}

func (m Machine) ID() string {
	return m.InstanceID
}

func (m Machine) IPAddress() string {
	if len(m.PublicIPAddresses) > 0 {
		return m.PublicIPAddresses[0]
	}
	return "unknown"
}

func (m Machine) State() string {
	return m.CurrentState.Name
}

func (m Machine) Start() error {
	return m.Env.StartVM(m.InstanceID)
}

func (m Machine) Stop() error {
	return m.Env.StopVM(m.InstanceID)
}

func (m Machine) Restart() error {
	return m.Env.RestartVM(m.InstanceID)
}

func (m Machine) ExecSSH(user string) error {
	path, _ := exec.LookPath("ssh")
	ip := m.IPAddress()

	args := []string{
		"ssh",
		"-o",
		"StrictHostKeyChecking=no",
		"-o",
		"UserKnownHostsFile=/dev/null",
		//"-p",
		//strconv.Itoa(m.SSHPort),
		fmt.Sprintf("%s@%s", user, ip),
	}

	//args = append(args, additionalArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(path, args, os.Environ())
}
