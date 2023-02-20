package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Caixetadev/gophimation/constants"
	"github.com/Caixetadev/gophimation/episode"
	mostwatched "github.com/Caixetadev/gophimation/mostWatched"
	"github.com/Caixetadev/gophimation/presence"
	"github.com/Caixetadev/gophimation/search"
	"github.com/Caixetadev/gophimation/selectVideo"
	"github.com/Caixetadev/gophimation/utils"
)

func init() {
	utils.Clear()

	_, error := os.Stat(constants.FILE_NAME)

	if os.IsNotExist(error) {
		utils.CreateFile(constants.FILE_NAME)
	} else {
		utils.ReadFile(constants.FILE_NAME)
	}
}

func main() {
	if len(os.Args) == 1 {
		seila := mostwatched.MostWatched()
		watchEpisode(seila)
	} else {
		watchEpisode("")
	}
}

func watchEpisode(previousSearch string) {
	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	var search2 string
	if previousSearch != "" {
		search2 = previousSearch
	} else {
		search2 = search.Search()
	}

	epsisae := episode.SelectEpisode(search2)
	selectVideo.SelectVideo(epsisae)

	var resp string

	var currentEpisode int
	regex := regexp.MustCompile(`animes\/(.*?)\/(\d+)`)
	matches := regex.FindStringSubmatch(epsisae)
	animeName := matches[1]
	currentEpisode, _ = strconv.Atoi(matches[2])

	for {
		utils.Clear()

		fmt.Println("(n) proximo")
		fmt.Println("(q) sair")
		fmt.Scanln(&resp)

		switch strings.ToLower(resp) {
		case "n":
			nextEpisode := currentEpisode + 1
			nextEpisodeURL := fmt.Sprintf("https://animefire.net/animes/%s/%d", animeName, nextEpisode)
			selectVideo.SelectVideo(nextEpisodeURL)
			currentEpisode = nextEpisode
		case "q":
			return
		default:
			fmt.Println("Opção inválida")
			time.Sleep(2 * time.Second)
		}
	}
}
