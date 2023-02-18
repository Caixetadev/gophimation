package utils

import (
	"log"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
