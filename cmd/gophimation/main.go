package main

import (
	"os"

	"github.com/Caixetadev/gophimation/pkg/episode"
	mostWatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/search"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
)

// func init() {
// 	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

// 	util.Clear()

// 	_, error := os.Stat(constants.FILE_NAME)

// 	if os.IsNotExist(error) {
// 		util.CreateFile(constants.FILE_NAME)
// 	} else {
// 		util.ReadFile(constants.FILE_NAME)
// 	}
// }

func main() {
	if len(os.Args) == 1 {
		animeSelected := mostWatched.MostWatched()
		watchEpisode(animeSelected)
	} else {
		watchEpisode("")
	}
}

func watchEpisode(previousSearch string) {
	var search2 string
	if previousSearch != "" {
		search2 = previousSearch
	} else {
		search2 = search.Search()
	}

	episodeSelected := episode.SelectEpisode(search2)

	selectVideo.SelectVideo(episodeSelected, "caixeta aiinnn")
}
