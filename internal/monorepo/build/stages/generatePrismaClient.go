package monorepo_build_stages

import (
	"path/filepath"

	utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Generates the prisma client so types can be copied.
func GeneratePrismaClient() {
	utils_cmd.PipeCommand(filepath.Join(viper.GetViper().GetString("devRoot"), "packages", "database"), "pnpm", "prisma", "generate")
	log.Info("Generated prisma client. Types can now be copied.")
}