package types

type MachineLaunchOptions struct {
	CloudInitTpl string `json:"cloud-init"`
	Image        Image  `json:"image"`
	Name         string `json:"name"`
}

type Environment interface {
	Name() string
	VMs() []VM
	Profile() Profile
	StartVM(string) error
	StopVM(string) error
	SuspendVM(string) error
	RestartVM(string) error
	ImageBuilds() ([]ImageBuild, error)
	Images() ([]Image, error)
	CreateVM(MachineLaunchOptions) error
	DestroyVM(string) error
	DestroyImage(string) error
}
