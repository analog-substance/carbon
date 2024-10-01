package types

type Image interface {
	ID() string
	Name() string
	CreatedAt() string
	Environment() Environment
	Launch() error
}
