package types

type ImageLaunchOptions struct {
	Name string
}

type Image interface {
	ID() string
	Name() string
	CreatedAt() string
	Environment() Environment
	Profile() Profile
	Provider() Provider
	Launch(imageLaunchOptions ImageLaunchOptions) error
	Destroy() error
}
