package main

// Fazer opção para voltar para a lista de animes.

import (
	"os"

	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/episode"
	mostWatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/random"
	"github.com/Caixetadev/gophimation/pkg/search"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
	"github.com/Caixetadev/gophimation/pkg/util"
)

func init() {
	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	_, error := os.Stat(constants.FILE_NAME)

	if os.IsNotExist(error) {
		util.CreateFile(constants.FILE_NAME)
	} else {
		util.ReadFile(constants.FILE_NAME)
	}
}

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "random":
		random.Random()

	case len(os.Args) == 1:
		watchEpisode(mostWatched.MostWatched())

	default:
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

	selectVideo.SelectVideo(episodeSelected)
}
