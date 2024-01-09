package main

// Fazer opção para voltar para a lista de animes.

import (
	"os"
	"time"

	"github.com/Caixetadev/gophimation/internal/presence"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/scrapers"
)

func init() {
	go utils.CleanCache(time.Hour * 3)
	go presence.Presence("https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	rootDir := utils.GetCacheDir("user")

	_, error := os.Stat(rootDir)

	if os.IsNotExist(error) {
		utils.CreateFile(rootDir, constants.USER_SETTINGS_FILE)
	} else {
		go utils.ReadFile(rootDir, constants.USER_SETTINGS_FILE)
	}
}

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "random":
		scrapers.Random()

	case len(os.Args) > 1 && os.Args[1] == "--delete-cache":
		go utils.CleanCache(time.Second * 1)
		animeMostWatched := scrapers.MostWatched()
		episodeSelected, nextEpisode := scrapers.SelectEpisode(animeMostWatched)
		videoSelected := scrapers.SelectVideo(episodeSelected)
		if nextEpisode != nil {
			go scrapers.SelectVideo(*nextEpisode)
		}
		utils.PlayVideo(videoSelected.Url, videoSelected.Name)
	case len(os.Args) > 1:
		animeSearch := scrapers.Search()
		episodeSelected, nextEpisode := scrapers.SelectEpisode(animeSearch)
		videoSelected := scrapers.SelectVideo(episodeSelected)
		if nextEpisode != nil {
			go scrapers.SelectVideo(*nextEpisode)
		}
		utils.PlayVideo(videoSelected.Url, videoSelected.Name)
	default:
		animeMostWatched := scrapers.MostWatched()
		episodeSelected, nextEpisode := scrapers.SelectEpisode(animeMostWatched)
		videoSelected := scrapers.SelectVideo(episodeSelected)
		if nextEpisode != nil {
			go scrapers.SelectVideo(*nextEpisode)
		}
		utils.PlayVideo(videoSelected.Url, videoSelected.Name)
	}
}
