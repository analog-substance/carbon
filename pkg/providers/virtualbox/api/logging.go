package api

import (
	"github.com/analog-substance/carbon/pkg/common"
	"log/slog"
)

var log *slog.Logger

func init() {
	log = common.WithGroup("vboxmanage-wrapper")
}
