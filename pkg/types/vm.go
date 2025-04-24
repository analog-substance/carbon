package types

import (
	"github.com/analog-substance/carbon/pkg/ssh_util"
	"time"
)

type MachineState struct {
	Name string `json:"name"`
}

var StateRunning = MachineState{"Running"}
var StateStopped = MachineState{"Stopped"}
var StateStarting = MachineState{"Starting"}
var StateStopping = MachineState{"Stopping"}
var StateSleeping = MachineState{"Sleeping"}
var StateTerminating = MachineState{"Terminating"}
var StateTerminated = MachineState{"Terminated"}
var StateUnknown = MachineState{"Unknown"}

// VM interface provides access to useful information and actions related to Virtual Machines
type VM interface {
	// Name returns the name of a virtual machine
	Name() string

	// ID returns the ID of the virtual machine
	ID() string

	// IPAddress returns the public IP address of the virtual machine
	IPAddress() string

	// PrivateIPAddress of the virtual machine
	PrivateIPAddress() string
	UpTime() time.Duration
	State() string
	Type() string

	Environment() Environment
	Profile() Profile
	Provider() Provider

	Destroy() error
	Start() error
	Stop() error
	Suspend() error
	Restart() error

	ExecSSH(string, bool, ...string) error
	StartVNC(user string, privateIP bool, killVNC bool) error
	StartRDPClient(user string, privateIP bool) error
	Cmd(string, bool, ...string) (string, error)
	NewSSHSession(string, bool) (*ssh_util.Session, error)
}

type ProjectMachine struct {
	Name       string `yaml:"name"`
	Image      string `yaml:"image,omitempty"`
	Type       string `yaml:"type,omitempty"`
	Profile    string `yaml:"profile,omitempty"`
	Purpose    string `yaml:"purpose,omitempty"`
	VolumeSize int    `yaml:"volume_size,omitempty"`
	Provider   string `yaml:"provider,omitempty"`
}

type ProjectConfig struct {
	IngressIPs []string          `yaml:"ingress_ips"`
	Bastions   []*ProjectMachine `yaml:"bastions"`

	Machines []*ProjectMachine `yaml:"machines"`
}
