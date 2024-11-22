package template_version_cmd

import (
	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	templateapp_to_workspace_version "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/version/template"
	"github.com/spf13/cobra"
)

var ToWorkspaceVersionCmd = &cobra.Command{
	Use:   "toWorkspace",
	Short: "Set template app internal versions to workspace:*",
	Long:  "Set template app internal versions to workspace:*",
	// Args:  cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		templateapp_to_workspace_version.SetTemplateAppToWorkspaceVersions()
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
}
