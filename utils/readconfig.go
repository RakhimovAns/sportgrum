package utils

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
}

func ReadConfig() (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("ReadConfig : %w", err)
	}
	return cfg, nil
}
