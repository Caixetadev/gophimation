package configs

import (
	"os/user"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

func Colly() *colly.Collector {
	// obter o diretório home do usuário
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", "gophimation")

	c := colly.NewCollector(
		colly.CacheDir(cacheDir),
		colly.AllowedDomains("animefire.net", "www.blogger.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	return c
}
