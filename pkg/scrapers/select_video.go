package scrapers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/gocolly/colly"
)

type PlayerInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// SelectVideo does the search for the url of the video
func SelectVideo(ep string) *PlayerInfo {
	var urlPlayer []PlayerInfo

	c := config.Colly()

	iframeURL, nameAnimeAndEpisode := utils.GetIframe(constants.URL_BASE + ep)

	setCollyCallbacksPlayer(c, &urlPlayer, nameAnimeAndEpisode)

	if err := c.Visit(iframeURL); err != nil {
		log.Fatal(err)
	}

	return &PlayerInfo{Name: urlPlayer[0].Name, Url: urlPlayer[0].Url}
}

func setCollyCallbacksPlayer(c *colly.Collector, player *[]PlayerInfo, nameAnimeAndEpisode string) {
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", constants.URL_BASE)
	})

	c.OnHTML("script:nth-of-type(4)", func(h *colly.HTMLElement) {
		re := regexp.MustCompile(`"file":"([^"]+)"`)
		match := re.FindStringSubmatch(h.Text)

		if len(match) > 1 {
			*player = append(*player, PlayerInfo{Name: nameAnimeAndEpisode, Url: strings.ReplaceAll(match[1], "\\", "")})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})
}
