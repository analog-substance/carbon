package models

import (
	"encoding/base64"
	"fmt"
	"github.com/NoF0rte/cmd-builder"
	"github.com/analog-substance/carbon/pkg/ssh_util"
	"github.com/analog-substance/carbon/pkg/types"
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

type Machine struct {
	InstanceName       string
	InstanceID         string
	CurrentUpTime      time.Duration
	PublicIPAddresses  []string
	PrivateIPAddresses []string
	CurrentState       types.MachineState
	Env                types.Environment
}

func (m Machine) Environment() types.Environment {
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
	sshSession, err := m.NewSSHSession(user)
	if err != nil {
		return err
	}

	vncPassFile, err := m.getVNCPasswd(sshSession)
	if err != nil {
		return err
	}

	vncCmd := "if ! ps aux | grep -v grep | grep -i vnc 2>&1 >/dev/null  ; then vncserver -localhost -PasswordFile ~/.vnc/passwd -xstartup xfce4-session 2>&1 >/dev/null; fi; lsof -i -n -o -P | grep -i vnc | grep 127 | cut -d : -f2 | awk '{print $1}'"

	vncPortStr, err := sshSession.Output(vncCmd)
	if err != nil {
		return err
	}

	vncPort, err := strconv.Atoi(vncPortStr)
	if err != nil {
		return err
	}

	localPort := 5901

	go func() {
		vnc_viewer.Start(vnc_viewer.Options{
			Delay:        3,
			Host:         fmt.Sprintf("127.0.0.1:%d", vncPort),
			PasswordFile: vncPassFile,
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

func (m Machine) getVNCPasswd(sshSession *ssh_util.Session) (string, error) {

	vncPasswordB64, err := sshSession.Output("if [ ! -f ~/.vnc/passwd ]; then mkdir -p ~/.vnc; echo -n carbon | vncpasswd -f > ~/.vnc/passwd; fi ; cat ~/.vnc/passwd | base64 -w0")
	if err != nil {
		return "", err
	}

	passwdBytes, err := base64.StdEncoding.DecodeString(vncPasswordB64)
	if err != nil {
		return "", err
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(path.Join(home, ".vnc"), 0700)
	if err != nil {
		return "", err
	}

	vncPasswdPath := path.Join(home, ".vnc", "carbon-passwd")
	err = os.WriteFile(vncPasswdPath, passwdBytes, 0600)
	if err != nil {
		return "", err
	}

	return vncPasswdPath, nil
}
