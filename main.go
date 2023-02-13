package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.anitube.site", "rr1---sn-gx5auxaxjvhxpgxap-btoe.googlevideo.com", "www.blogger.com"),
	)

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
				re := regexp.MustCompile(`https://rr1---sn-gx5auxaxjvhxpgxap-btoe[\S]+?"`)
				src := re.FindAllStringSubmatch(e.Text, 2)[1]

				url := strings.Join(src, "")

				cmd := exec.Command("/usr/bin/mpv", strings.TrimRight(url, `"`))
				fmt.Println(cmd)
				cmd.Stdout = os.Stdout

				cmd.Run()
			})
		}
	})

	c.Visit("https://www.anitube.site/918795/")
}
