package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet"
)

func NewAuthCommand() cli.Command {
	command := cli.Command{
		Name:      "auth",
		ShortName: "a",
		Usage:     "Update your access token",
		Action:    actionAuth,

		Flags: []cli.Flag{
			configFlag(),
		},
	}

	return command
}

func actionAuth(ctx *cli.Context) {
	configPath := ctx.String("config")

	var config *bullet.Config
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println("Update the existing config file.")
		c, err := loadConfigPath(configPath)
		if err != nil {
			printError(err)
			return
		}

		config = c
	} else {
		fmt.Println("Create a new config file.")
		config = bullet.NewConfig()
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	fmt.Print("access token: ")
	if scanned := scanner.Scan(); !scanned {
		return
	}

	if err := scanner.Err(); err != nil {
		printError(err)
		return
	}
	token := scanner.Text()

	config.SetToken(token)

	if err := dumpConfigPath(config, configPath); err != nil {
		printError(err)
		return
	}
}
