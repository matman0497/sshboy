package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, args []string) error {

		home, err := os.UserHomeDir()

		err = os.MkdirAll(fmt.Sprintf("%s/.sshboy", home), os.ModePerm)

		if err != nil {
			return fmt.Errorf("an error occurred while writing ~/.sshboy/inventory.yaml: %w", err)
		}

		_, err = os.Create(fmt.Sprintf("%s/.sshboy/inventory.yaml", home))

		if err != nil {
			return fmt.Errorf("an error occurred while writing ~/.sshboy/inventory.yaml: %w", err)
		}

		return nil
	},
}
