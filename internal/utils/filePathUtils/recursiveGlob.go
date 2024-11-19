package pathutils

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/gobwas/glob"
)

func RecursiveGlob(rootDir, globPattern string) ([]string, error) {
	log.WithPrefix("glob root").Debug(rootDir)
	g := glob.MustCompile(globPattern)
	files := []string{}
	err := filepath.Walk(rootDir, func(path string, f os.FileInfo, err error) error {
		if g.Match(path) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
