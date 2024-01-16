package utils

import (
	"os/user"
	"path/filepath"

	"github.com/Caixetadev/gophimation/pkg/constants"
)

// GetCacheDir retrieves the cache directory. It expects a parameter 'folderName' that specifies the desired folder name.
func GetCacheDir(folderName string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache", constants.FOLDER_NAME, folderName)

	return cacheDir
}
