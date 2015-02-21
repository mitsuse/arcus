package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/app"
	"github.com/mitsuse/bullet/commands"
)

func main() {
	cmd := initCmd()
	cmd.Run(os.Args)
}

func initCmd() *cli.App {
	cmd := cli.NewApp()

	cmd.Name = app.NAME
	cmd.Version = app.VERSION
	cmd.Usage = app.DESC
	cmd.Author = app.AUTHOR
	cmd.Email = app.AUTHOR_EMAIL

	cmd.Commands = []cli.Command{
		commands.NewSendCommand(),
		commands.NewListCommand(),
	}

	return cmd
}
