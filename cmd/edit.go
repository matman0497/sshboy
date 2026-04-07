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
	Run: func(cmd *cobra.Command, args []string) {
		server := config.GetServer(args[0])

		if server == nil {
			fmt.Printf("Server %s not found\n", args[0])
			os.Exit(1)
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
		fmt.Print("Saved!")
	},
}
