package commands

import (
	"github.com/codegangsta/cli"
)

func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "send text or files.",
		Action:    actionSend,
	}

	return command
}

func actionSend(ctx *cli.Context) {
	// TODO: Implement this.
}
