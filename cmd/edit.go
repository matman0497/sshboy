package cmd

import (
	"bufio"
	"fmt"
	"matman0497/sshboy/internal/config"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCommand)
}

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edit a servers configuration",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		server := config.GetServer(args[0])

		if server == nil {
			return fmt.Errorf("connect server: %w", fmt.Errorf("server %s not found", args[0]))
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Host (leave blank to keep existing): ")
		newHost, _ := reader.ReadString('\n')

		fmt.Print("User (leave blank to keep existing): ")
		newUser, _ := reader.ReadString('\n')

		newHost = strings.TrimSpace(newHost)
		newUser = strings.TrimSpace(newUser)

		if newHost != "" {
			server.Host = strings.TrimSpace(newHost)
		}

		if newUser != "" {
			server.User = strings.TrimSpace(newUser)
		}

		config.Save()
		cmd.Println("Server was saved.")

		return nil
	},
}
