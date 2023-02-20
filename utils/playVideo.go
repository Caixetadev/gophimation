package utils

import (
	"log"
	"os"
	"os/exec"
)

func PlayVideo(videoUrl string) {
	cmd := exec.Command("mpv", "--save-position-on-quit", "--no-terminal", "--fs", "--force-media-title=caixetachan", "--cache=yes", videoUrl)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
