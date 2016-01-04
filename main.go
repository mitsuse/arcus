package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/arcus/application"
)

const (
	NAME         = `arcus`
	VERSION      = "0.1.3"
	DESCRIPTION  = "A command-line tool to send a message to devices via Pushbullet."
	AUTHOR       = "Tomoya Kose (mitsuse)"
	AUTHOR_EMAIL = "tomoya@mitsuse.jp"

	variableToken = "ARCUS_ACCESS_TOKEN"
)

func main() {
	app := cli.NewApp()

	app.Name = NAME
	app.Version = VERSION
	app.Usage = DESCRIPTION
	app.Author = AUTHOR
	app.Email = AUTHOR_EMAIL

	app.Commands = []cli.Command{
		newSendCommand(),
		newListCommand(),
	}

	app.Run(os.Args)
}

func newListCommand() cli.Command {
	command := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "List devices that can be pushed to",

		Action: func(c *cli.Context) {
			token := getToken()

			devices, err := application.ListDevices(token)
			if err != nil {
				printError(err)
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
			token := getToken()
			title := c.String("title")
			message := c.String("message")
			location := c.String("location")
			device := c.String("device")

			if err := application.Send(token, title, message, location, device); err != nil {
				printError(err)
				return
			}
		},
	}

	return command
}

func getToken() string {
	return os.Getenv(variableToken)
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
}
