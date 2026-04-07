package cmd

import (
	"matman0497/sshboy/interactive"
	"matman0497/sshboy/internal/config"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(interactiveCommand)
}

type model struct {
	config   *config.Config
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

var interactiveCommand = &cobra.Command{
	Use:   "interactive",
	Short: "interactive",
	Run: func(cmd *cobra.Command, args []string) {
		interactive.Init()
	},
}
