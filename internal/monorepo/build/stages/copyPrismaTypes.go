package monorepo_build_stages

import (
	"path/filepath"

	utils_io "github.com/Uh-little-less-dum/dev-cli/internal/utils/ioUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/spf13/viper"
)

func CopyPrismaTypes() {
	rootDir := viper.GetViper().GetString("devRoot")
	inputPath := filepath.Join(rootDir, "node_modules", ".prisma", "client", "index.d.ts")
	outputPath := filepath.Join(rootDir, "packages", "types", "src", "database.d.ts")
	utils_io.CopyFile(inputPath, outputPath)

	utils_logger.LogStageSuccess("Copied prisma types to the types package.")
}
