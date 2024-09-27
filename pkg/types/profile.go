package types

type Profile interface {
	Environments(validNames ...string) []Environment
	Name() string
	Provider() Provider
}
