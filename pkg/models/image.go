package models

import (
	"encoding/json"
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

func (i *Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        i.ID(),
		Name:      i.Name(),
		CreatedAt: i.CreatedAt(),
	})
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

func (i *Image) Launch(imageLaunchOptions types.ImageLaunchOptions) error {
	launchOptions := types.MachineLaunchOptions{
		CloudInitTpl: "",
		Image:        i,
		Name:         imageLaunchOptions.Name,
	}

	return i.env.CreateVM(launchOptions)
}

func (i *Image) Destroy() error {
	return i.env.DestroyImage(i.imageID)
}

func (i *Image) Environment() types.Environment {
	return i.env
}

func (i *Image) Provider() types.Provider {
	return i.env.Profile().Provider()
}

func (i *Image) Profile() types.Profile {
	return i.env.Profile()
}
