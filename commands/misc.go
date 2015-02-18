package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func configFlag() cli.StringFlag {
	homePath := os.Getenv("HOME")
	configName := ".config.bullet"

	flag := cli.StringFlag{
		Name:  "config,c",
		Value: fmt.Sprintf("%s/%s", homePath, configName),
		Usage: "The path of your config file",
	}

	return flag
}
