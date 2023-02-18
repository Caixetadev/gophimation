package utils

import (
	"log"
	"os"
	"os/exec"
)

func PlayVideo(videoUrl string) {
	cmd := exec.Command("mpv", "--save-position-on-quit", videoUrl, "--no-terminal", "--fs")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
