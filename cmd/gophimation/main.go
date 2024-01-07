package main

// Fazer opção para voltar para a lista de animes.

import (
	"os"

	"github.com/Caixetadev/gophimation/internal/presence"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/scrapers"
)

func init() {
	go presence.Presence("https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	pathFile := utils.GetHomeDir(constants.FILE_NAME)

	_, error := os.Stat(pathFile)

	if os.IsNotExist(error) {
		utils.CreateFile(pathFile)
	} else {
		utils.ReadFile(pathFile)
	}
}

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "random":
		scrapers.Random()

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
