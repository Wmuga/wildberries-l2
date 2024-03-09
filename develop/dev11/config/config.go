package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Address string `json:"address"`
}

func ReadConfig(filename string) (cfg Config, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	json.Unmarshal(data, &cfg)
	return
}
