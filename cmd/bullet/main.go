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

	app.Name = "bullet"
	app.Version = "0.0.1"
	app.Usage = "A command-line tool to send text or files to Pushbullet."

	app.Commands = []cli.Command{
		commands.NewSendCommand(),
	}

	return app
}
