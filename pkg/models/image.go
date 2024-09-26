package models

import (
	"github.com/analog-substance/carbon/pkg/types"
	"time"
)

type Image struct {
	Name      string
	ID        string
	CreatedAt time.Time
	Env       types.Environment
}
