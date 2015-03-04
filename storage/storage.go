package storage

import (
	"errors"

	"github.com/jfbus/impressionist/config"
	"github.com/thoas/gostorages"
)

var (
	storages           = map[string]gostorages.Storage{}
	ErrStorageNotFound = errors.New("storage not found")
)

func Init(cfg []config.StorageConfig, cacheSize int) {
	for _, s := range cfg {
		switch s.Type {
		case "local":
			storages[s.Name] = gostorages.NewFileSystemStorage(s.Path, "")
		default:
			panic("no such type " + s.Type)
		}
	}
	InitCache(cacheSize)
}

func Get(name string) (gostorages.Storage, error) {
	if s, ok := storages[name]; ok {
		return s, nil
	}
	return nil, ErrStorageNotFound
}
