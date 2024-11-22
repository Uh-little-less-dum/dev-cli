package monorepo_utils

import (
	"os"

	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type PackageVersionData struct {
	Name    string
	Version string
}

func getVersionData(fp string) PackageVersionData {
	content, err := os.ReadFile(fp)
	utils_error.HandleError(err)
	j := gjson.ParseBytes(content)
	name := j.Get("name").Str
	version := j.Get("version").Str
	return PackageVersionData{
		Name:    name,
		Version: version,
	}
}

// Returns versions directly from package.json, not from the generated buildstaticdata.json in the utilities folder. This should make sure things are always current, without running into ordering issues.
func GetCurrentPackageVersions() []PackageVersionData {
	var versionData []PackageVersionData
	packagePaths := pathutils.GetInternalPackageJsonPaths()
	for _, item := range packagePaths {
		versionData = append(versionData, getVersionData(item))
	}
	return versionData
}

type InternalDepData struct {
	Name           string
	CurrentVersion string
	DependencyType string
}

func GetInternalDependencies(j gjson.Result) []InternalDepData {
	var data []InternalDepData
	devItems := j.Get("devDendencies.@ulld/*").Array()
	for _, item := range devItems {
		log.Info(item)
		//    data = append(data, InternalDepData{
		// 	Name: item
		// })
	}
	// items := j.Get("dependencies.@ulld/*").Array()
	// peerItems := j.Get("peerDendencies.@ulld/*").Array()

	return data
}
