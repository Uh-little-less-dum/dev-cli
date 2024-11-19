package monorepo_build_stages

import (
	"path/filepath"

	outpututils_generate_tunnel "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils/generateTunnel"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Creates tunnel files in the api package.
func GatherApiTypes() {
	rootDir := viper.GetViper().GetString("devRoot")
	if rootDir == "" {
		log.Fatal("No ULLD_DEV_ROOT variable found.")
	}
	targetPath := filepath.Join(rootDir, "packages", "api", "src", "types.ts")
	globPattern := filepath.Join(rootDir, "packages", "api", "src", "{trpcTypes,individualTypesForDocumentation}", "**", "*.{ts,tsx}")

	g := outpututils_generate_tunnel.NewGenerator(targetPath, globPattern, filepath.Join(rootDir, "packages", "api", "src"), "asType")
	g.Generate()
	log.Info("Wrote tunnel files in the api package.")
}
