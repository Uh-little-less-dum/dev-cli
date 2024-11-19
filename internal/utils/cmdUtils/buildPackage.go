package utils_cmd

import (
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Takes the path to the root directory of the package, relative to the monorepo root, and an optional typescript config file name. If the file name is an empty string, tsconfig.json is used.
func BuildPackage(packageName string, packageRootPath ...string) {
	log.WithPrefix("Building").Info(packageName)
	p := []string{viper.GetViper().GetString("devRoot")}
	p = append(p, packageRootPath...)
	fp := filepath.Join(p...)
	PipeCommand(fp, "pnpm", "build")
	log.Infof("Built %s successfully", packageName)
}
