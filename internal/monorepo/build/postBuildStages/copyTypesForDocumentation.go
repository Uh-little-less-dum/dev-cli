package monorepo_post_build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Uh-little-less-dum/dev-cli/internal/monorepo/commonpaths"
	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/Uh-little-less-dum/dev-cli/internal/utils_logger"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

// Copies typescript/tsx files into a txt file in the website app to be used for documentation using fumadocs and typedoc. Inputs are retrieved from the packages/utilities/src/utils/buildStaticData.json#pathsToCopyForDocumentation array.
func copyPathItemForDocumentation(devRoot string, item gjson.Result) {
	var pathItems = []string{devRoot}
	for _, k := range item.Get("globRoot").Array() {
		pathItems = append(pathItems, k.Str)
	}
	globRoot := filepath.Join(pathItems...)
	res, err := pathutils.RecursiveGlob(globRoot, item.Get("glob").Str, "", true)
	utils_error.HandleError(err)
	for _, r := range res {
		fileName := filepath.Base(r)
		fileName, _ = strings.CutSuffix(fileName, filepath.Ext(fileName))
		content, err := os.ReadFile(r)
		utils_error.HandleError(err)
		var lines []string
		for _, l := range strings.Split(string(content), "\n") {
			if (!strings.HasPrefix(l, "import")) && (strings.Trim(l, " ") != "") {
				lines = append(lines, l)
			}
		}
		err = os.WriteFile(filepath.Join(devRoot, "apps", "website", "content", "embeddedDocs", "generated", fmt.Sprintf("%s.txt", fileName)), []byte(strings.Join(lines, "\n")), 0666)
		utils_error.HandleError(err)
	}
}

func CopyTypesForDocumentation() {
	buildDataPath := commonpaths.InternalBuildStaticDataPath()
	devRoot := viper.GetViper().GetString("devRoot")
	data, err := os.ReadFile(buildDataPath)
	utils_error.HandleError(err)
	j := gjson.ParseBytes(data)
	pathItemsToCopy := j.Get("pathsToCopyForDocumentation").Array()
	var wg sync.WaitGroup
	for _, item := range pathItemsToCopy {
		wg.Add(1)
		go func() {
			defer wg.Done()
			copyPathItemForDocumentation(devRoot, item)
		}()
	}
	wg.Wait()
	utils_logger.LogStageSuccess("Wrote txt files to the website app to be used for documentation.")
}
