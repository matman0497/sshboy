package cmd

import (
	"fmt"
	"matman0497/sshboy/internal"
	"matman0497/sshboy/internal/config"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCommand)
}

var connectCommand = &cobra.Command{
	Use:   "connect [server name]",
	Short: "Connect to a server",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		server := config.GetServer(args[0])

		if server == nil {
			return fmt.Errorf("connect server: %w", fmt.Errorf("server %s not found", args[0]))
		}

		sshCmd := internal.Connect(server)

		sshCmd.Stdin = os.Stdin
		sshCmd.Stdout = os.Stdout
		sshCmd.Stderr = os.Stderr

		err := sshCmd.Run()
		if err != nil {
			panic(err)
		}

		return nil

	},
}
