package config

import (
	"net/http"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func Http() *http.Client {
	cache := diskcache.New("cache-directory")

	client := &http.Client{
		Transport: httpcache.NewTransport(cache),
	}

	return client
}
