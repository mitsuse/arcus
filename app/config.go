package app

import (
	"encoding/json"
	"io"
	"os"
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

func LoadConfigPath(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadConfig(file)
}

func DumpConfigPath(config *Config, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return DumpConfig(config, file)
}
