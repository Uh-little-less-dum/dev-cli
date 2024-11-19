package cmd

import (
	"os"

	init_helpers "github.com/Uh-little-less-dum/dev-cli/internal/initHelpers"
	outpututils_generate_tunnel "github.com/Uh-little-less-dum/dev-cli/internal/utils/outputUtils/generateTunnel"
	"github.com/charmbracelet/log"
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
		f := viper.GetViper().GetString("format")
		if f == "" {
			f = "asType"
		}
		globRoot := viper.GetViper().GetString("cwd")
		generator := outpututils_generate_tunnel.NewGenerator(fp, glob, globRoot, f)
		generator.Generate()
	},
}

func init() {
	cobra.OnInitialize(init_helpers.InitCmd())
	v := viper.GetViper()
	TunnelCmd.Flags().StringP("format", "f", "asType", "Indicates the format of the tunnel exports. Options: asType, basic")
	err := v.BindPFlag("format", TunnelCmd.Flags().Lookup("format"))
	if err != nil {
		log.Fatal(err)
	}
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	TunnelCmd.Flags().String("cwd", d, "Indicates the format of the tunnel exports. Options: asType, basic")
	err = v.BindPFlag("format", TunnelCmd.Flags().Lookup("format"))
	if err != nil {
		log.Fatal(err)
	}
	RootCmd.AddCommand(TunnelCmd)
}
