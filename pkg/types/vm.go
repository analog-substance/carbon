package types

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"github.com/analog-substance/carbon/pkg/ssh_util"
	"github.com/analog-substance/carbon/pkg/vnc_viewer"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
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
	ExecSSH(string, ...string) error
	StartVNC(string) error
	NewSSHSession(string) (*ssh_util.Session, error)
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

func (m Machine) ExecSSH(user string, additionalArgs ...string) error {
	sshPath, _ := exec.LookPath("ssh")
	ip := m.IPAddress()

	args := []string{
		"ssh",
		"-o",
		"StrictHostKeyChecking=no",
		"-o",
		"UserKnownHostsFile=/dev/null",
		fmt.Sprintf("%s@%s", user, ip),
	}

	args = append(args, additionalArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(sshPath, args, os.Environ())
}

func (m Machine) StartVNC(user string) error {
	vncCmd := "ps aux | grep -v grep | grep -i vnc 2>&1 >/dev/null || vncserver -localhost -PasswordFile ~/.vnc/passwd -xstartup xfce4-session 2>&1 >/dev/null; lsof -i -n -o -P | grep -i vnc | grep 127 | cut -d : -f2 | awk '{print $1}'"
	sshSession, err := m.NewSSHSession(user)
	if err != nil {
		return err
	}
	vncPortStr, err := sshSession.Output(vncCmd)
	if err != nil {
		return err
	}

	vncPort, err := strconv.Atoi(vncPortStr)
	if err != nil {
		return err
	}

	localPort := 5901

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	go func() {
		vnc_viewer.Start(vnc_viewer.Options{
			Delay:        3,
			Host:         fmt.Sprintf("127.0.0.1:%d", vncPort),
			PasswordFile: path.Join(home, ".vnc", "passwd"),
		})
	}()

	err = sshSession.ForwardLocalPort(localPort, vncPort)
	if err != nil {
		return err
	}

	return nil
}

func (m Machine) NewSSHSession(user string) (*ssh_util.Session, error) {
	session, err := ssh_util.NewSession()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			session.Close()
		}
	}()

	err = session.Connect(fmt.Sprintf("%s:%d", m.IPAddress(), 22), user)
	if err != nil {
		return nil, err
	}

	return session, nil
}