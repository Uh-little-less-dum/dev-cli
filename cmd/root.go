/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "godum",
	Short: "Useful utilities that power the ULLD monorepo and build process.",
	Long:  `Useful utilities that power the ULLD monorepo and build process.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Output more verbose logs.")
	viper.GetViper().BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
}
