package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/episode"
	mostwatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/search"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
	"github.com/Caixetadev/gophimation/pkg/util"
)

func init() {
	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")

	util.Clear()

	_, error := os.Stat(constants.FILE_NAME)

	if os.IsNotExist(error) {
		util.CreateFile(constants.FILE_NAME)
	} else {
		util.ReadFile(constants.FILE_NAME)
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
	var search2 string
	if previousSearch != "" {
		search2 = previousSearch
	} else {
		search2 = search.Search()
	}

	var currentEpisode int

	epsisae, nameEpisode, imageAnime, nameAnime, _, _ := episode.SelectEpisode(search2)

	var resp string

	regex := regexp.MustCompile(`animes\/(.*?)\/(\d+)`)
	matches := regex.FindStringSubmatch(epsisae)
	animeName := matches[1]
	currentEpisode, _ = strconv.Atoi(matches[2])

	episode := util.FormatEpisode(currentEpisode)

	selectVideo.SelectVideo(epsisae, nameAnime+" - "+episode)

	for {
		util.Clear()

		fmt.Println("(n) proximo")
		fmt.Println("(q) sair")
		fmt.Scanln(&resp)

		switch strings.ToLower(resp) {
		case "n":
			nextEpisode := currentEpisode + 1
			nextEpisodeURL := fmt.Sprintf("https://animefire.net/animes/%s/%d", animeName, nextEpisode)
			episode := util.FormatEpisode(nextEpisode)
			presence.Presence(nameEpisode, imageAnime, "Assistindo "+nameAnime, episode, "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")
			selectVideo.SelectVideo(nextEpisodeURL, nameAnime+" - "+episode)
			currentEpisode = nextEpisode
		case "q":
			return
		default:
			fmt.Println("Opção inválida")
			time.Sleep(2 * time.Second)
		}
	}
}
