package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func PlayVideo(videoUrl, nameEpisode string) {
	cmd := exec.Command("mpv", "--save-position-on-quit", "--no-terminal", "--fs", fmt.Sprintf("--force-media-title=%v", nameEpisode), "--cache=yes", videoUrl)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
