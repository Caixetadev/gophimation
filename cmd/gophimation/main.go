package main

// Fazer opção para voltar para a lista de animes.

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Caixetadev/gophimation/internal/entity"
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
		utils.ReadFile(rootDir, constants.USER_SETTINGS_FILE)
	}
}

func getSelectedVideo(episodes []entity.Anime, index int) {
	videoSelected := scrapers.SelectVideo(episodes[index].URL)
	utils.PlayVideo(videoSelected.Url, videoSelected.Name)
}

func semNome(animeSelected string) (int, []entity.Anime) {
	nextEpisode, currentEpIndex, episodes := scrapers.SelectEpisode(animeSelected)
	getSelectedVideo(episodes, currentEpIndex)
	if nextEpisode != nil {
		go scrapers.SelectVideo(*nextEpisode)
	}

	return currentEpIndex, episodes
}

func app() {
	var watchedEpisodes []entity.Anime
	var currentEpisodeIndex int
	var searchResult string

	switch {
	// case len(os.Args) > 1 && os.Args[1] == "random":
	// 	nextEpisode, currentEpIndex, episodes := scrapers.SelectEpisode("random")
	// 	getSelectedVideo(episodes, currentEpIndex)
	// 	if nextEpisode != nil {
	// 		go scrapers.SelectVideo(*nextEpisode)
	// 	}
	//
	// 	currentEpisodeIndex = currentEpIndex
	// 	watchedEpisodes = episodes
	// searchResult = animeMostWatched
	case len(os.Args) > 1 && os.Args[1] == "--delete-cache":
		go utils.CleanCache(time.Second * 1)
		animeMostWatched := scrapers.MostWatched()
		currentEpisodeIndex, watchedEpisodes = semNome(animeMostWatched)
	case len(os.Args) > 1:
		animeSearch := scrapers.Search()
		currentEpisodeIndex, watchedEpisodes = semNome(animeSearch)
	default:
		animeMostWatched := scrapers.MostWatched()
		currentEpisodeIndex, watchedEpisodes = semNome(animeMostWatched)
	}

	utils.Clear()

	for {
		var optionSelected string

		options := []string{
			"[n] - Play Next Episode\n",
			"[b] - Go Back to Previous Episode\n",
			"[r] - Replay Current Episode\n",
			"[s] - Return to Episode List\n",
			"[q] - Quit\n",
		}

		fmt.Println(strings.Join(options, ""))

		fmt.Scanln(&optionSelected)

		switch strings.ToLower(optionSelected) {
		case "n":
			currentEpisodeIndex++

			getSelectedVideo(watchedEpisodes, currentEpisodeIndex)

			if currentEpisodeIndex-1 <= len(watchedEpisodes) {
				go scrapers.SelectVideo(watchedEpisodes[(currentEpisodeIndex+1)-1].URL)
			}
		case "r":
			getSelectedVideo(watchedEpisodes, currentEpisodeIndex)
		case "s":
			nextEpisode, currentEpIndex, episodes := scrapers.SelectEpisode(searchResult)
			getSelectedVideo(episodes, currentEpIndex)
			if nextEpisode != nil {
				go scrapers.SelectVideo(*nextEpisode)
			}

			currentEpisodeIndex = currentEpIndex
		case "b":
			currentEpisodeIndex--

			getSelectedVideo(watchedEpisodes, currentEpisodeIndex)
		case "q":
			return
		default:
			utils.Clear()
			fmt.Print("Invalid option. Please try again.\n\n")
		}
	}
}

func main() {
	app()
}
