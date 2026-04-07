package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the current version number",
	Run: func(cmd *cobra.Command, args []string) {

		println("SSHBOY v0.1.0")

	},
}
