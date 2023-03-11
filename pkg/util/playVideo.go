package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// PlayVideo runs a command that opens a video player to play a video from a specified URL
func PlayVideo(videoUrl, nameEpisode string) {
	var executablesByOS = map[string]string{
		"windows": "mpv.exe",
		"linux":   "mpv",
	}

	cmd := exec.Command(executablesByOS[runtime.GOOS], "--save-position-on-quit", "--no-terminal", "--fs", fmt.Sprintf("--force-media-title=%v", nameEpisode), "--cache=yes", videoUrl)
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Abrindo player...")

	// Set Discord user presence to show default presence
	// presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")
}
