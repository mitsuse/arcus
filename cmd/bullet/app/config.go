package app

import (
	"encoding/json"
	"io"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfig(reader io.Reader) (*Config, error) {
	decoder := json.NewDecoder(reader)

	var config *Config
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func DumpConfig(c *Config, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(c)
}
