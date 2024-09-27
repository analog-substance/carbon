package types

type Provider interface {
	Profiles(validNames ...string) []Profile
	Name() string
	Type() string
	IsAvailable() bool
}
