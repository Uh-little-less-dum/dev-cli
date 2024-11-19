package monorepo_build

import (
	monorepo_build_stages "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/stages"
	utils_cmd "github.com/Uh-little-less-dum/dev-cli/internal/utils/cmdUtils"
)

func BuildMonorepo() {
	monorepo_build_stages.EnsureSynchronousBuild()
	// The generatePrismaClient type must come after the clearGeneratedTypes method, as the former clears the latter.
	monorepo_build_stages.ClearGeneratedTypes()
	monorepo_build_stages.GeneratePrismaClient()
	monorepo_build_stages.CopyPrismaTypes()
	monorepo_build_stages.GatherApiTypes()
	monorepo_build_stages.WriteCurrentPackageData() // Move this to the last method.
	utils_cmd.BuildPackage("@ulld/configschema", "packages", "configschema")
}
