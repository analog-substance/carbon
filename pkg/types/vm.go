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

type VM interface {
	Name() string
	ID() string
	IPAddress() string
	UpTime() time.Duration
	State() string
	Type() string

	Environment() Environment
	Profile() Profile
	Provider() Provider

	Destroy() error
	Start() error
	Stop() error
	Restart() error

	ExecSSH(string, ...string) error
	StartVNC(user string, killVNC bool) error
	Cmd(string, ...string) (string, error)
	NewSSHSession(string) (*ssh_util.Session, error)
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
	Machines []*ProjectMachine `yaml:"machines"`
}
