package cmd

import (
	"fmt"
	"matman0497/sshboy/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Print all the available servers",
	Run: func(cmd *cobra.Command, args []string) {

		store := config.Store{}
		prettyPrint(store.List())
	},
}

func prettyPrint(servers []config.Server) {
	var b strings.Builder

	b.WriteString("Hosts:\n")

	for _, s := range servers {
		fmt.Fprintf(&b,
			"   %s\n"+
				"      Host: %s\n"+
				"      User: %s\n",
			s.Name,
			s.Host,
			s.User,
		)
	}

	fmt.Println(b.String())
}
