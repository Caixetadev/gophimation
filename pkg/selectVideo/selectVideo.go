package selectVideo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

type PlayerInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// SelectVideo does the search for the url of the video
func SelectVideo(ep string) {
	c := configs.Colly()

	iframeURL, nameAnimeAndEpisode := util.GetIframe(constants.URL_BASE + ep)

	var urlPlayer []PlayerInfo

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", constants.URL_BASE)
	})

	c.OnHTML("script:nth-of-type(4)", func(h *colly.HTMLElement) {
		re := regexp.MustCompile(`"file":"([^"]+)"`)
		match := re.FindStringSubmatch(h.Text)

		if len(match) > 1 {
			urlPlayer = append(urlPlayer, PlayerInfo{Name: nameAnimeAndEpisode, Url: strings.ReplaceAll(match[1], "\\", "")})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Visit(iframeURL)

	util.PlayVideo(urlPlayer[0].Url, urlPlayer[0].Name)
}
