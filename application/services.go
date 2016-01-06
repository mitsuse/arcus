package application

import (
	"errors"
	"regexp"

	"github.com/mitsuse/arcus/domain/device"
	"github.com/mitsuse/arcus/domain/resource"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

// `ListDevices` obtains the list of devices connected to the account authorized with the given `token`.
func ListDevices(token string) (devices []*device.Type, err error) {
	if len(token) == 0 {
		message := "The environment variable \"ARCUS_ACCESS_TOKEN\" should not be empty."
		return nil, errors.New(message)
	}

	return device.ListPushable(token)
}

// `Send` pushes the  given data to devices connected to the account authorized with the give `token`.
func Send(token, title, message, location, deviceName string) error {
	if len(token) == 0 {
		message := "The environment variable \"ARCUS_ACCESS_TOKEN\" should not be empty."
		return errors.New(message)
	}

	var deviceId string
	if len(deviceName) == 0 {
		deviceId = ""
	} else {
		device, err := device.GetPushableByName(token, deviceName)
		if err != nil {
			return err
		}

		deviceId = device.Id()
	}

	// TODO: Print an error message easy to understand.
	return send(token, deviceId, title, message, location)
}

func send(token, deviceId, title, message, location string) error {
	if len(location) == 0 {
		push := requests.NewNote()
		push.Title = title
		push.Body = message
		push.DeviceIden = deviceId

		_, err := pushbullet.New(token).PostPushesNote(push)
		return err
	}

	if isLink(location) {
		push := requests.NewLink()
		push.Title = title
		push.Body = message
		push.Url = location
		push.DeviceIden = deviceId

		_, err := pushbullet.New(token).PostPushesLink(push)
		return err
	}

	return upload(token, deviceId, title, message, location)
}

func upload(token, deviceId, title, message, location string) error {
	l, err := resource.Read(location)
	if err != nil {
		return err
	}

	r, err := resource.Upload(token, l)
	if err != nil {
		return err
	}

	push := requests.NewFile()
	push.Title = title
	push.Body = message
	push.FileName = r.Name()
	push.FileType = r.Type()
	push.FileUrl = r.Url()
	push.DeviceIden = deviceId

	if _, err := pushbullet.New(token).PostPushesFile(push); err != nil {
		return err
	}

	return nil
}

func isLink(location string) bool {
	return regexp.MustCompile(`^https?://`).MatchString(location)
}
