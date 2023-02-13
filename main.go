package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Caixetadev/ani-go/episode"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.anitube.site", "rr1---sn-gx5auxaxjvhxpgxap-btoe.googlevideo.com", "www.blogger.com"),
	)

	ep := episode.SelectEpisode()

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	var triggerHTML bool

	c.OnHTML("#video .pagEpiAbasContainer", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`src="(.*?)"`)
		src := re.FindStringSubmatch(e.Text)[1]

		triggerHTML = true
		c.Visit(src)
	})

	c.OnResponse(func(r *colly.Response) {
		if triggerHTML {
			c.OnHTML("html", func(e *colly.HTMLElement) {
				res := regexp.MustCompile(`"https://rr[\S]+?"`)
				url := res.FindAllStringSubmatch(e.Text, -1)

				urlstring := strings.Join(res.FindAllStringSubmatch(e.Text, -1)[len(url)-1], "")

				cmd := exec.Command("mpv", strings.Replace(urlstring, `"`, "", -1))
				cmd.Stdout = os.Stdout
				cmd.Run()
			})
		}
	})

	c.Visit(ep)
}
