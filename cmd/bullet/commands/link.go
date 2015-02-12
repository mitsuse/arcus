package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewLinkCommand() cli.Command {
	command := cli.Command{
		Name:      "link",
		ShortName: "l",
		Usage:     "Send a link",
		Action:    actionLink,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config,c",
				Value: os.Getenv("HOME") + "/.config.bullet",
				Usage: "The path of your config file",
			},

			cli.StringFlag{
				Name:  "title,t",
				Value: "",
				Usage: "The title of a link to be sent",
			},

			cli.StringFlag{
				Name:  "body,b",
				Value: "",
				Usage: "The message associated with a link to be sent",
			},

			cli.StringFlag{
				Name:  "url,u",
				Value: "",
				Usage: "The url to be sent",
			},
		},
	}

	return command
}

func actionLink(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	title := ctx.String("title")
	body := ctx.String("body")
	url := ctx.String("url")
	link := pushes.NewLink(title, body, url)

	pb := pushbullet.New(config.Token())

	if err := pb.PostLink(link); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}
