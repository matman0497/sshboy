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
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new server configuration",
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Name: ")
		name, _ := reader.ReadString('\n')

		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Host: ")
		host, _ := reader.ReadString('\n')

		fmt.Print("User: ")
		user, _ := reader.ReadString('\n')

		host = strings.TrimSpace(host)
		user = strings.TrimSpace(user)
		name = strings.TrimSpace(name)

		if host == "" || user == "" || name == "" {
			return fmt.Errorf("add server: %w", fmt.Errorf("Host, User or Name must not be blank"))
		}

		err := config.Add(name, host, user)

		if err != nil {
			return fmt.Errorf("save config: %w", err)
		}

		config.Save()

		return nil
	},
}
