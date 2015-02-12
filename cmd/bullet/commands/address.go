package commands

import (
	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewAddressCommand() cli.Command {
	command := cli.Command{
		Name:      "address",
		ShortName: "a",
		Usage:     "Send an address",
		Action:    actionAddress,

		Flags: []cli.Flag{
			configFlag(),

			cli.StringFlag{
				Name:  "name,n",
				Value: "",
				Usage: "The place's name",
			},

			cli.StringFlag{
				Name:  "address,a",
				Value: "",
				Usage: "The place's address or a map search query.",
			},
		},
	}

	return command
}

func actionAddress(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	name := ctx.String("name")
	address := ctx.String("address")
	addr := pushes.NewAddress(name, address)

	pb := pushbullet.New(config.Token())

	if err := pb.PostAddress(addr); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}
