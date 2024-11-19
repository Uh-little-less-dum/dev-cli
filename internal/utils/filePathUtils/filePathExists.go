package pathutils

import (
	"errors"
	"os"
)

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
