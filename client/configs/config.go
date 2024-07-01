package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	User struct {
		UserID  int    `yaml:"id"`
		Balance int    `yaml:"balance"`
		URL     string `yaml:"url"`
		Login   string `yaml:"login"`
	} `yaml:"user"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}
