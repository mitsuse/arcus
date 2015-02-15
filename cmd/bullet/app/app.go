package app

import (
	"fmt"
	"os"
)

const (
	NAME    = "bullet"
	VERSION = "0.0.4"
	DESC    = "A command-line tool to send text or files to Pushbullet."
)

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
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
