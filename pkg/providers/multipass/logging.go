package multipass

import (
	"github.com/analog-substance/carbon/pkg/common"
	"log/slog"
)

var log *slog.Logger

func init() {
	log = common.WithGroup("providers.multipass")
}
