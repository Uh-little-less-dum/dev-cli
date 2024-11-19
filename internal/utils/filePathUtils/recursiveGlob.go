package pathutils

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/gobwas/glob"
)

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.IsDir()
}

// Returns a list of *absolute* paths to files matching the glob.
func RecursiveGlob(rootDir, globPattern, ignoreFile string, noDir bool) ([]string, error) {
	log.WithPrefix("glob root").Debug(rootDir)
	g := glob.MustCompile(globPattern)
	files := []string{}
	err := filepath.Walk(rootDir, func(path string, f os.FileInfo, err error) error {
		if (path != "") && (g.Match(path)) && (ignoreFile != path) {
			if (!noDir) || (!IsDirectory(path)) {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}
