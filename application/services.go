package application

import (
	"errors"
	"mime"
	"os"
	"path"
	"regexp"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
	"github.com/mitsuse/pushbullet-go/responses"
)

// `ListDevices` obtains the list of devices connected to the account authorized with the given `token`.
func ListDevices(token string) (devices []*responses.Device, err error) {
	if len(token) == 0 {
		message := "The environment variable \"ARCUS_ACCESS_TOKEN\" should not be empty."
		return nil, errors.New(message)
	}

	client := pushbullet.New(token)

	return client.GetDevices()
}

// `Send` pushes the  given data to devices connected to the account authorized with the give `token`.
func Send(token, title, message, location, device string) error {
	if len(token) == 0 {
		message := "The environment variable \"ARCUS_ACCESS_TOKEN\" should not be empty."
		return errors.New(message)
	}

	client := pushbullet.New(token)

	deviceId, err := getDeviceId(client, device)
	if err != nil {
		return err
	}

	// TODO: Print an error message easy to understand.
	return send(client, deviceId, title, message, location)
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
