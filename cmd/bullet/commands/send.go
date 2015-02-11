package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet"
	"github.com/mitsuse/bullet/pushbullet"
)

func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "send text or files.",
		Action:    actionSend,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config,c",
				Value: os.Getenv("HOME") + "/.config.bullet",
				Usage: "The path of your config file",
			},
		},
	}

	return command
}

func actionSend(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := loadConfigPath(configPath)
	if err != nil {
		printError(err)
		return
	}

	pb := pushbullet.New(config.Token())

	// TODO: Implement this.
	_ = pb
}

func loadConfigPath(path string) (*bullet.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bullet.LoadConfig(file)
}
