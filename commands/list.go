package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/arcus/application"
)

/*
Create "list" sub-command.
This sub-command lists the names of registered devices on Pushbullet.
*/
func NewListCommand() cli.Command {
	command := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "List devices that can be pushed to",
		Action:    actionList,
	}

	return command
}

func actionList(ctx *cli.Context) {
	token := os.Getenv("ARCUS_ACCESS_TOKEN")

	devices, err := application.ListDevices(token)
	if err != nil {
		application.ExitWith(err)
		return
	}

	for _, d := range devices {
		if !d.Pushable {
			continue
		}

		fmt.Println(d.Nickname)
	}
}
