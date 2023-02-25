package utils

import (
	"log"
	"os"
	"os/exec"
)

// Clear runs the command to clear the terminal
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
