package commands

import (
	"github.com/codegangsta/cli"
)

func NewAuthCommand() cli.Command {
	command := cli.Command{
		Name:      "auth",
		ShortName: "a",
		Usage:     "Update your access token",
		Action:    actionAuth,

		Flags: []cli.Flag{
			configFlag(),
		},
	}

	return command
}

func actionAuth(ctx *cli.Context) {
	// TODO: Implement this.
}
