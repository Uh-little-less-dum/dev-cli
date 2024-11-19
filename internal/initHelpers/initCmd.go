package init_helpers

import (
	init_logger "github.com/Uh-little-less-dum/dev-cli/internal/logger"
	init_viper "github.com/Uh-little-less-dum/dev-cli/internal/viper"
)

func InitCmd() func() {
	return func() {
		init_logger.InitLogger("")
		init_viper.InitViper()
	}
}
