package cmd

import (
	"fmt"
	"mattiamancina/sshboy/internal"
	"mattiamancina/sshboy/internal/config"
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
	Run: func(cmd *cobra.Command, args []string) {

		server := config.GetServer(args[0])

		if server == nil {
			fmt.Printf("Server %s not found\n", args[0])
			os.Exit(1)
		}

		sshCmd := internal.Connect(server)

		sshCmd.Stdin = os.Stdin
		sshCmd.Stdout = os.Stdout
		sshCmd.Stderr = os.Stderr

		err := sshCmd.Run()
		if err != nil {
			panic(err)
		}

	},
}
