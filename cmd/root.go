package cmd

import (
	"fmt"
	"matman0497/sshboy/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{

		Use:   "sshboy",
		Short: "sshboy is a simple tool to connect to a server via ssh",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {

	config.Init()
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

}
