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
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Machine struct {
	InstanceName       string             `json:"name"`
	InstanceID         string             `json:"id"`
	CurrentUpTime      time.Duration      `json:"up_time"`
	InstanceType       string             `json:"type"`
	PublicIPAddresses  []string           `json:"public_ip_addresses"`
	PrivateIPAddresses []string           `json:"private_ip_addresses"`
	CurrentState       types.MachineState `json:"current_state"`
	Env                types.Environment  `json:"-"`
}

func (m *Machine) Environment() types.Environment {
	return m.Env
}

func (m *Machine) Profile() types.Profile {
	return m.Env.Profile()
}

func (m *Machine) Provider() types.Provider {
	return m.Env.Profile().Provider()
}

func (m *Machine) Name() string {
	return m.InstanceName
}

func (m *Machine) ID() string {
	return m.InstanceID
}

func (m *Machine) IPAddress() string {
	if len(m.PublicIPAddresses) > 0 {
		return m.PublicIPAddresses[0]
	}
	return "unknown"
}

func (m *Machine) UpTime() time.Duration {
	return m.CurrentUpTime
}

func (m *Machine) State() string {
	return m.CurrentState.Name
}

func (m *Machine) Type() string {
	return m.InstanceType
}

func (m *Machine) Destroy() error {
	return m.Env.DestroyVM(m.InstanceID)
}

func (m *Machine) Start() error {
	return m.Env.StartVM(m.InstanceID)
}

func (m *Machine) Stop() error {
	return m.Env.StopVM(m.InstanceID)
}

func (m *Machine) Restart() error {
	return m.Env.RestartVM(m.InstanceID)
}

func (m *Machine) ExecSSH(user string, cmdArgs ...string) error {
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

	args = append(args, cmdArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(sshPath, args, os.Environ())
}

func (m *Machine) StartVNC(user string, killVNC bool) error {
	sshSession, err := m.NewSSHSession(user)
	if err != nil {
		return err
	}

	vncCmd := []string{}

	if killVNC {
		vncCmd = append(vncCmd, "killall vncserver > /dev/null 2>&1;")
	}

	vncCmd = append(vncCmd, "if ! ps aux | grep -v grep | grep -i vnc 2>&1 >/dev/null  ; then if [ ! -f ~/.vnc/passwd ]; then mkdir -p ~/.vnc; echo -n carbon | vncpasswd -f > ~/.vnc/passwd; fi ; vncserver -localhost -PasswordFile ~/.vnc/passwd -xstartup xfce4-session 2>&1 >/dev/null; fi; cat ~/.vnc/passwd | base64 -w0; echo ; lsof -i -n -o -P | grep -i vnc | grep 127 | cut -d : -f2 | awk '{print $1}' ; ")

	sshSession.Session.Stdout = nil
	vncConfig, err := sshSession.Output(strings.Join(vncCmd, " "))
	if err != nil {
		return err
	}
	log.Debug("vnc conf", "vncConfig", vncConfig, "machine", m.Name())

	vncConfigSlice := strings.Split(vncConfig, "\n")
	if len(vncConfigSlice) != 2 {
		return fmt.Errorf("invalid vnc config %s", vncConfig)
	}
	passwdB64 := vncConfigSlice[0]
	vncPortStr := vncConfigSlice[1]

	vncPassFile, err := m.setVNCPasswd(passwdB64)
	if err != nil {
		return err
	}

	vncPort, err := strconv.Atoi(vncPortStr)
	if err != nil {
		return err
	}

	localPort := 5901

	go func() {
		log.Debug("start vncviewer", "machine", m.Name())
		err := vnc_viewer.Start(vnc_viewer.Options{
			Delay:        3,
			Host:         fmt.Sprintf("127.0.0.1:%d", vncPort),
			PasswordFile: vncPassFile,
		})
		if err != nil {
			fmt.Println("unable start vncviewer. Is it installed?")
			fmt.Printf("try running\n\tvncviewer localhost:%d\n", localPort)
			fmt.Println("You can install it using your systems package manager or")
			fmt.Println("download a release from: https://github.com/TigerVNC/tigervnc/releases")
		}
	}()
	log.Debug("fwd port", "localPort", localPort, "vncPort", vncPort)
	err = sshSession.ForwardLocalPort(localPort, vncPort)
	if err != nil {
		return err
	}

	return nil
}

func (m *Machine) NewSSHSession(user string) (*ssh_util.Session, error) {
	// Get SSH_AUTH_SOCK
	// if it is not set, look at .ssh/config for IdentityAgent
	// use that value
	sshAuthSock := os.Getenv("SSH_AUTH_SOCK")
	if sshAuthSock == "" {
		log.Debug("no SSH_AUTH_SOCK environment variable detected")
		sshAuthSock = getAuthSockFromConfig()
		if sshAuthSock != "" {
			log.Debug("found SSH_AUTH_SOCK in ssh config", "sshAuthSock", sshAuthSock)
			err := os.Setenv("SSH_AUTH_SOCK", sshAuthSock)
			if err != nil {
				log.Debug("cant set SSH_AUTH_SOCK")
			}
		}
	}

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

func (m *Machine) Cmd(user string, cmdArgs ...string) (string, error) {
	sshSession, err := m.NewSSHSession(user)
	if err != nil {
		return "", err
	}
	defer sshSession.Close()

	//sshSession.Session.Stdout = nil
	return sshSession.Output(strings.Join(cmdArgs, " "))
}

func (m *Machine) setVNCPasswd(vncPasswordB64 string) (string, error) {

	passwdBytes, err := base64.StdEncoding.DecodeString(vncPasswordB64)
	if err != nil {
		return "", err
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Join(home, ".vnc"), 0700)
	if err != nil {
		return "", err
	}

	vncPasswdPath := filepath.Join(home, ".vnc", "carbon-passwd")
	err = os.WriteFile(vncPasswdPath, passwdBytes, 0600)
	if err != nil {
		return "", err
	}

	return vncPasswdPath, nil
}

func getAuthSockFromConfig() string {
	home, err := homedir.Dir()
	if err != nil {
		return ""
	}
	fileContents, err := os.ReadFile(filepath.Join(home, ".ssh", "config"))
	if err != nil {
		return ""
	}

	idRe := regexp.MustCompile(`\s*IdentityAgent "?([^"]+)"?`)
	homeDirRe := regexp.MustCompile(`^~`)
	lines := strings.Split(string(fileContents), "\n")
	for _, line := range lines {
		if strings.Contains(line, "IdentityAgent") {
			agentSock := idRe.ReplaceAllString(line, "$1")
			agentSock = homeDirRe.ReplaceAllString(agentSock, home)
			return agentSock
		}
	}
	return ""
}
