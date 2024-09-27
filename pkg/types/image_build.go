package types

type ImageBuild interface {
	Name() string
	ProviderType() string
	Provisioner() string
	Build() error
}
