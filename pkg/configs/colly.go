package configs

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gocolly/colly"
)

func Colly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("betteranime.net"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}

	c.SetCookieJar(cookieJar)

	url, err := url.Parse("https://betteranime.net")

	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:   "betteranime_session",
		Value:  "eyJpdiI6IncrYXdVUHVpTGpoRi9nOTdFazB4N3c9PSIsInZhbHVlIjoic0w4eXYrTEpJL1R6Nkx4TjRtaEZWeFFpamhNTUk3VStSd1ZzRXh5MUZodDBRZ1VZV2N2ejdlRlF4V293RlBydHZTaHpETnRXUEdCOGMyN0haM1dnTTgzc0YwNGVGRDhEYnRPT2JkS1dEbExDYlA1UlZCUmVOUER6VWxxUmNhRngiLCJtYWMiOiIyNmRkNGM5MTUyYTliMWEzN2I2ZjIzM2VlYWY2MzRlZmRiZWQ5ZmMxYmU2OTU1OWE5ZjJlNDJiZDMxNTE0OWJjIiwidGFnIjoiIn0%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "BetterQuality",
		Value:  "eyJpdiI6IlhNbnFZbVMwV2pHK2xNU1FBTUd1dEE9PSIsInZhbHVlIjoibmh4bVVBT1h2TVBJaEFXSzI3Q2xlUVlNMWxZTlpOSEZtUHBtcVlxa0hDSlRteHpQS083ZmpDQ2xramFOc29NciIsIm1hYyI6Ijg2NzA3M2ZkZDQ4MWY3YzgxNjIzZDQxZjViYzJhNGRjNGY5MTk3NzVhNzI2NWJkYzUxOGNiMDdmYzQyNTA2NGUiLCJ0YWciOiIifQ%3D%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie, cookie2})

	return c
}
