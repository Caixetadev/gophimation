package config

import (
	"github.com/gocolly/colly/v2"
)

func Colly() *colly.Collector {
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
		colly.AllowedDomains("www.anitube.site", "rr1---sn-gx5auxaxjvhxpgxap-btoe.googlevideo.com", "www.blogger.com", "animefire.net"),
	)

	return c
}
