package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/arcus/application"
)

/*
Create "send" sub-command.
This sub-command sends a "push" to devices such as "note", "link" or "file".
*/
func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "Send a message or a file",
		Action:    actionSend,

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
	}

	return command
}

func actionSend(ctx *cli.Context) {
	token := os.Getenv("ARCUS_ACCESS_TOKEN")
	title := ctx.String("title")
	message := ctx.String("message")
	location := ctx.String("location")
	device := ctx.String("device")

	if err := application.Send(token, title, message, location, device); err != nil {
		application.ExitWith(err)
		return
	}
}
