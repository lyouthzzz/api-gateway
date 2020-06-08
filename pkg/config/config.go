package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	Debug    bool `env:"DEBUG"`
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     int    `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PWD"`
		Name     string `env:"DB_NAME"`
	}
}

func LoadConfig(fileName string, config interface{}) (err error) {
	err = configor.Load(config, fileName)
	return
}
