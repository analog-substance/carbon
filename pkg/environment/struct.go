package environment

type Environment struct{}

func New() *Environment {
	return &Environment{}
}

func All() []*Environment {
	return []*Environment{}
}

func (e *Environment) AllVMs() {}
