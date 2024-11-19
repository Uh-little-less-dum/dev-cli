package monorepo_build_stages

import (
	"os"
	"path/filepath"

	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/spf13/viper"
)

// Clears the automatically generated types directory at packages/types/src/generated and prepares it to be regenerated.
func ClearGeneratedTypes() {
	devRoot := viper.GetViper().GetString("devRoot")
	p1 := filepath.Join(devRoot, "packages", "types", "src", "generated")
	os.RemoveAll(p1)
	err := os.Mkdir(p1, 0777)
	handleError(err)
	f, err := os.Create(filepath.Join(p1, "index.ts"))
	handleError(err)
	_, err = f.Write([]byte("export type JUST_TO_AVOID_ERROR = {}"))
	handleError(err)
	err = f.Close()
	handleError(err)
	err = os.MkdirAll(filepath.Join(p1, "configschema"), 0777)
	handleError(err)
	f, err = os.Create(filepath.Join(p1, "configschema", "index.ts"))
	handleError(err)
	_, err = f.Write([]byte("export type JUST_TO_AVOID_ERROR = {}"))
	handleError(err)
	f.Close()

	utils_logger.LogStageSuccess("Cleared automatically generated types at types/src/generated. This directory can now be populated.")
}
