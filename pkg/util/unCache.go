package util

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func UnCache(URL string) {
	log.Println("Trying to remove cached response for:", URL)

	// obter o diretório home do usuário
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", "gophimation")

	sum := sha1.Sum([]byte(URL))
	hash := hex.EncodeToString(sum[:])
	dir := path.Join(cacheDir, hash[:2])
	filename := path.Join(dir, hash)
	log.Println("Deleting cached file:", filename)
	if err := os.Remove(filename); err != nil {
		log.Fatal(err)
	}
}
