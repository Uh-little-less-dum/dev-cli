package cmd

import (
	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	monorepo_build "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build"
	"github.com/spf13/cobra"
)

var buildMonorepoCommand = &cobra.Command{
	Use:   "build",
	Short: "Build the entire ULLD ecosystem",
	Long:  `Builds the entire ULLD ecosystem. This depends on the ULLD_DEV_ROOT and ULLD_CLI_DEVROOT env variables to find each target directory.`,
	// Args:  cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		monorepo_build.BuildMonorepo()
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	RootCmd.AddCommand(buildMonorepoCommand)
}
