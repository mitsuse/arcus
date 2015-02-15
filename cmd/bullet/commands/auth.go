package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet"
	"github.com/mitsuse/bullet/cmd/bullet/app"
	"github.com/mitsuse/bullet/pushbullet"
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
		c, err := app.LoadConfigPath(configPath)
		if err != nil {
			app.PrintError(err)
			return
		}

		config = c
	} else {
		fmt.Println("Create a new config file.")
		config = bullet.NewConfig()
	}

	for {
		fmt.Print("access token: ")
		token, err := readToken()
		if err != nil {
			app.PrintError(err)
			return
		}
		config.SetToken(token)

		pb := pushbullet.New(config.Token())
		if _, err := pb.GetUsersMe(); err == nil {
			break
		}
		fmt.Println("The access token is invalid.")
	}

	if err := app.DumpConfigPath(config, configPath); err != nil {
		app.PrintError(err)
		return
	}
}

func readToken() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		app.PrintError(err)
		return "", err
	}

	return scanner.Text(), nil
}
