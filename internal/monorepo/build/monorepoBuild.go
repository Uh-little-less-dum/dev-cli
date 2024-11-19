package monorepo_build

import (
	monorepo_post_build "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/postBuildStages"
	monorepo_build_stages "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/stages"
	monorepo_template_app "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/syncTemplatePackage"
	utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
	// utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
)

func BuildMonorepo() {
	monorepo_build_stages.EnsureSynchronousBuild()
	// The generatePrismaClient type must come after the clearGeneratedTypes method, as the former clears the latter.
	monorepo_build_stages.ClearGeneratedTypes()
	monorepo_build_stages.GeneratePrismaClient()
	monorepo_build_stages.CopyPrismaTypes()
	monorepo_build_stages.GatherApiTypes()
	internalPackageItems := monorepo_build_stages.WriteCurrentPackageData()
	monorepo_build_stages.CreateUnifiedTypesExport(internalPackageItems)

	// RESUME: Come back here and handle the build of the rest of the transpiled packages here.
	utils_cmd.BuildPackage("@ulld/configschema", "packages", "configschema")

	monorepo_post_build.CopyTypesForDocumentation()
	monorepo_post_build.FormatPackageJsonFiles()
	monorepo_template_app.CopyDatabaseSchemaFile()
}
