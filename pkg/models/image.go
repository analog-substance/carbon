package models

import (
	"github.com/analog-substance/carbon/pkg/types"
	"time"
)

type Image struct {
	imageID   string
	imageName string
	createdAt time.Time
	env       types.Environment
}

func NewImage(imageID string, imageName string, createdAt time.Time, env types.Environment) types.Image {
	return &Image{
		imageID:   imageID,
		imageName: imageName,
		createdAt: createdAt,
		env:       env,
	}
}

func (i *Image) ID() string {
	return i.imageID
}

func (i *Image) Name() string {
	return i.imageName
}

func (i *Image) CreatedAt() string {
	return i.createdAt.Format(time.RFC3339Nano)
}

func (i *Image) Launch() error {
	return nil
}

func (i *Image) Environment() types.Environment {
	return i.env
}
