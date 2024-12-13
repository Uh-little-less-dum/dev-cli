package gather_cli_data

import (
	"os"

	monorepo_utils "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/utils"
	"github.com/charmbracelet/log"
)

func GatherCliData() {
	devRoot := os.Getenv("ULLD_DEV_ROOT")          // Path to monorepo root.
	cliDevRoot := os.Getenv("ULLD_BUILD_DEV_ROOT") // Path to build repo root. (Not the cli root.)
	if (devRoot == "") || (cliDevRoot == "") {
		log.Fatal("Cannot continue generating data for the ULLD cli. One or more environment variables are missing. Required: ULLD_DEV_ROOT, ULLD_BUILD_DEV_ROOT")
	}
	packageVersions := monorepo_utils.GetCurrentPackageVersions()
	WriteInternalPackageData(packageVersions)
}
