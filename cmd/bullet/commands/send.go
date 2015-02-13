package commands

import (
	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "Send a message or a file",
		Action:    actionSend,

		Flags: []cli.Flag{
			configFlag(),

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
				Name:  "link,l",
				Value: "",
				Usage: "The link to be sent",
			},
		},
	}

	return command
}

func actionSend(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	pb := pushbullet.New(config.Token())

	title := ctx.String("title")
	message := ctx.String("message")
	path := ctx.String("path")

	if err := send(pb, title, message, path); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}

func send(pb *pushbullet.Pushbullet, title, message, path string) error {
	if len(path) == 0 {
		note := pushes.NewNote(title, message)
		return pb.PostNote(note)
	} else {
		link := pushes.NewLink(title, message, path)
		return pb.PostLink(link)
	}
}
