package types

type ImageBuild interface {
	Name() string
	Environment() Environment
}
