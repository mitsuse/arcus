package commands

import (
	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewCheckCommand() cli.Command {
	command := cli.Command{
		Name:      "check",
		ShortName: "c",
		Usage:     "Send a checklist",
		Action:    actionCheck,

		Flags: []cli.Flag{
			configFlag(),
			titleFlag("The list's title"),
		},
	}

	return command
}

func actionCheck(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	title := ctx.String("title")
	check := pushes.NewCheck(title, ctx.Args()...)

	pb := pushbullet.New(config.Token())

	if err := pb.PostCheck(check); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}
