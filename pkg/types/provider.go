package types

type Provider interface {
	Platforms(validNames ...string) []Platform
	Name() string
	IsAvailable() bool
}
