package pathutils

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Returns package.json paths in either the apps directory or the packages directory.
func getDirContents(dirPath string) []string {
	var dirEntries []string
	ad, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range ad {
		if item.IsDir() {
			fp := filepath.Join(dirPath, item.Name(), "package.json")
			if Exists(fp) {
				dirEntries = append(dirEntries, fp)
			}
		}
	}
	return dirEntries
}

func GetInternalPackageJsonPaths() []string {
	rootDir := viper.GetViper().GetString("devRoot")
	var filePaths []string = getDirContents(filepath.Join(rootDir, "apps"))
	filePaths = append(filePaths, getDirContents(filepath.Join(rootDir, "packages"))...)
	return filePaths
}
