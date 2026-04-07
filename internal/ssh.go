package internal

import (
	"fmt"
	"matman0497/sshboy/internal/config"
	"os"
	"os/exec"
)

func Connect(server *config.Server) *exec.Cmd {

	if server == nil {
		fmt.Printf("Server not found\n")
		os.Exit(1)
	}

	sshCmd := exec.Command("ssh", fmt.Sprintf("%s@%s", server.User, server.Host))

	return sshCmd
}
