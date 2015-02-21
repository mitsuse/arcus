package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/commands"
)

func main() {
	cmd := initCmd()
	cmd.Run(os.Args)
}

func initCmd() *cli.App {
	cmd := cli.NewApp()

	cmd.Name = commands.NAME
	cmd.Version = commands.VERSION
	cmd.Usage = commands.DESC
	cmd.Author = commands.AUTHOR
	cmd.Email = commands.AUTHOR_EMAIL

	cmd.Commands = []cli.Command{
		commands.NewSendCommand(),
		commands.NewListCommand(),
	}

	return cmd
}
