package types

type Image interface {
	Name() string
	Environment() Environment
	Launch() error
}
