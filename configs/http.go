package configs

import (
	"net/http"

	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func Http() *http.Client {
	cacheDir := util.GetHomeDir()

	cache := diskcache.New(cacheDir)

	client := &http.Client{
		Transport: httpcache.NewTransport(cache),
	}

	return client
}
