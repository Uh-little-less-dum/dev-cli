package init_helpers

import (
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	init_viper "github.com/Uh-little-less-dum/dev-cli/internal/viper"
)

func InitCmd() func() {
	return func() {
		utils_logger.InitLogger("")
		init_viper.InitViper()
	}
}
