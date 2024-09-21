package types

type Environment interface {
	Name() string
	VMs() []VM
	Platform() Platform
	StartVM(string) error
	StopVM(string) error
	RestartVM(string) error
}
