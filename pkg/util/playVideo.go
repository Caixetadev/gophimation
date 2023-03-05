package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Caixetadev/gophimation/pkg/presence"
)

// PlayVideo runs a command that opens a video player to play a video from a specified URL
func PlayVideo(videoUrl, nameEpisode string) {
	cmd := exec.Command("mpv", "--save-position-on-quit", "--no-terminal", "--fs", fmt.Sprintf("--force-media-title=%v", nameEpisode), "--cache=yes", videoUrl)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	// Set Discord user presence to show default presence
	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")
}
