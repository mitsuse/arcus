package commands

import (
	"fmt"
	"os/user"
	"runtime"

	"github.com/codegangsta/cli"
)

func configFlag(u *user.User) cli.StringFlag {
	var configName string
	if runtime.GOOS == "windows" {
		configName = "_config.bullet"
	} else {
		configName = ".config.bullet"
	}

	flag := cli.StringFlag{
		Name:  "config,c",
		Value: fmt.Sprintf("%s/%s", u.HomeDir, configName),
		Usage: "The path of your config file",
	}

	return flag
}
