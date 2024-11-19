package monorepo_build_stages

import (
	"os"

	"github.com/charmbracelet/log"
)

func EnsureSynchronousBuild() {
	os.Setenv("WIREIT_PARALLEL", "1")
	log.Info("Creating a synchronous build.")
}
