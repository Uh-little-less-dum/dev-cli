package outpututils_generate_tunnel

import (
	"fmt"
	"path/filepath"
	"strings"

	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/charmbracelet/log"
)

type ExportFormat int

const (
	ExportAsType ExportFormat = iota
)

type TunnelGenerator struct {
	path         string
	glob         string
	globRoot     string
	exportFormat ExportFormat
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatRelativePath(s string) string {
	if strings.HasPrefix(s, ".") {
		return s
	} else {
		return fmt.Sprintf("./%s", s)
	}
}

func (t TunnelGenerator) getExportItemString(relPath string) string {
	if t.exportFormat == ExportAsType {
		return fmt.Sprintf("export type * from \"%s\";\n", formatRelativePath(relPath))
	}
	return ""
}

func (t TunnelGenerator) getOutputString(files []string) string {
	s := ""
	for _, f := range files {
		fp, err := filepath.Rel(filepath.Dir(t.path), f)
		handleError(err)
		s += t.getExportItemString(fp)
	}
	return s
}

func (t TunnelGenerator) Generate() {
	f, err := pathutils.EnsureFileExists(t.path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	handleError(err)
	files, err := pathutils.RecursiveGlob(t.globRoot, t.glob)
	handleError(err)
	content := t.getOutputString(files)
	_, err = f.WriteString(content)
	handleError(err)
	log.Infof("Generated tunnel file at %s", t.path)
}

func getExportFormatFromString(exportFormat string) ExportFormat {
	switch exportFormat {
	case "asType":
		return ExportAsType
	default:
		return ExportAsType
	}
}

func NewGenerator(targetPath, relativeGlob, globRoot, exportFormat string) TunnelGenerator {
	return TunnelGenerator{
		path:         targetPath,
		glob:         relativeGlob,
		globRoot:     globRoot,
		exportFormat: getExportFormatFromString(exportFormat),
	}
}
