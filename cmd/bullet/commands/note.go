package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewNoteCommand() cli.Command {
	command := cli.Command{
		Name:      "note",
		ShortName: "n",
		Usage:     "Send a note",
		Action:    actionNote,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config,c",
				Value: os.Getenv("HOME") + "/.config.bullet",
				Usage: "The path of your config file",
			},

			cli.StringFlag{
				Name:  "title,t",
				Value: "",
				Usage: "The title of a note to be sent",
			},

			cli.StringFlag{
				Name:  "body,b",
				Value: "",
				Usage: "The body of a note to be sent",
			},
		},
	}

	return command
}

func actionNote(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	title := ctx.String("title")
	body := ctx.String("body")
	note := pushes.NewNote(title, body)

	pb := pushbullet.New(config.Token())

	if err := pb.PostNote(note); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}

func loadConfigPath(path string) (*bullet.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bullet.LoadConfig(file)
}
