package types

type Platform interface {
	Environments(validNames ...string) []Environment
	Name() string
	Provider() Provider
}
