package monorepo_post_build

import (
	"os/exec"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/spf13/viper"
)

func FormatPackageJsonFiles() {
	cmd := exec.Command("pnpm", "syncpack", "format")
	devRoot := viper.GetViper().GetString("devRoot")
	cmd.Dir = devRoot
	err := cmd.Run()
	utils_error.HandleError(err)
	utils_logger.LogStageSuccess("Formatted all package.json files with syncpack.")
}
