package storage

import (
	"github.com/jfbus/impressionist/img"
	"github.com/jfbus/impressionist/log"
	"github.com/thoas/gokvstores"
)

var cache gokvstores.KVStoreConnection

func InitCache(size int) {
	if size > 0 {
		log.Infof("Initializing source cache with %d entries", size)
		cache = gokvstores.NewCacheKVStoreConnection(size)
	} else {
		log.Info("Source cache is disabled")
	}
}

func getCacheKey(storage, file string) string {
	return storage + "##" + file
}

func getFromCache(storage, file string) (img.Img, bool) {
	if cache == nil {
		return nil, false
	}
	k := getCacheKey(storage, file)
	if !cache.Exists(k) {
		return nil, false
	}
	i, ok := cache.Get(k).(img.Img)
	return i, ok
}

func setToCache(storage, file string, i img.Img) {
	if cache != nil {
		cache.Set(getCacheKey(storage, file), i)
	}
}
