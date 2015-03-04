package config

import (
	"encoding/json"
	"os"

	"github.com/imdario/mergo"
)

type HttpConfig struct {
	Port int    `json:"port"`
	Root string `json:"root"`
}

type FilterConfig struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type JPEGConfig struct {
	Quality int `json:"quality"`
}

type StorageConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type Config struct {
	Http     HttpConfig      `json:"http"`
	Filters  []FilterConfig  `json:"filters"`
	JPEG     JPEGConfig      `json:"jpeg"`
	Storages []StorageConfig `json:"storages"`
}

var defaults = Config{
	Http: HttpConfig{
		Port: 80,
		Root: "/impressionist",
	},
	JPEG: JPEGConfig{
		Quality: 80,
	},
}

func Load(file string) *Config {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(fd)
	user := Config{}
	err = dec.Decode(&user)
	if err != nil {
		panic(err)
	}
	cfg := Config{}
	if err := mergo.Merge(&cfg, defaults); err != nil {
		panic(err)
	}
	if err := mergo.Merge(&cfg, user); err != nil {
		panic(err)
	}
	return &cfg
}
