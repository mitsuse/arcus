package commands

import (
	"os"

	"github.com/codegangsta/cli"
)

func configFlag() cli.StringFlag {
	flag := cli.StringFlag{
		Name:  "config,c",
		Value: os.Getenv("HOME") + "/.config.bullet",
		Usage: "The path of your config file",
	}

	return flag
}
