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
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9,pt;q=0.8")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Host", "tirith.betterbunker.xyz")
		r.Headers.Set("Origin", "https://betteranime.net")
		r.Headers.Set("Referer", "https://betteranime.net/")
		r.Headers.Set("sec-ch-ua", "\"Not_A Brand\";v=\"99\", \"Brave\";v=\"109\", \"Chromium\";v=\"109\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Linux\"")
		r.Headers.Set("Sec-Fetch-Dest", "empty")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Sec-Fetch-Site", "cross-site")
		r.Headers.Set("Sec-GPC", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.3")
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
