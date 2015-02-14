package commands

import (
	"github.com/codegangsta/cli"
)

func NewListCommand() cli.Command {
	command := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "List devices that can be pushed to",
		Action:    actionList,

		Flags: []cli.Flag{
			configFlag(),
		},
	}

	return command
}

func actionList(ctx *cli.Context) {
	// TODO: Implement this.
}
