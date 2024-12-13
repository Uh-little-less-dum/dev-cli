package gather_cli_data

import (
	"fmt"
	"os"
	"path/filepath"

	monorepo_utils "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/utils"
	"github.com/charmbracelet/log"
)

func getFormattedList(items []monorepo_utils.PackageVersionData) string {
	s := ""
	for _, l := range items {
		s += fmt.Sprintf("        { Name: \"%s\", Version: \"%s\" },\n", l.Name, l.Version)
	}
	return s
}

func getOutputPath() string {
	buildDevRoot := os.Getenv("ULLD_BUILD_DEV_ROOT")
	if buildDevRoot == "" {
		log.Fatal("Cannot continue generating data for the ULLD cli. One or more environment variables are missing. Required: ULLD_DEV_ROOT, ULLD_BUILD_DEV_ROOT")
	}
	return filepath.Join(buildDevRoot, "pkg", "internalPackages", "internalPackageItems.go")
}

func WriteInternalPackageData(items []monorepo_utils.PackageVersionData) {

	outputPath := getOutputPath()

	s := fmt.Sprintf(`package internal_package_items 

func GetInternalPackageData() []InternalPackageItem {
	return []InternalPackageItem{
%s
	}
	}`, getFormattedList(items))

	err := os.WriteFile(outputPath, []byte(s), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
