package commonpaths

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func InternalBuildStaticDataPath() string {
	return filepath.Join(viper.GetViper().GetString("devRoot"), "packages", "utilities", "src", "utils", "buildStaticData.json")
}
