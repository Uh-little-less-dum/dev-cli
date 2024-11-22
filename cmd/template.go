package cmd

import (
	template_version_cmd "github.com/Uh-little-less-dum/dev-cli/cmd/templateInternalVersion"
	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	"github.com/spf13/cobra"
)

var TemplateCmd = &cobra.Command{
	Use:   "templateApp",
	Short: "Contains several template app sub commands",
	Long:  "Contains several template app sub commands",
	// Args:  cobra.MaximumNArgs(2),

	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	TemplateCmd.AddCommand(template_version_cmd.ToProductionVersionsCmd)
	TemplateCmd.AddCommand(template_version_cmd.ToWorkspaceVersionCmd)
	RootCmd.AddCommand(TemplateCmd)
}
