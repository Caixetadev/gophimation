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
		Value:  "eyJpdiI6ImNITkd2ZDdCQkUwTk4yOEptVUZ3RXc9PSIsInZhbHVlIjoiNVhzelQ5SFJtQUtjSkJIYkV5ZDVSK0Z4UDNpNEV3SVdIcGZhRVFJV1MrajBRMUdPN0NYejY5eSsvNkdPM05zcktyOUZWb2ZZMEpMdWdlMjdLK01HdUlaY2dMTW1QeG5uRWFJeTFhRCtKT2FDUWFLSnRxekFWSUJBY3lsMW1LOXEiLCJtYWMiOiIzYjhmNThkNWY1NjA3NDg5MDVkNDllZWZlZTVmMjUwODQ3NDIwYzE0MzVkODM0NDk2NGQ4YTU5OTRmZjJhMjM5IiwidGFnIjoiIn0%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "BetterQuality",
		Value:  "eyJpdiI6ImRlN2R6UjZTV052eUVKbEE0YzVwZ1E9PSIsInZhbHVlIjoiZ2UzNEF3STlZeEg3NHlvbU1BQzU0bHBOemNuQ0JWbzNXbXdqT2F6YzJlTnNRakc2dzJESEdNQjYxZk1FLzdPbyIsIm1hYyI6ImY0NTY3ZjljZWVjODI4NmYzYzA4OTVhNThkNjZkYmIxZWIwNjY4NTEzNTg2YmY3NmRmODU4MDkxZGEwNzY5MDkiLCJ0YWciOiIifQ%3D%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie, cookie2})

	return c
}
