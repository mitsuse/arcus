package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/arcus/application"
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
		newSendCommand(),
		newListCommand(),
	}

	return app
}

func newListCommand() cli.Command {
	command := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "List devices that can be pushed to",

		Action: func(c *cli.Context) {
			token := os.Getenv("ARCUS_ACCESS_TOKEN")

			devices, err := application.ListDevices(token)
			if err != nil {
				application.ExitWith(err)
				return
			}

			for _, d := range devices {
				if !d.Pushable {
					continue
				}

				fmt.Println(d.Nickname)
			}
		},
	}

	return command
}

func newSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "Send a message or a file",

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "device,d",
				Value: "",
				Usage: "The name of target device",
			},

			cli.StringFlag{
				Name:  "title,t",
				Value: "",
				Usage: "The title of the message or file to be sent",
			},

			cli.StringFlag{
				Name:  "message,m",
				Value: "",
				Usage: "The message to be sent",
			},

			cli.StringFlag{
				Name:  "location,l",
				Value: "",
				Usage: "The path of file or link to be sent",
			},
		},

		Action: func(c *cli.Context) {
			token := os.Getenv("ARCUS_ACCESS_TOKEN")
			title := c.String("title")
			message := c.String("message")
			location := c.String("location")
			device := c.String("device")

			if err := application.Send(token, title, message, location, device); err != nil {
				application.ExitWith(err)
				return
			}
		},
	}

	return command
}
