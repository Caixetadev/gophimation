package cache

import (
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/peterbourgon/diskv/v3"
)

type CacheManager struct {
	d   *diskv.Diskv
	key string
}

func NewCacheManager(key string) *CacheManager {
	d := diskv.New(diskv.Options{
		BasePath:     utils.GetHomeDir("gophimation"),
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 10 * 1024 * 1024,
	})

	return &CacheManager{key: key, d: d}
}

func (c *CacheManager) ReadFromCache() ([]byte, error) {
	data, err := c.d.Read(c.key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *CacheManager) WriteToCache(data []byte) error {
	err := c.d.Write(c.key, data)
	if err != nil {
		return err
	}

	return nil
}
