package templateapp_to_workspace_version

import (
	"os"
	"path/filepath"

	monorepo_utils "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/utils"
	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func SetTemplateAppToWorkspaceVersions() {
	packagePath := filepath.Join(viper.GetViper().GetString("devRoot"), "apps", "template", "package.json")
	content, err := os.ReadFile(packagePath)
	utils_error.HandleError(err)
	j := gjson.ParseBytes(content)
	monorepo_utils.GetInternalDependencies(j)

}
