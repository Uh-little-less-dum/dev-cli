package cmd

import (
	gather_cli_data "github.com/Uh-little-less-dum/dev-cli/internal/dataSharing/gatherCliData"
	"github.com/spf13/cobra"
)

var GatherCliDataCommand = &cobra.Command{
	Use:    "gatherData",
	Short:  "Gather's data from the monorepo and generates code in the CLI.",
	Hidden: true,

	Run: func(cmd *cobra.Command, args []string) {
		gather_cli_data.GatherCliData()
	},
}

func init() {
	RootCmd.AddCommand(GatherCliDataCommand)
}
