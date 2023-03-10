package util

import (
	"os/user"
	"path/filepath"
)

func GetHomeDir(fileName string) string {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", fileName)

	return cacheDir
}
