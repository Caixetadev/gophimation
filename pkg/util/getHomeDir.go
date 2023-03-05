package util

import (
	"os/user"
	"path/filepath"
)

func GetHomeDir() string {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", "gophimation")

	return cacheDir
}
