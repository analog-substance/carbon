package models

import (
	"github.com/analog-substance/carbon/pkg/types"
	"path"
)

type ImageBuild struct {
	Path string
	Env  types.Environment
}

func (b *ImageBuild) Name() string {
	return path.Base(b.Path)
}

func (b *ImageBuild) Environment() types.Environment {
	return b.Env
}
