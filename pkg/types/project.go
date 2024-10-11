package types

type Project interface {
	Name() string
	TerraformApply() error
	AddMachine(machine *ProjectMachine, noApply bool) error
}
