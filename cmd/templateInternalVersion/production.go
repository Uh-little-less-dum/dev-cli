package template_version_cmd

import (
	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// BUG: Fix both version manipulation methods. Move much of this functionality to the package.json struct.
var ToProductionVersionsCmd = &cobra.Command{
	Use:   "toProduction",
	Short: "Set template app versions to latest production versions.",
	Long:  "Set template app versions to latest production versions.",
	// Args:  cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("Handle this")
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	// cmd.TemplateCmd.AddCommand(ToProductionVersionsCmd)
}
