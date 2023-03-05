package util

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"os"
	"path"
)

func UnCache(URL string) {
	log.Println("Trying to remove cached response for:", URL)

	cacheDir := GetHomeDir()

	sum := sha1.Sum([]byte(URL))
	hash := hex.EncodeToString(sum[:])
	dir := path.Join(cacheDir, hash[:2])
	filename := path.Join(dir, hash)
	log.Println("Deleting cached file:", filename)
	if err := os.Remove(filename); err != nil {
		log.Fatal(err)
	}
}
