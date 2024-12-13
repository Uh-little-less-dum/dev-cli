/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	RootCmd.AddCommand(JsonSchemaToStructCmd)
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Output more verbose logs.")
	err := viper.GetViper().BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		log.Fatal(err)
	}
}
