package config

import (
	"encoding/json"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/imdario/mergo"
)

type HttpConfig struct {
	Port       int           `json:"port"`
	Root       string        `json:"root"`
	TimeOutStr string        `json:"timeout"`
	TimeOut    time.Duration `json:"-"`
	Workers    int           `json:"workers"`
}

type FilterConfig struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type ImageConfig struct {
	Quality int `json:"quality"`
}

type StorageConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type CacheConfig struct {
	Source int `json:"source"`
}

type Config struct {
	Http     HttpConfig      `json:"http"`
	Filters  []FilterConfig  `json:"filters"`
	Image    ImageConfig     `json:"image"`
	Storages []StorageConfig `json:"storages"`
	Cache    CacheConfig     `json:"caches"`
}

var cfg = &Config{
	Http: HttpConfig{
		Port:    80,
		Root:    "/impressionist",
		TimeOut: 30 * time.Second,
		Workers: 10,
	},
	Image: ImageConfig{
		Quality: 75,
	},
	Cache: CacheConfig{
		Source: 100,
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
	if user.Http.TimeOutStr != "" {
		user.Http.TimeOut, err = time.ParseDuration(user.Http.TimeOutStr)
		if err != nil {
			log.Warnf("Invalid timeout : %s, ignoring", user.Http.TimeOutStr)
		}
	}
	if err := mergo.Merge(cfg, user); err != nil {
		panic(err)
	}
	return cfg
}

func Get() *Config {
	return cfg
}
