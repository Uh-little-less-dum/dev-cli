package pathutils

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func EnsureDirExists(p string) error {
	err := os.MkdirAll(p, os.ModePerm)
	return err
}

// Creates a file or directory path if it doesn't exist. If it does exist, it does nothing.
func EnsureFileExists(p string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(p), os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(p)
	return f, err
}

// func GetRelativePath(fromPath string, toPath string) {
// 	filepath.Rel(fromPath, targPath)
// }

func GetDevRoot() string {
	return viper.GetViper().GetString("devRoot")
}
