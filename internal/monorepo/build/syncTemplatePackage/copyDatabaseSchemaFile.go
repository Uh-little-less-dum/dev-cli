package monorepo_template_app

import (
	"os"
	"path/filepath"
	"strings"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
)

// Copies database schema file to the template app. This does not necessarily have to come after the generation of the prisma client, as this is only applied to the pre-transpiled template app.
func CopyDatabaseSchemaFile() {
	devRoot := pathutils.GetDevRoot()
	schemaFilePath := filepath.Join(devRoot, "packages", "database", "prisma", "schema.prisma")
	b, err := os.ReadFile(schemaFilePath)
	utils_error.HandleError(err)
	allLines := strings.Split(string(b), "\n")
	var lines []string
	for _, l := range allLines {
		if strings.Contains(l, "<<CUT-HERE>>") {
			break
		} else {
			lines = append(lines, l)
		}
	}
	err = os.WriteFile(filepath.Join(devRoot, "apps", "template", "src", "database", "schema.prisma"), []byte(strings.Join(lines, "\n")), 0666)
	utils_error.HandleError(err)
	utils_logger.LogTemplateAppMsg("Copied schema file to the template package.")
}
