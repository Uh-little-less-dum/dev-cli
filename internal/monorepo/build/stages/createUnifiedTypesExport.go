package monorepo_build_stages

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	utils_output "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/spf13/viper"
)

func getInternalAppUnion(data []InternalPackageItem) string {
	var names []string
	for _, item := range data {
		names = append(names, fmt.Sprintf("\"%s\"", item.Name))
	}
	nameUnion := strings.Join(names, " | ")
	return fmt.Sprintf("export type InternalAppName = %s\n\nexport type InternalAppShortName = %s", nameUnion, strings.ReplaceAll(nameUnion, "@ulld/", ""))
}

// Creates various tunnel files within the types package. These generated tunnels are used by packages during the build process, but may also change after the build is complete. This function should be run near the end of the build script, when typescript transpilation and bundling is complete.
func CreateUnifiedTypesExport(internalPackageItems []InternalPackageItem) {
	devRoot := viper.GetViper().GetString("devRoot")
	typesRoot := filepath.Join(devRoot, "packages", "types", "src")
	generatedTypesRoot := filepath.Join(typesRoot, "generated")
	enumRoot := filepath.Join(typesRoot, "enums")
	manualTypesRoot := filepath.Join(typesRoot, "manualTypes")
	internalAppTypePath := filepath.Join(typesRoot, "internalAppNames.ts")
	err := os.WriteFile(internalAppTypePath, []byte(getInternalAppUnion(internalPackageItems)), 0666)
	utils_error.HandleError(err)
	enumTunnel := utils_output.NewGenerator(filepath.Join(enumRoot, "index.ts"), "**/*.{ts,tsx,json}", enumRoot, utils_output.ExportAllWithoutExtension)
	enumTunnel.Generate()

	manualTypesTunnel := utils_output.NewGenerator(filepath.Join(manualTypesRoot, "index.ts"), "**/*.{ts,tsx,json}", manualTypesRoot, utils_output.ExportAsType)
	manualTypesTunnel.Generate()

	generatedTypesTunnel := utils_output.NewGenerator(filepath.Join(generatedTypesRoot, "index.ts"), "**/*.{ts,tsx,json}", generatedTypesRoot, utils_output.ExportAsType)
	generatedTypesTunnel.Generate()
	utils_logger.LogStageSuccess("Generated tunnel files in the @ulld/types package succesfully.")
}
