package commands

import (
	"mime"
	"os"
	"path"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/pushbullet"
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "Send a message or a file",
		Action:    actionSend,

		Flags: []cli.Flag{
			configFlag(),

			cli.StringFlag{
				Name:  "title,t",
				Value: "",
				Usage: "The title of the message or file to be sent",
			},

			cli.StringFlag{
				Name:  "message,m",
				Value: "",
				Usage: "The message to be sent",
			},

			cli.StringFlag{
				Name:  "location,l",
				Value: "",
				Usage: "The path of file or link to be sent",
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

	title := ctx.String("title")
	message := ctx.String("message")
	location := ctx.String("location")

	if err := send(pb, title, message, location); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}

func send(pb *pushbullet.Pushbullet, title, message, location string) error {
	if len(location) == 0 {
		note := &pushes.Note{
			Type:  pushes.TYPE_NOTE,
			Title: title,
			Body:  message,
		}

		return pb.PostPushesNote(note)
	}

	if isLink(location) {
		link := &pushes.Link{
			Type:  pushes.TYPE_LINK,
			Title: title,
			Body:  message,
			Url:   location,
		}

		return pb.PostPushesLink(link)
	}

	return upload(pb, title, message, location)
}

func upload(pb *pushbullet.Pushbullet, title, message, location string) error {
	fileName := path.Base(location)
	fileType := mime.TypeByExtension(path.Ext(fileName))

	res, err := pb.PostUploadRequest(fileName, fileType)
	if err != nil {
		return err
	}

	file, err := os.Open(location)
	if err != nil {
		return err
	}

	if err := pushbullet.Upload(res, file); err != nil {
		return err
	}

	push := &pushes.File{
		Type:     pushes.TYPE_FILE,
		Title:    title,
		Body:     message,
		FileName: res.FileName,
		FileType: res.FileType,
		FileUrl:  res.FileUrl,
	}

	return pb.PostPushesFile(push)
}

func isLink(location string) bool {
	return regexp.MustCompile(`^https?://`).MatchString(location)
}
