package utils_io

import (
	"os"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
)

func CopyFile(inputPath, outputPath string) {
	data, err := os.ReadFile(inputPath)
	utils_error.HandleError(err)
	err = os.WriteFile(outputPath, data, 0666)
}
