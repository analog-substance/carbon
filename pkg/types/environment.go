package types

import "github.com/aws/aws-sdk-go-v2/service/ec2/types"

type MachineLaunchOptions struct {
	CloudInitTpl string      `json:"cloud-init"`
	Image        types.Image `json:"image"`
}

type Environment interface {
	Name() string
	VMs() []VM
	Profile() Profile
	StartVM(string) error
	StopVM(string) error
	RestartVM(string) error
	ImageBuilds() ([]ImageBuild, error)
	Images() ([]Image, error)
	CreateVM(MachineLaunchOptions) error
	DestroyVM(string) error
}
