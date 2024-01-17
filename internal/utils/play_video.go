package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// PlayVideo runs a command that opens a video player to play a video from a specified URL
func PlayVideo(videoUrl, nameEpisode string) {
	executablesByOS := map[string]string{
		"windows": "mpv.exe",
		"linux":   "mpv",
	}

	cmd := exec.Command(
		executablesByOS[runtime.GOOS],
		videoUrl,
		"--fs",
		"--force-window=immediate",
		"--no-terminal",
		fmt.Sprintf("--force-media-title=%v", nameEpisode),
		"--cache=yes",
	)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
