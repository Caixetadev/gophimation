package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Caixetadev/gophimation/presence"
)

func PlayVideo(videoUrl, nameEpisode string) {
	cmd := exec.Command("mpv", "--save-position-on-quit", "--no-terminal", "--fs", fmt.Sprintf("--force-media-title=%v", nameEpisode), "--cache=yes", videoUrl)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")
}
