package models

import (
	"github.com/analog-substance/carbon/pkg/types"
	"time"
)

type Image struct {
	Label     string
	ID        string
	CreatedAt time.Time
	Env       types.Environment
}

func NewImage(imageID string, env types.Environment) types.Image {
	return &Image{
		ID:  imageID,
		Env: env,
	}
}

func (i *Image) Name() string {
	return i.ID
}

func (i *Image) Launch() error {
	return nil
}

func (i *Image) Environment() types.Environment {
	return i.Env
}
