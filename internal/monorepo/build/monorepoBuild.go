package monorepo_build

import (
	"os"
	"os/exec"

	monorepo_post_build "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/postBuildStages"
	monorepo_build_stages "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/stages"
	monorepo_template_app "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/syncTemplatePackage"
	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	// utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
	// utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
)

func turboBuild() {
	devRoot := pathutils.GetDevRoot()
	cmd := exec.Command("pnpm", "turbo", "build", "--filter=!@ulld/template", "--filter=!@ulld/website", "--filter=!@ulld/internal-dev-cli", "--filter=!@ulld/cli", "--filter=!@ulld/developer-cli")
	cmd.Dir = devRoot
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	utils_error.HandleError(err)
}

func BuildMonorepo() {
	devRoot := pathutils.GetDevRoot()
	monorepo_build_stages.EnsureSynchronousBuild()
	monorepo_build_stages.ClearGeneratedTypes()
	monorepo_build_stages.GeneratePrismaClient()
	monorepo_build_stages.CopyPrismaTypes()
	monorepo_build_stages.GatherApiTypes()
	monorepo_build_stages.GenerateConfigSchemaTypes(devRoot)
	internalPackageItems := monorepo_build_stages.WriteCurrentPackageData()
	monorepo_build_stages.CreateUnifiedTypesExport(internalPackageItems)

	// Build all packages according to their wireit config using turbo for concurrency.
	turboBuild()

	// RESUME: Come back here and handle the build of the rest of the transpiled packages here.
	// utils_cmd.BuildPackage("@ulld/configschema", "packages", "configschema")

	monorepo_post_build.CopyTypesForDocumentation()
	monorepo_post_build.FormatPackageJsonFiles()
	monorepo_template_app.CopyDatabaseSchemaFile()
}
