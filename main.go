package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/arcus/application"
	"github.com/mitsuse/arcus/commands"
)

func main() {
	app := createApp()
	app.Run(os.Args)
}

func createApp() *cli.App {
	app := cli.NewApp()

	app.Name = application.NAME
	app.Version = application.VERSION
	app.Usage = application.DESCRIPTION
	app.Author = application.AUTHOR
	app.Email = application.AUTHOR_EMAIL

	app.Commands = []cli.Command{
		commands.NewSendCommand(),
		commands.NewListCommand(),
	}

	return app
}
