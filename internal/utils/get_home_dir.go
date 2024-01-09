package utils

import (
	"os/user"
	"path/filepath"

	"github.com/Caixetadev/gophimation/pkg/constants"
)

func GetCacheDir(folderName string) string {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", constants.FOLDER_NAME, folderName)

	return cacheDir
}
