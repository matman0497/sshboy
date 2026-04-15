package main

import (
	"log"
	"matman0497/sshboy/cmd"
	"os"
)

func main() {

	// Open file (create if not exists, append mode)
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set log output to file
	log.SetOutput(file)

	cmd.Execute()

}
