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
				Name:  "auth,a",
				Value: os.Getenv("HOME") + "/.auth.bullet",
				Usage: "The path of your authentication file",
			},
		},
	}

	return command
}

func actionSend(ctx *cli.Context) {
	authPath := ctx.String("auth")

	auth, err := loadAuthPath(authPath)
	if err != nil {
		printError(err)
		return
	}

	pb := pushbullet.New(auth.Token())

	// TODO: Implement this.
	_ = pb
}

func loadAuthPath(path string) (*bullet.Auth, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bullet.LoadAuth(file)
}
