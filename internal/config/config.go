package config

import (
	"log"
	"sync"

	"github.com/kkyr/fig"
)

var once = &sync.Once{}
var config *Config

type Config struct {
	DataSource string `fig:"data_source"`
}

func GetConfig() *Config {
	once.Do(func() {
		config = new(Config)
		err := fig.Load(config)
		if err != nil {
			log.Fatal(err)
		}
	})
	return config
}
