package monorepo_build_stages

import (
	"os/exec"
	"path/filepath"

	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
)

func GenerateConfigSchemaTypes(devRoot string) {
	filePath := filepath.Join(devRoot, "packages", "configschema", "tsconfigTypes.json")
	cmd := exec.Command("pnpm", "tsc", "--project", filePath, "--emitDeclarationOnly")
	cmd.Dir = devRoot
	cmd.Run()

	utils_logger.LogStageSuccess("Generated configschema types in the types package. The dependent packages can not be built successfully.")

}
