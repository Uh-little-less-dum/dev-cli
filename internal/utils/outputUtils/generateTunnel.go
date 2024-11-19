package utils_output

import (
	"fmt"
	"path/filepath"
	"strings"

	pathutils "github.com/Uh-little-less-dum/dev-cli/internal/utils/filePathUtils"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type ExportFormat int

const (
	ExportAsType ExportFormat = iota
	ExportAllWithoutExtension
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
	switch t.exportFormat {
	case ExportAsType:
		return fmt.Sprintf("export type * from \"%s\";\n", formatRelativePath(relPath))
	case ExportAllWithoutExtension:
		return fmt.Sprintf("export * from \"%s\";\n", pathutils.AddLeadingSlashAndRemoveExtension(relPath))
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
	files, err := pathutils.RecursiveGlob(t.globRoot, t.glob, t.path, true)
	handleError(err)
	content := t.getOutputString(files)
	_, err = f.WriteString(content)
	handleError(err)
	log.Infof("Generated tunnel file at %s", t.path)
}

func GetExportFormatFromString(exportFormat string) ExportFormat {
	switch exportFormat {
	case "asType":
		return ExportAsType
	default:
		return ExportAsType
	}
}

func GetTunnelFormatFromViper() ExportFormat {
	v := viper.GetViper()
	if v.GetBool("asType") {
		return ExportAsType
	}
	if v.GetBool("allWithoutExtension") {
		return ExportAllWithoutExtension
	}
	return ExportAsType
}

func NewGenerator(targetPath, relativeGlob, globRoot string, exportFormat ExportFormat) TunnelGenerator {
	return TunnelGenerator{
		path:         targetPath,
		glob:         relativeGlob,
		globRoot:     globRoot,
		exportFormat: exportFormat,
	}
}
