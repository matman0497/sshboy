package cmd

import (
	"bufio"
	"fmt"
	"mattiamancina/sshboy/internal/config"
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
	Run: func(cmd *cobra.Command, args []string) {

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

		err := config.Add(name, host, user)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		config.Save()
	},
}
