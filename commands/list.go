package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/app"
	"github.com/mitsuse/bullet/pushbullet"
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
	token := os.Getenv("BULLET_ACCESS_TOKEN")
	if len(token) == 0 {
		// TODO: Set an error message.
		app.PrintError(errors.New(""))
		return
	}

	pb := pushbullet.New(token)

	res, err := pb.GetDevices()
	if err != nil {
		app.PrintError(err)
		return
	}

	for _, device := range res {
		if !device.Pushable {
			continue
		}

		fmt.Println(device.Nickname)
	}
}
