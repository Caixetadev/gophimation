package main

// Fazer opção para voltar para a lista de animes.

import (
	"os"

	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/episode"
	"github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/random"
	"github.com/Caixetadev/gophimation/pkg/search"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
	"github.com/Caixetadev/gophimation/pkg/util"
)

func init() {
	go presence.Presence("https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	pathFile := util.GetHomeDir(constants.FILE_NAME)

	_, error := os.Stat(pathFile)

	if os.IsNotExist(error) {
		util.CreateFile(pathFile)
	} else {
		util.ReadFile(pathFile)
	}
}

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "random":
		random.Random()

	case len(os.Args) > 1:
		animeSearch := search.Search()
		episodeSelected, nextEpisode := episode.SelectEpisode(animeSearch)
		videoSelected := selectVideo.SelectVideo(episodeSelected)
		if nextEpisode != nil {
			go selectVideo.SelectVideo(*nextEpisode)
		}
		util.PlayVideo(videoSelected.Url, videoSelected.Name)
	default:
		animeMostWatched := mostWatched.MostWatched()
		episodeSelected, nextEpisode := episode.SelectEpisode(animeMostWatched)
		videoSelected := selectVideo.SelectVideo(episodeSelected)
		if nextEpisode != nil {
			go selectVideo.SelectVideo(*nextEpisode)
		}
		util.PlayVideo(videoSelected.Url, videoSelected.Name)
	}

}
