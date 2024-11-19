package monorepo_build_stages

import (
	"path/filepath"

	utils_output "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Creates tunnel files in the api package.
func GatherApiTypes() {
	rootDir := viper.GetViper().GetString("devRoot")
	if rootDir == "" {
		log.Fatal("No ULLD_DEV_ROOT variable found.")
	}
	targetPath := filepath.Join(rootDir, "packages", "api", "src", "types.ts")
	globPattern := filepath.Join(rootDir, "packages", "api", "src", "{trpcTypes,individualTypesForDocumentation}", "**", "*.{ts,tsx}")

	g := utils_output.NewGenerator(targetPath, globPattern, filepath.Join(rootDir, "packages", "api", "src"), utils_output.ExportAsType)
	g.Generate()
	utils_logger.LogStageSuccess("Wrote tunnel files in the api package.")
}
