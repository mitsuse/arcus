package main

import (
	"os"

	"github.com/codegangsta/cli"
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

	return app
}
