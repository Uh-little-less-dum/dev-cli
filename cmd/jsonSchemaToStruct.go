package cmd

import (
	"os/exec"

	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	_ "github.com/a-h/generate"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func copyToClipboard(val []byte) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, val)
}

var JsonSchemaToStructCmd = &cobra.Command{
	Use:   "jsonSchemaToStruct",
	Short: "Generate a go struct from a given json schema",
	Long:  "Generate a go struct from a given json schema",

	Run: func(cmd *cobra.Command, args []string) {
		inputPath := cmd.Flags().Lookup("input")
		outputPath := cmd.Flags().Lookup("output")
		c := exec.Command("schema-generate", "-o", outputPath.Value.String(), inputPath.Value.String())
		copyToClipboard([]byte(c.String()))
		err := c.Run()
		if err != nil {
			log.Fatal("Error: ", err)
		}
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	JsonSchemaToStructCmd.Flags().StringP("input", "i", "", "--input=some/input/path.json")
	JsonSchemaToStructCmd.Flags().StringP("output", "o", "", "--output=some/output/path.go")
	JsonSchemaToStructCmd.Flags().StringP("package", "p", "main", "--package=schemas_package_json")
}
