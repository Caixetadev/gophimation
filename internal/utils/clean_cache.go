package utils

import (
	"os"
	"time"
)

// CleanCache clears the cache, expecting a parameter named cacheDuration of type time.Duration
// that determines the maximum duration a cache file is allowed to exist before being deleted.
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
