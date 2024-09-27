package types

import (
	"github.com/analog-substance/carbon/pkg/ssh_util"
)

type MachineState struct{ Name string }

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
	State() string
	Start() error
	Stop() error
	Restart() error
	Environment() Environment
	ExecSSH(string, ...string) error
	StartVNC(user string, killVNC bool) error
	NewSSHSession(string) (*ssh_util.Session, error)
}
