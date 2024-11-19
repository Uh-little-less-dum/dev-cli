package cmd

import (
	"os"

	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	utils_error "github.com/Uh-little-less-dum/dev-cli/internal/utils/errorHandler"
	utils_output "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TunnelCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "Generate a tunnel file",
	Long: `Generates a tunel file. Takes two positional arguments:
	1. The target file
	2. The glob pattern used to create the tunnel. (default: **/*.{ts,tsx})`,
	Args: cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		fp := args[0]
		var glob string
		if len(args) == 2 {
			glob = args[1]
		} else {
			glob = "**/*.{ts,tsx}"
		}
		globRoot := viper.GetViper().GetString("cwd")
		format := utils_output.GetTunnelFormatFromViper()
		generator := utils_output.NewGenerator(fp, glob, globRoot, format)
		generator.Generate()
	},
}

func handleBoolFlag(v *viper.Viper, viperKey, flagString, desc string, defaultVal bool) {
	TunnelCmd.Flags().Bool(flagString, defaultVal, desc)
	err := v.BindPFlag(viperKey, TunnelCmd.Flags().Lookup(flagString))
	utils_error.HandleError(err)
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	v := viper.GetViper()
	handleBoolFlag(v, "asType", "asType", "Generate tunnel exports as type only exports.", false)
	handleBoolFlag(v, "allWithoutExtension", "allWithoutExtension", "Generate tunnel exports using the '* from' syntax, but without a file extension.", false)
	d, err := os.Getwd()
	utils_error.HandleError(err)

	TunnelCmd.Flags().String("cwd", d, "Indicates the format of the tunnel exports. Options: asType, basic")
	RootCmd.AddCommand(TunnelCmd)
}
