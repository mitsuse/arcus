package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet"
)

const (
	APP_NAME    = "bullet"
	APP_VERSION = "0.0.4"
	APP_DESC    = "A command-line tool to send text or files to Pushbullet."
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", APP_NAME, err)
}

func loadConfigPath(path string) (*bullet.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bullet.LoadConfig(file)
}

func dumpConfigPath(config *bullet.Config, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return bullet.DumpConfig(config, file)
}

func configFlag() cli.StringFlag {
	flag := cli.StringFlag{
		Name:  "config,c",
		Value: os.Getenv("HOME") + "/.config.bullet",
		Usage: "The path of your config file",
	}

	return flag
}
