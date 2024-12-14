package monorepo_build_stages

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/Uh-little-less-dum/dev-cli/internal/monorepo/commonpaths"
	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	schemas_package_json "github.com/Uh-little-less-dum/go-utils/pkg/schemastructs/packageJson"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"github.com/tidwall/sjson"
)

type InternalPackageType string

const (
	AppRepo     InternalPackageType = "app"
	PackageRepo InternalPackageType = "package"
)

type SourceLocation string

const (
	SrcDir  SourceLocation = "src"
	DistDir SourceLocation = "dist"
)

type InternalPackageItem struct {
	CurrentVersion  string              `json:"currentVersion"`
	RelativeDirPath string              `json:"relativeDirPath"`
	Name            string              `json:"name"`
	RepoType        InternalPackageType `json:"type"`
	HasPluginConfig bool                `json:"hasPluginConfig"`
	IsTranspiled    bool                `json:"isTranspiled"`
	SourceLocation  SourceLocation      `json:"sourceLocation"`
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type HasFileType struct {
	src  bool
	dist bool
}

func fatalIfBothExports(fileTypes HasFileType, packageName string) {
	if (fileTypes.dist) && (fileTypes.src) {
		log.Fatalf("Attempting to gather the InternalData from %s. This package appears to have both src and dist directory exports.", packageName)
	}
}

func handleImportPath(p string, fileTypes *HasFileType) {
	if p != "" {
		if strings.HasPrefix(p, "./src") {
			fileTypes.src = true
		} else {
			fileTypes.dist = true
		}
	}
}

func getSourceType(data schemas_package_json.PackageJsonSchema) SourceLocation {
	log.Debugf("Gathering the InternalPackageItem for %s", data.Name)
	fileTypes := HasFileType{src: false, dist: false}
	filesField := data.Files
	stringType := reflect.TypeOf("")
	for _, f := range filesField {
		if f == "src" {
			fileTypes.src = true
		} else if f == "dist" {
			fileTypes.dist = true
		}
	}
	fatalIfBothExports(fileTypes, data.Name)
	var sourceType SourceLocation
	if fileTypes.dist {
		sourceType = DistDir
	} else {
		sourceType = SrcDir
	}
	var exportKeys []reflect.Value
	if data.Exports != nil {
		if reflect.TypeOf(data.Exports) == stringType {
			exportVal := fmt.Sprintf("%s", data.Exports)
			log.Info(exportVal)
			if strings.HasPrefix(exportVal, "./src") {
				return SrcDir
			} else {
				return DistDir
			}
		}
		v := reflect.ValueOf(data.Exports)
		exportKeys = v.MapKeys()
		for _, k := range exportKeys {
			keyString := fmt.Sprintf("%s", k)
			exportVal := data.Exports.(map[string]any)[keyString]
			exportType := reflect.TypeOf(exportVal)
			isStringExport := exportType == stringType
			if isStringExport {
				handleImportPath(fmt.Sprintf("%s", exportVal), &fileTypes)
			} else {
				for _, u := range reflect.ValueOf(exportVal).MapKeys() {
					us := u.String()
					handleImportPath(fmt.Sprintf("%s", exportVal.(map[string]any)[us]), &fileTypes)
				}
			}
		}
	}
	if fileTypes.dist {
		return DistDir
	}
	return sourceType
}

func getPackageItem(devRoot, packageJsonPath string) InternalPackageItem {
	b, err := os.ReadFile(packageJsonPath)
	handleError(err)
	fileData := schemas_package_json.PackageJsonSchema{}
	err = json.Unmarshal(b, &fileData)
	handleError(err)
	relPath, err := filepath.Rel(devRoot, filepath.Dir(packageJsonPath))
	handleError(err)

	var repoType InternalPackageType
	if strings.HasPrefix(relPath, "apps") {
		repoType = AppRepo
	} else {
		repoType = PackageRepo
	}

	srcData := getSourceType(fileData)

	return InternalPackageItem{
		Name:            fileData.Name,
		CurrentVersion:  fileData.Version,
		RelativeDirPath: relPath,
		RepoType:        repoType,
		HasPluginConfig: pathutils.Exists(filepath.Join(filepath.Dir(packageJsonPath), "pluginConfig.ulld.json")),
		SourceLocation:  srcData,
		IsTranspiled:    srcData == DistDir,
	}
}

// Writes current package data to utilities/src/utils/buildStaticData.json based on their current package.json.
func WriteCurrentPackageData() []InternalPackageItem {
	rootDir := viper.GetViper().GetString("devRoot")
	files := pathutils.GetInternalPackageJsonPaths()

	var wg sync.WaitGroup
	var items []InternalPackageItem
	for _, f := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := getPackageItem(rootDir, f)
			items = append(items, data)
		}()
	}
	wg.Wait()

	outputPath := commonpaths.InternalBuildStaticDataPath()
	targetContent, err := os.ReadFile(outputPath)
	handleError(err)

	newData, err := sjson.SetBytes(targetContent, "internalPackageData", items)
	handleError(err)

	err = os.WriteFile(outputPath, newData, 0666)
	handleError(err)

	utils_logger.LogStageSuccess("Wrote internalPackageData to utilities/src/utils/buildStaticData.json successfully.")
	return items
}
