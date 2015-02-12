package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/cmd/bullet/commands"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = commands.APP_NAME
	app.Version = commands.APP_VERSION
	app.Usage = commands.APP_DESC

	app.Commands = []cli.Command{
		commands.NewNoteCommand(),
		commands.NewLinkCommand(),
	}

	return app
}
