package utils

import (
	"os"
	"time"
)

func CleanCache(cacheDuration time.Duration) {
	cacheDir := GetCacheDir("anime")
	folders, err := os.ReadDir(cacheDir)

	if err == nil {
		for _, folder := range folders {
			info, _ := folder.Info()
			if time.Since(info.ModTime()) > cacheDuration {
				go os.RemoveAll(cacheDir + "/" + folder.Name())
			}
		}
	}
}
