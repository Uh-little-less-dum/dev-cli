package monorepo_build_stages

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	utils_output "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/charmbracelet/log"
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
	logger := log.WithPrefix("Tunnel Output")
	devRoot := viper.GetViper().GetString("devRoot")
	typesRoot := filepath.Join(devRoot, "packages", "types", "src")
	generatedTypesRoot := filepath.Join(typesRoot, "generated")
	enumRoot := filepath.Join(typesRoot, "enums")
	manualTypesRoot := filepath.Join(typesRoot, "manualTypes")
	internalAppTypePath := filepath.Join(typesRoot, "internalAppNames.ts")
	err := os.WriteFile(internalAppTypePath, []byte(getInternalAppUnion(internalPackageItems)), 0666)
	utils_error.HandleError(err)
	fp := filepath.Join(enumRoot, "index.ts")
	enumTunnel := utils_output.NewGenerator(fp, "**/*.{ts,tsx,json,d.ts}", enumRoot, utils_output.ExportAllWithoutExtension)
	enumTunnel.Generate()
	logger.Debugf("Wrote tunnel file to %s", fp)
	fp = filepath.Join(manualTypesRoot, "index.ts")

	manualTypesTunnel := utils_output.NewGenerator(fp, "**/*.{ts,tsx,json,d.ts}", manualTypesRoot, utils_output.ExportAsType)
	manualTypesTunnel.Generate()
	logger.Debugf("Wrote tunnel file to %s", fp)
	fp = filepath.Join(generatedTypesRoot, "index.ts")

	generatedTypesTunnel := utils_output.NewGenerator(fp, "**/*.{ts,tsx,json,d.ts}", generatedTypesRoot, utils_output.ExportAsType)
	generatedTypesTunnel.Generate()
	logger.Debugf("Wrote tunnel file to %s", fp)

	utils_logger.LogStageSuccess("Generated tunnel files in the @ulld/types package succesfully.")
}
