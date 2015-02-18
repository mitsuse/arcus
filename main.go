package main

import (
	"os"
	"os/user"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/app"
	"github.com/mitsuse/bullet/commands"
)

func main() {
	cmd, err := initCmd()
	if err != nil {
		app.PrintError(err)
		return
	}

	cmd.Run(os.Args)
}

func initCmd() (*cli.App, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	cmd := cli.NewApp()

	cmd.Name = app.NAME
	cmd.Version = app.VERSION
	cmd.Usage = app.DESC
	cmd.Author = app.AUTHOR
	cmd.Email = app.AUTHOR_EMAIL

	cmd.Commands = []cli.Command{
		commands.NewAuthCommand(u),
		commands.NewSendCommand(u),
		commands.NewListCommand(u),
	}

	return cmd, nil
}
