package cmd

import (
	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	monorepo_build_stages "github.com/Uh-little-less-dum/dev-cli/internal/monorepo/build/stages"
	"github.com/spf13/cobra"
)

var TestRouteCommand = &cobra.Command{
	Use:    "test",
	Short:  "Test command route. You will regret using this.",
	Long:   `Test individual parts of the build script. This should never be used apart from during internal development, because this will just execute the most random shit.`,
	Hidden: true,
	Args:   cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		monorepo_build_stages.WriteCurrentPackageData()
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	RootCmd.AddCommand(TestRouteCommand)
}
