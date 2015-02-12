package commands

import (
	"github.com/codegangsta/cli"
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
			configFlag(),
			titleFlag("The title of a note to be sent"),
			bodyFlag("The body of a note to be sent"),
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
