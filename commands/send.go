package commands

import (
	"errors"
	"mime"
	"os"
	"path"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

/*
Create "send" sub-command.
This sub-command sends a "push" to devices such as "note", "link" or "file".
*/
func NewSendCommand() cli.Command {
	command := cli.Command{
		Name:      "send",
		ShortName: "s",
		Usage:     "Send a message or a file",
		Action:    actionSend,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "device,d",
				Value: "",
				Usage: "The name of target device",
			},

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
	token := os.Getenv("ARCUS_ACCESS_TOKEN")
	if len(token) == 0 {
		message := "The environment variable \"ARCUS_ACCESS_TOKEN\" should not be empty."
		printError(errors.New(message))
		return
	}

	pb := pushbullet.New(token)

	title := ctx.String("title")
	message := ctx.String("message")
	location := ctx.String("location")
	device := ctx.String("device")

	deviceId, err := getDeviceId(pb, device)
	if err != nil {
		printError(err)
		return
	}

	if err := send(pb, deviceId, title, message, location); err != nil {
		// TODO: Print an error message easy to understand.
		printError(err)
		return
	}
}

func send(pb *pushbullet.Pushbullet, deviceId, title, message, location string) error {
	if len(location) == 0 {
		push := requests.NewNote()
		push.Title = title
		push.Body = message
		push.DeviceIden = deviceId

		_, err := pb.PostPushesNote(push)
		return err
	}

	if isLink(location) {
		push := requests.NewLink()
		push.Title = title
		push.Body = message
		push.Url = location
		push.DeviceIden = deviceId

		_, err := pb.PostPushesLink(push)
		return err
	}

	return upload(pb, deviceId, title, message, location)
}

func upload(pb *pushbullet.Pushbullet, deviceId, title, message, location string) error {
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

	if err := pushbullet.Upload(pb.Client(), res, file); err != nil {
		return err
	}

	push := requests.NewFile()
	push.Title = title
	push.Body = message
	push.FileName = res.FileName
	push.FileType = res.FileType
	push.FileUrl = res.FileUrl
	push.DeviceIden = deviceId

	if _, err := pb.PostPushesFile(push); err != nil {
		return err
	}

	return nil
}

func getDeviceId(pb *pushbullet.Pushbullet, name string) (string, error) {
	if len(name) > 0 {
		res, err := pb.GetDevices()
		if err != nil {
			return "", err
		}

		for _, device := range res {
			if device.Nickname == name {
				return device.Iden, nil
			}
		}
	} else {
		return "", nil
	}

	return "", errors.New("No such device.")
}

func isLink(location string) bool {
	return regexp.MustCompile(`^https?://`).MatchString(location)
}
